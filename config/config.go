package config

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperorchidlab/pirate_contract/contract"
)

type EthConfig struct {
	NetworkID int            `json:"id"`
	EthApiUrl string         `json:"apiUrl"`
	Token     common.Address `json:"token"`
	Market    common.Address `json:"market"`
}

type PlatEthConfig struct {
	EthConfig
	DirectNetFunc func()
}

//change this to use mainnet
var SysEthConfig = &PlatEthConfig{
	EthConfig: EthConfig{
		NetworkID: 3,
		EthApiUrl: "https://ropsten.infura.io/v3/8533ef82c9744d38801f512fdd004133",
		Token:     common.HexToAddress("0xad44c8493de3fe2b070f33927a315b50da9a0e25"),
		Market:    common.HexToAddress("0xaF0FFac8b5674032c5f1DC2ce6b571a47AFCd896"),
	},
}

func (ec *EthConfig)String() string  {
	return fmt.Sprintf("\n++++++++++++++++++++++++++++++++++++++++++++++++++++\n"+
		"+NetworkID:\t%d+\n"+
		"+EthApiUrl:\t%s+\n"+
		"+MicroPaySys:\t%s+\n"+
		"+Token:\t%s+\n"+
		"++++++++++++++++++++++++++++++++++++++++++++++++++++\n",
		ec.NetworkID,
		ec.EthApiUrl,
		ec.Market.String(),
		ec.Token.String())
}

type MarketClient struct {
	client *ethclient.Client
	*contract.TrafficMarket
}

type TokenClient struct {
	client *ethclient.Client
	*contract.Token
}

func (mc *MarketClient) GetClient() *ethclient.Client {
	return mc.client
}

func (mc *MarketClient) Close() {
	if mc.client == nil {
		return
	}
	mc.client.Close()
	mc.client = nil
}

func (ec *PlatEthConfig) NewClient() (*MarketClient, error) {

	if ec.DirectNetFunc != nil{
		ec.DirectNetFunc()
	}

	client, err := ethclient.Dial(ec.EthApiUrl)
	if err != nil {
		return nil, err
	}

	var market *contract.TrafficMarket
	market, err = contract.NewTrafficMarket(ec.Market, client)
	if err != nil {
		client.Close()
		return nil, err
	}

	return &MarketClient{client: client, TrafficMarket: market}, nil
}


func (ec *PlatEthConfig)NewEthClient() (*ethclient.Client, error)  {
	if ec.DirectNetFunc != nil{
		ec.DirectNetFunc()
	}
	return ethclient.Dial(ec.EthApiUrl)
}


func (ec *PlatEthConfig)NewTokenClient() (*TokenClient,error)  {
	if ec.DirectNetFunc != nil{
		ec.DirectNetFunc()
	}
	client, err := ethclient.Dial(ec.EthApiUrl)
	if err != nil {
		return nil, err
	}
	var token *contract.Token
	token, err = contract.NewToken(ec.Token,client)
	if err!=nil{
		client.Close()
		return nil, err
	}
	return &TokenClient{client: client,Token:token},nil
}

func (tc *TokenClient)Close()  {
	if tc.client == nil {
		return
	}
	tc.client.Close()
	tc.client = nil

}



func test(user common.Address) {
	mc, _ := SysEthConfig.NewClient()
	if mc == nil {
		return
	}
	defer mc.Close()

	hop, eth, ap, _ := mc.TokenBalance(nil, user)
	fmt.Println(hop.String(), eth.String(), ap.String())
}
