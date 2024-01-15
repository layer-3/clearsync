// SPDX-License-Identifier: MIT
// OpenZeppelin Contracts (last updated v4.7.0) (token/ERC20/ERC20.sol)

pragma solidity 0.8.20;

import {Context} from '@openzeppelin/contracts/utils/Context.sol';

/**
 * @dev Copy-pasted from Openzeppelin ERC20 contract, but with the inheritance from IERC20 interface removed and no return value for transferFrom.
 */
contract BadToken is Context {
	function transferFrom(address from, address to, uint256 value) public virtual {
		address spender = _msgSender();
		_spendAllowance(from, spender, value);
		_transfer(from, to, value);
		// return true; // purposefully omitted because bad token
	}

	constructor(address owner) {
		_name = 'BadToken';
		_symbol = 'BTKN';
		_mint(owner, 10_000_000_000);
	}

	// The rest of the file is a merge of OpenZeppelin files, specifically events from IERC20.sol,
	// errors from IERC20Errors and all code from ERC20.sol with all comments removed.
	event Transfer(address indexed from, address indexed to, uint256 value);

	event Approval(address indexed owner, address indexed spender, uint256 value);

	error ERC20InsufficientBalance(address sender, uint256 balance, uint256 needed);

	error ERC20InvalidSender(address sender);

	error ERC20InvalidReceiver(address receiver);

	error ERC20InsufficientAllowance(address spender, uint256 allowance, uint256 needed);

	error ERC20InvalidApprover(address approver);

	error ERC20InvalidSpender(address spender);

	mapping(address account => uint256) private _balances;

	mapping(address account => mapping(address spender => uint256)) private _allowances;

	uint256 private _totalSupply;

	string private _name;
	string private _symbol;

	function name() public view virtual returns (string memory) {
		return _name;
	}

	function symbol() public view virtual returns (string memory) {
		return _symbol;
	}

	function decimals() public view virtual returns (uint8) {
		return 18;
	}

	function totalSupply() public view virtual returns (uint256) {
		return _totalSupply;
	}

	function balanceOf(address account) public view virtual returns (uint256) {
		return _balances[account];
	}

	function transfer(address to, uint256 value) public virtual returns (bool) {
		address owner = _msgSender();
		_transfer(owner, to, value);
		return true;
	}

	function allowance(address owner, address spender) public view virtual returns (uint256) {
		return _allowances[owner][spender];
	}

	function approve(address spender, uint256 value) public virtual returns (bool) {
		address owner = _msgSender();
		_approve(owner, spender, value);
		return true;
	}

	function _transfer(address from, address to, uint256 value) internal {
		if (from == address(0)) {
			revert ERC20InvalidSender(address(0));
		}
		if (to == address(0)) {
			revert ERC20InvalidReceiver(address(0));
		}
		_update(from, to, value);
	}

	function _update(address from, address to, uint256 value) internal virtual {
		if (from == address(0)) {
			_totalSupply += value;
		} else {
			uint256 fromBalance = _balances[from];
			if (fromBalance < value) {
				revert ERC20InsufficientBalance(from, fromBalance, value);
			}
			unchecked {
				_balances[from] = fromBalance - value;
			}
		}

		if (to == address(0)) {
			unchecked {
				_totalSupply -= value;
			}
		} else {
			unchecked {
				_balances[to] += value;
			}
		}

		emit Transfer(from, to, value);
	}

	function _mint(address account, uint256 value) internal {
		if (account == address(0)) {
			revert ERC20InvalidReceiver(address(0));
		}
		_update(address(0), account, value);
	}

	function _burn(address account, uint256 value) internal {
		if (account == address(0)) {
			revert ERC20InvalidSender(address(0));
		}
		_update(account, address(0), value);
	}

	function _approve(address owner, address spender, uint256 value) internal {
		_approve(owner, spender, value, true);
	}

	function _approve(
		address owner,
		address spender,
		uint256 value,
		bool emitEvent
	) internal virtual {
		if (owner == address(0)) {
			revert ERC20InvalidApprover(address(0));
		}
		if (spender == address(0)) {
			revert ERC20InvalidSpender(address(0));
		}
		_allowances[owner][spender] = value;
		if (emitEvent) {
			emit Approval(owner, spender, value);
		}
	}

	function _spendAllowance(address owner, address spender, uint256 value) internal virtual {
		uint256 currentAllowance = allowance(owner, spender);
		if (currentAllowance != type(uint256).max) {
			if (currentAllowance < value) {
				revert ERC20InsufficientAllowance(spender, currentAllowance, value);
			}
			unchecked {
				_approve(owner, spender, currentAllowance - value, false);
			}
		}
	}
}
