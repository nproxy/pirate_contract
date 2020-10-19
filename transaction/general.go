package transaction

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperorchidlab/pirate_contract"
	"github.com/hyperorchidlab/pirate_contract/contract"
	"math/big"
)

func approve(no int64, tokenAddr, spender common.Address, priKey *ecdsa.PrivateKey) *types.Transaction {
	client, err := ethclient.Dial(pirate_contract.CurConfig.EthApiUrl)
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
	return approve(no, pirate_contract.CurConfig.Token, pirate_contract.CurConfig.Market, priKey)
}
