package trafficmarket

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hyperorchidlab/pirate_contract/config"
	"github.com/hyperorchidlab/pirate_contract/storageService"
	"github.com/hyperorchidlab/pirate_contract/util"
	"math/big"
)

func Charge(userAddr, poolAddr common.Address, no float64, priKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	mc, err := config.SysEthConfig.NewClient()
	if err != nil {
		return nil, err
	}
	defer mc.Close()

	transactor := bind.NewKeyedTransactor(priKey)

	var index int
	index, err = storageService.GetPoolIndex(poolAddr)
	if err != nil {
		fmt.Println("no pool in contract")
		return nil, err
	}

	tokenNo := util.BalanceEth(no)

	tx, err := mc.Charge(transactor, userAddr, tokenNo, poolAddr, big.NewInt(int64(index)))

	if err != nil {
		fmt.Println("can't charge", err)
		return nil, err
	}

	return tx, nil
}
