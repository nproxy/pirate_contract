package logService

import (
	"errors"
	"github.com/btcsuite/goleveldb/leveldb"
	"github.com/btcsuite/goleveldb/leveldb/opt"
	"github.com/btcsuite/goleveldb/leveldb/util"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hyperorchidlab/pirate_contract/config"
)

type LogConf struct {
	db  *leveldb.DB
	cfg *config.EthConfig
}

var logConf *LogConf

type BlockPos struct {
	BlockNumber uint64 `json:"block_number"`
	TxIndex     uint   `json:"tx_index"`
}

func (bp *BlockPos) Max(log types.Log) {
	bp.Max2(log.BlockNumber, log.TxIndex)
}

func (bp *BlockPos) Max2(blkNum uint64, txIdx uint) {
	if bp.BlockNumber < blkNum {
		bp.BlockNumber = blkNum
		bp.TxIndex = txIdx
	} else if bp.BlockNumber == blkNum {
		if bp.TxIndex < txIdx {
			bp.TxIndex = txIdx
		}
	}
}

func (bp *BlockPos) SetBlkNum(n uint64) {
	bp.BlockNumber = n
	bp.TxIndex = 0
}

type EventPos struct {
	lastAccessPos  *BlockPos
	nextPos        *BlockPos
	curBlockNumber uint64
}

func (ep *EventPos) LastMax(log types.Log) {
	ep.LastMax2(log.BlockNumber, log.TxIndex)
}

func (ep *EventPos) LastMax2(blkNum uint64, txIdx uint) {
	if ep.lastAccessPos == nil {
		ep.lastAccessPos = &BlockPos{}
	}
	bp := ep.lastAccessPos

	bp.Max2(blkNum, txIdx)
}

func (ep *EventPos) NextMax(log types.Log) {
	if ep.nextPos == nil {
		ep.nextPos = &BlockPos{}
	}
	bp := ep.nextPos
	bp.Max(log)
}

func (ep *EventPos) LastBlkNum() {
	if ep.lastAccessPos == nil {
		ep.lastAccessPos = &BlockPos{}
	}
	bp := ep.lastAccessPos

	if ep.curBlockNumber > bp.BlockNumber {
		bp.SetBlkNum(ep.curBlockNumber)
	}
}

func (ep *EventPos) SetCurrent(blockNumber uint64) {
	ep.curBlockNumber = blockNumber
}

func (ep *EventPos) IsDone() bool {

	if ep.lastAccessPos == nil || ep.nextPos == nil {
		return false
	}

	last := ep.lastAccessPos
	next := ep.nextPos

	if next.BlockNumber > last.BlockNumber {
		return false
	}

	if next.BlockNumber == last.BlockNumber && next.TxIndex > last.TxIndex {
		return false
	}

	if ep.curBlockNumber > last.BlockNumber {
		return false
	}

	return true
}

func (ep *EventPos) IsFirst() bool {
	if ep.lastAccessPos == nil {
		return true
	}
	return false
}

func (ep *EventPos) StartPos() uint64 {
	if ep.lastAccessPos == nil {
		return 0
	} else {
		return ep.lastAccessPos.BlockNumber
	}
}

func (lc *LogConf) NewMarketClient() (*config.MarketClient, error) {
	if lc.cfg == nil {
		panic("no config")
	}
	return lc.cfg.NewClient()
}

func GetLogConf() *LogConf {
	return logConf
}

func Configure(db *leveldb.DB, cfg *config.EthConfig) error {
	if db == nil || cfg == nil {
		return errors.New("need db and contract cfg")
	}

	if logConf == nil {
		logConf = &LogConf{}
	}
	logConf.cfg = cfg
	logConf.db = db

	return nil
}

func (lc *LogConf) Save(key, value []byte) error {
	if lc.db == nil {
		panic("save failed " + string(key))
	}

	o := &opt.WriteOptions{Sync: true}

	return lc.db.Put(key, value, o)
}

func (lc *LogConf) Delete(key []byte) error {
	if lc.db == nil {
		panic("delete key failed " + string(key))
	}
	o := &opt.WriteOptions{Sync: true}

	return lc.db.Delete(key, o)
}

func (lc *LogConf) Get(key []byte) (v []byte, err error) {
	if lc.db == nil {
		panic("get value failed " + string(key))
	}

	return lc.db.Get(key, nil)
}

type DBKV struct {
	key   []byte
	vaule []byte
}

func (lc *LogConf) BatchGet(patternStart, patternEnd []byte) []*DBKV {
	r := &util.Range{Start: patternStart, Limit: patternEnd}

	iter := lc.db.NewIterator(r, nil)
	defer iter.Release()

	var dbkv []*DBKV

	for iter.Next() {
		kv := &DBKV{key: iter.Key(), vaule: iter.Value()}
		dbkv = append(dbkv, kv)
	}

	return dbkv
}
