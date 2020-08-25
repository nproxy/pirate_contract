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

func Charge(userAddr, poolAddr common.Address, no int64, priKey *ecdsa.PrivateKey) *types.Transaction {
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

	t := big.NewInt(no)
	tokenNo := t.Mul(t, big.NewInt(1e18))

	tx, err := market.Charge(transactor,userAddr, tokenNo,poolAddr, big.NewInt(int64(index)))

	if err != nil {
		fmt.Println("can't charge", err)
		return nil
	}

	return tx
}
