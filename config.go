package pirate_contract

import (
	"github.com/BPassword/bpassword-lib/eth"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Config struct {
	NetworkID int            `json:"id"`
	EthApiUrl string         `json:"apiUrl"`
	Token     common.Address `json:"token"`
	Market    common.Address `json:"market"`
}

//change this to use mainnet
var CurConfig = &Config{
	NetworkID: 3,
	EthApiUrl:	"https://ropsten.infura.io/v3/8533ef82c9744d38801f512fdd004133",
	Token: common.HexToAddress("0xad44c8493de3fe2b070f33927a315b50da9a0e25"),
	Market: common.HexToAddress("0x18b3e987a65c18E9518580f820fb08015375D025"),
}


var GasPrice *big.Int

func SuggestedGasPrice() float64 {
	price := eth.GetGasPrice()
	return eth.ConvertByDecimal(price)
}

func SetGasPrice(price float64) {
	valF := big.NewFloat(price)
	tn := new(big.Int)
	valF.Int(tn)
	eth.GasPrice = tn
}