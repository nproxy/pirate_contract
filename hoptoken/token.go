package hoptoken

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/hyperorchidlab/pirate_contract/config"
	"github.com/hyperorchidlab/pirate_contract/util"
	"math/big"
)

func HopBalance(user common.Address) (*big.Int, error) {
	t, err := config.SysEthConfig.NewTokenClient()
	if err != nil {
		fmt.Println("[QueryApproved]: tokenConn err:", err.Error())
		return nil, err
	}
	defer t.Close()

	h, e := t.BalanceOf(nil, user)
	if e != nil {
		return nil, err
	}

	return h, nil
}

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

	var nid *big.Int
	nid,err = t.GetClient().ChainID(context.TODO())
	if err!=nil{
		fmt.Println(err)
		return "", err
	}

	var transactOpts *bind.TransactOpts
	transactOpts,err = bind.NewKeyedTransactorWithChainID(key,nid)
	if err!=nil{
		fmt.Println(err)
		return "", err
	}

	val := util.BalanceEth(tokenNo)

	fmt.Printf("\n----->%.2f", util.BalanceHuman(val))

	tx, err := t.Transfer(transactOpts, common.HexToAddress(target), val)
	if err != nil {
		fmt.Println("[TransferERCToken]: Transfer err:", err.Error())
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func Approve(no float64, tokenAddr, spender common.Address, priKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	client, err := config.SysEthConfig.NewTokenClient()

	if err != nil {
		fmt.Println("can't connect to ethereum")
		return nil, err
	}

	defer client.Close()


	var nid *big.Int
	nid,err = client.GetClient().ChainID(context.TODO())
	if err!=nil{
		fmt.Println(err)
		return nil, err
	}

	var transactor *bind.TransactOpts
	transactor,err = bind.NewKeyedTransactorWithChainID(priKey,nid)
	if err!=nil{
		fmt.Println(err)
		return nil, err
	}


	tokenNo := util.BalanceEth(no)

	tx, err := client.Approve(transactor, spender, tokenNo)

	if err != nil {
		return nil, err
	}
	return tx, nil
}

func ApproveToMarket(no float64, priKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	return Approve(no, config.SysEthConfig.Token, config.SysEthConfig.Market, priKey)
}
