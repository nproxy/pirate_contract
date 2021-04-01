package logService

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperorchidlab/pirate_contract/cabinet"
	"github.com/hyperorchidlab/pirate_contract/contract"
	"github.com/hyperorchidlab/pirate_contract/util"
	"math/big"
	"strings"
	"sync"
)

type UserChargeHistory struct {
	BlockPos
	BlockTime 	  int64		`json:"block_time,omitempty"`
	TokenAmount   *big.Int `json:"token_amount"`
	TrafficAmount *big.Int `json:"traffic_amount"`
}

type UserCharge struct {
	User          common.Address
	TokenAmount   *big.Int             `json:"token_amount"`
	TrafficAmount *big.Int             `json:"traffic_amount"`
	History       []*UserChargeHistory `json:"history"`
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

func GetUserChargeCount() int {
	usersInPool.lock.Lock()
	defer usersInPool.lock.Unlock()

	cnt:=0

	u:=make(map[common.Address]struct{})

	for _,pooluser:=range usersInPool.users{
		for k,_:=range pooluser{
			if _,ok:=u[k];!ok{
				cnt ++
				u[k] = struct{}{}
			}
		}
	}

	return cnt

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

var UserChargeNotify func(user common.Address, pool common.Address, tokenAmount *big.Int, lastAmount *big.Int,firstTime bool,chargeTime int64) error

var RecoverChargeNotify func(user common.Address,pool common.Address,tokenAmount *big.Int, lastAmout *big.Int, firstTime bool,chargeTime int64) error

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

func _addNewUserChargeHistory(pool, user common.Address, l types.Log, tokenAmount *big.Int, trafficAmount *big.Int,nowtime int64) (bool,bool,*big.Int) {
	usersInPool.lock.Lock()
	defer usersInPool.lock.Unlock()

	var firstTime bool
	lastAmount := &big.Int{}

	v, ok := usersInPool.users[pool]
	if !ok {
		usersInPool.users[pool] = make(map[common.Address]*UserCharge)
		v = usersInPool.users[pool]
	}
	_, ok = v[user]
	if !ok {
		v[user] = &UserCharge{User: user, TokenAmount: &big.Int{}, TrafficAmount: &big.Int{}}
	}

	for _, history := range v[user].History {
		if history.BlockNumber == l.BlockNumber && history.TxIndex == l.TxIndex {
			return false,false,lastAmount
		}
	}
	lastAmount = v[user].TokenAmount

	if len(v[user].History) == 0{
		firstTime = true
	}

	fmt.Println("add user-----", v[user].TrafficAmount.String(), v[user].TokenAmount.String(), tokenAmount.String(), trafficAmount.String())

	v[user].TrafficAmount = util.MaxBigInt(trafficAmount, v[user].TrafficAmount)
	v[user].TokenAmount = util.MaxBigInt(tokenAmount, v[user].TokenAmount)

	fmt.Println("UserCharge: ", pool.String(), user.String(), v[user].TokenAmount, v[user].TrafficAmount)

	uc := v[user]
	h := &UserChargeHistory{BlockPos: BlockPos{BlockNumber: l.BlockNumber, TxIndex: l.TxIndex}, TokenAmount: tokenAmount,BlockTime: nowtime}

	uc.History = append(uc.History, h)

	key := getUserChargeKey(pool, user, &h.BlockPos)
	dbv, _ := json.Marshal(*h)

	fmt.Println("UserCharge: Save to db", key, string(dbv))

	GetLogConf().Save([]byte(key), dbv)

	return true,firstTime,lastAmount
}

func addNewUserChargeHistory(pool, user common.Address, l types.Log, tokenAmount *big.Int, trafficAmount *big.Int,nowTime int64) {

	n,firstTimne,lastAmount := _addNewUserChargeHistory(pool, user, l, tokenAmount, trafficAmount,nowTime)

	if n && UserChargeNotify != nil {
		fmt.Println("UserCharge Notify:", pool.String(), user.String(), tokenAmount.String(), trafficAmount.String())
		UserChargeNotify(user, pool, tokenAmount, lastAmount,firstTimne,nowTime)
	}
}

func getBlockTime(c *ethclient.Client,hash common.Hash) int64  {
	if blk,err:= c.BlockByHash(context.TODO(),hash);err!=nil{
		return 0
	}else{
		return int64(blk.Time())
	}
}

func getBlockTime2(c *ethclient.Client, num uint64) int64  {
	bignum := &big.Int{}
	bignum.SetUint64(num)

	if blk,err:=c.BlockByNumber(context.TODO(),bignum);err!=nil{
		return 0
	}else{
		return int64(blk.Time())
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
		nowtime:=getBlockTime(mc.GetClient(),iter.Event.Raw.BlockHash)
		addNewUserChargeHistory(ev.Pool, ev.User, ev.Raw, ev.TokenAmount, ev.TrafficAmount,nowtime)
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
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!watch user charge failed!!!!!!!!!!!!!!!!!!!!!!!=====================")
		return e
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
	c,err:=GetLogConf().cfg.NewEthClient()
	if err!=nil{
		return err
	}
	defer c.Close()

	usersInPool.lock.Lock()
	defer usersInPool.lock.Unlock()
	for i := 0; i < len(alluc); i++ {

		uc := alluc[i]
		fmt.Println("UserCharge: Recover from Db", string(uc.key), string(uc.vaule))
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
			v[user] = &UserCharge{User: user, TokenAmount: &big.Int{}, TrafficAmount: &big.Int{}}
		}

		dbv := &UserChargeHistory{}
		json.Unmarshal(uc.vaule, dbv)

		if dbv.BlockTime == 0{
			dbv.BlockTime = getBlockTime2(c,dbv.BlockNumber)
			if dbv.BlockTime > 0{
				data,_:=json.Marshal(dbv)
				GetLogConf().Save(uc.key,data)
			}
		}

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

		if RecoverChargeNotify != nil{
			//RecoverChargeNotify(user,pool,dbv.TokenAmount,lastAmount,firstTime,dbv.BlockTime)
			recoverNotify(user,pool,dbv.TokenAmount,dbv.BlockTime)
		}

		muc := v[user]
		fmt.Println(dbv.TrafficAmount.String(), dbv.TokenAmount.String())
		muc.TokenAmount = util.MaxBigInt(dbv.TokenAmount, muc.TokenAmount)
		muc.TrafficAmount = util.MaxBigInt(dbv.TrafficAmount, muc.TrafficAmount)
		fmt.Println("UserCharge: Recover from db to mem", pool.String(), user.String(), muc.TokenAmount, muc.TrafficAmount)
		userChargeEvent.LastMax2(dbv.BlockNumber, dbv.TxIndex)

		muc.History = append(muc.History, dbv)

	}

	recoverNotifyCommit()

	return nil

}

type NotifyItem struct {
	TokenAmount *big.Int
	BlockTime int64
}

type ChargeNotify struct {
	User common.Address
	Pool common.Address
	Items []*NotifyItem
}

var globalChargeNotifyMem *ChargeNotify

func insert(items []*NotifyItem,item *NotifyItem) []*NotifyItem {

	idx:=-1

	var outs []*NotifyItem

	for i:=0;i<len(items);i++{
		if item.BlockTime < items[i].BlockTime{
			idx = i
			break
		}
	}

	if idx == -1 || idx == 0{
		outs = append(outs,item)
		outs = append(outs,items...)
	}else{
		outs = append(outs,items[:idx]...)
		outs = append(outs,item)
		outs = append(outs,items[idx:]...)
	}

	return outs
}

func recoverNotifyCommit()  {
	var items []*NotifyItem
	if len(globalChargeNotifyMem.Items)>1{
		for i:=0;i<len(globalChargeNotifyMem.Items);i++{
			item:=globalChargeNotifyMem.Items[i]
			items = insert(items,item)
		}
	}else{
		items = globalChargeNotifyMem.Items
	}

	lastAmount := &big.Int{}
	firstTime := true

	for i:=0;i<len(items);i++{
		m:=items[i]
		RecoverChargeNotify(globalChargeNotifyMem.User,globalChargeNotifyMem.Pool,m.TokenAmount,lastAmount,firstTime,m.BlockTime)
		lastAmount = m.TokenAmount
		firstTime = false
	}

}

func recoverNotify(user,pool common.Address, amount *big.Int, blkTime int64)  {

	if globalChargeNotifyMem == nil || user != globalChargeNotifyMem.User{
		if globalChargeNotifyMem != nil{
			recoverNotifyCommit()
		}

		globalChargeNotifyMem = &ChargeNotify{
			User: user,
			Pool: pool,
		}
	}

	item:=&NotifyItem{
		TokenAmount: amount,
		BlockTime: blkTime,
	}

	globalChargeNotifyMem.Items = append(globalChargeNotifyMem.Items,item)

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
