default: proto contracts

.PHONY: contracts
contracts:
	make -C contracts compile
	make -C contracts bindings

.PHONY: proto
proto:
	make -C proto proto
