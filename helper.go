package pirate_contract

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperorchidlab/pirate_contract/contract"
	"math/big"
)


//this method is very dangerous, please avoid expose your secret key !!!
func PrivateKeyFromStr(keyStr string) ecdsa.PrivateKey{
	var pri ecdsa.PrivateKey
	pri.D, _ = new(big.Int).SetString(keyStr, 16)
	pri.PublicKey.Curve = crypto.S256()
	pri.PublicKey.X, pri.PublicKey.Y = pri.PublicKey.Curve.ScalarBaseMult(pri.D.Bytes())
	return pri
}


func RecoverMarket() (*ethclient.Client, *contract.TrafficMarket){
	client, err := ethclient.Dial(CurConfig.EthApiUrl)
	if err != nil {
		fmt.Println("can't connect to ethereum")
		return nil, nil
	}

	market, err := contract.NewTrafficMarket(CurConfig.Market, client)

	if err != nil {
		fmt.Println("can't recover market")
		return nil, nil
	}

	return client, market
}