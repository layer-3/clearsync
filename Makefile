default: proto sc

sc:
	make -C pkg/sc compile
	make -C pkg/sc bindings

.PHONY: proto
proto:
	make -C proto proto
