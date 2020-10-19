package notifier

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hyperorchidlab/pirate_contract"
	"github.com/hyperorchidlab/pirate_contract/contract"
)

//this is an example of notification
func notify(claimChan chan *contract.TrafficMarketPoolClaim) error {
	client, market := pirate_contract.RecoverMarket()
	if client == nil {

	}

	sub, err := market.WatchPoolClaim(nil, claimChan, []common.Address{}, []common.Address{})

	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case cc := <-claimChan:
				fmt.Println(cc.User.String())
			case err := <-sub.Err():
				fmt.Println("notify error", err)
				sub.Unsubscribe()
				client.Close()
				break
			}
		}
	}()

	return nil
}
