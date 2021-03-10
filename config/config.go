package config

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hyperorchidlab/pirate_contract/contract"
	"context"
	"math/big"
	"path"
	"strings"
)

type EthConfig struct {
	NetworkID int            `json:"id"`
	EthApiUrl string         `json:"apiUrl"`
	Token     common.Address `json:"token"`
	Market    common.Address `json:"market"`
}

type PlatEthConfig struct {
	EthConfig
	DirectNetFunc func() `json:"-"`
}

//change this to use mainnet
var SysEthConfig = &PlatEthConfig{
	EthConfig: EthConfig{
		NetworkID: 3,
		EthApiUrl: "https://ropsten.infura.io/v3/df97d0caa3514b3d99e94bc7764cffa0",
		Token:     common.HexToAddress("0xfdeaf9536c2374c3b7a1a7a427a83aea1811462c"),
		Market:    common.HexToAddress("0x09381f7a2aff9b1fcbc0689a4e6607653651920a"),
	},
}

func (ec *EthConfig) String() string {
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

func (tc *TokenClient) GetClient() *ethclient.Client {
	return tc.client
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

	if ec.DirectNetFunc != nil {
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

func (ec *PlatEthConfig) NewWSClient() (*MarketClient, error) {
	url := ec.EthApiUrl
	arr := strings.Split(url, "://")
	host := arr[1]
	arr1 := strings.Split(host, "/")
	var arr2 []string
	for i := 0; i < len(arr1); i++ {
		arr2 = append(arr2, arr1[i])
		if i == 0 {
			arr2 = append(arr2, "ws")
		}
	}
	newhost := path.Join(arr2...)
	dialurl := "wss://" + newhost

	client, err := ethclient.Dial(dialurl)
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

func (ec *PlatEthConfig) NewMarketClient(client *ethclient.Client) (*MarketClient, error) {
	//if ec.DirectNetFunc!=nil{
	//	ec.DirectNetFunc()
	//}
	//
	//var market *contract.TrafficMarket
	market, err := contract.NewTrafficMarket(ec.Market, client)
	if err != nil {
		return nil, err
	}

	return &MarketClient{client: client, TrafficMarket: market}, nil
}

func (ec *PlatEthConfig) NewEthClient() (*ethclient.Client, error) {
	if ec.DirectNetFunc != nil {
		ec.DirectNetFunc()
	}
	return ethclient.Dial(ec.EthApiUrl)
}

func (ec *PlatEthConfig) NewTokenClient() (*TokenClient, error) {
	if ec.DirectNetFunc != nil {
		ec.DirectNetFunc()
	}
	client, err := ethclient.Dial(ec.EthApiUrl)
	if err != nil {
		return nil, err
	}
	var token *contract.Token
	token, err = contract.NewToken(ec.Token, client)
	if err != nil {
		client.Close()
		return nil, err
	}
	return &TokenClient{client: client, Token: token}, nil
}

func (tc *TokenClient) Close() {
	if tc.client == nil {
		return
	}
	tc.client.Close()
	tc.client = nil

}

func (tc *TokenClient)GetChainId() (*big.Int,error)  {
	return tc.client.ChainID(context.TODO())
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
