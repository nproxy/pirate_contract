package logService

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hyperorchidlab/go-miner-pool/account"
	"context"
	"github.com/hyperorchidlab/pirate_contract/contract"
	"strings"
	"sync"
)

type MinerInPoolDataHistory struct {
	MinerSubAddr [32]byte       `json:"miner_sub_addr"`
	PoolAddrFrom common.Address `json:"pool_address"`
	PoolAddrTo   common.Address `json:"pool_addr_to"`
	PayAddr      common.Address `json:"pay_addr"`
	OpAct        int            `json:"op_act"`
	BlockPos
}

type MinerInPoolData struct {
	History      []*MinerInPoolDataHistory
}


const(
	ActJoin int = iota
	ActChange
	ActQuit
)


var minerActHistory map[[32]byte]*MinerInPoolData


var MinerActNotify func(addr [32]byte, h *MinerInPoolDataHistory) error

type MinerInPool struct {
	poolMiner map[common.Address]map[[32]byte]common.Address
	lock sync.Mutex
}


var MinerInPoolInst *MinerInPool

func (mip *MinerInPool)GetPayAddr(miner [32]byte,pool common.Address) (*common.Address,error) {
	mip.lock.Lock()
	defer mip.lock.Unlock()

	v,ok:= mip.poolMiner[pool]
	if !ok{
		return nil,errors.New("no pool in mem")
	}
	var addr common.Address
	addr,ok = v[miner]

	if !ok{
		return nil,errors.New("no miner in mem")
	}

	return &addr,nil
}

func (mip *MinerInPool)GetMinerList(pool common.Address) ([][32]byte,error)  {
	mip.lock.Lock()
	defer mip.lock.Unlock()

	v,ok:= mip.poolMiner[pool]
	if !ok{
		return nil,errors.New("no pool in mem")
	}

	var ms [][32]byte
	for k,_:=range v{
		ms = append(ms,k)
	}

	return ms,nil
}


func (mip *MinerInPool)join(miner [32]byte, pool,payAddr common.Address) {
	mip.lock.Lock()
	defer mip.lock.Unlock()

	mip._join(miner,pool,payAddr)
}

func (mip *MinerInPool)_join(miner [32]byte,pool,payAddr common.Address)  {
	v,ok:=mip.poolMiner[pool]
	if  !ok{
		mip.poolMiner[pool] = make(map[[32]byte]common.Address)
		v = mip.poolMiner[pool]
	}

	v[miner] = payAddr
}


func (mip *MinerInPool)quit(miner [32]byte, pool common.Address) {
	mip.lock.Lock()
	defer mip.lock.Unlock()

	mip._quit(miner,pool)
}

func (mip *MinerInPool)_quit(miner [32]byte, pool common.Address)  {
	_,ok:=mip.poolMiner[pool]
	if  !ok{
		return
	}

	delete(mip.poolMiner[pool],miner)
}

func (mip *MinerInPool)Change(miner [32]byte,poolFrom,poolTo,payAddr common.Address)  {
	mip.lock.Lock()
	defer mip.lock.Unlock()

	mip._quit(miner,poolFrom)
	mip._join(miner,poolTo,payAddr)
}


var minerActEventPos EventPos

var minerActKeyHead = "miner_act_"

var minerActKey = minerActKeyHead + "%s_%d_%d"

func getMinerActKey(miner [32]byte,pos *BlockPos) string  {
	return fmt.Sprintf(minerActKey,account.ConvertToID2(miner[:]).String(),pos.BlockNumber,pos.TxIndex)
}

func minerAct2Address(key []byte) (miner [32]byte,err error)  {
	ks := string(key)

	ksarr := strings.Split(ks,"_")

	if len(ksarr) < 5{
		return miner,errors.New("key error")
	}

	miner = account.ID(ksarr[2]).ToArray()

	return miner,nil
}

func addNewMinerActHistory(miner [32]byte,poolFrom, poolTo,payAddr common.Address, act int, pos types.Log)  {
	mip,ok:=minerActHistory[miner]
	if !ok{
		minerActHistory[miner] = &MinerInPoolData{}
		mip = minerActHistory[miner]
	}

	for _,history:=range mip.History{
		if history.BlockNumber == pos.BlockNumber && history.TxIndex == pos.TxIndex{
			return
		}
	}

	h:=&MinerInPoolDataHistory{MinerSubAddr: miner,
		PoolAddrFrom: poolFrom,PoolAddrTo: poolTo,OpAct: act, PayAddr: payAddr,BlockPos:BlockPos{pos.BlockNumber,pos.TxIndex}}

	mip.History = append(mip.History,h)

	k:=getMinerActKey(miner,&h.BlockPos)

	dbv,_:=json.Marshal(*h)

	GetLogConf().Save([]byte(k),dbv)

	if act == ActJoin{
		MinerInPoolInst.join(miner,poolFrom,payAddr)
	}else if act == ActChange{
		MinerInPoolInst.Change(miner,poolFrom,poolTo,payAddr)
	}else if act == ActQuit{
		MinerInPoolInst.quit(miner,poolFrom)
	}

	if MinerActNotify != nil{
		MinerActNotify(miner,h)
	}

}

func batchMinerInPool() error  {
	if minerActEventPos.IsDone(){
		return nil
	}

	mc,err:=GetLogConf().NewMarketClient()
	if err!=nil{
		return err
	}
	defer mc.Close()

	var opt *bind.FilterOpts

	if !minerActEventPos.IsFirst(){
		opt = &bind.FilterOpts{
			Start: minerActEventPos.StartPos(),
			End: nil,
			Context: context.TODO(),
		}
	}

	iter, err:=mc.FilterMinerEvent(opt,nil,nil)
	if err!=nil{
		fmt.Print("batch miner action failed")
		return err
	}

	for iter.Next(){
		ev:=iter.Event
		addNewMinerActHistory(ev.SubAddr,ev.Addr1,ev.Addr2, ev.PayAddr,int(ev.EventType),ev.Raw)
		minerActEventPos.LastMax(ev.Raw)
	}

	minerActEventPos.LastBlkNum()

	return nil
}

var logMinerActSrvItem *LogServiceItem

func watchMinerInPool(batch *chan *LogServiceItem) error  {
	mc,err:=GetLogConf().NewMarketClient()
	if err!=nil{
		return err
	}
	defer mc.Close()

	c:=make(chan *contract.TrafficMarketMinerEvent,1024)

	sub, e:=mc.WatchMinerEvent(nil,c,nil,nil)
	if e!= nil{
		return e
	}

	for{
		select {
		case pc:=<-c:
			minerActEventPos.NextMax(pc.Raw)
			ch := *batch

			ch <- logMinerActSrvItem
		case err = <-sub.Err():
			return err
		}
	}
}

func curMinerActBlockN(n uint64)  {
	logMinerActSrvItem.evPos.SetCurrent(n)
}

func recoverMinerAct() error  {
	allma:=GetLogConf().BatchGet([]byte(minerActKeyHead),nil)

	for i:=0;i<len(allma); i++{
		ma:=allma[i]
		maddr, err:=minerAct2Address(ma.key)
		if err!=nil{
			fmt.Println("key",string(ma.key),err)
			continue
		}

		mip,ok:=minerActHistory[maddr]
		if !ok{
			minerActHistory[maddr] = &MinerInPoolData{}
			mip = minerActHistory[maddr]
		}

		h:=&MinerInPoolDataHistory{}

		json.Unmarshal(ma.vaule,h)

		found := false

		for _,history:=range mip.History{
			if history.BlockNumber == h.BlockNumber && history.TxIndex == h.TxIndex{
				found = true
				break
			}
		}

		if found {
			continue
		}

		if h.OpAct == ActJoin{
			MinerInPoolInst.join(h.MinerSubAddr,h.PoolAddrFrom, h.PayAddr)
		}else if h.OpAct == ActChange{
			MinerInPoolInst.Change(h.MinerSubAddr,h.PoolAddrFrom,h.PoolAddrTo,h.PayAddr)
		}else if h.OpAct == ActQuit{
			MinerInPoolInst.quit(h.MinerSubAddr,h.PoolAddrFrom)
		}

		minerActEventPos.LastMax2(h.BlockNumber,h.TxIndex)

		mip.History = append(mip.History,h)
	}

	return nil

}



func init()  {
	MinerInPoolInst = &MinerInPool{
		poolMiner: make(map[common.Address]map[[32]byte]common.Address),
	}

	minerActHistory = make(map[[32]byte]*MinerInPoolData)

	logMinerActSrvItem = &LogServiceItem{}
	logMinerActSrvItem.name = "MinerEvent"
	logMinerActSrvItem.stop = make(chan struct{})
	logMinerActSrvItem.watch = watchMinerInPool
	logMinerActSrvItem.batch = batchMinerInPool
	logMinerActSrvItem.recover = recoverMinerAct
	logMinerActSrvItem.CurBlockNum = curMinerActBlockN

	GetLogService().RegLogSrv(logMinerActSrvItem)

}
