package logService

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hyperorchidlab/pirate_contract/contract"
	"github.com/hyperorchidlab/pirate_contract/storageService"
	"math/big"
	"strings"
	"sync"
)

type PoolClaimData struct {
	PoolAddr    common.Address `json:"pool_addr"`
	UserAddr    common.Address `json:"user_addr"`
	UsedTraffic *big.Int
	History     []*PoolClaimHistory
}

type PoolClaimHistory struct {
	MinerUsedPacket *big.Int
	MinerPacket     *big.Int
	ClaimedBalance  *big.Int
	PoolTotalPacket *big.Int
	BlockPos
}

var poolClaimEventPos EventPos

type PoolClaimStore struct {
	claims map[common.Address]map[common.Address]*PoolClaimData
	lock   sync.Mutex
}

//pool addr          user addr
var poolClaimUser *PoolClaimStore

var PoolClaimNotify func(pool, user common.Address, pch *PoolClaimHistory) error

var poolClaimKeyHead = "pool_claim_"
var poolClaimKey = poolClaimKeyHead + "%s_%s_%d_%d"
var poolClaimKeyPattenEnd = poolClaimKeyHead + "0xffffffffffffffffffff"

func GetPoolClaim(pool common.Address, user common.Address) *PoolClaimHistory {
	poolClaimUser.lock.Lock()
	defer poolClaimUser.lock.Unlock()

	v, ok := poolClaimUser.claims[pool]
	if !ok {
		return nil
	}

	var pd *PoolClaimData
	pd, ok = v[user]
	if !ok {
		return nil
	}

	if len(pd.History) == 0 {
		return nil
	}

	maxh := pd.History[0]

	for i := 1; i < len(pd.History); i++ {
		h := pd.History[i]
		if maxh.BlockNumber < h.BlockNumber {
			maxh = h
		}
		if maxh.BlockNumber == h.BlockNumber {
			if maxh.TxIndex < h.TxIndex {
				maxh = h
			}
		}
	}

	return maxh
}

func GetUsedTrafficFromPoolClaim(pool, user common.Address) (*big.Int, error) {
	poolClaimUser.lock.Lock()
	defer poolClaimUser.lock.Unlock()

	v, ok := poolClaimUser.claims[pool]
	if !ok {
		return nil, errors.New("no pool")
	}

	var pd *PoolClaimData
	pd, ok = v[user]
	if !ok {
		return nil, errors.New("no user")
	}

	used := pd.UsedTraffic
	if used == nil {
		used = &big.Int{}
	}

	return used, nil

}

func getPoolClaimKey(pool, user common.Address, pos *BlockPos) string {
	return fmt.Sprintf(poolClaimKey, pool.String(), user.String(), pos.BlockNumber, pos.TxIndex)
}

func poolClaimKey2Address(key []byte) (pool common.Address, user common.Address, err error) {
	ks := string(key)

	ksarr := strings.Split(ks, "_")

	if len(ksarr) < 6 {
		return pool, user, errors.New("key error")
	}

	pool = common.HexToAddress(ksarr[2])
	user = common.HexToAddress(ksarr[3])

	return
}

func getUsedTraffic(expectAdd *big.Int, left *big.Int, realAdd *big.Int, curUsed *big.Int) *big.Int {
	if left == nil {
		left = &big.Int{}
	}

	if realAdd == nil {
		realAdd = &big.Int{}
	}

	if left.Cmp(realAdd) > 0 {
		return curUsed
	}

	if expectAdd == nil {
		expectAdd = &big.Int{}
	}

	if curUsed == nil {
		curUsed = &big.Int{}
	}
	z := &big.Int{}

	z = z.Sub(expectAdd, realAdd)

	u := &big.Int{}

	t, _ := storageService.Balance2Traffic(z)

	u = u.Sub(curUsed, t)

	return u
}

func _addNewPoolClaimnHistory(pool, user common.Address, l types.Log, minerUsedPacket, minerPacket, claimedBalance, poolTotalPacket *big.Int) (bool, *PoolClaimHistory) {
	poolClaimUser.lock.Lock()
	defer poolClaimUser.lock.Unlock()

	v, ok := poolClaimUser.claims[pool]
	if !ok {
		poolClaimUser.claims[pool] = make(map[common.Address]*PoolClaimData)
		v = poolClaimUser.claims[pool]
	}

	_, ok = v[user]
	if !ok {
		v[user] = &PoolClaimData{PoolAddr: pool, UserAddr: user, UsedTraffic: &big.Int{}}
	}

	for _, history := range v[user].History {
		if history.BlockNumber == l.BlockNumber && history.TxIndex == l.TxIndex {
			return false, nil
		}
	}

	poolhistory := v[user]
	h := &PoolClaimHistory{MinerUsedPacket: minerUsedPacket, MinerPacket: minerPacket, ClaimedBalance: claimedBalance, PoolTotalPacket: poolTotalPacket,
		BlockPos: BlockPos{BlockNumber: l.BlockNumber, TxIndex: l.TxIndex}}
	poolhistory.History = append(poolhistory.History, h)

	used := getUsedTraffic(minerUsedPacket, minerPacket, claimedBalance, poolTotalPacket)
	if v[user].UsedTraffic.Cmp(used) < 0 {
		v[user].UsedTraffic = used
	}

	k := getPoolClaimKey(pool, user, &h.BlockPos)

	dbv, _ := json.Marshal(*h)

	GetLogConf().Save([]byte(k), dbv)

	return true, h
}

func addNewPoolClaimnHistory(pool, user common.Address, l types.Log, traffic, token, microNonce, claimNonce *big.Int) {
	n, h := _addNewPoolClaimnHistory(pool, user, l, traffic, token, microNonce, claimNonce)
	//GetLogConf().db.Put([]byte(k),dbv,nil)
	if n && PoolClaimNotify != nil {
		PoolClaimNotify(pool, user, h)
	}

}

func batchPoolClaim() error {

	if poolClaimEventPos.IsDone() {
		return nil
	}

	mc, err := GetLogConf().NewMarketClient()
	if err != nil {
		return err
	}
	defer mc.Close()

	var opt *bind.FilterOpts

	if !poolClaimEventPos.IsFirst() {
		opt = &bind.FilterOpts{
			Start:   poolClaimEventPos.StartPos(),
			End:     nil,
			Context: context.TODO(),
		}
	}

	iter, e := mc.FilterPoolClaim(opt, nil, nil)
	if e != nil {
		fmt.Print("batch Pool Claim failed")
		return e
	}

	for iter.Next() {
		ev := iter.Event
		addNewPoolClaimnHistory(ev.Pool, ev.User, ev.Raw, ev.MinerUsedPacket, ev.MinerPacket, ev.ClaimedBalance, ev.PoolTotalPacket)
		poolClaimEventPos.LastMax(ev.Raw)
	}
	poolClaimEventPos.LastBlkNum()

	return nil
}

var logPoolClaimSrvItem *LogServiceItem

func watchPoolClaim(batch *chan *LogServiceItem) error {
	//fmt.Println("poolclaim xxxxxxxx--->,",GetLogConf().String())
	mc, err := GetLogConf().NewWSMarketClient()
	if err != nil {
		return err
	}
	defer mc.Close()

	c := make(chan *contract.TrafficMarketPoolClaim, 1024)

	sub, e := mc.WatchPoolClaim(nil, c, nil, nil)
	if e != nil {
		fmt.Println("====>", e)
		return e
	}

	for {
		select {
		case pc := <-c:
			poolClaimEventPos.NextMax(pc.Raw)
			ch := *batch

			ch <- logPoolClaimSrvItem
		case err = <-sub.Err():
			return err
		}
	}
}

func curPoolClaimBlockN(n uint64) {
	logPoolClaimSrvItem.evPos.SetCurrent(n)
}

func recoverPoolClaim() error {
	allcm := GetLogConf().BatchGet([]byte(poolClaimKeyHead), []byte(poolClaimKeyPattenEnd))

	poolClaimUser.lock.Lock()
	defer poolClaimUser.lock.Unlock()

	for i := 0; i < len(allcm); i++ {
		cm := allcm[i]
		pool, user, err := poolClaimKey2Address(cm.key)
		if err != nil {
			fmt.Println("key", string(cm.key), err)
			continue
		}

		v, ok := poolClaimUser.claims[pool]
		if !ok {
			poolClaimUser.claims[pool] = make(map[common.Address]*PoolClaimData)
			v = poolClaimUser.claims[pool]
		}

		_, ok = v[user]
		if !ok {
			v[user] = &PoolClaimData{PoolAddr: pool, UserAddr: user, UsedTraffic: &big.Int{}}
		}

		dbv := &PoolClaimHistory{}
		json.Unmarshal(cm.vaule, dbv)

		found := false

		for _, history := range v[user].History {
			if history.BlockNumber == dbv.BlockNumber && history.TxIndex == dbv.TxIndex {
				found = true
				break
			}
		}

		if found {
			continue
		}

		used := getUsedTraffic(dbv.MinerUsedPacket, dbv.MinerPacket, dbv.ClaimedBalance, dbv.PoolTotalPacket)
		if v[user].UsedTraffic.Cmp(used) < 0 {
			v[user].UsedTraffic = used
		}

		poolhistory := v[user]

		poolClaimEventPos.LastMax2(dbv.BlockNumber, dbv.TxIndex)

		poolhistory.History = append(poolhistory.History, dbv)

	}
	return nil

}

func init() {

	store := make(map[common.Address]map[common.Address]*PoolClaimData)
	poolClaimUser = &PoolClaimStore{claims: store}

	logPoolClaimSrvItem = &LogServiceItem{}
	logPoolClaimSrvItem.name = "PoolClaim"
	logPoolClaimSrvItem.stop = make(chan struct{})
	logPoolClaimSrvItem.watch = watchPoolClaim
	logPoolClaimSrvItem.batch = batchPoolClaim
	logPoolClaimSrvItem.recover = recoverPoolClaim
	logPoolClaimSrvItem.CurBlockNum = curPoolClaimBlockN

	GetLogService().RegLogSrv(logPoolClaimSrvItem)
}
