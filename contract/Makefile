SHELL=PATH='$(PATH)' /bin/sh

.PHONY: all

all:
	rm -rf *.abi
	npx poa-solidity-flattener $(shell pwd)/TrafficMarket.sol
	mv out/TrafficMarket_flat.sol ./
	rmdir out
	solcjs TrafficMarket_flat.sol --abi --overwrite -o ./

	abigen --abi TrafficMarket_flat_sol_TrafficMarket.abi --pkg contract --type TrafficMarket --out TrafficMarket.go
