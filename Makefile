default: proto sc

.PHONY: contracts
contracts:
	make -C pkg/contracts compile
	make -C pkg/contracts bindings

.PHONY: proto
proto:
	make -C proto proto
