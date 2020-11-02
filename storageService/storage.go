package storageService

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hyperorchidlab/pirate_contract/cabinet"
	"github.com/hyperorchidlab/pirate_contract/config"
	"github.com/kprc/nbsnetwork/tools"
	"math/big"
	"sync"
)

var (
	poolsLock sync.Mutex
	pools     []common.Address

	pes *cabinet.PirateEthSetting
	lastTime int64
)

func GetPirateEthSettings() (*cabinet.PirateEthSetting, error) {
	mc, err := config.SysEthConfig.NewClient()
	if err != nil {
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

func getCacheSetting(now int64) error {
	var err error
	pes,err = GetPirateEthSettings()
	if err!=nil{
		return err
	}
	lastTime = now
	return nil
}

func GetCacheSetting() (*cabinet.PirateEthSetting,error)  {
	now:=tools.GetNowMsTime()
	if pes == nil{
		if err:=getCacheSetting(now);err!=nil{
			return nil, err
		}
	}else{
		if now - lastTime > 300000{
			go getCacheSetting(now)
		}
	}

	return pes,nil
}



func Traffic2Balance(traffic *big.Int) (balance *big.Int,err error)  {
	GetCacheSetting()
	if pes == nil || pes.MBytesPerToken.Cmp(&big.Int{}) == 0{
		return &big.Int{},errors.New("system error")
	}

	var z *big.Int
	x,y:=&big.Int{},&big.Int{}
	x.SetInt64(10)
	y.SetInt64(12)

	z = z.Exp(x,y,nil)
	z = z.Mul(z,traffic)
	z = z.Div(z,pes.MBytesPerToken)

	return z,nil
}

func Balance2Traffic(balance *big.Int) (traffic *big.Int,err error)  {
	GetCacheSetting()
	if pes == nil{
		return &big.Int{},errors.New("system error")
	}

	var z *big.Int
	x,y:=&big.Int{},&big.Int{}
	x.SetInt64(10)
	y.SetInt64(12)

	z = z.Exp(x,y,nil)
	t:=*balance
	tt:=&t
	tt = tt.Mul(tt,pes.MBytesPerToken)
	z = tt.Div(tt,z)

	return z,nil
}



func GetPoolsList() ([]common.Address, error) {
	mc, err := config.SysEthConfig.NewClient()
	if err != nil {
		return nil, err
	}
	defer mc.Close()

	return mc.GetPoolList(nil)
}

func findIndex(pool common.Address) (int, error) {
	for i := 0; i < len(pools); i++ {
		if pool == pools[i] {
			return i, nil
		}
	}

	return 0, errors.New("not found")
}

func GetPoolIndex(pool common.Address) (int, error) {
	poolsLock.Lock()
	defer poolsLock.Unlock()

	var (
		idx      int
		err      error
		poolstmp []common.Address
	)

	idx, err = findIndex(pool)
	if err == nil {
		return idx, nil
	}

	poolstmp, err = GetPoolsList()
	if err != nil {
		return 0, err
	}
	pools = poolstmp

	return findIndex(pool)
}

func GetUserData(pool common.Address, user common.Address) (*cabinet.PirateEthUserData, error) {
	mc, err := config.SysEthConfig.NewClient()
	if err != nil {
		return nil, err
	}
	defer mc.Close()

	ud, e := mc.UserData(nil, pool, user)
	if e != nil {
		return nil, e
	}

	peud := &cabinet.PirateEthUserData{
		ChargeBalance: ud.TotalChargeBalance,
		TotalTraffic:  ud.TotalTraffic,
	}

	return peud, nil
}

func GetPayForMiner(pool common.Address, miner [32]byte) (payAddr common.Address, err error) {
	mc, err1 := config.SysEthConfig.NewClient()
	if err1 != nil {
		return payAddr, err
	}
	defer mc.Close()

	payAddr, err = mc.PayerForMiner(nil, pool, miner)
	if err != nil {
		return payAddr, err
	}

	return payAddr, nil

}

func TokenBalance2(user common.Address)  (hop *big.Int, eth *big.Int, apr *big.Int)  {
	mc, err := config.SysEthConfig.NewClient()
	if err != nil {
		return
	}
	defer mc.Close()

	hop, eth, apr, _ = mc.TokenBalance(nil, user)

	return
}

func TokenBalance(userAddr string) (hop *big.Int, eth *big.Int, apr *big.Int) {
	user:=common.HexToAddress(userAddr)

	return TokenBalance2(user)

}
