SHELL=PATH='$(PATH)' /bin/sh

.PHONY: all

all:
	rm -rf *.abi
	solc TrafficMarket.sol --abi --overwrite -o ./
	solc hyperorchidprotocol.sol --abi --overwrite -o ./
	solc ExtendToken.sol --abi --overwrite -o ./
	solc ApplyToken.sol --abi --overwrite -o ./
	solc HopShop.sol --abi --overwrite -o ./

	abigen --abi TrafficMarket.abi --pkg contract --type TrafficMarket --out TrafficMarket.go
	abigen --abi hyperorchidprotocol.abi --pkg contract --type Token --out Token.go
	abigen --abi ExtendToken.abi --pkg contract --type ExtendToken --out ExtendToken.go
	abigen --abi ApplyToken.abi --pkg contract --type ApplyToken --out ApplyToken.go
	abigen --abi HopShop.abi --pkg contract --type HopShop --out HopShop.go
