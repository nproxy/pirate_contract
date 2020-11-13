package logService

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hyperorchidlab/pirate_contract/cabinet"
	"github.com/hyperorchidlab/pirate_contract/contract"
	"github.com/hyperorchidlab/pirate_contract/util"
	"math/big"
	"strings"
	"sync"
)

type UserChargeHistory struct {
	BlockPos
	TokenAmount   *big.Int `json:"token_amount"`
	TrafficAmount *big.Int `json:"traffic_amount"`
}

type UserCharge struct {
	User          common.Address
	TokenAmount   *big.Int `json:"token_amount"`
	TrafficAmount *big.Int `json:"traffic_amount"`
	History       []*UserChargeHistory
}

var userChargeEvent EventPos

type UserStore struct {
	users map[common.Address]map[common.Address]*UserCharge
	lock  sync.Mutex
}

var usersInPool *UserStore

func GetUserList(pool common.Address) ([]common.Address, error) {

	usersInPool.lock.Lock()
	defer usersInPool.lock.Unlock()

	v, ok := usersInPool.users[pool]
	if !ok {
		return nil, errors.New("no pool in mem")
	}

	var uas []common.Address

	for k, _ := range v {
		uas = append(uas, k)
	}

	return uas, nil
}

func GetSubPools(user common.Address) []common.Address {
	usersInPool.lock.Lock()
	defer usersInPool.lock.Unlock()

	var pools []common.Address

	for k, v := range usersInPool.users {
		if _, ok := v[user]; ok {
			pools = append(pools, k)
		}
	}

	return pools
}

func GetUserData(pool, user common.Address) *cabinet.PirateEthUserData {
	usersInPool.lock.Lock()
	defer usersInPool.lock.Unlock()

	v, ok := usersInPool.users[pool]
	if !ok {
		return nil
	}
	var uc *UserCharge
	uc, ok = v[user]
	if !ok {
		return nil
	}

	return &cabinet.PirateEthUserData{TotalTraffic: uc.TrafficAmount, ChargeBalance: uc.TokenAmount}

}

var UserChargeNotify func(user common.Address, pool common.Address, tokenAmount *big.Int, trafficAmount *big.Int) error

var userChargeKeyHead = "user_charge_"
var userChargeKey = userChargeKeyHead + "%s_%s_%d_%d"
var userChargeKeyPatternEnd = userChargeKeyHead + "0xffffffffffffffffffff"

func getUserChargeKey(pool, user common.Address, pos *BlockPos) string {
	return fmt.Sprintf(userChargeKey, pool.String(), user.String(), pos.BlockNumber, pos.TxIndex)
}

func userChargeKey2Address(key []byte) (pool, user common.Address, err error) {
	ks := string(key)

	ksarr := strings.Split(ks, "_")
	if len(ksarr) < 6 {
		return pool, user, errors.New("key error")
	}

	pool = common.HexToAddress(ksarr[2])
	user = common.HexToAddress(ksarr[3])

	return
}

func _addNewUserChargeHistory(pool, user common.Address, l types.Log, tokenAmount *big.Int, trafficAmount *big.Int) bool {
	usersInPool.lock.Lock()
	defer usersInPool.lock.Unlock()

	v, ok := usersInPool.users[pool]
	if !ok {
		usersInPool.users[pool] = make(map[common.Address]*UserCharge)
		v = usersInPool.users[pool]
	}
	_, ok = v[user]
	if !ok {
		v[user] = &UserCharge{User: user}
	}

	for _, history := range v[user].History {
		if history.BlockNumber == l.BlockNumber && history.TxIndex == l.TxIndex {
			return false
		}
	}

	v[user].TrafficAmount = util.MaxBigInt(trafficAmount, v[user].TrafficAmount)
	v[user].TokenAmount = util.MaxBigInt(tokenAmount, v[user].TokenAmount)

	uc := v[user]
	h := &UserChargeHistory{BlockPos: BlockPos{BlockNumber: l.BlockNumber, TxIndex: l.TxIndex}, TokenAmount: tokenAmount}

	uc.History = append(uc.History, h)

	key := getUserChargeKey(pool, user, &h.BlockPos)
	dbv, _ := json.Marshal(*h)

	GetLogConf().Save([]byte(key), dbv)

	return true
}

func addNewUserChargeHistory(pool, user common.Address, l types.Log, tokenAmount *big.Int, trafficAmount *big.Int) {

	n := _addNewUserChargeHistory(pool, user, l, tokenAmount, trafficAmount)

	if n && UserChargeNotify != nil {
		UserChargeNotify(user, pool, tokenAmount, trafficAmount)
	}
}

func batchUserCharge() error {
	if userChargeEvent.IsDone() {
		return nil
	}

	mc, err := GetLogConf().NewMarketClient()
	if err != nil {
		return err
	}
	defer mc.Close()

	var opt *bind.FilterOpts

	if !userChargeEvent.IsFirst() {
		opt = &bind.FilterOpts{
			Start:   userChargeEvent.StartPos(),
			End:     nil,
			Context: context.TODO(),
		}
	}

	iter, e := mc.FilterCharge(opt, nil, nil)
	if e != nil {
		fmt.Print("batch user charge failed")
		return e
	}

	for iter.Next() {
		ev := iter.Event
		addNewUserChargeHistory(ev.Pool, ev.User, ev.Raw, ev.TokenAmount, ev.TrafficAmount)
		userChargeEvent.LastMax(ev.Raw)
	}

	userChargeEvent.LastBlkNum()

	return nil
}

func watchUserCharge(batch *chan *LogServiceItem) error {
	mc, err := GetLogConf().NewWSMarketClient()
	if err != nil {
		return err
	}
	defer mc.Close()

	c := make(chan *contract.TrafficMarketCharge, 1024)

	sub, e := mc.WatchCharge(nil, c, nil, nil)
	if e != nil {
		return err
	}
	for {
		select {
		case pc := <-c:
			userChargeEvent.NextMax(pc.Raw)
			ch := *batch

			ch <- logUserChargeSrvItem
		case err = <-sub.Err():
			return err
		}
	}
}

func recoverUserCharge() error {
	alluc := GetLogConf().BatchGet([]byte(userChargeKeyHead), []byte(userChargeKeyPatternEnd))

	usersInPool.lock.Lock()
	defer usersInPool.lock.Unlock()
	for i := 0; i < len(alluc); i++ {
		uc := alluc[i]

		pool, user, err := userChargeKey2Address(uc.key)
		if err != nil {
			fmt.Println("key", string(uc.key), err)
			continue
		}
		v, ok := usersInPool.users[pool]
		if !ok {
			usersInPool.users[pool] = make(map[common.Address]*UserCharge)
			v = usersInPool.users[pool]
		}
		_, ok = v[user]
		if !ok {
			v[user] = &UserCharge{User: user}
		}
		dbv := &UserChargeHistory{}
		json.Unmarshal(uc.vaule, dbv)

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

		muc := v[user]
		muc.TokenAmount = util.MaxBigInt(dbv.TokenAmount, muc.TokenAmount)
		muc.TrafficAmount = util.MaxBigInt(dbv.TrafficAmount, muc.TrafficAmount)

		userChargeEvent.LastMax2(dbv.BlockNumber, dbv.TxIndex)

		muc.History = append(muc.History, dbv)

	}

	return nil

}

var logUserChargeSrvItem *LogServiceItem

func curUserChargeBlockN(n uint64) {
	logUserChargeSrvItem.evPos.SetCurrent(n)
}

func init() {

	store := make(map[common.Address]map[common.Address]*UserCharge)
	usersInPool = &UserStore{users: store}

	logUserChargeSrvItem = &LogServiceItem{}
	logUserChargeSrvItem.name = "Charge"
	logUserChargeSrvItem.stop = make(chan struct{})
	logUserChargeSrvItem.watch = watchUserCharge
	logUserChargeSrvItem.batch = batchUserCharge
	logUserChargeSrvItem.recover = recoverUserCharge
	logUserChargeSrvItem.CurBlockNum = curUserChargeBlockN

	GetLogService().RegLogSrv(logUserChargeSrvItem)

}
