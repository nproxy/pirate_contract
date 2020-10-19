package storageService

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hyperorchidlab/pirate_contract/cabinet"
	"github.com/hyperorchidlab/pirate_contract/config"
	"sync"
)


var (
	poolsLock sync.Mutex
	pools []common.Address
)

func GetPirateEthSettings() (*cabinet.PirateEthSetting, error) {
	mc, err := config.SysEthConfig.NewClient()
	if err!=nil{
		return nil, err
	}
	defer mc.Close()

	pes := &cabinet.PirateEthSetting{}

	pes.MBytesPerToken, pes.PoolDeposit, pes.MinerDeposit, err = mc.BlockChainSettings(nil)
	if err != nil {
		return nil, err
	}

	return pes, nil

}

func GetPoolsList() ([]common.Address, error) {
	mc, err := config.SysEthConfig.NewClient()
	if err!=nil{
		return nil, err
	}
	defer mc.Close()

	return mc.GetPoolList(nil)
}

func findIndex(pool common.Address) (int,error) {
	for i:=0;i<len(pools);i++{
		if pool == pools[i]{
			return i,nil
		}
	}

	return 0,errors.New("not found")
}

func GetPoolIndex(pool common.Address) (int,error)  {
	poolsLock.Lock()
	defer poolsLock.Unlock()

	var (
		idx int
		err error
		poolstmp []common.Address
	)

	idx,err = findIndex(pool)
	if err==nil{
		return idx,nil
	}

	poolstmp,err = GetPoolsList()
	if err!=nil{
		return 0, err
	}
	pools = poolstmp

	return findIndex(pool)
}


func GetUserData(pool common.Address, user common.Address) (*cabinet.PirateEthUserData, error) {
	mc, err := config.SysEthConfig.NewClient()
	if err!=nil{
		return nil, err
	}
	defer mc.Close()

	ud, e := mc.UserData(nil, pool, user)
	if e != nil {
		return nil, e
	}

	peud := &cabinet.PirateEthUserData{
		ChargeBalance: ud.ChargeBalance,
		Epoch:         ud.Epoch,
	}

	return peud, nil
}

func GetPayForMiner(pool common.Address,miner [32]byte) (payAddr common.Address,err error)  {
	mc,err:=config.SysEthConfig.NewClient()
	if err!=nil{
		return payAddr,err
	}
	defer mc.Close()

	payAddr,e:=mc.PayerForMiner(nil,pool,miner)
	if e!=nil{
		return payAddr,e
	}

	return payAddr,nil

}
