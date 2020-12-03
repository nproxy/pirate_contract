package hoptoken

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hyperorchidlab/pirate_contract/config"
	"github.com/hyperorchidlab/pirate_contract/util"
)

func TransferEth(target string, tokenNo float64, privateKey *ecdsa.PrivateKey) (string, error) {

	client, err := config.SysEthConfig.NewEthClient()
	if err != nil {
		fmt.Println("[TransferEth]: Dial err:", err.Error())
		return "", err
	}
	defer client.Close()

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("[TransferEth]: cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return "", err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println("[TransferEth]: PendingNonceAt err:", err.Error())
		return "", err
	}

	value := util.BalanceEth(tokenNo) // in wei (1 eth)
	gasLimit := uint64(21000)         // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("[TransferEth]: SuggestGasPrice err:", err.Error())
		return "", err
	}

	var data []byte
	tx := types.NewTransaction(nonce, common.HexToAddress(target), value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println("[TransferEth]: NetworkID err:", err.Error())
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Println("[TransferEth]: SignTx err:", err.Error())
		return "", err
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println("[TransferEth]: SendTransaction err:", err.Error())
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}

func TxStatus(tx common.Hash) bool {

	client, err := config.SysEthConfig.NewEthClient()
	if err != nil {
		fmt.Println("[BuyPacket]: connect err:", err.Error())
		return false
	}
	defer client.Close()

	receipt, err := client.TransactionReceipt(context.Background(), tx)
	if err != nil {
		fmt.Println("[BuyPacket]: query receipt err:", err.Error())
		return false
	}
	return receipt.Status == 1
}

const (
	pending = iota
	success
	fail
)

func TxStatus2(tx common.Hash) int {

	client, err := config.SysEthConfig.NewEthClient()
	if err != nil {
		fmt.Println("[BuyPacket]: connect err:", err.Error())
		return fail
	}
	defer client.Close()

	receipt, err := client.TransactionReceipt(context.Background(), tx)
	if err != nil || receipt == nil{
		fmt.Println("[BuyPacket]: query receipt err:", err.Error())
		return fail
	}
	if receipt.Status == 1{
		return success
	}else{
		return fail
	}
}
