package hoptoken

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hyperorchidlab/pirate_contract/config"
	"github.com/hyperorchidlab/pirate_contract/contract"
	"github.com/hyperorchidlab/pirate_contract/util"
	"math/big"
)

func QueryApproved(address common.Address) *big.Int {
	t, err := config.SysEthConfig.NewTokenClient()
	if err != nil {
		fmt.Println("[QueryApproved]: tokenConn err:", err.Error())
		return nil
	}
	defer t.Close()

	a, e := t.Allowance(nil, address, config.SysEthConfig.Market)
	if e != nil {
		return nil
	}
	fmt.Println(a.String())
	return a
}

func TransferERCToken(target string, tokenNo float64, key *ecdsa.PrivateKey) (string, error) {

	t, err := config.SysEthConfig.NewTokenClient()
	if err != nil {
		fmt.Println("[TransferERCToken]: tokenConn err:", err.Error())
		return "", err
	}
	defer t.Close()

	opts := bind.NewKeyedTransactor(key)
	val := util.BalanceEth(tokenNo)

	fmt.Printf("\n----->%.2f", util.BalanceHuman(val))

	tx, err := t.Transfer(opts, common.HexToAddress(target), val)
	if err != nil {
		fmt.Println("[TransferERCToken]: Transfer err:", err.Error())
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func approve(no int64, tokenAddr, spender common.Address, priKey *ecdsa.PrivateKey) *types.Transaction {
	client, err := config.SysEthConfig.NewEthClient()
	defer client.Close()

	if err != nil {
		fmt.Println("can't connect to ethereum")
		return nil
	}

	token, err := contract.NewToken(tokenAddr, client)

	if err != nil {
		fmt.Println("can't recover token")
	}

	transactor := bind.NewKeyedTransactor(priKey)

	t := big.NewInt(no)
	tokenNo := t.Mul(t, big.NewInt(1e18))

	tx, err := token.Approve(transactor, spender, tokenNo)

	if err != nil {
		fmt.Println("approve err:", err)
	}
	return tx
}

func ApproveToMarket(no int64, priKey *ecdsa.PrivateKey) *types.Transaction {
	return approve(no, config.SysEthConfig.Token, config.SysEthConfig.Market, priKey)
}
