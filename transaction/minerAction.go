package transaction

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hyperorchidlab/pirate_contract"
	"github.com/hyperorchidlab/pirate_contract/reader"
	"math/big"
)

func JoinPool(poolAddr common.Address, subAddr [32]byte, priKey *ecdsa.PrivateKey) *types.Transaction {
	client, market := pirate_contract.RecoverMarket()
	if client == nil {
		return nil
	}
	defer client.Close()

	transactor := bind.NewKeyedTransactor(priKey)

	index := reader.GetPoolIndex(poolAddr)

	tx, err := market.JoinPool(transactor, poolAddr, big.NewInt(int64(index)), subAddr)

	if err != nil {
		fmt.Println("can't join pool,", err)
		return nil
	}
	return tx
}

func ChangePool(from, to common.Address, subAddr [32]byte, priKey *ecdsa.PrivateKey) *types.Transaction {
	client, market := pirate_contract.RecoverMarket()
	if client == nil {
		return nil
	}
	defer client.Close()

	transactor := bind.NewKeyedTransactor(priKey)

	tx, err := market.ChangePool(transactor, from, to, subAddr)

	if err != nil {
		fmt.Println("can't change pool", err)
		return nil
	}

	return tx
}

func RetireFromPool(from common.Address, subAddr [32]byte, priKey *ecdsa.PrivateKey) *types.Transaction {
	client, market := pirate_contract.RecoverMarket()
	if client == nil {
		return nil
	}
	defer client.Close()

	transactor := bind.NewKeyedTransactor(priKey)

	tx, err := market.RetireFromPool(transactor, from, subAddr)

	if err != nil {
		fmt.Println("can't retire", err)
		return nil
	}

	return tx
}
