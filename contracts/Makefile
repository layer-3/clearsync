OS = $(shell uname -o | tr A-Z a-z)
ARCH = $(shell uname -m)

ROOT_DIR = ..
BIN_DIR = ../tmp/bin
CACHE_DIR = ../tmp/cache
GO_PKG_DIR = ../pkg

JQ_BIN = ${BIN_DIR}/gojq
ABIGEN_BIN = ${BIN_DIR}/abigen
ABIGEN_VERSION ?= 1.12.2

BINDINGS = \
	abi/itoken/ierc20.go \
	abi/ierc1271/ierc1271.go \
	abi/iadjudicator/iyellow_adjudicator.go \
	abi/iuniswap_v3_factory/iuniswap_v3_factory.go \
	abi/iuniswap_v3_pool/iuniswap_v3_pool.go \
	abi/iuniswap_v3_router/iuniswap_v3_router.go \
	abi/isyncswap_factory/isyncswap_factory.go \
	abi/isyncswap_pool/isyncswap_pool.go \
	artifacts/test_erc20/test_erc20.go \
	artifacts/adjudicator/yellow_adjudicator.go \
	artifacts/yellow_token/yellow_token.go \
	artifacts/consensus_app/consensus_app.go \
	artifacts/margin_app/margin_app_v1.go \
	artifacts/escrow_app/escrow_app.go \
	artifacts/syncswap_router/syncswap_router.go

default: deps solc bindings

${ABIGEN_BIN}:
	@mkdir -p ${BIN_DIR}
	@GOBIN=$$(pwd)/${BIN_DIR} go install github.com/ethereum/go-ethereum/cmd/abigen@v${ABIGEN_VERSION}

${JQ_BIN}:
	@mkdir -p ${BIN_DIR}
	@GOBIN=$$(pwd)/${BIN_DIR} go install github.com/itchyny/gojq/cmd/gojq@latest

$(ROOT_DIR)/node_modules: $(ROOT_DIR)/contracts
	@cd $(ROOT_DIR) && npm ci

${ROOT_DIR}/artifacts: $(ROOT_DIR)/node_modules
	@cd $(ROOT_DIR) && npx hardhat compile

compile: ${ROOT_DIR}/artifacts

# <go_file>_GO = <struct_name> <path_to_artifact>
ierc20_GO = IERC20 @openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol/IERC20Metadata.json
ierc1271_GO = IERC1271 @openzeppelin/contracts/interfaces/IERC1271.sol/IERC1271.json
iyellow_adjudicator_GO = IYellowAdjudicator contracts/interfaces/IYellowAdjudicator.sol/IYellowAdjudicator.json
iuniswap_v3_factory_GO = IUniswapV3Factory @uniswap/v3-core/contracts/interfaces/IUniswapV3Factory.sol/IUniswapV3Factory.json
iuniswap_v3_pool_GO = IUniswapV3Pool @uniswap/v3-core/contracts/interfaces/IUniswapV3Pool.sol/IUniswapV3Pool.json
iuniswap_v3_router_GO = IUniswapV3Router @uniswap/v3-periphery/contracts/interfaces/ISwapRouter.sol/ISwapRouter.json
isyncswap_factory_GO = ISyncSwapFactory @syncswap/core-contracts/contracts/interfaces/factory/IBasePoolFactory.sol/IBasePoolFactory.json
isyncswap_pool_GO = ISyncSwapPool @syncswap/core-contracts/contracts/interfaces/pool/IBasePool.sol/IBasePool.json

test_erc20_GO = TestERC20 contracts/test/TestERC20.sol/TestERC20.json
yellow_adjudicator_GO = YellowAdjudicator contracts/clearing/YellowAdjudicator.sol/YellowAdjudicator.json
yellow_token_GO = YellowToken contracts/YellowToken.sol/YellowToken.json
consensus_app_GO = ConsensusApp contracts/nitro/ConsensusApp.sol/ConsensusApp.json
margin_app_v1_GO = MarginAppV1 contracts/clearing/MarginAppV1.sol/MarginAppV1.json
escrow_app_GO = EscrowApp contracts/clearing/EscrowApp.sol/EscrowApp.json
syncswap_router_GO = SyncSwapRouter @syncswap/core-contracts/contracts/SyncSwapRouter.sol/SyncSwapRouter.json
define gen_binding
	$(eval STRUCT_NAME = $(word 1,$($(3)_GO)))
	$(eval ARTIFACT_PATH = $(word 2,$($(3)_GO)))
	@mkdir -p $(1)
	@mkdir -p $(CACHE_DIR)
	@cat ${ROOT_DIR}/artifacts/${ARTIFACT_PATH} | ${JQ_BIN} -r '.abi' > $(CACHE_DIR)/${STRUCT_NAME}.abi
	@cat ${ROOT_DIR}/artifacts/${ARTIFACT_PATH} | ${JQ_BIN} -r '.bytecode' > $(CACHE_DIR)/${STRUCT_NAME}.bin
	${ABIGEN_BIN} \
		--abi $(CACHE_DIR)/${STRUCT_NAME}.abi \
		--bin $(CACHE_DIR)/${STRUCT_NAME}.bin \
		--pkg $(2) \
		--type ${STRUCT_NAME} \
		--out $(1)/$(3).go
endef
${BINDINGS:%=${GO_PKG_DIR}/%}: ${ABIGEN_BIN} ${JQ_BIN}
	$(eval VALUES = $(shell echo $@ | sed -r 's#((.*/)?(.*))/(.*).go#\1 \3 \4#'))
	$(eval PKG_PATH = $(word 1,${VALUES}))
	$(eval PKG = $(word 2,${VALUES}))
	$(eval FILE_NAME = $(word 3,${VALUES}))
	$(call gen_binding,${PKG_PATH},${PKG},${FILE_NAME})

bindings: ${BINDINGS:%=${GO_PKG_DIR}/%}

clean-bin:
	rm -rf ${ABIGEN_BIN}

clean-cache:
	rm -rf ${CACHE_DIR}

clean-bindings:
	for binding in ${BINDINGS}; do \
		rm -rf $$binding; \
	done

clean: clean-bin clean-bindings
