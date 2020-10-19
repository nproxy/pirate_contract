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
	"math/big"
	"strings"
)

type UserChargeHistory struct {
	BlockPos
	TokenAmount *big.Int `json:"token_amount"`
}

type UserCharge struct {
	User    common.Address
	History []*UserChargeHistory
}

var userChargeEvent EventPos

var usersInPool map[common.Address]map[common.Address]*UserCharge

func GetUserList(pool common.Address) ([]common.Address,error)  {
	v,ok:=usersInPool[pool]
	if !ok{
		return nil,errors.New("no pool in mem")
	}

	var uas []common.Address

	for k,_:=range v{
		uas = append(uas,k)
	}

	return uas,nil
}

var UserChargeNotify func(user common.Address, pool common.Address, tokenAmount *big.Int) error

var userChargeKeyHead = "user_charge_"
var userChargeKey = userChargeKeyHead + "%s_%s_%d_%d"

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

func addNewUserChargeHistory(pool, user common.Address, l types.Log, tokenAmount *big.Int) {
	v, ok := usersInPool[pool]
	if !ok {
		usersInPool[pool] = make(map[common.Address]*UserCharge)
		v = usersInPool[pool]
	}
	_, ok = v[user]
	if !ok {
		v[user] = &UserCharge{User: user}
	}

	for _, history := range v[user].History {
		if history.BlockNumber == l.BlockNumber && history.TxIndex == l.TxIndex {
			return
		}
	}

	uc := v[user]
	h := &UserChargeHistory{BlockPos: BlockPos{BlockNumber: l.BlockNumber, TxIndex: l.TxIndex}, TokenAmount: tokenAmount}

	uc.History = append(uc.History, h)

	key := getUserChargeKey(pool, user, &h.BlockPos)
	dbv, _ := json.Marshal(*h)

	GetLogConf().Save([]byte(key), dbv)

	if UserChargeNotify != nil{
		UserChargeNotify(user, pool, tokenAmount)
	}

}

func batchUserCharge() error {
	if userChargeEvent.IsDone() {
		return nil
	}

	mc := GetLogConf().NewMarketClient()
	defer mc.Close()

	var opt *bind.FilterOpts

	if !userChargeEvent.IsFirst() {
		opt = &bind.FilterOpts{
			Start:   userChargeEvent.StartPos(),
			End:     nil,
			Context: context.TODO(),
		}
	}

	iter, err := mc.Market.FilterCharge(opt, nil, nil)
	if err != nil {
		fmt.Print("batch user charge failed")
		return err
	}

	for iter.Next() {
		ev := iter.Event
		addNewUserChargeHistory(ev.Pool, ev.User, ev.Raw, ev.TokenAmount)
		userChargeEvent.LastMax(ev.Raw)
	}

	userChargeEvent.LastBlkNum()

	return nil
}

func watchUserCharge(batch *chan *LogServiceItem) error {
	ec := GetLogConf().NewMarketClient()
	defer ec.Close()

	c := make(chan *contract.TrafficMarketCharge, 1024)

	sub, err := ec.Market.WatchCharge(nil, c, nil, nil)
	if err != nil {
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
	alluc := GetLogConf().BatchGet([]byte(userChargeKeyHead), nil)

	for i := 0; i < len(alluc); i++ {
		uc := alluc[i]

		pool, user, err := userChargeKey2Address(uc.key)
		if err != nil {
			fmt.Println("key", string(uc.key), err)
			continue
		}
		v, ok := usersInPool[pool]
		if !ok {
			usersInPool[pool] = make(map[common.Address]*UserCharge)
			v = usersInPool[pool]
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

	usersInPool = make(map[common.Address]map[common.Address]*UserCharge)

	logUserChargeSrvItem = &LogServiceItem{}
	logUserChargeSrvItem.name = "Charge"
	logUserChargeSrvItem.stop = make(chan struct{})
	logUserChargeSrvItem.watch = watchUserCharge
	logUserChargeSrvItem.batch = batchUserCharge
	logUserChargeSrvItem.recover = recoverUserCharge
	logUserChargeSrvItem.CurBlockNum = curUserChargeBlockN

	GetLogService().RegLogSrv(logUserChargeSrvItem)

}
