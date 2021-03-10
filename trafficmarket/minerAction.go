package trafficmarket

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hyperorchidlab/pirate_contract/config"
	"github.com/hyperorchidlab/pirate_contract/storageService"
	"math/big"
)

func JoinPool(poolAddr common.Address, subAddr [32]byte, priKey *ecdsa.PrivateKey) *types.Transaction {
	mc, err := config.SysEthConfig.NewClient()
	if err != nil {
		return nil
	}
	defer mc.Close()

	var nid *big.Int
	nid,err = mc.GetClient().ChainID(context.TODO())
	if err!=nil{
		fmt.Println(err)
		return nil
	}

	var transactor *bind.TransactOpts
	transactor,err = bind.NewKeyedTransactorWithChainID(priKey,nid)
	if err!=nil{
		fmt.Println(err)
		return nil
	}

	var index int
	index, err = storageService.GetPoolIndex(poolAddr)
	if err != nil {
		fmt.Println("no pool in contract")
		return nil
	}

	tx, err := mc.JoinPool(transactor, poolAddr, big.NewInt(int64(index)), subAddr)

	if err != nil {
		fmt.Println("can't join pool,", err)
		return nil
	}
	return tx
}

func ChangePool(from, to common.Address, subAddr [32]byte, priKey *ecdsa.PrivateKey) *types.Transaction {
	mc, err := config.SysEthConfig.NewClient()
	if err != nil {
		return nil
	}
	defer mc.Close()

	var nid *big.Int
	nid,err = mc.GetClient().ChainID(context.TODO())
	if err!=nil{
		fmt.Println(err)
		return nil
	}

	var transactor *bind.TransactOpts
	transactor,err = bind.NewKeyedTransactorWithChainID(priKey,nid)
	if err!=nil{
		fmt.Println(err)
		return nil
	}

	tx, err := mc.ChangePool(transactor, from, to, subAddr)

	if err != nil {
		fmt.Println("can't change pool", err)
		return nil
	}

	return tx
}

func RetireFromPool(from common.Address, subAddr [32]byte, priKey *ecdsa.PrivateKey) *types.Transaction {
	mc, err := config.SysEthConfig.NewClient()
	if err != nil {
		return nil
	}
	defer mc.Close()

	var nid *big.Int
	nid,err = mc.GetClient().ChainID(context.TODO())
	if err!=nil{
		fmt.Println(err)
		return nil
	}

	var transactor *bind.TransactOpts
	transactor,err = bind.NewKeyedTransactorWithChainID(priKey,nid)
	if err!=nil{
		fmt.Println(err)
		return nil
	}


	tx, err := mc.RetireFromPool(transactor, from, subAddr)

	if err != nil {
		fmt.Println("can't retire", err)
		return nil
	}

	return tx
}
