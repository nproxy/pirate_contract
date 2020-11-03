package trafficmarket

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

func RegPool(priKey *ecdsa.PrivateKey) *types.Transaction {

	mc, err := config.SysEthConfig.NewClient()
	if err != nil {
		return nil
	}
	defer mc.Close()

	transactor := bind.NewKeyedTransactor(priKey)

	tx, err := mc.RegPool(transactor)

	if err != nil {
		fmt.Println("can't reg pool:", err)
		return nil
	}

	return tx
}

func DestroyPool(poolAddr common.Address, priKey *ecdsa.PrivateKey) *types.Transaction {
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

	tx, err := mc.DestroyPool(transactor, big.NewInt(int64(index)))

	if err != nil {
		fmt.Println("can't destroy pool:", err)
	}

	return tx
}

func PClaim(userAddr, poolAddr common.Address, credit, amount, cn *big.Int, sig []byte, priKey *ecdsa.PrivateKey) *types.Transaction {
	mc, err := config.SysEthConfig.NewClient()
	if err != nil {
		return nil
	}
	defer mc.Close()

	transactor := bind.NewKeyedTransactor(priKey)

	tx, err := mc.Pclaim(transactor, userAddr, poolAddr, credit, amount, cn, sig)

	if err != nil {
		fmt.Println("can't claim", err)
	}

	return tx
}
