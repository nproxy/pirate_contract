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

func RegPool(priKey *ecdsa.PrivateKey) *types.Transaction {

	client, market := pirate_contract.RecoverMarket()
	if client == nil{
		return nil
	}
	defer client.Close()

	transactor := bind.NewKeyedTransactor(priKey)

	tx, err := market.RegPool(transactor)

	if err != nil {
		fmt.Println("can't reg pool:", err)
		return nil
	}

	return tx
}


func DestroyPool(poolAddr common.Address, priKey *ecdsa.PrivateKey) *types.Transaction {
	client, market := pirate_contract.RecoverMarket()
	if client == nil{
		return nil
	}
	defer client.Close()

	transactor := bind.NewKeyedTransactor(priKey)

	index := reader.GetPoolIndex(poolAddr)

	if index == -1 {
		fmt.Println("can't find given pool address")
	}


	tx, err := market.DestroyPool(transactor, big.NewInt(int64(index)))

	if err != nil {
		fmt.Println("can't destroy pool:", err)
	}

	return tx
}

func Claim(userAddr, poolAddr common.Address, credit, amount, micNonce, cn *big.Int, sig []byte, priKey *ecdsa.PrivateKey) *types.Transaction {
	client, market := pirate_contract.RecoverMarket()
	if client == nil{
		return nil
	}
	defer client.Close()

	transactor := bind.NewKeyedTransactor(priKey)

	tx, err := market.Claim(transactor, userAddr, poolAddr, credit, amount, micNonce, cn, sig)

	if err != nil {
		fmt.Println("can't claim", err)
	}

	return tx
}