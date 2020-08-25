package reader

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperorchidlab/pirate_contract"
	"github.com/hyperorchidlab/pirate_contract/contract"
	"math/big"
	"sync"
)

//@return MBytesPerToken, PoolDeposit, MinerDeposit
func ContractSetting() (*big.Int, *big.Int, *big.Int){
	client, err := ethclient.Dial(pirate_contract.CurConfig.EthApiUrl)
	defer client.Close()
	if err!= nil {
		fmt.Println("can't connect to ethereum")
		return nil,nil,nil
	}
	market, err := contract.NewTrafficMarket(pirate_contract.CurConfig.Market, client)
	if err != nil {
		fmt.Println("can't recover market")
		return nil, nil, nil
	}

	var wg sync.WaitGroup
	wg.Add(3)
	var ratio, poolDps, minerDps *big.Int
	market.MBytesPerToken(nil)
	go func() {
		ratio, _ = market.MBytesPerToken(nil)
		wg.Done()
	}()
	go func() {
		poolDps, _ = market.PoolDeposit(nil)
		wg.Done()
	}()
	go func() {
		minerDps, _ = market.MinerDeposit(nil)
		wg.Done()
	}()
	wg.Wait()
	return ratio, poolDps, minerDps
}

//get account list for pools
func Pools() []common.Address {
	client, market := pirate_contract.RecoverMarket()
	if client == nil{
		return nil
	}
	defer client.Close()

	list, err := market.GetPoolList(nil)

	if err != nil {
		return nil
	}

	return list
}

func GetPoolIndex(poolAddr common.Address) int{
	list := Pools()
	for i:=0;i<len(list);i++{
		if list[i] == poolAddr {
			return i
		}
	}
	return -1
}

func UserData(poolAddr, userAddr common.Address) (*big.Int, *big.Int, error){
	client, market := pirate_contract.RecoverMarket()
	if client == nil{
		return nil, nil, errors.New("query user data failed")
	}
	defer client.Close()
	result, err := market.UserData(nil, poolAddr, userAddr)
	if err != nil {
		return nil, nil, err
	}

	return result.ChargeBalance, result.Epoch, nil
}

func PayerForMiner(poolAddr common.Address, subAddr [32]byte) (common.Address, error) {
	client, market := pirate_contract.RecoverMarket()
	if client == nil {
		return nil, errors.New("query miner address failed")
	}
	defer client.Close()
	result, err := market.PayerForMiner(nil,poolAddr, subAddr)
	if err != nil {
		return [20]byte{}, err
	}
	return result, nil
}
