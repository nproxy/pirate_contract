package reader

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/hyperorchidlab/pirate_contract"
)

func UsersOfPool(poolAddr common.Address) ([]common.Address, error) {
	var list []common.Address
	client, market := pirate_contract.RecoverMarket()
	if client == nil{
		return list, errors.New("query users of pool err")
	}
	defer client.Close()

	it, err := market.FilterCharge(nil,[]common.Address{}, []common.Address{poolAddr})

	if err != nil {
		return list, err
	}
	for it.Next(){
		list = append(list, it.Event.User)
	}

	return list, nil
}

func MinerPoolMap() (map[common.Address][][32]byte, error) {
	m := make(map[common.Address][][32]byte)
	client, market := pirate_contract.RecoverMarket()
	if client == nil{
		return m, errors.New("query users of pool err")
	}
	defer client.Close()

	it, err := market.FilterMinerEvent(nil,[][32]byte{}, []common.Address{})

	if err != nil {
		return nil, err
	}

	for it.Next(){
		switch it.Event.EventType {
		case 0:
			m[it.Event.Addr1] = append(m[it.Event.Addr1], it.Event.SubAddr)
		case 1:
			v := m[it.Event.Addr1]
			for i:=1;i<len(v);i++ {
				if v[i] == it.Event.SubAddr {
					m[it.Event.Addr1] = append(v[:i], v[i+1:]...)
					break
				}
			}
			m[it.Event.Addr2] = append(m[it.Event.Addr2], it.Event.SubAddr)
		case 2:
			v := m[it.Event.Addr1]
			for i:=1;i<len(v);i++ {
				if v[i] == it.Event.SubAddr {
					m[it.Event.Addr1] = append(v[:i], v[i+1:]...)
					break
				}
			}
		}
	}

	return m, nil
}