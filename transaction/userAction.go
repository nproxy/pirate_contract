package transaction

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hyperorchidlab/pirate_contract/config"
	"github.com/hyperorchidlab/pirate_contract/storageService"
	"math/big"
)

func Charge(userAddr, poolAddr common.Address, no int64, priKey *ecdsa.PrivateKey) *types.Transaction {
	mc, err := config.SysEthConfig.NewClient()
	if err != nil {
		return nil
	}
	defer mc.Close()

	transactor := bind.NewKeyedTransactor(priKey)

	var index int
	index, err = storageService.GetPoolIndex(poolAddr)
	if err != nil {
		fmt.Println("no pool in contract")
		return nil
	}

	t := big.NewInt(no)
	tokenNo := t.Mul(t, big.NewInt(1e18))

	tx, err := mc.Charge(transactor, userAddr, tokenNo, poolAddr, big.NewInt(int64(index)))

	if err != nil {
		fmt.Println("can't charge", err)
		return nil
	}

	return tx
}
