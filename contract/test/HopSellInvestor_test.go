package test

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperorchidlab/pirate_contract/contract"
	"github.com/hyperorchidlab/pirate_contract/util"
	"math/big"
	"testing"
)

const(
	dialerUrl = "https://mainnet.infura.io/v3/d64d364124684359ace20feae1f9ac20"
	investorContract = "0xDaC511E4eF3Ead99f90716e246282662A3e36732"
)


var _cipherTxt = `{
		"cipher": "aes-128-ctr",
		"ciphertext": "3f84fa9dcf9637ee531cd972fa7fcda976e1361f9cee6ee9f5222e2b3d59807d",
		"cipherparams": {
			"iv": "308f21f38eb5f3379168664f1b6a278e"
		},
		"kdf": "scrypt",
		"kdfparams": {
			"dklen": 32,
			"n": 262144,
			"p": 1,
			"r": 8,
			"salt": "e770746f6cf346bd1150b83d5c0a915a9bd6a5e4a6dff4506c6befa97fc5c3d6"
		},
		"mac": "c0881f242b339a17ab71eca7bb4556f1b7e07e8d3fefe16c2af5494841b241ab"
	}`


func GetPrivKey() *ecdsa.PrivateKey {
	j := &keystore.CryptoJSON{}
	json.Unmarshal([]byte(_cipherTxt), j)
	key, err := keystore.DecryptDataV3(*j, "123")

	if err != nil {
		fmt.Println("err is ", err)
		return nil
	}
	priKey := crypto.ToECDSAUnsafe(key)

	return priKey
}

func getRandomId() (rid [32]byte) {
	rand.Read(rid[:])
	return
}

func toPubKeyString(priv *ecdsa.PrivateKey) string {
	pubkey := priv.PublicKey
	return crypto.PubkeyToAddress(pubkey).String()
}

func TestPrintEthAddr(t *testing.T)  {
	addr:=toPubKeyString(GetPrivKey())
	fmt.Println(addr)


	cli,err:=ethclient.Dial(dialerUrl)
	if err!=nil{
		panic(err)
	}
	defer cli.Close()

	b,err:=cli.BalanceAt(context.TODO(),common.HexToAddress(addr),nil)
	if err!=nil{
		panic(err)
	}

	f := util.BalanceHuman(b)


	fmt.Printf("%.6f",f)

}

func TestReleaseFor(t *testing.T)  {
	cli,err:=ethclient.Dial(dialerUrl)
	if err!=nil{
		panic(err)
	}
	defer cli.Close()

	ctr,err:=contract.NewHopSellInvestor(common.HexToAddress(investorContract),cli)
	if err!=nil{
		panic(err)
	}

	chainId,err:=cli.ChainID(context.TODO())
	if err!=nil{
		panic(err)
	}

	transopt, err := bind.NewKeyedTransactorWithChainID(GetPrivKey(),chainId)
	if err!=nil {
		panic(err)
	}

	amount := &big.Int{}
	amount.SetString("24095400000000000000000",10)

	tx,err:=ctr.ReleaseFor(transopt,common.HexToAddress("0xdBc379F98F3786Df825F8B4b636179647880f093"),amount)
	if err!=nil{
		panic(err)
	}

	fmt.Println(tx.Hash().String())
}

func TestBalance(t *testing.T)  {
	cli,err:=ethclient.Dial(dialerUrl)
	if err!=nil{
		panic(err)
	}
	defer cli.Close()

	ctr,err:=contract.NewHopSellInvestor(common.HexToAddress(investorContract),cli)
	if err!=nil{
		panic(err)
	}

	//chainId,err:=cli.ChainID(context.TODO())
	//if err!=nil{
	//	panic(err)
	//}

	//transopt, err := bind.NewKeyedTransactorWithChainID(GetPrivKey(),chainId)
	//if err!=nil {
	//	panic(err)
	//}

	n,err:=ctr.BalanceDetail(nil,common.HexToAddress("0xdBc379F98F3786Df825F8B4b636179647880f093"))
	if err!=nil{
		panic(err)
	}
	fmt.Println("balance:",n.Balance.String())
	fmt.Println("balance:",n.Mutiplier.String())
	fmt.Println("balance:",n.Claimable.String())

}

