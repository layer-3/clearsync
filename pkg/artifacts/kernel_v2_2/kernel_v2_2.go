// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package kernel_v2_2

import (
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// Call is an auto generated low-level Go binding around an user-defined struct.
type Call struct {
	To    common.Address
	Value *big.Int
	Data  []byte
}

// ExecutionDetail is an auto generated low-level Go binding around an user-defined struct.
type ExecutionDetail struct {
	ValidAfter *big.Int
	ValidUntil *big.Int
	Executor   common.Address
	Validator  common.Address
}

// UserOperation is an auto generated low-level Go binding around an user-defined struct.
type UserOperation struct {
	Sender               common.Address
	Nonce                *big.Int
	InitCode             []byte
	CallData             []byte
	CallGasLimit         *big.Int
	VerificationGasLimit *big.Int
	PreVerificationGas   *big.Int
	MaxFeePerGas         *big.Int
	MaxPriorityFeePerGas *big.Int
	PaymasterAndData     []byte
	Signature            []byte
}

// KernelMetaData contains all meta data concerning the Kernel contract.
var KernelMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"_entryPoint\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DeprecatedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DisabledMode\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAuthorizedCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEntryPoint\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldValidator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newValidator\",\"type\":\"address\"}],\"name\":\"DefaultValidatorChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"selector\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ExecutionChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_disableFlag\",\"type\":\"bytes4\"}],\"name\":\"disableMode\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"entryPoint\",\"outputs\":[{\"internalType\":\"contractIEntryPoint\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"enumOperation\",\"name\":\"_operation\",\"type\":\"uint8\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structCall[]\",\"name\":\"calls\",\"type\":\"tuple[]\"}],\"name\":\"executeBatch\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDefaultValidator\",\"outputs\":[{\"internalType\":\"contractIKernelValidator\",\"name\":\"validator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDisabledMode\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"disabled\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"}],\"name\":\"getExecution\",\"outputs\":[{\"components\":[{\"internalType\":\"ValidAfter\",\"name\":\"validAfter\",\"type\":\"uint48\"},{\"internalType\":\"ValidUntil\",\"name\":\"validUntil\",\"type\":\"uint48\"},{\"internalType\":\"address\",\"name\":\"executor\",\"type\":\"address\"},{\"internalType\":\"contractIKernelValidator\",\"name\":\"validator\",\"type\":\"address\"}],\"internalType\":\"structExecutionDetail\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastDisabledTime\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint192\",\"name\":\"key\",\"type\":\"uint192\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIKernelValidator\",\"name\":\"_defaultValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIKernelValidator\",\"name\":\"_defaultValidator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"setDefaultValidator\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_selector\",\"type\":\"bytes4\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"},{\"internalType\":\"contractIKernelValidator\",\"name\":\"_validator\",\"type\":\"address\"},{\"internalType\":\"ValidUntil\",\"name\":\"_validUntil\",\"type\":\"uint48\"},{\"internalType\":\"ValidAfter\",\"name\":\"_validAfter\",\"type\":\"uint48\"},{\"internalType\":\"bytes\",\"name\":\"_enableData\",\"type\":\"bytes\"}],\"name\":\"setExecution\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"validateSignature\",\"outputs\":[{\"internalType\":\"ValidationData\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"initCode\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"callGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"verificationGasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"preVerificationGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPriorityFeePerGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"paymasterAndData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"internalType\":\"structUserOperation\",\"name\":\"_userOp\",\"type\":\"tuple\"},{\"internalType\":\"bytes32\",\"name\":\"userOpHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"missingAccountFunds\",\"type\":\"uint256\"}],\"name\":\"validateUserOp\",\"outputs\":[{\"internalType\":\"ValidationData\",\"name\":\"validationData\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x61014034620001b757601f62002f5438819003918201601f19168301916001600160401b03831184841017620001bc57808492602094604052833981010312620001b757516001600160a01b0381168103620001b757306080524660a05260a062000069620001d2565b600681526005602082016512d95c9b995b60d21b815260206200008b620001d2565b838152019264181719171960d91b845251902091208160c0528060e052604051917f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f83526020830152604082015246606082015230608082015220906101009182526101209081527f439ffe7df606b78489639bc0b827913bd09e1246fa6802968a5b3694c53e0dd96a010000000000000000000080600160f01b031982541617905560405190612d619283620001f3843960805183612650015260a05183612673015260c051836126e5015260e0518361270b0152518261262f01525181818161066901528181610b0401528181610c8a01528181610ec601528181611061015281816113d9015281816115740152818161174f015281816118e6015281816119ce0152611d8f0152f35b600080fd5b634e487b7160e01b600052604160045260246000fd5b60408051919082016001600160401b03811183821017620001bc5760405256fe6080604052600436101561001d575b3661187b5761001b612c6d565b005b60003560e01c806306fdde031461018d5780630b3dc35414610188578063150b7a02146101835780631626ba7e1461017e57806329f8b17414610179578063333daf921461017457806334fcd5be1461016f5780633659cfe61461016a5780633a871cdd146101655780633e1b08121461016057806351166ba01461015b578063519454471461015657806354fd4d501461015157806355b14f501461014c57806357b750471461014757806384b0196e1461014257806388e7fd061461013d578063b0d691fe14610138578063bc197c8114610133578063d087d2881461012e578063d1f5789414610129578063d5416221146101245763f23a6e610361000e576117ea565b611706565b6115b8565b6114f5565b61142e565b61138e565b611329565b611247565b6111c7565b61103e565b610fad565b610e4c565b610cfd565b610be7565b610b86565b610ab5565b610995565b6108f1565b6105c8565b610546565b610463565b610398565b610346565b600091031261019d57565b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b67ffffffffffffffff81116101e557604052565b6101a2565b6060810190811067ffffffffffffffff8211176101e557604052565b6080810190811067ffffffffffffffff8211176101e557604052565b6040810190811067ffffffffffffffff8211176101e557604052565b90601f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0910116810190811067ffffffffffffffff8211176101e557604052565b6040519061028c82610206565b565b60405190610160820182811067ffffffffffffffff8211176101e557604052565b604051906102bc82610222565b600682527f4b65726e656c00000000000000000000000000000000000000000000000000006020830152565b919082519283825260005b8481106103325750507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f8460006020809697860101520116010190565b6020818301810151848301820152016102f3565b3461019d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019d576103946103806102af565b6040519182916020835260208301906102e8565b0390f35b3461019d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019d5760207f439ffe7df606b78489639bc0b827913bd09e1246fa6802968a5b3694c53e0dd95460501c73ffffffffffffffffffffffffffffffffffffffff60405191168152f35b73ffffffffffffffffffffffffffffffffffffffff81160361019d57565b359061028c8261040c565b9181601f8401121561019d5782359167ffffffffffffffff831161019d576020838186019501011161019d57565b3461019d5760807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019d5761049d60043561040c565b6104a860243561040c565b60643567ffffffffffffffff811161019d576104c8903690600401610435565b505060206040517f150b7a02000000000000000000000000000000000000000000000000000000008152f35b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc83011261019d57600435916024359067ffffffffffffffff821161019d5761054291600401610435565b9091565b3461019d57602061055f610559366104f4565b91612741565b7fffffffff0000000000000000000000000000000000000000000000000000000060405191168152f35b600435907fffffffff000000000000000000000000000000000000000000000000000000008216820361019d57565b65ffffffffffff81160361019d57565b60c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019d576105fa610589565b6024356106068161040c565b604435916106138361040c565b60643592610620846105b8565b6084359361062d856105b8565b60a43567ffffffffffffffff811161019d5761064d903690600401610435565b95909273ffffffffffffffffffffffffffffffffffffffff92837f000000000000000000000000000000000000000000000000000000000000000016331415806108e7575b6108bd5783926106c76107fd926106b86106aa61027f565b65ffffffffffff9094168452565b65ffffffffffff166020830152565b73ffffffffffffffffffffffffffffffffffffffff8816604082015273ffffffffffffffffffffffffffffffffffffffff83166060820152610754877fffffffff00000000000000000000000000000000000000000000000000000000166000527f439ffe7df606b78489639bc0b827913bd09e1246fa6802968a5b3694c53e0dda602052604060002090565b81516020830151604084015160309190911b6bffffffffffff0000000000001665ffffffffffff9290921691909117606091821b7fffffffffffffffffffffffffffffffffffffffff00000000000000000000000016178255919091015160019190910180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b1693843b1561019d57604051927f0c95955600000000000000000000000000000000000000000000000000000000845283806108406000998a946004840161221b565b038183895af19283156108b8577fffffffff000000000000000000000000000000000000000000000000000000009361089f575b501691167fed03d2572564284398470d3f266a693e29ddfff3eba45fc06c5e91013d3213538480a480f35b806108ac6108b2926101d1565b80610192565b38610874565b611d69565b60046040517f7046c88d000000000000000000000000000000000000000000000000000000008152fd5b5030331415610692565b3461019d57602061090a610904366104f4565b91612b48565b604051908152f35b92919267ffffffffffffffff82116101e5576040519161095a60207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116018461023e565b82948184528183011161019d578281602093846000960137010152565b9080601f8301121561019d5781602061099293359101610912565b90565b6020807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019d5767ffffffffffffffff60043581811161019d573660238201121561019d578060040135918083116101e5578260051b90604090815194610a028785018761023e565b855285850191602480948601019436861161019d57848101935b868510610a2c5761001b886119b7565b843584811161019d57820160607fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffdc823603011261019d57835191610a6f836101ea565b87820135610a7c8161040c565b835260448201358b84015260648201359286841161019d57610aa68c94938a869536920101610977565b86820152815201940193610a1c565b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019d57600435610aeb8161040c565b73ffffffffffffffffffffffffffffffffffffffff90817f00000000000000000000000000000000000000000000000000000000000000001633141580610b7c575b6108bd57807f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b600080a2005b5030331415610b2d565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc60608136011261019d576004359067ffffffffffffffff821161019d5761016090823603011261019d5761090a6020916044359060243590600401611d75565b3461019d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019d5760043577ffffffffffffffffffffffffffffffffffffffffffffffff811680910361019d57604051907f35567e1a000000000000000000000000000000000000000000000000000000008252306004830152602482015260208160448173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa80156108b85761039491600091610ccf575b506040519081529081906020820190565b610cf0915060203d8111610cf6575b610ce8818361023e565b810190611c5a565b38610cbe565b503d610cde565b3461019d5760207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019d57610394610dac610d3a610589565b60006060604051610d4a81610206565b82815282602082015282604082015201527fffffffff00000000000000000000000000000000000000000000000000000000166000527f439ffe7df606b78489639bc0b827913bd09e1246fa6802968a5b3694c53e0dda602052604060002090565b73ffffffffffffffffffffffffffffffffffffffff600160405192610dd084610206565b805465ffffffffffff80821686528160301c16602086015260601c60408501520154166060820152604051918291829190916060608082019365ffffffffffff80825116845260208201511660208401528173ffffffffffffffffffffffffffffffffffffffff91826040820151166040860152015116910152565b60807ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019d57600435610e828161040c565b60443567ffffffffffffffff811161019d57610ea2903690600401610977565b90606435600281101561019d5773ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001633141580610f6a575b80610f55575b6108bd57610f028161193b565b610f2b576000828193926020839451920190602435905af13d82803e15610f27573d90f35b3d90fd5b60046040517f67ce7759000000000000000000000000000000000000000000000000000000008152fd5b50610f65610f61612871565b1590565b610ef5565b5030331415610eef565b60405190610f8182610222565b600582527f302e322e320000000000000000000000000000000000000000000000000000006020830152565b3461019d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019d57610394610380610f74565b9060407ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc83011261019d5760043561101e8161040c565b916024359067ffffffffffffffff821161019d5761054291600401610435565b61104736610fe7565b909173ffffffffffffffffffffffffffffffffffffffff807f000000000000000000000000000000000000000000000000000000000000000016331415806111bd575b6108bd57807f439ffe7df606b78489639bc0b827913bd09e1246fa6802968a5b3694c53e0dd95460501c169161112a817f439ffe7df606b78489639bc0b827913bd09e1246fa6802968a5b3694c53e0dd9907fffff0000000000000000000000000000000000000000ffffffffffffffffffff7dffffffffffffffffffffffffffffffffffffffff0000000000000000000083549260501b169116179055565b1692836040519360009586947fa35f5cdc5fbabb614b4cd5064ce5543f43dc8fab0e4da41255230eb8aba2531c8680a3813b156111b957838561119781959382947f0c9595560000000000000000000000000000000000000000000000000000000084526004840161221b565b03925af180156108b8576111a9575080f35b806108ac6111b6926101d1565b80f35b8380fd5b503033141561108a565b3461019d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019d5760207f439ffe7df606b78489639bc0b827913bd09e1246fa6802968a5b3694c53e0dd95460e01b7fffffffff0000000000000000000000000000000000000000000000000000000060405191168152f35b3461019d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019d576112d76112816102af565b611289610f74565b906040519283927f0f0000000000000000000000000000000000000000000000000000000000000084526112c960209360e08587015260e08601906102e8565b9084820360408601526102e8565b90466060840152306080840152600060a084015282820360c08401528060605192838152019160809160005b82811061131257505050500390f35b835185528695509381019392810192600101611303565b3461019d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019d5760207f439ffe7df606b78489639bc0b827913bd09e1246fa6802968a5b3694c53e0dd95465ffffffffffff60405191831c168152f35b3461019d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019d57602060405173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000168152f35b9181601f8401121561019d5782359167ffffffffffffffff831161019d576020808501948460051b01011161019d57565b3461019d5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019d5761146860043561040c565b61147360243561040c565b67ffffffffffffffff60443581811161019d576114949036906004016113fd565b505060643581811161019d576114ae9036906004016113fd565b505060843590811161019d576114c8903690600401610435565b50506040517fbc197c81000000000000000000000000000000000000000000000000000000008152602090f35b3461019d5760007ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019d576040517f35567e1a0000000000000000000000000000000000000000000000000000000081523060048201526000602482015260208160448173ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000165afa80156108b85761039491600091610ccf57506040519081529081906020820190565b6115c136610fe7565b909173ffffffffffffffffffffffffffffffffffffffff90817f439ffe7df606b78489639bc0b827913bd09e1246fa6802968a5b3694c53e0dd95460501c166116dc577f439ffe7df606b78489639bc0b827913bd09e1246fa6802968a5b3694c53e0dd980547fffff0000000000000000000000000000000000000000ffffffffffffffffffff16605083901b7dffffffffffffffffffffffffffffffffffffffff00000000000000000000161790551691823b1561019d576116bf92600092836040518096819582947f0c9595560000000000000000000000000000000000000000000000000000000084526020600485015260248401916121dc565b03925af180156108b8576116cf57005b806108ac61001b926101d1565b60046040517f0dc149f0000000000000000000000000000000000000000000000000000000008152fd5b60207ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019d57611738610589565b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016331415806117e0575b6108bd577f439ffe7df606b78489639bc0b827913bd09e1246fa6802968a5b3694c53e0dd99081547fffffffffffffffffffffffffffffffffffffffffffff0000000000000000000069ffffffffffff000000004260201b169260e01c911617179055600080f35b5030331415611778565b3461019d5760a07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc36011261019d5761182460043561040c565b61182f60243561040c565b60843567ffffffffffffffff811161019d5761184f903690600401610435565b505060206040517ff23a6e61000000000000000000000000000000000000000000000000000000008152f35b600080357fffffffff000000000000000000000000000000000000000000000000000000001681527f439ffe7df606b78489639bc0b827913bd09e1246fa6802968a5b3694c53e0dda602052604081205460601c73ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000163314158061192c575b6108bd57818091368280378136915af43d82803e15610f27573d90f35b50611935612871565b1561190f565b6002111561194557565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b80518210156119885760209160051b010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b73ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001633141580611a5e575b6108bd5780519060005b828110611a0d57505050565b600080611a1a8385611974565b51805173ffffffffffffffffffffffffffffffffffffffff166020916040838201519101519283519301915af13d6000803e15611a5957600101611a01565b3d6000fd5b50611a6a610f61612871565b6119f7565b9060041161019d5790600490565b909291928360041161019d57831161019d57600401917ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc0190565b9060241161019d5760100190601490565b9060581161019d5760380190602090565b9060241161019d5760040190602090565b9060381161019d5760240190601490565b90600a1161019d5760040190600690565b9060101161019d57600a0190600690565b9093929384831161019d57841161019d578101920390565b7fffffffff000000000000000000000000000000000000000000000000000000009035818116939260048110611b6b57505050565b60040360031b82901b16169150565b91906101608382031261019d57611b8f61028e565b92611b998161042a565b84526020810135602085015260408101359167ffffffffffffffff9283811161019d5781611bc8918401610977565b6040860152606082013583811161019d5781611be5918401610977565b60608601526080820135608086015260a082013560a086015260c082013560c086015260e082013560e086015261010080830135908601526101208083013584811161019d5782611c37918501610977565b90860152610140928383013590811161019d57611c549201610977565b90830152565b9081602091031261019d575190565b611d5f6040929594939560608352611c9a60608401825173ffffffffffffffffffffffffffffffffffffffff169052565b6020810151608084015283810151611cc0610160918260a08701526101c08601906102e8565b90611d4c611cfe6060850151937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa094858983030160c08a01526102e8565b608085015160e088015260a085015192610100938489015260c08601519061012091828a015260e08701519461014095868b01528701519089015285015184888303016101808901526102e8565b92015190848303016101a08501526102e8565b9460208201520152565b6040513d6000823e3d90fd5b73ffffffffffffffffffffffffffffffffffffffff9392917f0000000000000000000000000000000000000000000000000000000000000000851633036120af57600494853592836101448101350191876024840193013594611de1611ddb8786611a6f565b90611b36565b927fffffffff0000000000000000000000000000000000000000000000000000000080851691821561208c57611e18903690611b7a565b94611e447f439ffe7df606b78489639bc0b827913bd09e1246fa6802968a5b3694c53e0dd95460e01b90565b161615611e7457896040517ffc2f51c5000000000000000000000000000000000000000000000000000000008152fd5b979896977c0100000000000000000000000000000000000000000000000000000000810361202257509060209596611f99611f41611f18611ec7611ddb8760646000990135016024878201359101611a6f565b7fffffffff00000000000000000000000000000000000000000000000000000000166000527f439ffe7df606b78489639bc0b827913bd09e1246fa6802968a5b3694c53e0dda602052604060002090565b9980611f3b60018d015473ffffffffffffffffffffffffffffffffffffffff1690565b98611a7d565b99547fffffffffffff000000000000000000000000000000000000000000000000000079ffffffffffff00000000000000000000000000000000000000008260701b169160d01b1617995b8b612014575b3691610912565b610140850152611fd6604051998a97889687947f3a871cdd0000000000000000000000000000000000000000000000000000000086528501611c69565b0393165af19081156108b85761099292600092611ff4575b50612c9e565b61200d91925060203d8111610cf657610ce8818361023e565b9038611fee565b348080808f335af150611f92565b90959391907c02000000000000000000000000000000000000000000000000000000000361207f57612075611f9994600093612070611ddb8a606460209c01350160248d8201359101611a6f565b61222c565b9199929691611f8c565b5050505050505050600190565b9697505050505050506109929394508215612a24573434343486335af150612a24565b60046040517fd663742a000000000000000000000000000000000000000000000000000000008152fd5b7fffffffffffffffffffffffffffffffffffffffff000000000000000000000000903581811693926014811061210e57505050565b60140360031b82901b16169150565b35906020811061212b575090565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9060200360031b1b1690565b7fffffffffffff0000000000000000000000000000000000000000000000000000903581811693926006811061218d57505050565b60060360031b82901b16169150565b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f6020938084528060008686013760008582860101520116010190565b601f82602094937fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0938186528686013760008582860101520116010190565b9160206109929381815201916121dc565b91906122388282611ab8565b612241916120d9565b60601c9361224f8383611ac9565b6122589161211d565b605883016078820194858360580190612272918388611b1e565b61227b9161211d565b6122858287611ada565b61228e9161211d565b6122988388611aeb565b6122a1916120d9565b60601c6122af368787610912565b80519060200120908a604051928392602084019561234393879094939273ffffffffffffffffffffffffffffffffffffffff906080937fffffffff0000000000000000000000000000000000000000000000000000000060a08501987f3ce406685c1b3551d706d85a68afdaa49ac4e07b451ad9b8ff8b58c3ee964176865216602085015260408401521660608201520152565b037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe081018252612373908261023e565b51902061237f9061262d565b9084019660788801612392918489611b1e565b9061239c92612b48565b6123a68287611ada565b6123af9161211d565b7fffffffffffffffffffffffff0000000000000000000000000000000000000000166123da91612c9e565b968686016078019682037fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8801956124118382611afc565b61241a91612158565b60d01c926124288183611b0d565b61243191612158565b60d01c9161243f8282611aeb565b612448916120d9565b60601c9161245591611ab8565b61245e916120d9565b60601c9161246a61027f565b65ffffffffffff909516855265ffffffffffff16602085015273ffffffffffffffffffffffffffffffffffffffff16604084015273ffffffffffffffffffffffffffffffffffffffff16606083015261250e907fffffffff00000000000000000000000000000000000000000000000000000000166000527f439ffe7df606b78489639bc0b827913bd09e1246fa6802968a5b3694c53e0dda602052604060002090565b81516020830151604084015160309190911b6bffffffffffff0000000000001665ffffffffffff9290921691909117606091821b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001617825590910151600190910180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905573ffffffffffffffffffffffffffffffffffffffff871691823b1561019d5761260f92600092836040518096819582947f0c9595560000000000000000000000000000000000000000000000000000000084526004840161221b565b03925af180156108b8576126205750565b806108ac61028c926101d1565b7f00000000000000000000000000000000000000000000000000000000000000007f000000000000000000000000000000000000000000000000000000000000000030147f0000000000000000000000000000000000000000000000000000000000000000461416156126ba575b671901000000000000600052601a52603a526042601820906000603a52565b5060a06040517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f81527f000000000000000000000000000000000000000000000000000000000000000060208201527f000000000000000000000000000000000000000000000000000000000000000060408201524660608201523060808201522061269b565b9061274c9291612b48565b65ffffffffffff808260a01c16908115600114612824575b428360d01c116127fd57429116106127d85773ffffffffffffffffffffffffffffffffffffffff166127b4577f1626ba7e0000000000000000000000000000000000000000000000000000000090565b7fffffffff0000000000000000000000000000000000000000000000000000000090565b507fffffffff0000000000000000000000000000000000000000000000000000000090565b5050507fffffffff0000000000000000000000000000000000000000000000000000000090565b905080612764565b9081602091031261019d5751801515810361019d5790565b60409073ffffffffffffffffffffffffffffffffffffffff6109929493168152816020820152019061219c565b61287b3633612be5565b612a1f57600080357fffffffff000000000000000000000000000000000000000000000000000000001681527f439ffe7df606b78489639bc0b827913bd09e1246fa6802968a5b3694c53e0dda6020526040902061290c6128f3600183015473ffffffffffffffffffffffffffffffffffffffff1690565b73ffffffffffffffffffffffffffffffffffffffff1690565b9073ffffffffffffffffffffffffffffffffffffffff8216159081156129e3575b81156129b8575b50156129405750600090565b602060405180927f9ea9bd59000000000000000000000000000000000000000000000000000000008252818061297a363360048401612844565b03915afa9081156108b857600091612990575090565b610992915060203d81116129b1575b6129a9818361023e565b81019061282c565b503d61299f565b546129d2915065ffffffffffff165b65ffffffffffff1690565b65ffffffffffff4291161138612934565b905065ffffffffffff612a036129c7835465ffffffffffff9060301c1690565b168015159081612a15575b509061292d565b9050421138612a0e565b600190565b9091612a303683611b7a565b6101409283810135907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe18136030182121561019d57019384359467ffffffffffffffff861161019d5760200193853603851361019d57612a99611f9287612b1a98602098611a7d565b908301526000612ae36128f37f439ffe7df606b78489639bc0b827913bd09e1246fa6802968a5b3694c53e0dd95460501c73ffffffffffffffffffffffffffffffffffffffff1690565b92604051968795869485937f3a871cdd00000000000000000000000000000000000000000000000000000000855260048501611c69565b03925af19081156108b857600091612b30575090565b610992915060203d8111610cf657610ce8818361023e565b90602091612bcf9373ffffffffffffffffffffffffffffffffffffffff7f439ffe7df606b78489639bc0b827913bd09e1246fa6802968a5b3694c53e0dd95460501c1691604051958694859384937f333daf9200000000000000000000000000000000000000000000000000000000855260048501526040602485015260448401916121dc565b03915afa9081156108b857600091612b30575090565b61297a9160209173ffffffffffffffffffffffffffffffffffffffff807f439ffe7df606b78489639bc0b827913bd09e1246fa6802968a5b3694c53e0dd95460501c1691604051958694859384937f9ea9bd5900000000000000000000000000000000000000000000000000000000855216600484015260406024840152604483019061219c565b7f88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f8852587460408051338152346020820152a1565b73ffffffffffffffffffffffffffffffffffffffff8282181615600114612cc6575050600190565b7fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff908281831692169079ffffffffffff0000000000000000000000000000000000000000908116938415612d4b575b81168015612d44575b848110908518028085189414612d3c575b5081811190821802181790565b925038612d2f565b5080612d1e565b93508093612d1556fea164736f6c6343000812000a",
}

// KernelABI is the input ABI used to generate the binding from.
// Deprecated: Use KernelMetaData.ABI instead.
var KernelABI = KernelMetaData.ABI

// KernelBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KernelMetaData.Bin instead.
var KernelBin = KernelMetaData.Bin

// DeployKernel deploys a new Ethereum contract, binding an instance of Kernel to it.
func DeployKernel(auth *bind.TransactOpts, backend bind.ContractBackend, _entryPoint common.Address) (common.Address, *types.Transaction, *Kernel, error) {
	parsed, err := KernelMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KernelBin), backend, _entryPoint)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Kernel{KernelCaller: KernelCaller{contract: contract}, KernelTransactor: KernelTransactor{contract: contract}, KernelFilterer: KernelFilterer{contract: contract}}, nil
}

// Kernel is an auto generated Go binding around an Ethereum contract.
type Kernel struct {
	KernelCaller     // Read-only binding to the contract
	KernelTransactor // Write-only binding to the contract
	KernelFilterer   // Log filterer for contract events
}

// KernelCaller is an auto generated read-only Go binding around an Ethereum contract.
type KernelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KernelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KernelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KernelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KernelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KernelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KernelSession struct {
	Contract     *Kernel           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KernelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KernelCallerSession struct {
	Contract *KernelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// KernelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KernelTransactorSession struct {
	Contract     *KernelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KernelRaw is an auto generated low-level Go binding around an Ethereum contract.
type KernelRaw struct {
	Contract *Kernel // Generic contract binding to access the raw methods on
}

// KernelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KernelCallerRaw struct {
	Contract *KernelCaller // Generic read-only contract binding to access the raw methods on
}

// KernelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KernelTransactorRaw struct {
	Contract *KernelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKernel creates a new instance of Kernel, bound to a specific deployed contract.
func NewKernel(address common.Address, backend bind.ContractBackend) (*Kernel, error) {
	contract, err := bindKernel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Kernel{KernelCaller: KernelCaller{contract: contract}, KernelTransactor: KernelTransactor{contract: contract}, KernelFilterer: KernelFilterer{contract: contract}}, nil
}

// NewKernelCaller creates a new read-only instance of Kernel, bound to a specific deployed contract.
func NewKernelCaller(address common.Address, caller bind.ContractCaller) (*KernelCaller, error) {
	contract, err := bindKernel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KernelCaller{contract: contract}, nil
}

// NewKernelTransactor creates a new write-only instance of Kernel, bound to a specific deployed contract.
func NewKernelTransactor(address common.Address, transactor bind.ContractTransactor) (*KernelTransactor, error) {
	contract, err := bindKernel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KernelTransactor{contract: contract}, nil
}

// NewKernelFilterer creates a new log filterer instance of Kernel, bound to a specific deployed contract.
func NewKernelFilterer(address common.Address, filterer bind.ContractFilterer) (*KernelFilterer, error) {
	contract, err := bindKernel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KernelFilterer{contract: contract}, nil
}

// bindKernel binds a generic wrapper to an already deployed contract.
func bindKernel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KernelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kernel *KernelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kernel.Contract.KernelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kernel *KernelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kernel.Contract.KernelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kernel *KernelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kernel.Contract.KernelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Kernel *KernelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Kernel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Kernel *KernelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kernel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Kernel *KernelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Kernel.Contract.contract.Transact(opts, method, params...)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Kernel *KernelCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Kernel.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Kernel *KernelSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Kernel.Contract.Eip712Domain(&_Kernel.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Kernel *KernelCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Kernel.Contract.Eip712Domain(&_Kernel.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Kernel *KernelCaller) EntryPoint(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kernel.contract.Call(opts, &out, "entryPoint")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Kernel *KernelSession) EntryPoint() (common.Address, error) {
	return _Kernel.Contract.EntryPoint(&_Kernel.CallOpts)
}

// EntryPoint is a free data retrieval call binding the contract method 0xb0d691fe.
//
// Solidity: function entryPoint() view returns(address)
func (_Kernel *KernelCallerSession) EntryPoint() (common.Address, error) {
	return _Kernel.Contract.EntryPoint(&_Kernel.CallOpts)
}

// GetDefaultValidator is a free data retrieval call binding the contract method 0x0b3dc354.
//
// Solidity: function getDefaultValidator() view returns(address validator)
func (_Kernel *KernelCaller) GetDefaultValidator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Kernel.contract.Call(opts, &out, "getDefaultValidator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetDefaultValidator is a free data retrieval call binding the contract method 0x0b3dc354.
//
// Solidity: function getDefaultValidator() view returns(address validator)
func (_Kernel *KernelSession) GetDefaultValidator() (common.Address, error) {
	return _Kernel.Contract.GetDefaultValidator(&_Kernel.CallOpts)
}

// GetDefaultValidator is a free data retrieval call binding the contract method 0x0b3dc354.
//
// Solidity: function getDefaultValidator() view returns(address validator)
func (_Kernel *KernelCallerSession) GetDefaultValidator() (common.Address, error) {
	return _Kernel.Contract.GetDefaultValidator(&_Kernel.CallOpts)
}

// GetDisabledMode is a free data retrieval call binding the contract method 0x57b75047.
//
// Solidity: function getDisabledMode() view returns(bytes4 disabled)
func (_Kernel *KernelCaller) GetDisabledMode(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Kernel.contract.Call(opts, &out, "getDisabledMode")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// GetDisabledMode is a free data retrieval call binding the contract method 0x57b75047.
//
// Solidity: function getDisabledMode() view returns(bytes4 disabled)
func (_Kernel *KernelSession) GetDisabledMode() ([4]byte, error) {
	return _Kernel.Contract.GetDisabledMode(&_Kernel.CallOpts)
}

// GetDisabledMode is a free data retrieval call binding the contract method 0x57b75047.
//
// Solidity: function getDisabledMode() view returns(bytes4 disabled)
func (_Kernel *KernelCallerSession) GetDisabledMode() ([4]byte, error) {
	return _Kernel.Contract.GetDisabledMode(&_Kernel.CallOpts)
}

// GetExecution is a free data retrieval call binding the contract method 0x51166ba0.
//
// Solidity: function getExecution(bytes4 _selector) view returns((uint48,uint48,address,address))
func (_Kernel *KernelCaller) GetExecution(opts *bind.CallOpts, _selector [4]byte) (ExecutionDetail, error) {
	var out []interface{}
	err := _Kernel.contract.Call(opts, &out, "getExecution", _selector)

	if err != nil {
		return *new(ExecutionDetail), err
	}

	out0 := *abi.ConvertType(out[0], new(ExecutionDetail)).(*ExecutionDetail)

	return out0, err

}

// GetExecution is a free data retrieval call binding the contract method 0x51166ba0.
//
// Solidity: function getExecution(bytes4 _selector) view returns((uint48,uint48,address,address))
func (_Kernel *KernelSession) GetExecution(_selector [4]byte) (ExecutionDetail, error) {
	return _Kernel.Contract.GetExecution(&_Kernel.CallOpts, _selector)
}

// GetExecution is a free data retrieval call binding the contract method 0x51166ba0.
//
// Solidity: function getExecution(bytes4 _selector) view returns((uint48,uint48,address,address))
func (_Kernel *KernelCallerSession) GetExecution(_selector [4]byte) (ExecutionDetail, error) {
	return _Kernel.Contract.GetExecution(&_Kernel.CallOpts, _selector)
}

// GetLastDisabledTime is a free data retrieval call binding the contract method 0x88e7fd06.
//
// Solidity: function getLastDisabledTime() view returns(uint48)
func (_Kernel *KernelCaller) GetLastDisabledTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kernel.contract.Call(opts, &out, "getLastDisabledTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastDisabledTime is a free data retrieval call binding the contract method 0x88e7fd06.
//
// Solidity: function getLastDisabledTime() view returns(uint48)
func (_Kernel *KernelSession) GetLastDisabledTime() (*big.Int, error) {
	return _Kernel.Contract.GetLastDisabledTime(&_Kernel.CallOpts)
}

// GetLastDisabledTime is a free data retrieval call binding the contract method 0x88e7fd06.
//
// Solidity: function getLastDisabledTime() view returns(uint48)
func (_Kernel *KernelCallerSession) GetLastDisabledTime() (*big.Int, error) {
	return _Kernel.Contract.GetLastDisabledTime(&_Kernel.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0x3e1b0812.
//
// Solidity: function getNonce(uint192 key) view returns(uint256)
func (_Kernel *KernelCaller) GetNonce(opts *bind.CallOpts, key *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Kernel.contract.Call(opts, &out, "getNonce", key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x3e1b0812.
//
// Solidity: function getNonce(uint192 key) view returns(uint256)
func (_Kernel *KernelSession) GetNonce(key *big.Int) (*big.Int, error) {
	return _Kernel.Contract.GetNonce(&_Kernel.CallOpts, key)
}

// GetNonce is a free data retrieval call binding the contract method 0x3e1b0812.
//
// Solidity: function getNonce(uint192 key) view returns(uint256)
func (_Kernel *KernelCallerSession) GetNonce(key *big.Int) (*big.Int, error) {
	return _Kernel.Contract.GetNonce(&_Kernel.CallOpts, key)
}

// GetNonce0 is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Kernel *KernelCaller) GetNonce0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Kernel.contract.Call(opts, &out, "getNonce0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce0 is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Kernel *KernelSession) GetNonce0() (*big.Int, error) {
	return _Kernel.Contract.GetNonce0(&_Kernel.CallOpts)
}

// GetNonce0 is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Kernel *KernelCallerSession) GetNonce0() (*big.Int, error) {
	return _Kernel.Contract.GetNonce0(&_Kernel.CallOpts)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4)
func (_Kernel *KernelCaller) IsValidSignature(opts *bind.CallOpts, hash [32]byte, signature []byte) ([4]byte, error) {
	var out []interface{}
	err := _Kernel.contract.Call(opts, &out, "isValidSignature", hash, signature)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4)
func (_Kernel *KernelSession) IsValidSignature(hash [32]byte, signature []byte) ([4]byte, error) {
	return _Kernel.Contract.IsValidSignature(&_Kernel.CallOpts, hash, signature)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x1626ba7e.
//
// Solidity: function isValidSignature(bytes32 hash, bytes signature) view returns(bytes4)
func (_Kernel *KernelCallerSession) IsValidSignature(hash [32]byte, signature []byte) ([4]byte, error) {
	return _Kernel.Contract.IsValidSignature(&_Kernel.CallOpts, hash, signature)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Kernel *KernelCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Kernel.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Kernel *KernelSession) Name() (string, error) {
	return _Kernel.Contract.Name(&_Kernel.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Kernel *KernelCallerSession) Name() (string, error) {
	return _Kernel.Contract.Name(&_Kernel.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Kernel *KernelCaller) OnERC1155BatchReceived(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Kernel.contract.Call(opts, &out, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Kernel *KernelSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Kernel.Contract.OnERC1155BatchReceived(&_Kernel.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) pure returns(bytes4)
func (_Kernel *KernelCallerSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) ([4]byte, error) {
	return _Kernel.Contract.OnERC1155BatchReceived(&_Kernel.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Kernel *KernelCaller) OnERC1155Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Kernel.contract.Call(opts, &out, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Kernel *KernelSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Kernel.Contract.OnERC1155Received(&_Kernel.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) pure returns(bytes4)
func (_Kernel *KernelCallerSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) ([4]byte, error) {
	return _Kernel.Contract.OnERC1155Received(&_Kernel.CallOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Kernel *KernelCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Kernel.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Kernel *KernelSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Kernel.Contract.OnERC721Received(&_Kernel.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Kernel *KernelCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Kernel.Contract.OnERC721Received(&_Kernel.CallOpts, arg0, arg1, arg2, arg3)
}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 hash, bytes signature) view returns(uint256)
func (_Kernel *KernelCaller) ValidateSignature(opts *bind.CallOpts, hash [32]byte, signature []byte) (*big.Int, error) {
	var out []interface{}
	err := _Kernel.contract.Call(opts, &out, "validateSignature", hash, signature)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 hash, bytes signature) view returns(uint256)
func (_Kernel *KernelSession) ValidateSignature(hash [32]byte, signature []byte) (*big.Int, error) {
	return _Kernel.Contract.ValidateSignature(&_Kernel.CallOpts, hash, signature)
}

// ValidateSignature is a free data retrieval call binding the contract method 0x333daf92.
//
// Solidity: function validateSignature(bytes32 hash, bytes signature) view returns(uint256)
func (_Kernel *KernelCallerSession) ValidateSignature(hash [32]byte, signature []byte) (*big.Int, error) {
	return _Kernel.Contract.ValidateSignature(&_Kernel.CallOpts, hash, signature)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Kernel *KernelCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Kernel.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Kernel *KernelSession) Version() (string, error) {
	return _Kernel.Contract.Version(&_Kernel.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Kernel *KernelCallerSession) Version() (string, error) {
	return _Kernel.Contract.Version(&_Kernel.CallOpts)
}

// DisableMode is a paid mutator transaction binding the contract method 0xd5416221.
//
// Solidity: function disableMode(bytes4 _disableFlag) payable returns()
func (_Kernel *KernelTransactor) DisableMode(opts *bind.TransactOpts, _disableFlag [4]byte) (*types.Transaction, error) {
	return _Kernel.contract.Transact(opts, "disableMode", _disableFlag)
}

// DisableMode is a paid mutator transaction binding the contract method 0xd5416221.
//
// Solidity: function disableMode(bytes4 _disableFlag) payable returns()
func (_Kernel *KernelSession) DisableMode(_disableFlag [4]byte) (*types.Transaction, error) {
	return _Kernel.Contract.DisableMode(&_Kernel.TransactOpts, _disableFlag)
}

// DisableMode is a paid mutator transaction binding the contract method 0xd5416221.
//
// Solidity: function disableMode(bytes4 _disableFlag) payable returns()
func (_Kernel *KernelTransactorSession) DisableMode(_disableFlag [4]byte) (*types.Transaction, error) {
	return _Kernel.Contract.DisableMode(&_Kernel.TransactOpts, _disableFlag)
}

// Execute is a paid mutator transaction binding the contract method 0x51945447.
//
// Solidity: function execute(address to, uint256 value, bytes data, uint8 _operation) payable returns()
func (_Kernel *KernelTransactor) Execute(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, _operation uint8) (*types.Transaction, error) {
	return _Kernel.contract.Transact(opts, "execute", to, value, data, _operation)
}

// Execute is a paid mutator transaction binding the contract method 0x51945447.
//
// Solidity: function execute(address to, uint256 value, bytes data, uint8 _operation) payable returns()
func (_Kernel *KernelSession) Execute(to common.Address, value *big.Int, data []byte, _operation uint8) (*types.Transaction, error) {
	return _Kernel.Contract.Execute(&_Kernel.TransactOpts, to, value, data, _operation)
}

// Execute is a paid mutator transaction binding the contract method 0x51945447.
//
// Solidity: function execute(address to, uint256 value, bytes data, uint8 _operation) payable returns()
func (_Kernel *KernelTransactorSession) Execute(to common.Address, value *big.Int, data []byte, _operation uint8) (*types.Transaction, error) {
	return _Kernel.Contract.Execute(&_Kernel.TransactOpts, to, value, data, _operation)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) payable returns()
func (_Kernel *KernelTransactor) ExecuteBatch(opts *bind.TransactOpts, calls []Call) (*types.Transaction, error) {
	return _Kernel.contract.Transact(opts, "executeBatch", calls)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) payable returns()
func (_Kernel *KernelSession) ExecuteBatch(calls []Call) (*types.Transaction, error) {
	return _Kernel.Contract.ExecuteBatch(&_Kernel.TransactOpts, calls)
}

// ExecuteBatch is a paid mutator transaction binding the contract method 0x34fcd5be.
//
// Solidity: function executeBatch((address,uint256,bytes)[] calls) payable returns()
func (_Kernel *KernelTransactorSession) ExecuteBatch(calls []Call) (*types.Transaction, error) {
	return _Kernel.Contract.ExecuteBatch(&_Kernel.TransactOpts, calls)
}

// Initialize is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _defaultValidator, bytes _data) payable returns()
func (_Kernel *KernelTransactor) Initialize(opts *bind.TransactOpts, _defaultValidator common.Address, _data []byte) (*types.Transaction, error) {
	return _Kernel.contract.Transact(opts, "initialize", _defaultValidator, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _defaultValidator, bytes _data) payable returns()
func (_Kernel *KernelSession) Initialize(_defaultValidator common.Address, _data []byte) (*types.Transaction, error) {
	return _Kernel.Contract.Initialize(&_Kernel.TransactOpts, _defaultValidator, _data)
}

// Initialize is a paid mutator transaction binding the contract method 0xd1f57894.
//
// Solidity: function initialize(address _defaultValidator, bytes _data) payable returns()
func (_Kernel *KernelTransactorSession) Initialize(_defaultValidator common.Address, _data []byte) (*types.Transaction, error) {
	return _Kernel.Contract.Initialize(&_Kernel.TransactOpts, _defaultValidator, _data)
}

// SetDefaultValidator is a paid mutator transaction binding the contract method 0x55b14f50.
//
// Solidity: function setDefaultValidator(address _defaultValidator, bytes _data) payable returns()
func (_Kernel *KernelTransactor) SetDefaultValidator(opts *bind.TransactOpts, _defaultValidator common.Address, _data []byte) (*types.Transaction, error) {
	return _Kernel.contract.Transact(opts, "setDefaultValidator", _defaultValidator, _data)
}

// SetDefaultValidator is a paid mutator transaction binding the contract method 0x55b14f50.
//
// Solidity: function setDefaultValidator(address _defaultValidator, bytes _data) payable returns()
func (_Kernel *KernelSession) SetDefaultValidator(_defaultValidator common.Address, _data []byte) (*types.Transaction, error) {
	return _Kernel.Contract.SetDefaultValidator(&_Kernel.TransactOpts, _defaultValidator, _data)
}

// SetDefaultValidator is a paid mutator transaction binding the contract method 0x55b14f50.
//
// Solidity: function setDefaultValidator(address _defaultValidator, bytes _data) payable returns()
func (_Kernel *KernelTransactorSession) SetDefaultValidator(_defaultValidator common.Address, _data []byte) (*types.Transaction, error) {
	return _Kernel.Contract.SetDefaultValidator(&_Kernel.TransactOpts, _defaultValidator, _data)
}

// SetExecution is a paid mutator transaction binding the contract method 0x29f8b174.
//
// Solidity: function setExecution(bytes4 _selector, address _executor, address _validator, uint48 _validUntil, uint48 _validAfter, bytes _enableData) payable returns()
func (_Kernel *KernelTransactor) SetExecution(opts *bind.TransactOpts, _selector [4]byte, _executor common.Address, _validator common.Address, _validUntil *big.Int, _validAfter *big.Int, _enableData []byte) (*types.Transaction, error) {
	return _Kernel.contract.Transact(opts, "setExecution", _selector, _executor, _validator, _validUntil, _validAfter, _enableData)
}

// SetExecution is a paid mutator transaction binding the contract method 0x29f8b174.
//
// Solidity: function setExecution(bytes4 _selector, address _executor, address _validator, uint48 _validUntil, uint48 _validAfter, bytes _enableData) payable returns()
func (_Kernel *KernelSession) SetExecution(_selector [4]byte, _executor common.Address, _validator common.Address, _validUntil *big.Int, _validAfter *big.Int, _enableData []byte) (*types.Transaction, error) {
	return _Kernel.Contract.SetExecution(&_Kernel.TransactOpts, _selector, _executor, _validator, _validUntil, _validAfter, _enableData)
}

// SetExecution is a paid mutator transaction binding the contract method 0x29f8b174.
//
// Solidity: function setExecution(bytes4 _selector, address _executor, address _validator, uint48 _validUntil, uint48 _validAfter, bytes _enableData) payable returns()
func (_Kernel *KernelTransactorSession) SetExecution(_selector [4]byte, _executor common.Address, _validator common.Address, _validUntil *big.Int, _validAfter *big.Int, _enableData []byte) (*types.Transaction, error) {
	return _Kernel.Contract.SetExecution(&_Kernel.TransactOpts, _selector, _executor, _validator, _validUntil, _validAfter, _enableData)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address _newImplementation) payable returns()
func (_Kernel *KernelTransactor) UpgradeTo(opts *bind.TransactOpts, _newImplementation common.Address) (*types.Transaction, error) {
	return _Kernel.contract.Transact(opts, "upgradeTo", _newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address _newImplementation) payable returns()
func (_Kernel *KernelSession) UpgradeTo(_newImplementation common.Address) (*types.Transaction, error) {
	return _Kernel.Contract.UpgradeTo(&_Kernel.TransactOpts, _newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address _newImplementation) payable returns()
func (_Kernel *KernelTransactorSession) UpgradeTo(_newImplementation common.Address) (*types.Transaction, error) {
	return _Kernel.Contract.UpgradeTo(&_Kernel.TransactOpts, _newImplementation)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) _userOp, bytes32 userOpHash, uint256 missingAccountFunds) payable returns(uint256 validationData)
func (_Kernel *KernelTransactor) ValidateUserOp(opts *bind.TransactOpts, _userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Kernel.contract.Transact(opts, "validateUserOp", _userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) _userOp, bytes32 userOpHash, uint256 missingAccountFunds) payable returns(uint256 validationData)
func (_Kernel *KernelSession) ValidateUserOp(_userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Kernel.Contract.ValidateUserOp(&_Kernel.TransactOpts, _userOp, userOpHash, missingAccountFunds)
}

// ValidateUserOp is a paid mutator transaction binding the contract method 0x3a871cdd.
//
// Solidity: function validateUserOp((address,uint256,bytes,bytes,uint256,uint256,uint256,uint256,uint256,bytes,bytes) _userOp, bytes32 userOpHash, uint256 missingAccountFunds) payable returns(uint256 validationData)
func (_Kernel *KernelTransactorSession) ValidateUserOp(_userOp UserOperation, userOpHash [32]byte, missingAccountFunds *big.Int) (*types.Transaction, error) {
	return _Kernel.Contract.ValidateUserOp(&_Kernel.TransactOpts, _userOp, userOpHash, missingAccountFunds)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Kernel *KernelTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Kernel.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Kernel *KernelSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Kernel.Contract.Fallback(&_Kernel.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Kernel *KernelTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Kernel.Contract.Fallback(&_Kernel.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Kernel *KernelTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Kernel.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Kernel *KernelSession) Receive() (*types.Transaction, error) {
	return _Kernel.Contract.Receive(&_Kernel.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Kernel *KernelTransactorSession) Receive() (*types.Transaction, error) {
	return _Kernel.Contract.Receive(&_Kernel.TransactOpts)
}

// KernelDefaultValidatorChangedIterator is returned from FilterDefaultValidatorChanged and is used to iterate over the raw logs and unpacked data for DefaultValidatorChanged events raised by the Kernel contract.
type KernelDefaultValidatorChangedIterator struct {
	Event *KernelDefaultValidatorChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KernelDefaultValidatorChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KernelDefaultValidatorChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KernelDefaultValidatorChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KernelDefaultValidatorChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KernelDefaultValidatorChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KernelDefaultValidatorChanged represents a DefaultValidatorChanged event raised by the Kernel contract.
type KernelDefaultValidatorChanged struct {
	OldValidator common.Address
	NewValidator common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDefaultValidatorChanged is a free log retrieval operation binding the contract event 0xa35f5cdc5fbabb614b4cd5064ce5543f43dc8fab0e4da41255230eb8aba2531c.
//
// Solidity: event DefaultValidatorChanged(address indexed oldValidator, address indexed newValidator)
func (_Kernel *KernelFilterer) FilterDefaultValidatorChanged(opts *bind.FilterOpts, oldValidator []common.Address, newValidator []common.Address) (*KernelDefaultValidatorChangedIterator, error) {

	var oldValidatorRule []interface{}
	for _, oldValidatorItem := range oldValidator {
		oldValidatorRule = append(oldValidatorRule, oldValidatorItem)
	}
	var newValidatorRule []interface{}
	for _, newValidatorItem := range newValidator {
		newValidatorRule = append(newValidatorRule, newValidatorItem)
	}

	logs, sub, err := _Kernel.contract.FilterLogs(opts, "DefaultValidatorChanged", oldValidatorRule, newValidatorRule)
	if err != nil {
		return nil, err
	}
	return &KernelDefaultValidatorChangedIterator{contract: _Kernel.contract, event: "DefaultValidatorChanged", logs: logs, sub: sub}, nil
}

// WatchDefaultValidatorChanged is a free log subscription operation binding the contract event 0xa35f5cdc5fbabb614b4cd5064ce5543f43dc8fab0e4da41255230eb8aba2531c.
//
// Solidity: event DefaultValidatorChanged(address indexed oldValidator, address indexed newValidator)
func (_Kernel *KernelFilterer) WatchDefaultValidatorChanged(opts *bind.WatchOpts, sink chan<- *KernelDefaultValidatorChanged, oldValidator []common.Address, newValidator []common.Address) (event.Subscription, error) {

	var oldValidatorRule []interface{}
	for _, oldValidatorItem := range oldValidator {
		oldValidatorRule = append(oldValidatorRule, oldValidatorItem)
	}
	var newValidatorRule []interface{}
	for _, newValidatorItem := range newValidator {
		newValidatorRule = append(newValidatorRule, newValidatorItem)
	}

	logs, sub, err := _Kernel.contract.WatchLogs(opts, "DefaultValidatorChanged", oldValidatorRule, newValidatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KernelDefaultValidatorChanged)
				if err := _Kernel.contract.UnpackLog(event, "DefaultValidatorChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDefaultValidatorChanged is a log parse operation binding the contract event 0xa35f5cdc5fbabb614b4cd5064ce5543f43dc8fab0e4da41255230eb8aba2531c.
//
// Solidity: event DefaultValidatorChanged(address indexed oldValidator, address indexed newValidator)
func (_Kernel *KernelFilterer) ParseDefaultValidatorChanged(log types.Log) (*KernelDefaultValidatorChanged, error) {
	event := new(KernelDefaultValidatorChanged)
	if err := _Kernel.contract.UnpackLog(event, "DefaultValidatorChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KernelExecutionChangedIterator is returned from FilterExecutionChanged and is used to iterate over the raw logs and unpacked data for ExecutionChanged events raised by the Kernel contract.
type KernelExecutionChangedIterator struct {
	Event *KernelExecutionChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KernelExecutionChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KernelExecutionChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KernelExecutionChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KernelExecutionChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KernelExecutionChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KernelExecutionChanged represents a ExecutionChanged event raised by the Kernel contract.
type KernelExecutionChanged struct {
	Selector  [4]byte
	Executor  common.Address
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExecutionChanged is a free log retrieval operation binding the contract event 0xed03d2572564284398470d3f266a693e29ddfff3eba45fc06c5e91013d321353.
//
// Solidity: event ExecutionChanged(bytes4 indexed selector, address indexed executor, address indexed validator)
func (_Kernel *KernelFilterer) FilterExecutionChanged(opts *bind.FilterOpts, selector [][4]byte, executor []common.Address, validator []common.Address) (*KernelExecutionChangedIterator, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Kernel.contract.FilterLogs(opts, "ExecutionChanged", selectorRule, executorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &KernelExecutionChangedIterator{contract: _Kernel.contract, event: "ExecutionChanged", logs: logs, sub: sub}, nil
}

// WatchExecutionChanged is a free log subscription operation binding the contract event 0xed03d2572564284398470d3f266a693e29ddfff3eba45fc06c5e91013d321353.
//
// Solidity: event ExecutionChanged(bytes4 indexed selector, address indexed executor, address indexed validator)
func (_Kernel *KernelFilterer) WatchExecutionChanged(opts *bind.WatchOpts, sink chan<- *KernelExecutionChanged, selector [][4]byte, executor []common.Address, validator []common.Address) (event.Subscription, error) {

	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}
	var executorRule []interface{}
	for _, executorItem := range executor {
		executorRule = append(executorRule, executorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Kernel.contract.WatchLogs(opts, "ExecutionChanged", selectorRule, executorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KernelExecutionChanged)
				if err := _Kernel.contract.UnpackLog(event, "ExecutionChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionChanged is a log parse operation binding the contract event 0xed03d2572564284398470d3f266a693e29ddfff3eba45fc06c5e91013d321353.
//
// Solidity: event ExecutionChanged(bytes4 indexed selector, address indexed executor, address indexed validator)
func (_Kernel *KernelFilterer) ParseExecutionChanged(log types.Log) (*KernelExecutionChanged, error) {
	event := new(KernelExecutionChanged)
	if err := _Kernel.contract.UnpackLog(event, "ExecutionChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KernelReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Kernel contract.
type KernelReceivedIterator struct {
	Event *KernelReceived // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KernelReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KernelReceived)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KernelReceived)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KernelReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KernelReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KernelReceived represents a Received event raised by the Kernel contract.
type KernelReceived struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address sender, uint256 amount)
func (_Kernel *KernelFilterer) FilterReceived(opts *bind.FilterOpts) (*KernelReceivedIterator, error) {

	logs, sub, err := _Kernel.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &KernelReceivedIterator{contract: _Kernel.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address sender, uint256 amount)
func (_Kernel *KernelFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *KernelReceived) (event.Subscription, error) {

	logs, sub, err := _Kernel.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KernelReceived)
				if err := _Kernel.contract.UnpackLog(event, "Received", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReceived is a log parse operation binding the contract event 0x88a5966d370b9919b20f3e2c13ff65706f196a4e32cc2c12bf57088f88525874.
//
// Solidity: event Received(address sender, uint256 amount)
func (_Kernel *KernelFilterer) ParseReceived(log types.Log) (*KernelReceived, error) {
	event := new(KernelReceived)
	if err := _Kernel.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KernelUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Kernel contract.
type KernelUpgradedIterator struct {
	Event *KernelUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *KernelUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KernelUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(KernelUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *KernelUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KernelUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KernelUpgraded represents a Upgraded event raised by the Kernel contract.
type KernelUpgraded struct {
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed newImplementation)
func (_Kernel *KernelFilterer) FilterUpgraded(opts *bind.FilterOpts, newImplementation []common.Address) (*KernelUpgradedIterator, error) {

	var newImplementationRule []interface{}
	for _, newImplementationItem := range newImplementation {
		newImplementationRule = append(newImplementationRule, newImplementationItem)
	}

	logs, sub, err := _Kernel.contract.FilterLogs(opts, "Upgraded", newImplementationRule)
	if err != nil {
		return nil, err
	}
	return &KernelUpgradedIterator{contract: _Kernel.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed newImplementation)
func (_Kernel *KernelFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *KernelUpgraded, newImplementation []common.Address) (event.Subscription, error) {

	var newImplementationRule []interface{}
	for _, newImplementationItem := range newImplementation {
		newImplementationRule = append(newImplementationRule, newImplementationItem)
	}

	logs, sub, err := _Kernel.contract.WatchLogs(opts, "Upgraded", newImplementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KernelUpgraded)
				if err := _Kernel.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed newImplementation)
func (_Kernel *KernelFilterer) ParseUpgraded(log types.Log) (*KernelUpgraded, error) {
	event := new(KernelUpgraded)
	if err := _Kernel.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
