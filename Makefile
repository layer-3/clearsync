default: proto compile-contracts bindings

compile-contracts:
	make -C pkg/contracts compile-contracts

bindings:
	make -C pkg/contracts deps
	make -C pkg/contracts bindings

.PHONY: proto
proto:
	make -C proto proto
