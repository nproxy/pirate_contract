package transaction

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperorchidlab/pirate_contract/config"
	"time"
)

const (
	FAILED = iota
	SUCCESS
	PENDING
	NOTFOUND
	SYSERROR
)

//todo use transaction by hash
func CheckTransactionStatus(hash common.Hash) (int, error) {
	client, err := ethclient.Dial(config.SysEthConfig.EthApiUrl)
	if err != nil {
		fmt.Println("[CheckTransactionStatus] err:", err.Error())
		return SYSERROR, err
	}
	receipt, err := client.TransactionReceipt(context.Background(), hash)
	// not confirmed, not to check pending pool
	if receipt == nil {
		//check if transaction is pending
		isPending, err := IsPending(hash)
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

func IsPending(hash common.Hash) (bool, error) {
	client, err := ethclient.Dial(config.SysEthConfig.EthApiUrl)
	if err != nil {
		return false, err
	}
	_, isPending, err := client.TransactionByHash(context.Background(), hash)
	return isPending, err
}

func NextOnce(addr common.Address) (uint64, error) {
	if client, err := ethclient.Dial(config.SysEthConfig.EthApiUrl); err != nil {
		return 0, err
	} else {
		return client.NonceAt(context.Background(), addr, nil)
	}
}

func WaitMined(hash common.Hash) error {
	client, err := ethclient.Dial(config.SysEthConfig.EthApiUrl)
	if err != nil {
		fmt.Println("[WaitMined]: ClientConn err:", err.Error())
		return err
	}
	return modifiedWaitMined(context.Background(), client, hash)
}

func modifiedWaitMined(ctx context.Context, b *ethclient.Client, hash common.Hash) error {
	queryTicker := time.NewTicker(time.Second)
	defer queryTicker.Stop()

	for {
		receipt, _ := b.TransactionReceipt(ctx, hash)
		//transaction confirmed
		if receipt != nil {
			return nil
		}
		// Wait for the next round.
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-queryTicker.C:
		}
	}
}
