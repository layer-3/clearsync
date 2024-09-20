// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iclaim_executor

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
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

// ClaimExecutorMetaData contains all meta data concerning the ClaimExecutor contract.
var ClaimExecutorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"COOLDOWN_DAYS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"router\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usersData\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"latestClaimTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"streakCount\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"points\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"streakCount\",\"type\":\"uint128\",\"indexed\":false,\"internalType\":\"uint128\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Executed\",\"inputs\":[{\"name\":\"beneficiary\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CallerIsNotRouter\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidPayload\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RewardOnCooldown\",\"inputs\":[{\"name\":\"cooldownDeadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161070638038061070683398101604081905261002f91610059565b6001600081905580546001600160a01b0319166001600160a01b0392909216919091179055610089565b60006020828403121561006b57600080fd5b81516001600160a01b038116811461008257600080fd5b9392505050565b61066e806100986000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630560ab69146100515780631cff79cd146100b857806323f32c05146100cd578063f887ea40146100e3575b600080fd5b61008c61005f36600461040b565b6002602052600090815260409020546001600160401b03811690600160401b90046001600160801b031682565b604080516001600160401b0390931683526001600160801b039091166020830152015b60405180910390f35b6100cb6100c636600461042d565b61010e565b005b6100d5600181565b6040519081526020016100af565b6001546100f6906001600160a01b031681565b6040516001600160a01b0390911681526020016100af565b61011661019e565b6001546001600160a01b0316331461014157604051630693060f60e01b815260040160405180910390fd5b61014c8383836101c8565b826001600160a01b03167f0dcf4ffb3b85ab072df6f9ed79b0382ec6c9619a98f36f4538d3b2e87fd3fd1183836040516101879291906104af565b60405180910390a26101996001600055565b505050565b6002600054036101c157604051633ee5aeb560e01b815260040160405180910390fd5b6002600055565b6001600160a01b038316600090815260026020526040812080549091906101fc9062015180906001600160401b031661050a565b6001600160401b03166102126201518042610530565b61021c9190610544565b905060018161ffff1610156102855781546102449062015180906001600160401b031661050a565b61024f90600161055d565b61025c9062015180610584565b604051638f26d9a160e01b81526001600160401b03909116600482015260240160405180910390fd5b6000610293848601866105af565b80519091506001600160401b03166000036102c157604051637c6953f960e01b815260040160405180910390fd5b82546102da9062015180906001600160401b031661050a565b6001600160401b03166102f06201518042610530565b6102fa9190610544565b600103610349578254600160401b90046001600160801b031683600861031f83610612565b91906101000a8154816001600160801b0302191690836001600160801b031602179055505061036f565b825477ffffffffffffffffffffffffffffffff00000000000000001916600160401b1783555b825467ffffffffffffffff1916426001600160401b0390811691909117808555825160408051919093168152600160401b9091046001600160801b031660208201526001600160a01b038816917f9d79a1988c377d51e2ebd376a251db4f9ab5f1435fd3617f80f4c1291a4ede19910160405180910390a2505050505050565b80356001600160a01b038116811461040657600080fd5b919050565b60006020828403121561041d57600080fd5b610426826103ef565b9392505050565b60008060006040848603121561044257600080fd5b61044b846103ef565b925060208401356001600160401b038082111561046757600080fd5b818601915086601f83011261047b57600080fd5b81358181111561048a57600080fd5b87602082850101111561049c57600080fd5b6020830194508093505050509250925092565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60006001600160401b0380841680610524576105246104de565b92169190910492915050565b60008261053f5761053f6104de565b500490565b81810381811115610557576105576104f4565b92915050565b6001600160401b0381811683821601908082111561057d5761057d6104f4565b5092915050565b6001600160401b038181168382160280821691908281146105a7576105a76104f4565b505092915050565b6000602082840312156105c157600080fd5b604051602081016001600160401b0382821081831117156105f257634e487b7160e01b600052604160045260246000fd5b8160405284359150808216821461060857600080fd5b5081529392505050565b60006001600160801b0380831681810361062e5761062e6104f4565b600101939250505056fea2646970667358221220abb45fa06b57da0485e87ff6d9af0c499d8abe69b5867ed0214af950ae6ba80064736f6c63430008170033",
}

// ClaimExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use ClaimExecutorMetaData.ABI instead.
var ClaimExecutorABI = ClaimExecutorMetaData.ABI

// ClaimExecutorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ClaimExecutorMetaData.Bin instead.
var ClaimExecutorBin = ClaimExecutorMetaData.Bin

// DeployClaimExecutor deploys a new Ethereum contract, binding an instance of ClaimExecutor to it.
func DeployClaimExecutor(auth *bind.TransactOpts, backend bind.ContractBackend, router common.Address) (common.Address, *types.Transaction, *ClaimExecutor, error) {
	parsed, err := ClaimExecutorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ClaimExecutorBin), backend, router)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ClaimExecutor{ClaimExecutorCaller: ClaimExecutorCaller{contract: contract}, ClaimExecutorTransactor: ClaimExecutorTransactor{contract: contract}, ClaimExecutorFilterer: ClaimExecutorFilterer{contract: contract}}, nil
}

// ClaimExecutor is an auto generated Go binding around an Ethereum contract.
type ClaimExecutor struct {
	ClaimExecutorCaller     // Read-only binding to the contract
	ClaimExecutorTransactor // Write-only binding to the contract
	ClaimExecutorFilterer   // Log filterer for contract events
}

// ClaimExecutorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimExecutorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimExecutorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimExecutorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimExecutorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClaimExecutorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimExecutorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimExecutorSession struct {
	Contract     *ClaimExecutor    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClaimExecutorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimExecutorCallerSession struct {
	Contract *ClaimExecutorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ClaimExecutorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimExecutorTransactorSession struct {
	Contract     *ClaimExecutorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ClaimExecutorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimExecutorRaw struct {
	Contract *ClaimExecutor // Generic contract binding to access the raw methods on
}

// ClaimExecutorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimExecutorCallerRaw struct {
	Contract *ClaimExecutorCaller // Generic read-only contract binding to access the raw methods on
}

// ClaimExecutorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimExecutorTransactorRaw struct {
	Contract *ClaimExecutorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClaimExecutor creates a new instance of ClaimExecutor, bound to a specific deployed contract.
func NewClaimExecutor(address common.Address, backend bind.ContractBackend) (*ClaimExecutor, error) {
	contract, err := bindClaimExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClaimExecutor{ClaimExecutorCaller: ClaimExecutorCaller{contract: contract}, ClaimExecutorTransactor: ClaimExecutorTransactor{contract: contract}, ClaimExecutorFilterer: ClaimExecutorFilterer{contract: contract}}, nil
}

// NewClaimExecutorCaller creates a new read-only instance of ClaimExecutor, bound to a specific deployed contract.
func NewClaimExecutorCaller(address common.Address, caller bind.ContractCaller) (*ClaimExecutorCaller, error) {
	contract, err := bindClaimExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimExecutorCaller{contract: contract}, nil
}

// NewClaimExecutorTransactor creates a new write-only instance of ClaimExecutor, bound to a specific deployed contract.
func NewClaimExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*ClaimExecutorTransactor, error) {
	contract, err := bindClaimExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimExecutorTransactor{contract: contract}, nil
}

// NewClaimExecutorFilterer creates a new log filterer instance of ClaimExecutor, bound to a specific deployed contract.
func NewClaimExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*ClaimExecutorFilterer, error) {
	contract, err := bindClaimExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClaimExecutorFilterer{contract: contract}, nil
}

// bindClaimExecutor binds a generic wrapper to an already deployed contract.
func bindClaimExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClaimExecutorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimExecutor *ClaimExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClaimExecutor.Contract.ClaimExecutorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimExecutor *ClaimExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimExecutor.Contract.ClaimExecutorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimExecutor *ClaimExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimExecutor.Contract.ClaimExecutorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClaimExecutor *ClaimExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClaimExecutor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClaimExecutor *ClaimExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClaimExecutor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClaimExecutor *ClaimExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClaimExecutor.Contract.contract.Transact(opts, method, params...)
}

// COOLDOWNDAYS is a free data retrieval call binding the contract method 0x23f32c05.
//
// Solidity: function COOLDOWN_DAYS() view returns(uint256)
func (_ClaimExecutor *ClaimExecutorCaller) COOLDOWNDAYS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClaimExecutor.contract.Call(opts, &out, "COOLDOWN_DAYS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COOLDOWNDAYS is a free data retrieval call binding the contract method 0x23f32c05.
//
// Solidity: function COOLDOWN_DAYS() view returns(uint256)
func (_ClaimExecutor *ClaimExecutorSession) COOLDOWNDAYS() (*big.Int, error) {
	return _ClaimExecutor.Contract.COOLDOWNDAYS(&_ClaimExecutor.CallOpts)
}

// COOLDOWNDAYS is a free data retrieval call binding the contract method 0x23f32c05.
//
// Solidity: function COOLDOWN_DAYS() view returns(uint256)
func (_ClaimExecutor *ClaimExecutorCallerSession) COOLDOWNDAYS() (*big.Int, error) {
	return _ClaimExecutor.Contract.COOLDOWNDAYS(&_ClaimExecutor.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_ClaimExecutor *ClaimExecutorCaller) Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClaimExecutor.contract.Call(opts, &out, "router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_ClaimExecutor *ClaimExecutorSession) Router() (common.Address, error) {
	return _ClaimExecutor.Contract.Router(&_ClaimExecutor.CallOpts)
}

// Router is a free data retrieval call binding the contract method 0xf887ea40.
//
// Solidity: function router() view returns(address)
func (_ClaimExecutor *ClaimExecutorCallerSession) Router() (common.Address, error) {
	return _ClaimExecutor.Contract.Router(&_ClaimExecutor.CallOpts)
}

// UsersData is a free data retrieval call binding the contract method 0x0560ab69.
//
// Solidity: function usersData(address user) view returns(uint64 latestClaimTimestamp, uint128 streakCount)
func (_ClaimExecutor *ClaimExecutorCaller) UsersData(opts *bind.CallOpts, user common.Address) (struct {
	LatestClaimTimestamp uint64
	StreakCount          *big.Int
}, error) {
	var out []interface{}
	err := _ClaimExecutor.contract.Call(opts, &out, "usersData", user)

	outstruct := new(struct {
		LatestClaimTimestamp uint64
		StreakCount          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LatestClaimTimestamp = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.StreakCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UsersData is a free data retrieval call binding the contract method 0x0560ab69.
//
// Solidity: function usersData(address user) view returns(uint64 latestClaimTimestamp, uint128 streakCount)
func (_ClaimExecutor *ClaimExecutorSession) UsersData(user common.Address) (struct {
	LatestClaimTimestamp uint64
	StreakCount          *big.Int
}, error) {
	return _ClaimExecutor.Contract.UsersData(&_ClaimExecutor.CallOpts, user)
}

// UsersData is a free data retrieval call binding the contract method 0x0560ab69.
//
// Solidity: function usersData(address user) view returns(uint64 latestClaimTimestamp, uint128 streakCount)
func (_ClaimExecutor *ClaimExecutorCallerSession) UsersData(user common.Address) (struct {
	LatestClaimTimestamp uint64
	StreakCount          *big.Int
}, error) {
	return _ClaimExecutor.Contract.UsersData(&_ClaimExecutor.CallOpts, user)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address beneficiary, bytes data) returns()
func (_ClaimExecutor *ClaimExecutorTransactor) Execute(opts *bind.TransactOpts, beneficiary common.Address, data []byte) (*types.Transaction, error) {
	return _ClaimExecutor.contract.Transact(opts, "execute", beneficiary, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address beneficiary, bytes data) returns()
func (_ClaimExecutor *ClaimExecutorSession) Execute(beneficiary common.Address, data []byte) (*types.Transaction, error) {
	return _ClaimExecutor.Contract.Execute(&_ClaimExecutor.TransactOpts, beneficiary, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address beneficiary, bytes data) returns()
func (_ClaimExecutor *ClaimExecutorTransactorSession) Execute(beneficiary common.Address, data []byte) (*types.Transaction, error) {
	return _ClaimExecutor.Contract.Execute(&_ClaimExecutor.TransactOpts, beneficiary, data)
}

// ClaimExecutorClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the ClaimExecutor contract.
type ClaimExecutorClaimedIterator struct {
	Event *ClaimExecutorClaimed // Event containing the contract specifics and raw log

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
func (it *ClaimExecutorClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimExecutorClaimed)
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
		it.Event = new(ClaimExecutorClaimed)
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
func (it *ClaimExecutorClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimExecutorClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimExecutorClaimed represents a Claimed event raised by the ClaimExecutor contract.
type ClaimExecutorClaimed struct {
	User        common.Address
	Points      *big.Int
	StreakCount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x9d79a1988c377d51e2ebd376a251db4f9ab5f1435fd3617f80f4c1291a4ede19.
//
// Solidity: event Claimed(address indexed user, uint256 points, uint128 streakCount)
func (_ClaimExecutor *ClaimExecutorFilterer) FilterClaimed(opts *bind.FilterOpts, user []common.Address) (*ClaimExecutorClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ClaimExecutor.contract.FilterLogs(opts, "Claimed", userRule)
	if err != nil {
		return nil, err
	}
	return &ClaimExecutorClaimedIterator{contract: _ClaimExecutor.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x9d79a1988c377d51e2ebd376a251db4f9ab5f1435fd3617f80f4c1291a4ede19.
//
// Solidity: event Claimed(address indexed user, uint256 points, uint128 streakCount)
func (_ClaimExecutor *ClaimExecutorFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *ClaimExecutorClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ClaimExecutor.contract.WatchLogs(opts, "Claimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimExecutorClaimed)
				if err := _ClaimExecutor.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x9d79a1988c377d51e2ebd376a251db4f9ab5f1435fd3617f80f4c1291a4ede19.
//
// Solidity: event Claimed(address indexed user, uint256 points, uint128 streakCount)
func (_ClaimExecutor *ClaimExecutorFilterer) ParseClaimed(log types.Log) (*ClaimExecutorClaimed, error) {
	event := new(ClaimExecutorClaimed)
	if err := _ClaimExecutor.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClaimExecutorExecutedIterator is returned from FilterExecuted and is used to iterate over the raw logs and unpacked data for Executed events raised by the ClaimExecutor contract.
type ClaimExecutorExecutedIterator struct {
	Event *ClaimExecutorExecuted // Event containing the contract specifics and raw log

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
func (it *ClaimExecutorExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClaimExecutorExecuted)
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
		it.Event = new(ClaimExecutorExecuted)
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
func (it *ClaimExecutorExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClaimExecutorExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClaimExecutorExecuted represents a Executed event raised by the ClaimExecutor contract.
type ClaimExecutorExecuted struct {
	Beneficiary common.Address
	Data        []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExecuted is a free log retrieval operation binding the contract event 0x0dcf4ffb3b85ab072df6f9ed79b0382ec6c9619a98f36f4538d3b2e87fd3fd11.
//
// Solidity: event Executed(address indexed beneficiary, bytes data)
func (_ClaimExecutor *ClaimExecutorFilterer) FilterExecuted(opts *bind.FilterOpts, beneficiary []common.Address) (*ClaimExecutorExecutedIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _ClaimExecutor.contract.FilterLogs(opts, "Executed", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &ClaimExecutorExecutedIterator{contract: _ClaimExecutor.contract, event: "Executed", logs: logs, sub: sub}, nil
}

// WatchExecuted is a free log subscription operation binding the contract event 0x0dcf4ffb3b85ab072df6f9ed79b0382ec6c9619a98f36f4538d3b2e87fd3fd11.
//
// Solidity: event Executed(address indexed beneficiary, bytes data)
func (_ClaimExecutor *ClaimExecutorFilterer) WatchExecuted(opts *bind.WatchOpts, sink chan<- *ClaimExecutorExecuted, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _ClaimExecutor.contract.WatchLogs(opts, "Executed", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClaimExecutorExecuted)
				if err := _ClaimExecutor.contract.UnpackLog(event, "Executed", log); err != nil {
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

// ParseExecuted is a log parse operation binding the contract event 0x0dcf4ffb3b85ab072df6f9ed79b0382ec6c9619a98f36f4538d3b2e87fd3fd11.
//
// Solidity: event Executed(address indexed beneficiary, bytes data)
func (_ClaimExecutor *ClaimExecutorFilterer) ParseExecuted(log types.Log) (*ClaimExecutorExecuted, error) {
	event := new(ClaimExecutorExecuted)
	if err := _ClaimExecutor.contract.UnpackLog(event, "Executed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
