package pirate_contract

import (
	"github.com/ethereum/go-ethereum/common"
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
	Market: common.HexToAddress("0x7Df3Fa0dc8abf2787b96857BF8bEa268cAbA1474"),
}