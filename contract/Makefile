SHELL=PATH='$(PATH)' /bin/sh

.PHONY: all
all:
	abigen --abi TrafficMarket.abi --pkg contract --type TrafficMarket --out TrafficMarket.go
	abigen --abi PirateDeposit.abi --pkg contract --type PirateDeposit --out PirateDeposit.go
	abigen --abi HyperOrchidProtocol.abi --pkg contract --type Token --out Token.go
	abigen --abi ContentNetworkDeposit.abi --pkg contract --type ContentNetworkDeposit --out ContentNetworkDeposit.go
	abigen --abi ApplyToken.abi --pkg contract --type ApplyToken --out ApplyToken.go
	abigen --abi HopShop.abi --pkg contract --type HopShop --out HopShop.go

abi:
	solc --abi --overwrite TrafficMarket.sol -o ./
	solc --abi --overwrite PirateDeposit.sol -o ./
	solc --abi --overwrite hyperorchidprotocol.sol -o ./
	solc --abi --overwrite ContentNetworkDeposit.sol -o ./
	solc --abi --overwrite ApplyToken.sol -o ./
	solc --abi --overwrite HopShop.sol -o ./

clean:
#	rm -f *.abi
	rm -f *.go
