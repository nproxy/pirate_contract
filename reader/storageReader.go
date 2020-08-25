package reader

import (
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
	client, err := ethclient.Dial(pirate_contract.CurConfig.EthApiUrl)
	defer client.Close()
	if err!= nil{
		return nil
	}
	market, err := contract.NewTrafficMarket(pirate_contract.CurConfig.Market, client)
	if err != nil {
		return nil
	}

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