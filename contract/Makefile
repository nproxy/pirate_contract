SHELL=PATH='$(PATH)' /bin/sh

.PHONY: all
all:
	abigen --abi TrafficMarket.abi --pkg contract --type TrafficMarket --out TrafficMarket.go
	abigen --abi PirateDeposit.abi --pkg contract --type PirateDeposit --out PirateDeposit.go
	abigen --abi HyperOrchidProtocol.abi --pkg contract --type Token --out Token.go
	#abigen --abi ExtendToken.abi --pkg contract --type ExtendToken --out ExtendToken.go
	abigen --abi ApplyToken.abi --pkg contract --type ApplyToken --out ApplyToken.go
	abigen --abi HopShop.abi --pkg contract --type HopShop --out HopShop.go
go:
	abigen --abi HopSellInvestor.abi --pkg contract --type HopSellInvestor --out HopSellInvestor.go

abi:
	solc --abi --overwrite TrafficMarket.sol -o ./
	solc --abi --overwrite PirateDeposit.sol -o ./
	solc --abi --overwrite hyperorchidprotocol.sol -o ./
	#solc --abi --overwrite ExtendToken.sol -o ./
	solc --abi --overwrite ApplyToken.sol -o ./
	solc --abi --overwrite HopShop.sol -o ./
abi1:
	solc --abi --overwrite ./hop_sell/inverstor.sol -o ./

clean:
#	rm -f *.abi
	rm -f *.go
