package pirate_contract

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	FAILED = iota
	SUCCESS
	PENDING
	NOTFOUND
	SYSERROR
)

func ClientConn() (*ethclient.Client, error) {
	conn, err := ethclient.Dial(CurConfig.EthApiUrl)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func CheckTransactionStatus(hash common.Hash) (int, error) {
	client, err := ClientConn()
	if err != nil {
		fmt.Println("[CheckTransactionStatus] err:", err.Error())
		return SYSERROR, err
	}
	receipt, err := client.TransactionReceipt(context.Background(), hash)
	// not confirmed, not to check pending pool
	if receipt == nil {
		//check if transaction is pending
		isPending, err := TransactionByHash(hash)
		if err == nil {
			if isPending {
				return PENDING, nil
			} else {
				//may be tx just get confirmed, in that case, tx status should be finalized,
				//but we should be careful, let caller code to check again
				return PENDING, nil
			}
		} else {
			//may be not broadcast or dropped by miner
			return NOTFOUND, err
		}
	} else {
		if receipt.Status == 1 {
			return SUCCESS, nil
		} else {
			return FAILED, nil
		}
	}
}

func TransactionByHash(hash common.Hash) (bool, error) {
	client, err := ClientConn()
	if err != nil {
		return false, err
	}
	_, isPending, err := client.TransactionByHash(context.Background(), hash)
	return isPending, err
}

func NextOnce(addr common.Address) (uint64, error) {
	if client, err := ClientConn(); err != nil {
		return 0, err
	} else {
		return client.NonceAt(context.Background(), addr, nil)
	}
}