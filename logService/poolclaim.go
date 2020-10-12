package logService

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/btcsuite/goleveldb/leveldb/util"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hyperorchidlab/pirate_contract/contract"
	"math/big"
	"strings"
)


type PoolClaimData struct {
	PoolAddr common.Address `json:"pool_addr"`
	UserAddr common.Address `json:"user_addr"`
	History []*PoolClaimHistory
}

type PoolClaimHistory struct {
	Traffic *big.Int
	Token *big.Int
	MicroNonce *big.Int
	ClaimNonce *big.Int
	BlockPos
}

var eventPos EventPos

//                     pool addr          user addr
var poolClaimUser map[common.Address]map[common.Address]*PoolClaimData

var PoolClaimNotify func(pool,user common.Address, pch *PoolClaimHistory) error

var poolClaimKeyHead = "pool_claim_"
var poolClaimKey = poolClaimKeyHead+"%s_%s_%d_%d"

func GetPoolClaim(pool common.Address,user common.Address) *PoolClaimHistory  {
	v,ok:=poolClaimUser[pool]
	if !ok{
		return nil
	}

	var pd *PoolClaimData
	pd,ok = v[user]
	if !ok{
		return nil
	}

	if len(pd.History) == 0{
		return nil
	}

	maxh := pd.History[0]

	for i:=1;i<len(pd.History);i++{
		h:=pd.History[i]
		if maxh.BlockNumber < h.BlockNumber{
			maxh = h
		}
		if maxh.BlockNumber == h.BlockNumber{
			if maxh.TxIndex < h.TxIndex{
				maxh = h
			}
		}
	}

	return maxh
}

func getPoolClaimKey(pool,user common.Address,pos *BlockPos) string {
	return fmt.Sprintf(poolClaimKey,pool.String(),user.String(),pos.BlockNumber,pos.TxIndex)
}

func poolClaimKey2Address(key []byte) (pool common.Address,user common.Address,err error)  {
	ks:=string(key)

	ksarr:=strings.Split(ks,"_")

	if len(ksarr) < 6{
		return pool,user,errors.New("key error")
	}

	pool = common.HexToAddress(ksarr[2])
	user = common.HexToAddress(ksarr[3])

	return
}


func addNewPoolClaimnHistory(pool,user common.Address, l types.Log,traffic,token,microNonce,claimNonce *big.Int)  {
	v, ok:=poolClaimUser[pool]
	if !ok{
		poolClaimUser[pool] = make(map[common.Address]*PoolClaimData)
		v = poolClaimUser[pool]
	}

	_, ok = v[user]
	if !ok{
		v[user] = &PoolClaimData{PoolAddr: pool,UserAddr: user}
	}

	for _,history := range v[user].History{
		if history.BlockNumber == l.BlockNumber && history.TxIndex == l.TxIndex{
			return
		}
	}

	poolhistory := v[user]
	h:=&PoolClaimHistory{Traffic: traffic,Token: token,MicroNonce: microNonce,ClaimNonce: claimNonce,
		BlockPos:BlockPos{BlockNumber: l.BlockNumber,TxIndex: l.TxIndex}}
	poolhistory.History = append(poolhistory.History,h)

	k:=getPoolClaimKey(pool,user,&h.BlockPos)
	dbv,_:=json.Marshal(*h)

	GetLogConf().Save([]byte(k),dbv)
	//GetLogConf().db.Put([]byte(k),dbv,nil)

	PoolClaimNotify(pool,user,h)

}

func BatchPoolClaim() error  {

	if eventPos.IsDone(){
		return nil
	}

	mc := GetLogConf().NewMarketClient()
	defer mc.Close()

	var opt *bind.FilterOpts

	if !eventPos.IsFirst(){
		opt=&bind.FilterOpts{
			Start: eventPos.StartPos(),
			End: nil,
			Context: context.TODO(),
		}
	}

	iter, err:=mc.Market.FilterPoolClaim(opt,nil,nil)
	if err!=nil{
		fmt.Print("Batch Pool Claim failed")
		return err
	}

	for iter.Next(){
		ev:=iter.Event
		addNewPoolClaimnHistory(ev.Pool,ev.User,ev.Raw,ev.Packet,ev.Tonken,ev.MicrNonce,ev.ClaimNonce)
		eventPos.LastMax(ev.Raw)
	}
	eventPos.LastBlkNum()

	return nil
}

var logPoolClaimSrvItem *LogServiceItem

func watchPoolClaim(batch *chan *LogServiceItem) error {
	ec:=GetLogConf().NewMarketClient()
	defer ec.Close()

	c:=make(chan *contract.TrafficMarketPoolClaim,64)

	sub, err:=ec.Market.WatchPoolClaim(nil,c,nil,nil)
	if err!=nil{
		return err
	}

	for{
		select {
		case pc:=<-c:
			eventPos.NextMax(pc.Raw)
			ch:=*batch

			ch <- logPoolClaimSrvItem
		case err = <-sub.Err():
			return err
		}
	}
}

func curBlockN(n uint64)  {
	logPoolClaimSrvItem.evPos.SetCurrent(n)
}
//
//func recover() error {
//	iter:=GetLogConf().db.NewIterator(&util.Range{Start: []byte(poolClaimKeyHead)},nil)
//
//	for iter.Next(){
//		pool,user,err:=poolClaimKey2Address(iter.Key())
//		if err!=nil{
//			fmt.Println("key",string(iter.Key()),err)
//		}
//		v, ok:=poolClaimUser[pool]
//		if !ok{
//			poolClaimUser[pool] = make(map[common.Address]*PoolClaimData)
//			v = poolClaimUser[pool]
//		}
//
//		_, ok = v[user]
//		if !ok{
//			v[user] = &PoolClaimData{PoolAddr: pool,UserAddr: user}
//		}
//
//		dbv:=&PoolClaimHistory{}
//		json.Unmarshal(iter.Value(),dbv)
//
//		found:=false
//
//		for _,history := range v[user].History{
//			if history.BlockNumber == dbv.BlockNumber && history.TxIndex == dbv.TxIndex{
//				found = true
//				break
//			}
//		}
//
//		if found{
//			continue
//		}
//
//		poolhistory := v[user]
//
//		eventPos.LastMax2(dbv.BlockNumber,dbv.TxIndex)
//
//		poolhistory.History = append(poolhistory.History,dbv)
//	}
//
//	return nil
//}

func recover1() error  {
	allcm:=GetLogConf().BatchGet([]byte(poolClaimKeyHead),nil)

	for i:=0;i<len(allcm);i++{
		cm:=allcm[i]
		pool,user,err:=poolClaimKey2Address(cm.key)
		if err!=nil{
			fmt.Println("key",string(cm.key),err)
		}
		v, ok:=poolClaimUser[pool]
		if !ok{
			poolClaimUser[pool] = make(map[common.Address]*PoolClaimData)
			v = poolClaimUser[pool]
		}

		_, ok = v[user]
		if !ok{
			v[user] = &PoolClaimData{PoolAddr: pool,UserAddr: user}
		}

		dbv:=&PoolClaimHistory{}
		json.Unmarshal(cm.vaule,dbv)

		found:=false

		for _,history := range v[user].History{
			if history.BlockNumber == dbv.BlockNumber && history.TxIndex == dbv.TxIndex{
				found = true
				break
			}
		}

		if found{
			continue
		}

		poolhistory := v[user]

		eventPos.LastMax2(dbv.BlockNumber,dbv.TxIndex)

		poolhistory.History = append(poolhistory.History,dbv)

	}

}


func init()  {
	logPoolClaimSrvItem = &LogServiceItem{}
	logPoolClaimSrvItem.name = "PoolClaim"
	logPoolClaimSrvItem.stop = make(chan struct{})
	logPoolClaimSrvItem.watch = watchPoolClaim
	logPoolClaimSrvItem.batch = BatchPoolClaim
	logPoolClaimSrvItem.recover = recover1
	logPoolClaimSrvItem.CurBlockNum = curBlockN

	GetLogService().RegLogSrv(logPoolClaimSrvItem)
}

