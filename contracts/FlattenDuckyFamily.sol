// Sources flattened with hardhat v2.14.0 https://hardhat.org

// File @openzeppelin/contracts-upgradeable/utils/introspection/IERC165Upgradeable.sol@v4.8.3

// OpenZeppelin Contracts v4.4.1 (utils/introspection/IERC165.sol)

pragma solidity ^0.8.0;

/**
 * @dev Interface of the ERC165 standard, as defined in the
 * https://eips.ethereum.org/EIPS/eip-165[EIP].
 *
 * Implementers can declare support of contract interfaces, which can then be
 * queried by others ({ERC165Checker}).
 *
 * For an implementation, see {ERC165}.
 */
interface IERC165Upgradeable {
    /**
     * @dev Returns true if this contract implements the interface defined by
     * `interfaceId`. See the corresponding
     * https://eips.ethereum.org/EIPS/eip-165#how-interfaces-are-identified[EIP section]
     * to learn more about how these ids are created.
     *
     * This function call must use less than 30 000 gas.
     */
    function supportsInterface(bytes4 interfaceId) external view returns (bool);
}


// File @openzeppelin/contracts-upgradeable/token/ERC721/IERC721Upgradeable.sol@v4.8.3

// OpenZeppelin Contracts (last updated v4.8.0) (token/ERC721/IERC721.sol)

pragma solidity ^0.8.0;

/**
 * @dev Required interface of an ERC721 compliant contract.
 */
interface IERC721Upgradeable is IERC165Upgradeable {
    /**
     * @dev Emitted when `tokenId` token is transferred from `from` to `to`.
     */
    event Transfer(address indexed from, address indexed to, uint256 indexed tokenId);

    /**
     * @dev Emitted when `owner` enables `approved` to manage the `tokenId` token.
     */
    event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId);

    /**
     * @dev Emitted when `owner` enables or disables (`approved`) `operator` to manage all of its assets.
     */
    event ApprovalForAll(address indexed owner, address indexed operator, bool approved);

    /**
     * @dev Returns the number of tokens in ``owner``'s account.
     */
    function balanceOf(address owner) external view returns (uint256 balance);

    /**
     * @dev Returns the owner of the `tokenId` token.
     *
     * Requirements:
     *
     * - `tokenId` must exist.
     */
    function ownerOf(uint256 tokenId) external view returns (address owner);

    /**
     * @dev Safely transfers `tokenId` token from `from` to `to`.
     *
     * Requirements:
     *
     * - `from` cannot be the zero address.
     * - `to` cannot be the zero address.
     * - `tokenId` token must exist and be owned by `from`.
     * - If the caller is not `from`, it must be approved to move this token by either {approve} or {setApprovalForAll}.
     * - If `to` refers to a smart contract, it must implement {IERC721Receiver-onERC721Received}, which is called upon a safe transfer.
     *
     * Emits a {Transfer} event.
     */
    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId,
        bytes calldata data
    ) external;

    /**
     * @dev Safely transfers `tokenId` token from `from` to `to`, checking first that contract recipients
     * are aware of the ERC721 protocol to prevent tokens from being forever locked.
     *
     * Requirements:
     *
     * - `from` cannot be the zero address.
     * - `to` cannot be the zero address.
     * - `tokenId` token must exist and be owned by `from`.
     * - If the caller is not `from`, it must have been allowed to move this token by either {approve} or {setApprovalForAll}.
     * - If `to` refers to a smart contract, it must implement {IERC721Receiver-onERC721Received}, which is called upon a safe transfer.
     *
     * Emits a {Transfer} event.
     */
    function safeTransferFrom(
        address from,
        address to,
        uint256 tokenId
    ) external;

    /**
     * @dev Transfers `tokenId` token from `from` to `to`.
     *
     * WARNING: Note that the caller is responsible to confirm that the recipient is capable of receiving ERC721
     * or else they may be permanently lost. Usage of {safeTransferFrom} prevents loss, though the caller must
     * understand this adds an external call which potentially creates a reentrancy vulnerability.
     *
     * Requirements:
     *
     * - `from` cannot be the zero address.
     * - `to` cannot be the zero address.
     * - `tokenId` token must be owned by `from`.
     * - If the caller is not `from`, it must be approved to move this token by either {approve} or {setApprovalForAll}.
     *
     * Emits a {Transfer} event.
     */
    function transferFrom(
        address from,
        address to,
        uint256 tokenId
    ) external;

    /**
     * @dev Gives permission to `to` to transfer `tokenId` token to another account.
     * The approval is cleared when the token is transferred.
     *
     * Only a single account can be approved at a time, so approving the zero address clears previous approvals.
     *
     * Requirements:
     *
     * - The caller must own the token or be an approved operator.
     * - `tokenId` must exist.
     *
     * Emits an {Approval} event.
     */
    function approve(address to, uint256 tokenId) external;

    /**
     * @dev Approve or remove `operator` as an operator for the caller.
     * Operators can call {transferFrom} or {safeTransferFrom} for any token owned by the caller.
     *
     * Requirements:
     *
     * - The `operator` cannot be the caller.
     *
     * Emits an {ApprovalForAll} event.
     */
    function setApprovalForAll(address operator, bool _approved) external;

    /**
     * @dev Returns the account approved for `tokenId` token.
     *
     * Requirements:
     *
     * - `tokenId` must exist.
     */
    function getApproved(uint256 tokenId) external view returns (address operator);

    /**
     * @dev Returns if the `operator` is allowed to manage all of the assets of `owner`.
     *
     * See {setApprovalForAll}
     */
    function isApprovedForAll(address owner, address operator) external view returns (bool);
}


// File @openzeppelin/contracts/access/IAccessControl.sol@v4.8.3

// OpenZeppelin Contracts v4.4.1 (access/IAccessControl.sol)

pragma solidity ^0.8.0;

/**
 * @dev External interface of AccessControl declared to support ERC165 detection.
 */
interface IAccessControl {
    /**
     * @dev Emitted when `newAdminRole` is set as ``role``'s admin role, replacing `previousAdminRole`
     *
     * `DEFAULT_ADMIN_ROLE` is the starting admin for all roles, despite
     * {RoleAdminChanged} not being emitted signaling this.
     *
     * _Available since v3.1._
     */
    event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole);

    /**
     * @dev Emitted when `account` is granted `role`.
     *
     * `sender` is the account that originated the contract call, an admin role
     * bearer except when using {AccessControl-_setupRole}.
     */
    event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender);

    /**
     * @dev Emitted when `account` is revoked `role`.
     *
     * `sender` is the account that originated the contract call:
     *   - if using `revokeRole`, it is the admin role bearer
     *   - if using `renounceRole`, it is the role bearer (i.e. `account`)
     */
    event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender);

    /**
     * @dev Returns `true` if `account` has been granted `role`.
     */
    function hasRole(bytes32 role, address account) external view returns (bool);

    /**
     * @dev Returns the admin role that controls `role`. See {grantRole} and
     * {revokeRole}.
     *
     * To change a role's admin, use {AccessControl-_setRoleAdmin}.
     */
    function getRoleAdmin(bytes32 role) external view returns (bytes32);

    /**
     * @dev Grants `role` to `account`.
     *
     * If `account` had not been already granted `role`, emits a {RoleGranted}
     * event.
     *
     * Requirements:
     *
     * - the caller must have ``role``'s admin role.
     */
    function grantRole(bytes32 role, address account) external;

    /**
     * @dev Revokes `role` from `account`.
     *
     * If `account` had been granted `role`, emits a {RoleRevoked} event.
     *
     * Requirements:
     *
     * - the caller must have ``role``'s admin role.
     */
    function revokeRole(bytes32 role, address account) external;

    /**
     * @dev Revokes `role` from the calling account.
     *
     * Roles are often managed via {grantRole} and {revokeRole}: this function's
     * purpose is to provide a mechanism for accounts to lose their privileges
     * if they are compromised (such as when a trusted device is misplaced).
     *
     * If the calling account had been granted `role`, emits a {RoleRevoked}
     * event.
     *
     * Requirements:
     *
     * - the caller must be `account`.
     */
    function renounceRole(bytes32 role, address account) external;
}


// File @openzeppelin/contracts/utils/Context.sol@v4.8.3

// OpenZeppelin Contracts v4.4.1 (utils/Context.sol)

pragma solidity ^0.8.0;

/**
 * @dev Provides information about the current execution context, including the
 * sender of the transaction and its data. While these are generally available
 * via msg.sender and msg.data, they should not be accessed in such a direct
 * manner, since when dealing with meta-transactions the account sending and
 * paying for execution may not be the actual sender (as far as an application
 * is concerned).
 *
 * This contract is only required for intermediate, library-like contracts.
 */
abstract contract Context {
    function _msgSender() internal view virtual returns (address) {
        return msg.sender;
    }

    function _msgData() internal view virtual returns (bytes calldata) {
        return msg.data;
    }
}


// File @openzeppelin/contracts/utils/introspection/IERC165.sol@v4.8.3

// OpenZeppelin Contracts v4.4.1 (utils/introspection/IERC165.sol)

pragma solidity ^0.8.0;

/**
 * @dev Interface of the ERC165 standard, as defined in the
 * https://eips.ethereum.org/EIPS/eip-165[EIP].
 *
 * Implementers can declare support of contract interfaces, which can then be
 * queried by others ({ERC165Checker}).
 *
 * For an implementation, see {ERC165}.
 */
interface IERC165 {
    /**
     * @dev Returns true if this contract implements the interface defined by
     * `interfaceId`. See the corresponding
     * https://eips.ethereum.org/EIPS/eip-165#how-interfaces-are-identified[EIP section]
     * to learn more about how these ids are created.
     *
     * This function call must use less than 30 000 gas.
     */
    function supportsInterface(bytes4 interfaceId) external view returns (bool);
}


// File @openzeppelin/contracts/utils/introspection/ERC165.sol@v4.8.3

// OpenZeppelin Contracts v4.4.1 (utils/introspection/ERC165.sol)

pragma solidity ^0.8.0;

/**
 * @dev Implementation of the {IERC165} interface.
 *
 * Contracts that want to implement ERC165 should inherit from this contract and override {supportsInterface} to check
 * for the additional interface id that will be supported. For example:
 *
 * ```solidity
 * function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
 *     return interfaceId == type(MyInterface).interfaceId || super.supportsInterface(interfaceId);
 * }
 * ```
 *
 * Alternatively, {ERC165Storage} provides an easier to use but more expensive implementation.
 */
abstract contract ERC165 is IERC165 {
    /**
     * @dev See {IERC165-supportsInterface}.
     */
    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
        return interfaceId == type(IERC165).interfaceId;
    }
}


// File @openzeppelin/contracts/utils/math/Math.sol@v4.8.3

// OpenZeppelin Contracts (last updated v4.8.0) (utils/math/Math.sol)

pragma solidity ^0.8.0;

/**
 * @dev Standard math utilities missing in the Solidity language.
 */
library Math {
    enum Rounding {
        Down, // Toward negative infinity
        Up, // Toward infinity
        Zero // Toward zero
    }

    /**
     * @dev Returns the largest of two numbers.
     */
    function max(uint256 a, uint256 b) internal pure returns (uint256) {
        return a > b ? a : b;
    }

    /**
     * @dev Returns the smallest of two numbers.
     */
    function min(uint256 a, uint256 b) internal pure returns (uint256) {
        return a < b ? a : b;
    }

    /**
     * @dev Returns the average of two numbers. The result is rounded towards
     * zero.
     */
    function average(uint256 a, uint256 b) internal pure returns (uint256) {
        // (a + b) / 2 can overflow.
        return (a & b) + (a ^ b) / 2;
    }

    /**
     * @dev Returns the ceiling of the division of two numbers.
     *
     * This differs from standard division with `/` in that it rounds up instead
     * of rounding down.
     */
    function ceilDiv(uint256 a, uint256 b) internal pure returns (uint256) {
        // (a + b - 1) / b can overflow on addition, so we distribute.
        return a == 0 ? 0 : (a - 1) / b + 1;
    }

    /**
     * @notice Calculates floor(x * y / denominator) with full precision. Throws if result overflows a uint256 or denominator == 0
     * @dev Original credit to Remco Bloemen under MIT license (https://xn--2-umb.com/21/muldiv)
     * with further edits by Uniswap Labs also under MIT license.
     */
    function mulDiv(
        uint256 x,
        uint256 y,
        uint256 denominator
    ) internal pure returns (uint256 result) {
        unchecked {
            // 512-bit multiply [prod1 prod0] = x * y. Compute the product mod 2^256 and mod 2^256 - 1, then use
            // use the Chinese Remainder Theorem to reconstruct the 512 bit result. The result is stored in two 256
            // variables such that product = prod1 * 2^256 + prod0.
            uint256 prod0; // Least significant 256 bits of the product
            uint256 prod1; // Most significant 256 bits of the product
            assembly {
                let mm := mulmod(x, y, not(0))
                prod0 := mul(x, y)
                prod1 := sub(sub(mm, prod0), lt(mm, prod0))
            }

            // Handle non-overflow cases, 256 by 256 division.
            if (prod1 == 0) {
                return prod0 / denominator;
            }

            // Make sure the result is less than 2^256. Also prevents denominator == 0.
            require(denominator > prod1);

            ///////////////////////////////////////////////
            // 512 by 256 division.
            ///////////////////////////////////////////////

            // Make division exact by subtracting the remainder from [prod1 prod0].
            uint256 remainder;
            assembly {
                // Compute remainder using mulmod.
                remainder := mulmod(x, y, denominator)

                // Subtract 256 bit number from 512 bit number.
                prod1 := sub(prod1, gt(remainder, prod0))
                prod0 := sub(prod0, remainder)
            }

            // Factor powers of two out of denominator and compute largest power of two divisor of denominator. Always >= 1.
            // See https://cs.stackexchange.com/q/138556/92363.

            // Does not overflow because the denominator cannot be zero at this stage in the function.
            uint256 twos = denominator & (~denominator + 1);
            assembly {
                // Divide denominator by twos.
                denominator := div(denominator, twos)

                // Divide [prod1 prod0] by twos.
                prod0 := div(prod0, twos)

                // Flip twos such that it is 2^256 / twos. If twos is zero, then it becomes one.
                twos := add(div(sub(0, twos), twos), 1)
            }

            // Shift in bits from prod1 into prod0.
            prod0 |= prod1 * twos;

            // Invert denominator mod 2^256. Now that denominator is an odd number, it has an inverse modulo 2^256 such
            // that denominator * inv = 1 mod 2^256. Compute the inverse by starting with a seed that is correct for
            // four bits. That is, denominator * inv = 1 mod 2^4.
            uint256 inverse = (3 * denominator) ^ 2;

            // Use the Newton-Raphson iteration to improve the precision. Thanks to Hensel's lifting lemma, this also works
            // in modular arithmetic, doubling the correct bits in each step.
            inverse *= 2 - denominator * inverse; // inverse mod 2^8
            inverse *= 2 - denominator * inverse; // inverse mod 2^16
            inverse *= 2 - denominator * inverse; // inverse mod 2^32
            inverse *= 2 - denominator * inverse; // inverse mod 2^64
            inverse *= 2 - denominator * inverse; // inverse mod 2^128
            inverse *= 2 - denominator * inverse; // inverse mod 2^256

            // Because the division is now exact we can divide by multiplying with the modular inverse of denominator.
            // This will give us the correct result modulo 2^256. Since the preconditions guarantee that the outcome is
            // less than 2^256, this is the final result. We don't need to compute the high bits of the result and prod1
            // is no longer required.
            result = prod0 * inverse;
            return result;
        }
    }

    /**
     * @notice Calculates x * y / denominator with full precision, following the selected rounding direction.
     */
    function mulDiv(
        uint256 x,
        uint256 y,
        uint256 denominator,
        Rounding rounding
    ) internal pure returns (uint256) {
        uint256 result = mulDiv(x, y, denominator);
        if (rounding == Rounding.Up && mulmod(x, y, denominator) > 0) {
            result += 1;
        }
        return result;
    }

    /**
     * @dev Returns the square root of a number. If the number is not a perfect square, the value is rounded down.
     *
     * Inspired by Henry S. Warren, Jr.'s "Hacker's Delight" (Chapter 11).
     */
    function sqrt(uint256 a) internal pure returns (uint256) {
        if (a == 0) {
            return 0;
        }

        // For our first guess, we get the biggest power of 2 which is smaller than the square root of the target.
        //
        // We know that the "msb" (most significant bit) of our target number `a` is a power of 2 such that we have
        // `msb(a) <= a < 2*msb(a)`. This value can be written `msb(a)=2**k` with `k=log2(a)`.
        //
        // This can be rewritten `2**log2(a) <= a < 2**(log2(a) + 1)`
        // → `sqrt(2**k) <= sqrt(a) < sqrt(2**(k+1))`
        // → `2**(k/2) <= sqrt(a) < 2**((k+1)/2) <= 2**(k/2 + 1)`
        //
        // Consequently, `2**(log2(a) / 2)` is a good first approximation of `sqrt(a)` with at least 1 correct bit.
        uint256 result = 1 << (log2(a) >> 1);

        // At this point `result` is an estimation with one bit of precision. We know the true value is a uint128,
        // since it is the square root of a uint256. Newton's method converges quadratically (precision doubles at
        // every iteration). We thus need at most 7 iteration to turn our partial result with one bit of precision
        // into the expected uint128 result.
        unchecked {
            result = (result + a / result) >> 1;
            result = (result + a / result) >> 1;
            result = (result + a / result) >> 1;
            result = (result + a / result) >> 1;
            result = (result + a / result) >> 1;
            result = (result + a / result) >> 1;
            result = (result + a / result) >> 1;
            return min(result, a / result);
        }
    }

    /**
     * @notice Calculates sqrt(a), following the selected rounding direction.
     */
    function sqrt(uint256 a, Rounding rounding) internal pure returns (uint256) {
        unchecked {
            uint256 result = sqrt(a);
            return result + (rounding == Rounding.Up && result * result < a ? 1 : 0);
        }
    }

    /**
     * @dev Return the log in base 2, rounded down, of a positive value.
     * Returns 0 if given 0.
     */
    function log2(uint256 value) internal pure returns (uint256) {
        uint256 result = 0;
        unchecked {
            if (value >> 128 > 0) {
                value >>= 128;
                result += 128;
            }
            if (value >> 64 > 0) {
                value >>= 64;
                result += 64;
            }
            if (value >> 32 > 0) {
                value >>= 32;
                result += 32;
            }
            if (value >> 16 > 0) {
                value >>= 16;
                result += 16;
            }
            if (value >> 8 > 0) {
                value >>= 8;
                result += 8;
            }
            if (value >> 4 > 0) {
                value >>= 4;
                result += 4;
            }
            if (value >> 2 > 0) {
                value >>= 2;
                result += 2;
            }
            if (value >> 1 > 0) {
                result += 1;
            }
        }
        return result;
    }

    /**
     * @dev Return the log in base 2, following the selected rounding direction, of a positive value.
     * Returns 0 if given 0.
     */
    function log2(uint256 value, Rounding rounding) internal pure returns (uint256) {
        unchecked {
            uint256 result = log2(value);
            return result + (rounding == Rounding.Up && 1 << result < value ? 1 : 0);
        }
    }

    /**
     * @dev Return the log in base 10, rounded down, of a positive value.
     * Returns 0 if given 0.
     */
    function log10(uint256 value) internal pure returns (uint256) {
        uint256 result = 0;
        unchecked {
            if (value >= 10**64) {
                value /= 10**64;
                result += 64;
            }
            if (value >= 10**32) {
                value /= 10**32;
                result += 32;
            }
            if (value >= 10**16) {
                value /= 10**16;
                result += 16;
            }
            if (value >= 10**8) {
                value /= 10**8;
                result += 8;
            }
            if (value >= 10**4) {
                value /= 10**4;
                result += 4;
            }
            if (value >= 10**2) {
                value /= 10**2;
                result += 2;
            }
            if (value >= 10**1) {
                result += 1;
            }
        }
        return result;
    }

    /**
     * @dev Return the log in base 10, following the selected rounding direction, of a positive value.
     * Returns 0 if given 0.
     */
    function log10(uint256 value, Rounding rounding) internal pure returns (uint256) {
        unchecked {
            uint256 result = log10(value);
            return result + (rounding == Rounding.Up && 10**result < value ? 1 : 0);
        }
    }

    /**
     * @dev Return the log in base 256, rounded down, of a positive value.
     * Returns 0 if given 0.
     *
     * Adding one to the result gives the number of pairs of hex symbols needed to represent `value` as a hex string.
     */
    function log256(uint256 value) internal pure returns (uint256) {
        uint256 result = 0;
        unchecked {
            if (value >> 128 > 0) {
                value >>= 128;
                result += 16;
            }
            if (value >> 64 > 0) {
                value >>= 64;
                result += 8;
            }
            if (value >> 32 > 0) {
                value >>= 32;
                result += 4;
            }
            if (value >> 16 > 0) {
                value >>= 16;
                result += 2;
            }
            if (value >> 8 > 0) {
                result += 1;
            }
        }
        return result;
    }

    /**
     * @dev Return the log in base 10, following the selected rounding direction, of a positive value.
     * Returns 0 if given 0.
     */
    function log256(uint256 value, Rounding rounding) internal pure returns (uint256) {
        unchecked {
            uint256 result = log256(value);
            return result + (rounding == Rounding.Up && 1 << (result * 8) < value ? 1 : 0);
        }
    }
}


// File @openzeppelin/contracts/utils/Strings.sol@v4.8.3

// OpenZeppelin Contracts (last updated v4.8.0) (utils/Strings.sol)

pragma solidity ^0.8.0;

/**
 * @dev String operations.
 */
library Strings {
    bytes16 private constant _SYMBOLS = "0123456789abcdef";
    uint8 private constant _ADDRESS_LENGTH = 20;

    /**
     * @dev Converts a `uint256` to its ASCII `string` decimal representation.
     */
    function toString(uint256 value) internal pure returns (string memory) {
        unchecked {
            uint256 length = Math.log10(value) + 1;
            string memory buffer = new string(length);
            uint256 ptr;
            /// @solidity memory-safe-assembly
            assembly {
                ptr := add(buffer, add(32, length))
            }
            while (true) {
                ptr--;
                /// @solidity memory-safe-assembly
                assembly {
                    mstore8(ptr, byte(mod(value, 10), _SYMBOLS))
                }
                value /= 10;
                if (value == 0) break;
            }
            return buffer;
        }
    }

    /**
     * @dev Converts a `uint256` to its ASCII `string` hexadecimal representation.
     */
    function toHexString(uint256 value) internal pure returns (string memory) {
        unchecked {
            return toHexString(value, Math.log256(value) + 1);
        }
    }

    /**
     * @dev Converts a `uint256` to its ASCII `string` hexadecimal representation with fixed length.
     */
    function toHexString(uint256 value, uint256 length) internal pure returns (string memory) {
        bytes memory buffer = new bytes(2 * length + 2);
        buffer[0] = "0";
        buffer[1] = "x";
        for (uint256 i = 2 * length + 1; i > 1; --i) {
            buffer[i] = _SYMBOLS[value & 0xf];
            value >>= 4;
        }
        require(value == 0, "Strings: hex length insufficient");
        return string(buffer);
    }

    /**
     * @dev Converts an `address` with fixed length of 20 bytes to its not checksummed ASCII `string` hexadecimal representation.
     */
    function toHexString(address addr) internal pure returns (string memory) {
        return toHexString(uint256(uint160(addr)), _ADDRESS_LENGTH);
    }
}


// File @openzeppelin/contracts/access/AccessControl.sol@v4.8.3

// OpenZeppelin Contracts (last updated v4.8.0) (access/AccessControl.sol)

pragma solidity ^0.8.0;




/**
 * @dev Contract module that allows children to implement role-based access
 * control mechanisms. This is a lightweight version that doesn't allow enumerating role
 * members except through off-chain means by accessing the contract event logs. Some
 * applications may benefit from on-chain enumerability, for those cases see
 * {AccessControlEnumerable}.
 *
 * Roles are referred to by their `bytes32` identifier. These should be exposed
 * in the external API and be unique. The best way to achieve this is by
 * using `public constant` hash digests:
 *
 * ```
 * bytes32 public constant MY_ROLE = keccak256("MY_ROLE");
 * ```
 *
 * Roles can be used to represent a set of permissions. To restrict access to a
 * function call, use {hasRole}:
 *
 * ```
 * function foo() public {
 *     require(hasRole(MY_ROLE, msg.sender));
 *     ...
 * }
 * ```
 *
 * Roles can be granted and revoked dynamically via the {grantRole} and
 * {revokeRole} functions. Each role has an associated admin role, and only
 * accounts that have a role's admin role can call {grantRole} and {revokeRole}.
 *
 * By default, the admin role for all roles is `DEFAULT_ADMIN_ROLE`, which means
 * that only accounts with this role will be able to grant or revoke other
 * roles. More complex role relationships can be created by using
 * {_setRoleAdmin}.
 *
 * WARNING: The `DEFAULT_ADMIN_ROLE` is also its own admin: it has permission to
 * grant and revoke this role. Extra precautions should be taken to secure
 * accounts that have been granted it.
 */
abstract contract AccessControl is Context, IAccessControl, ERC165 {
    struct RoleData {
        mapping(address => bool) members;
        bytes32 adminRole;
    }

    mapping(bytes32 => RoleData) private _roles;

    bytes32 public constant DEFAULT_ADMIN_ROLE = 0x00;

    /**
     * @dev Modifier that checks that an account has a specific role. Reverts
     * with a standardized message including the required role.
     *
     * The format of the revert reason is given by the following regular expression:
     *
     *  /^AccessControl: account (0x[0-9a-f]{40}) is missing role (0x[0-9a-f]{64})$/
     *
     * _Available since v4.1._
     */
    modifier onlyRole(bytes32 role) {
        _checkRole(role);
        _;
    }

    /**
     * @dev See {IERC165-supportsInterface}.
     */
    function supportsInterface(bytes4 interfaceId) public view virtual override returns (bool) {
        return interfaceId == type(IAccessControl).interfaceId || super.supportsInterface(interfaceId);
    }

    /**
     * @dev Returns `true` if `account` has been granted `role`.
     */
    function hasRole(bytes32 role, address account) public view virtual override returns (bool) {
        return _roles[role].members[account];
    }

    /**
     * @dev Revert with a standard message if `_msgSender()` is missing `role`.
     * Overriding this function changes the behavior of the {onlyRole} modifier.
     *
     * Format of the revert message is described in {_checkRole}.
     *
     * _Available since v4.6._
     */
    function _checkRole(bytes32 role) internal view virtual {
        _checkRole(role, _msgSender());
    }

    /**
     * @dev Revert with a standard message if `account` is missing `role`.
     *
     * The format of the revert reason is given by the following regular expression:
     *
     *  /^AccessControl: account (0x[0-9a-f]{40}) is missing role (0x[0-9a-f]{64})$/
     */
    function _checkRole(bytes32 role, address account) internal view virtual {
        if (!hasRole(role, account)) {
            revert(
                string(
                    abi.encodePacked(
                        "AccessControl: account ",
                        Strings.toHexString(account),
                        " is missing role ",
                        Strings.toHexString(uint256(role), 32)
                    )
                )
            );
        }
    }

    /**
     * @dev Returns the admin role that controls `role`. See {grantRole} and
     * {revokeRole}.
     *
     * To change a role's admin, use {_setRoleAdmin}.
     */
    function getRoleAdmin(bytes32 role) public view virtual override returns (bytes32) {
        return _roles[role].adminRole;
    }

    /**
     * @dev Grants `role` to `account`.
     *
     * If `account` had not been already granted `role`, emits a {RoleGranted}
     * event.
     *
     * Requirements:
     *
     * - the caller must have ``role``'s admin role.
     *
     * May emit a {RoleGranted} event.
     */
    function grantRole(bytes32 role, address account) public virtual override onlyRole(getRoleAdmin(role)) {
        _grantRole(role, account);
    }

    /**
     * @dev Revokes `role` from `account`.
     *
     * If `account` had been granted `role`, emits a {RoleRevoked} event.
     *
     * Requirements:
     *
     * - the caller must have ``role``'s admin role.
     *
     * May emit a {RoleRevoked} event.
     */
    function revokeRole(bytes32 role, address account) public virtual override onlyRole(getRoleAdmin(role)) {
        _revokeRole(role, account);
    }

    /**
     * @dev Revokes `role` from the calling account.
     *
     * Roles are often managed via {grantRole} and {revokeRole}: this function's
     * purpose is to provide a mechanism for accounts to lose their privileges
     * if they are compromised (such as when a trusted device is misplaced).
     *
     * If the calling account had been revoked `role`, emits a {RoleRevoked}
     * event.
     *
     * Requirements:
     *
     * - the caller must be `account`.
     *
     * May emit a {RoleRevoked} event.
     */
    function renounceRole(bytes32 role, address account) public virtual override {
        require(account == _msgSender(), "AccessControl: can only renounce roles for self");

        _revokeRole(role, account);
    }

    /**
     * @dev Grants `role` to `account`.
     *
     * If `account` had not been already granted `role`, emits a {RoleGranted}
     * event. Note that unlike {grantRole}, this function doesn't perform any
     * checks on the calling account.
     *
     * May emit a {RoleGranted} event.
     *
     * [WARNING]
     * ====
     * This function should only be called from the constructor when setting
     * up the initial roles for the system.
     *
     * Using this function in any other way is effectively circumventing the admin
     * system imposed by {AccessControl}.
     * ====
     *
     * NOTE: This function is deprecated in favor of {_grantRole}.
     */
    function _setupRole(bytes32 role, address account) internal virtual {
        _grantRole(role, account);
    }

    /**
     * @dev Sets `adminRole` as ``role``'s admin role.
     *
     * Emits a {RoleAdminChanged} event.
     */
    function _setRoleAdmin(bytes32 role, bytes32 adminRole) internal virtual {
        bytes32 previousAdminRole = getRoleAdmin(role);
        _roles[role].adminRole = adminRole;
        emit RoleAdminChanged(role, previousAdminRole, adminRole);
    }

    /**
     * @dev Grants `role` to `account`.
     *
     * Internal function without access restriction.
     *
     * May emit a {RoleGranted} event.
     */
    function _grantRole(bytes32 role, address account) internal virtual {
        if (!hasRole(role, account)) {
            _roles[role].members[account] = true;
            emit RoleGranted(role, account, _msgSender());
        }
    }

    /**
     * @dev Revokes `role` from `account`.
     *
     * Internal function without access restriction.
     *
     * May emit a {RoleRevoked} event.
     */
    function _revokeRole(bytes32 role, address account) internal virtual {
        if (hasRole(role, account)) {
            _roles[role].members[account] = false;
            emit RoleRevoked(role, account, _msgSender());
        }
    }
}


// File @openzeppelin/contracts/token/ERC20/IERC20.sol@v4.8.3

// OpenZeppelin Contracts (last updated v4.6.0) (token/ERC20/IERC20.sol)

pragma solidity ^0.8.0;

/**
 * @dev Interface of the ERC20 standard as defined in the EIP.
 */
interface IERC20 {
    /**
     * @dev Emitted when `value` tokens are moved from one account (`from`) to
     * another (`to`).
     *
     * Note that `value` may be zero.
     */
    event Transfer(address indexed from, address indexed to, uint256 value);

    /**
     * @dev Emitted when the allowance of a `spender` for an `owner` is set by
     * a call to {approve}. `value` is the new allowance.
     */
    event Approval(address indexed owner, address indexed spender, uint256 value);

    /**
     * @dev Returns the amount of tokens in existence.
     */
    function totalSupply() external view returns (uint256);

    /**
     * @dev Returns the amount of tokens owned by `account`.
     */
    function balanceOf(address account) external view returns (uint256);

    /**
     * @dev Moves `amount` tokens from the caller's account to `to`.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transfer(address to, uint256 amount) external returns (bool);

    /**
     * @dev Returns the remaining number of tokens that `spender` will be
     * allowed to spend on behalf of `owner` through {transferFrom}. This is
     * zero by default.
     *
     * This value changes when {approve} or {transferFrom} are called.
     */
    function allowance(address owner, address spender) external view returns (uint256);

    /**
     * @dev Sets `amount` as the allowance of `spender` over the caller's tokens.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * IMPORTANT: Beware that changing an allowance with this method brings the risk
     * that someone may use both the old and the new allowance by unfortunate
     * transaction ordering. One possible solution to mitigate this race
     * condition is to first reduce the spender's allowance to 0 and set the
     * desired value afterwards:
     * https://github.com/ethereum/EIPs/issues/20#issuecomment-263524729
     *
     * Emits an {Approval} event.
     */
    function approve(address spender, uint256 amount) external returns (bool);

    /**
     * @dev Moves `amount` tokens from `from` to `to` using the
     * allowance mechanism. `amount` is then deducted from the caller's
     * allowance.
     *
     * Returns a boolean value indicating whether the operation succeeded.
     *
     * Emits a {Transfer} event.
     */
    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) external returns (bool);
}


// File @openzeppelin/contracts/token/ERC20/extensions/IERC20Metadata.sol@v4.8.3

// OpenZeppelin Contracts v4.4.1 (token/ERC20/extensions/IERC20Metadata.sol)

pragma solidity ^0.8.0;

/**
 * @dev Interface for the optional metadata functions from the ERC20 standard.
 *
 * _Available since v4.1._
 */
interface IERC20Metadata is IERC20 {
    /**
     * @dev Returns the name of the token.
     */
    function name() external view returns (string memory);

    /**
     * @dev Returns the symbol of the token.
     */
    function symbol() external view returns (string memory);

    /**
     * @dev Returns the decimals places of the token.
     */
    function decimals() external view returns (uint8);
}


// File @openzeppelin/contracts/token/ERC20/ERC20.sol@v4.8.3

// OpenZeppelin Contracts (last updated v4.8.0) (token/ERC20/ERC20.sol)

pragma solidity ^0.8.0;



/**
 * @dev Implementation of the {IERC20} interface.
 *
 * This implementation is agnostic to the way tokens are created. This means
 * that a supply mechanism has to be added in a derived contract using {_mint}.
 * For a generic mechanism see {ERC20PresetMinterPauser}.
 *
 * TIP: For a detailed writeup see our guide
 * https://forum.openzeppelin.com/t/how-to-implement-erc20-supply-mechanisms/226[How
 * to implement supply mechanisms].
 *
 * We have followed general OpenZeppelin Contracts guidelines: functions revert
 * instead returning `false` on failure. This behavior is nonetheless
 * conventional and does not conflict with the expectations of ERC20
 * applications.
 *
 * Additionally, an {Approval} event is emitted on calls to {transferFrom}.
 * This allows applications to reconstruct the allowance for all accounts just
 * by listening to said events. Other implementations of the EIP may not emit
 * these events, as it isn't required by the specification.
 *
 * Finally, the non-standard {decreaseAllowance} and {increaseAllowance}
 * functions have been added to mitigate the well-known issues around setting
 * allowances. See {IERC20-approve}.
 */
contract ERC20 is Context, IERC20, IERC20Metadata {
    mapping(address => uint256) private _balances;

    mapping(address => mapping(address => uint256)) private _allowances;

    uint256 private _totalSupply;

    string private _name;
    string private _symbol;

    /**
     * @dev Sets the values for {name} and {symbol}.
     *
     * The default value of {decimals} is 18. To select a different value for
     * {decimals} you should overload it.
     *
     * All two of these values are immutable: they can only be set once during
     * construction.
     */
    constructor(string memory name_, string memory symbol_) {
        _name = name_;
        _symbol = symbol_;
    }

    /**
     * @dev Returns the name of the token.
     */
    function name() public view virtual override returns (string memory) {
        return _name;
    }

    /**
     * @dev Returns the symbol of the token, usually a shorter version of the
     * name.
     */
    function symbol() public view virtual override returns (string memory) {
        return _symbol;
    }

    /**
     * @dev Returns the number of decimals used to get its user representation.
     * For example, if `decimals` equals `2`, a balance of `505` tokens should
     * be displayed to a user as `5.05` (`505 / 10 ** 2`).
     *
     * Tokens usually opt for a value of 18, imitating the relationship between
     * Ether and Wei. This is the value {ERC20} uses, unless this function is
     * overridden;
     *
     * NOTE: This information is only used for _display_ purposes: it in
     * no way affects any of the arithmetic of the contract, including
     * {IERC20-balanceOf} and {IERC20-transfer}.
     */
    function decimals() public view virtual override returns (uint8) {
        return 18;
    }

    /**
     * @dev See {IERC20-totalSupply}.
     */
    function totalSupply() public view virtual override returns (uint256) {
        return _totalSupply;
    }

    /**
     * @dev See {IERC20-balanceOf}.
     */
    function balanceOf(address account) public view virtual override returns (uint256) {
        return _balances[account];
    }

    /**
     * @dev See {IERC20-transfer}.
     *
     * Requirements:
     *
     * - `to` cannot be the zero address.
     * - the caller must have a balance of at least `amount`.
     */
    function transfer(address to, uint256 amount) public virtual override returns (bool) {
        address owner = _msgSender();
        _transfer(owner, to, amount);
        return true;
    }

    /**
     * @dev See {IERC20-allowance}.
     */
    function allowance(address owner, address spender) public view virtual override returns (uint256) {
        return _allowances[owner][spender];
    }

    /**
     * @dev See {IERC20-approve}.
     *
     * NOTE: If `amount` is the maximum `uint256`, the allowance is not updated on
     * `transferFrom`. This is semantically equivalent to an infinite approval.
     *
     * Requirements:
     *
     * - `spender` cannot be the zero address.
     */
    function approve(address spender, uint256 amount) public virtual override returns (bool) {
        address owner = _msgSender();
        _approve(owner, spender, amount);
        return true;
    }

    /**
     * @dev See {IERC20-transferFrom}.
     *
     * Emits an {Approval} event indicating the updated allowance. This is not
     * required by the EIP. See the note at the beginning of {ERC20}.
     *
     * NOTE: Does not update the allowance if the current allowance
     * is the maximum `uint256`.
     *
     * Requirements:
     *
     * - `from` and `to` cannot be the zero address.
     * - `from` must have a balance of at least `amount`.
     * - the caller must have allowance for ``from``'s tokens of at least
     * `amount`.
     */
    function transferFrom(
        address from,
        address to,
        uint256 amount
    ) public virtual override returns (bool) {
        address spender = _msgSender();
        _spendAllowance(from, spender, amount);
        _transfer(from, to, amount);
        return true;
    }

    /**
     * @dev Atomically increases the allowance granted to `spender` by the caller.
     *
     * This is an alternative to {approve} that can be used as a mitigation for
     * problems described in {IERC20-approve}.
     *
     * Emits an {Approval} event indicating the updated allowance.
     *
     * Requirements:
     *
     * - `spender` cannot be the zero address.
     */
    function increaseAllowance(address spender, uint256 addedValue) public virtual returns (bool) {
        address owner = _msgSender();
        _approve(owner, spender, allowance(owner, spender) + addedValue);
        return true;
    }

    /**
     * @dev Atomically decreases the allowance granted to `spender` by the caller.
     *
     * This is an alternative to {approve} that can be used as a mitigation for
     * problems described in {IERC20-approve}.
     *
     * Emits an {Approval} event indicating the updated allowance.
     *
     * Requirements:
     *
     * - `spender` cannot be the zero address.
     * - `spender` must have allowance for the caller of at least
     * `subtractedValue`.
     */
    function decreaseAllowance(address spender, uint256 subtractedValue) public virtual returns (bool) {
        address owner = _msgSender();
        uint256 currentAllowance = allowance(owner, spender);
        require(currentAllowance >= subtractedValue, "ERC20: decreased allowance below zero");
        unchecked {
            _approve(owner, spender, currentAllowance - subtractedValue);
        }

        return true;
    }

    /**
     * @dev Moves `amount` of tokens from `from` to `to`.
     *
     * This internal function is equivalent to {transfer}, and can be used to
     * e.g. implement automatic token fees, slashing mechanisms, etc.
     *
     * Emits a {Transfer} event.
     *
     * Requirements:
     *
     * - `from` cannot be the zero address.
     * - `to` cannot be the zero address.
     * - `from` must have a balance of at least `amount`.
     */
    function _transfer(
        address from,
        address to,
        uint256 amount
    ) internal virtual {
        require(from != address(0), "ERC20: transfer from the zero address");
        require(to != address(0), "ERC20: transfer to the zero address");

        _beforeTokenTransfer(from, to, amount);

        uint256 fromBalance = _balances[from];
        require(fromBalance >= amount, "ERC20: transfer amount exceeds balance");
        unchecked {
            _balances[from] = fromBalance - amount;
            // Overflow not possible: the sum of all balances is capped by totalSupply, and the sum is preserved by
            // decrementing then incrementing.
            _balances[to] += amount;
        }

        emit Transfer(from, to, amount);

        _afterTokenTransfer(from, to, amount);
    }

    /** @dev Creates `amount` tokens and assigns them to `account`, increasing
     * the total supply.
     *
     * Emits a {Transfer} event with `from` set to the zero address.
     *
     * Requirements:
     *
     * - `account` cannot be the zero address.
     */
    function _mint(address account, uint256 amount) internal virtual {
        require(account != address(0), "ERC20: mint to the zero address");

        _beforeTokenTransfer(address(0), account, amount);

        _totalSupply += amount;
        unchecked {
            // Overflow not possible: balance + amount is at most totalSupply + amount, which is checked above.
            _balances[account] += amount;
        }
        emit Transfer(address(0), account, amount);

        _afterTokenTransfer(address(0), account, amount);
    }

    /**
     * @dev Destroys `amount` tokens from `account`, reducing the
     * total supply.
     *
     * Emits a {Transfer} event with `to` set to the zero address.
     *
     * Requirements:
     *
     * - `account` cannot be the zero address.
     * - `account` must have at least `amount` tokens.
     */
    function _burn(address account, uint256 amount) internal virtual {
        require(account != address(0), "ERC20: burn from the zero address");

        _beforeTokenTransfer(account, address(0), amount);

        uint256 accountBalance = _balances[account];
        require(accountBalance >= amount, "ERC20: burn amount exceeds balance");
        unchecked {
            _balances[account] = accountBalance - amount;
            // Overflow not possible: amount <= accountBalance <= totalSupply.
            _totalSupply -= amount;
        }

        emit Transfer(account, address(0), amount);

        _afterTokenTransfer(account, address(0), amount);
    }

    /**
     * @dev Sets `amount` as the allowance of `spender` over the `owner` s tokens.
     *
     * This internal function is equivalent to `approve`, and can be used to
     * e.g. set automatic allowances for certain subsystems, etc.
     *
     * Emits an {Approval} event.
     *
     * Requirements:
     *
     * - `owner` cannot be the zero address.
     * - `spender` cannot be the zero address.
     */
    function _approve(
        address owner,
        address spender,
        uint256 amount
    ) internal virtual {
        require(owner != address(0), "ERC20: approve from the zero address");
        require(spender != address(0), "ERC20: approve to the zero address");

        _allowances[owner][spender] = amount;
        emit Approval(owner, spender, amount);
    }

    /**
     * @dev Updates `owner` s allowance for `spender` based on spent `amount`.
     *
     * Does not update the allowance amount in case of infinite allowance.
     * Revert if not enough allowance is available.
     *
     * Might emit an {Approval} event.
     */
    function _spendAllowance(
        address owner,
        address spender,
        uint256 amount
    ) internal virtual {
        uint256 currentAllowance = allowance(owner, spender);
        if (currentAllowance != type(uint256).max) {
            require(currentAllowance >= amount, "ERC20: insufficient allowance");
            unchecked {
                _approve(owner, spender, currentAllowance - amount);
            }
        }
    }

    /**
     * @dev Hook that is called before any transfer of tokens. This includes
     * minting and burning.
     *
     * Calling conditions:
     *
     * - when `from` and `to` are both non-zero, `amount` of ``from``'s tokens
     * will be transferred to `to`.
     * - when `from` is zero, `amount` tokens will be minted for `to`.
     * - when `to` is zero, `amount` of ``from``'s tokens will be burned.
     * - `from` and `to` are never both zero.
     *
     * To learn more about hooks, head to xref:ROOT:extending-contracts.adoc#using-hooks[Using Hooks].
     */
    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 amount
    ) internal virtual {}

    /**
     * @dev Hook that is called after any transfer of tokens. This includes
     * minting and burning.
     *
     * Calling conditions:
     *
     * - when `from` and `to` are both non-zero, `amount` of ``from``'s tokens
     * has been transferred to `to`.
     * - when `from` is zero, `amount` tokens have been minted for `to`.
     * - when `to` is zero, `amount` of ``from``'s tokens have been burned.
     * - `from` and `to` are never both zero.
     *
     * To learn more about hooks, head to xref:ROOT:extending-contracts.adoc#using-hooks[Using Hooks].
     */
    function _afterTokenTransfer(
        address from,
        address to,
        uint256 amount
    ) internal virtual {}
}


// File @openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol@v4.8.3

// OpenZeppelin Contracts (last updated v4.5.0) (token/ERC20/extensions/ERC20Burnable.sol)

pragma solidity ^0.8.0;


/**
 * @dev Extension of {ERC20} that allows token holders to destroy both their own
 * tokens and those that they have an allowance for, in a way that can be
 * recognized off-chain (via event analysis).
 */
abstract contract ERC20Burnable is Context, ERC20 {
    /**
     * @dev Destroys `amount` tokens from the caller.
     *
     * See {ERC20-_burn}.
     */
    function burn(uint256 amount) public virtual {
        _burn(_msgSender(), amount);
    }

    /**
     * @dev Destroys `amount` tokens from `account`, deducting from the caller's
     * allowance.
     *
     * See {ERC20-_burn} and {ERC20-allowance}.
     *
     * Requirements:
     *
     * - the caller must have allowance for ``accounts``'s tokens of at least
     * `amount`.
     */
    function burnFrom(address account, uint256 amount) public virtual {
        _spendAllowance(account, _msgSender(), amount);
        _burn(account, amount);
    }
}


// File @openzeppelin/contracts/utils/cryptography/ECDSA.sol@v4.8.3

// OpenZeppelin Contracts (last updated v4.8.0) (utils/cryptography/ECDSA.sol)

pragma solidity ^0.8.0;

/**
 * @dev Elliptic Curve Digital Signature Algorithm (ECDSA) operations.
 *
 * These functions can be used to verify that a message was signed by the holder
 * of the private keys of a given address.
 */
library ECDSA {
    enum RecoverError {
        NoError,
        InvalidSignature,
        InvalidSignatureLength,
        InvalidSignatureS,
        InvalidSignatureV // Deprecated in v4.8
    }

    function _throwError(RecoverError error) private pure {
        if (error == RecoverError.NoError) {
            return; // no error: do nothing
        } else if (error == RecoverError.InvalidSignature) {
            revert("ECDSA: invalid signature");
        } else if (error == RecoverError.InvalidSignatureLength) {
            revert("ECDSA: invalid signature length");
        } else if (error == RecoverError.InvalidSignatureS) {
            revert("ECDSA: invalid signature 's' value");
        }
    }

    /**
     * @dev Returns the address that signed a hashed message (`hash`) with
     * `signature` or error string. This address can then be used for verification purposes.
     *
     * The `ecrecover` EVM opcode allows for malleable (non-unique) signatures:
     * this function rejects them by requiring the `s` value to be in the lower
     * half order, and the `v` value to be either 27 or 28.
     *
     * IMPORTANT: `hash` _must_ be the result of a hash operation for the
     * verification to be secure: it is possible to craft signatures that
     * recover to arbitrary addresses for non-hashed data. A safe way to ensure
     * this is by receiving a hash of the original message (which may otherwise
     * be too long), and then calling {toEthSignedMessageHash} on it.
     *
     * Documentation for signature generation:
     * - with https://web3js.readthedocs.io/en/v1.3.4/web3-eth-accounts.html#sign[Web3.js]
     * - with https://docs.ethers.io/v5/api/signer/#Signer-signMessage[ethers]
     *
     * _Available since v4.3._
     */
    function tryRecover(bytes32 hash, bytes memory signature) internal pure returns (address, RecoverError) {
        if (signature.length == 65) {
            bytes32 r;
            bytes32 s;
            uint8 v;
            // ecrecover takes the signature parameters, and the only way to get them
            // currently is to use assembly.
            /// @solidity memory-safe-assembly
            assembly {
                r := mload(add(signature, 0x20))
                s := mload(add(signature, 0x40))
                v := byte(0, mload(add(signature, 0x60)))
            }
            return tryRecover(hash, v, r, s);
        } else {
            return (address(0), RecoverError.InvalidSignatureLength);
        }
    }

    /**
     * @dev Returns the address that signed a hashed message (`hash`) with
     * `signature`. This address can then be used for verification purposes.
     *
     * The `ecrecover` EVM opcode allows for malleable (non-unique) signatures:
     * this function rejects them by requiring the `s` value to be in the lower
     * half order, and the `v` value to be either 27 or 28.
     *
     * IMPORTANT: `hash` _must_ be the result of a hash operation for the
     * verification to be secure: it is possible to craft signatures that
     * recover to arbitrary addresses for non-hashed data. A safe way to ensure
     * this is by receiving a hash of the original message (which may otherwise
     * be too long), and then calling {toEthSignedMessageHash} on it.
     */
    function recover(bytes32 hash, bytes memory signature) internal pure returns (address) {
        (address recovered, RecoverError error) = tryRecover(hash, signature);
        _throwError(error);
        return recovered;
    }

    /**
     * @dev Overload of {ECDSA-tryRecover} that receives the `r` and `vs` short-signature fields separately.
     *
     * See https://eips.ethereum.org/EIPS/eip-2098[EIP-2098 short signatures]
     *
     * _Available since v4.3._
     */
    function tryRecover(
        bytes32 hash,
        bytes32 r,
        bytes32 vs
    ) internal pure returns (address, RecoverError) {
        bytes32 s = vs & bytes32(0x7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff);
        uint8 v = uint8((uint256(vs) >> 255) + 27);
        return tryRecover(hash, v, r, s);
    }

    /**
     * @dev Overload of {ECDSA-recover} that receives the `r and `vs` short-signature fields separately.
     *
     * _Available since v4.2._
     */
    function recover(
        bytes32 hash,
        bytes32 r,
        bytes32 vs
    ) internal pure returns (address) {
        (address recovered, RecoverError error) = tryRecover(hash, r, vs);
        _throwError(error);
        return recovered;
    }

    /**
     * @dev Overload of {ECDSA-tryRecover} that receives the `v`,
     * `r` and `s` signature fields separately.
     *
     * _Available since v4.3._
     */
    function tryRecover(
        bytes32 hash,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) internal pure returns (address, RecoverError) {
        // EIP-2 still allows signature malleability for ecrecover(). Remove this possibility and make the signature
        // unique. Appendix F in the Ethereum Yellow paper (https://ethereum.github.io/yellowpaper/paper.pdf), defines
        // the valid range for s in (301): 0 < s < secp256k1n ÷ 2 + 1, and for v in (302): v ∈ {27, 28}. Most
        // signatures from current libraries generate a unique signature with an s-value in the lower half order.
        //
        // If your library generates malleable signatures, such as s-values in the upper range, calculate a new s-value
        // with 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFEBAAEDCE6AF48A03BBFD25E8CD0364141 - s1 and flip v from 27 to 28 or
        // vice versa. If your library also generates signatures with 0/1 for v instead 27/28, add 27 to v to accept
        // these malleable signatures as well.
        if (uint256(s) > 0x7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF5D576E7357A4501DDFE92F46681B20A0) {
            return (address(0), RecoverError.InvalidSignatureS);
        }

        // If the signature is valid (and not malleable), return the signer address
        address signer = ecrecover(hash, v, r, s);
        if (signer == address(0)) {
            return (address(0), RecoverError.InvalidSignature);
        }

        return (signer, RecoverError.NoError);
    }

    /**
     * @dev Overload of {ECDSA-recover} that receives the `v`,
     * `r` and `s` signature fields separately.
     */
    function recover(
        bytes32 hash,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) internal pure returns (address) {
        (address recovered, RecoverError error) = tryRecover(hash, v, r, s);
        _throwError(error);
        return recovered;
    }

    /**
     * @dev Returns an Ethereum Signed Message, created from a `hash`. This
     * produces hash corresponding to the one signed with the
     * https://eth.wiki/json-rpc/API#eth_sign[`eth_sign`]
     * JSON-RPC method as part of EIP-191.
     *
     * See {recover}.
     */
    function toEthSignedMessageHash(bytes32 hash) internal pure returns (bytes32) {
        // 32 is the length in bytes of hash,
        // enforced by the type signature above
        return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));
    }

    /**
     * @dev Returns an Ethereum Signed Message, created from `s`. This
     * produces hash corresponding to the one signed with the
     * https://eth.wiki/json-rpc/API#eth_sign[`eth_sign`]
     * JSON-RPC method as part of EIP-191.
     *
     * See {recover}.
     */
    function toEthSignedMessageHash(bytes memory s) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n", Strings.toString(s.length), s));
    }

    /**
     * @dev Returns an Ethereum Signed Typed Data, created from a
     * `domainSeparator` and a `structHash`. This produces hash corresponding
     * to the one signed with the
     * https://eips.ethereum.org/EIPS/eip-712[`eth_signTypedData`]
     * JSON-RPC method as part of EIP-712.
     *
     * See {recover}.
     */
    function toTypedDataHash(bytes32 domainSeparator, bytes32 structHash) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked("\x19\x01", domainSeparator, structHash));
    }
}


// File contracts/duckies/games/Genome.sol

pragma solidity 0.8.18;

/**
 * @title Genome
 *
 * @notice The library to work with NFT genomes.
 *
 * Genome is a number with a special structure that defines Duckling genes.
 * All genes are packed consequently in the reversed order in the Genome, meaning the first gene being stored in the last Genome bits.
 * Each gene takes up the block of 8 bits in genome, thus having 256 possible values.
 *
 * Example of genome, following genes Rarity, Head and Body are defined:
 *
 * 00000001|00000010|00000011
 *   Body    Head     Rarity
 *
 * This genome can be represented in uint24 as 66051.
 * Genes have the following values: Body = 1, Head = 2, Rarity = 3.
 */
library Genome {
	/// @notice Number of bits each gene constitutes. Thus, each gene can have 2^8 = 256 possible values.
	uint8 public constant BITS_PER_GENE = 8;

	uint8 public constant COLLECTION_GENE_IDX = 0;

	// Flags
	/// @notice Reserve 30th gene for bool flags, which are stored as a bit field.
	uint8 public constant FLAGS_GENE_IDX = 30;
	uint8 public constant FLAG_TRANSFERABLE = 1; // 0b0000_0001

	// Magic number
	/// @notice Reserve 31th gene for magic number, which is used as an extension for genomes.
	/// Genomes with wrong extension are considered invalid.
	uint8 public constant MAGIC_NUMBER_GENE_IDX = 31;
	uint8 public constant BASE_MAGIC_NUMBER = 209; // Ð
	uint8 public constant MYTHIC_MAGIC_NUMBER = 210; // Ð + 1

	/**
	 * @notice Read flags gene from genome.
	 * @dev Read flags gene from genome.
	 * @param self Genome to get flags gene from.
	 * @return flags Flags gene.
	 */
	function getFlags(uint256 self) internal pure returns (uint8) {
		return getGene(self, FLAGS_GENE_IDX);
	}

	/**
	 * @notice Read specific bit mask flag from genome.
	 * @dev Read specific bit mask flag from genome.
	 * @param self Genome to read flag from.
	 * @param flag Bit mask flag to read.
	 * @return value Value of the flag.
	 */
	function getFlag(uint256 self, uint8 flag) internal pure returns (bool) {
		return getGene(self, FLAGS_GENE_IDX) & flag > 0;
	}

	/**
	 * @notice Set specific bit mask flag in genome.
	 * @dev Set specific bit mask flag in genome.
	 * @param self Genome to set flag in.
	 * @param flag Bit mask flag to set.
	 * @param value Value of the flag.
	 * @return genome Genome with the flag set.
	 */
	function setFlag(uint256 self, uint8 flag, bool value) internal pure returns (uint256) {
		uint8 flags = getGene(self, FLAGS_GENE_IDX);
		if (value) {
			flags |= flag;
		} else {
			flags &= ~flag;
		}
		return setGene(self, FLAGS_GENE_IDX, flags);
	}

	/**
	 * @notice Set `value` to `gene` in genome.
	 * @dev Set `value` to `gene` in genome.
	 * @param self Genome to set gene in.
	 * @param gene Gene to set.
	 * @param value Value to set.
	 * @return genome Genome with the gene set.
	 */
	function setGene(
		uint256 self,
		uint8 gene,
		// by specifying uint8 we set maxCap for gene values, which is 256
		uint8 value
	) internal pure returns (uint256) {
		// number of bytes from genome's rightmost and geneBlock's rightmost
		// NOTE: maximum index of a gene is actually uint5
		uint8 shiftingBy = gene * BITS_PER_GENE;

		// remember genes we will shift off
		uint256 shiftedPart = self & ((1 << shiftingBy) - 1);

		// shift right so that genome's rightmost bit is the geneBlock's rightmost
		self >>= shiftingBy;

		// clear previous gene value by shifting it off
		self >>= BITS_PER_GENE;
		self <<= BITS_PER_GENE;

		// update gene's value
		self += value;

		// reserve space for restoring previously shifted off values
		self <<= shiftingBy;

		// restore previously shifted off values
		self += shiftedPart;

		return self;
	}

	/**
	 * @notice Get `gene` value from genome.
	 * @dev Get `gene` value from genome.
	 * @param self Genome to get gene from.
	 * @param gene Gene to get.
	 * @return geneValue Gene value.
	 */
	function getGene(uint256 self, uint8 gene) internal pure returns (uint8) {
		// number of bytes from genome's rightmost and geneBlock's rightmost
		// NOTE: maximum index of a gene is actually uint5
		uint8 shiftingBy = gene * BITS_PER_GENE;

		uint256 temp = self >> shiftingBy;
		return uint8(temp & ((1 << BITS_PER_GENE) - 1));
	}

	/**
	 * @notice Get largest value of a `gene` in `genomes`.
	 * @dev Get largest value of a `gene` in `genomes`.
	 * @param genomes Genomes to get gene from.
	 * @param gene Gene to get.
	 * @return maxValue Largest value of a `gene` in `genomes`.
	 */
	function _maxGene(uint256[] memory genomes, uint8 gene) internal pure returns (uint8) {
		uint8 maxValue = 0;

		for (uint256 i = 0; i < genomes.length; i++) {
			uint8 geneValue = getGene(genomes[i], gene);
			if (maxValue < geneValue) {
				maxValue = geneValue;
			}
		}

		return maxValue;
	}

	/**
	 * @notice Check if values of `gene` in `genomes` are equal.
	 * @dev Check if values of `gene` in `genomes` are equal.
	 * @param genomes Genomes to check.
	 * @param gene Gene to check.
	 * @return isEqual True if values of `gene` in `genomes` are equal, false otherwise.
	 */
	function _geneValuesAreEqual(
		uint256[] memory genomes,
		uint8 gene
	) internal pure returns (bool) {
		uint8 geneValue = getGene(genomes[0], gene);

		for (uint256 i = 1; i < genomes.length; i++) {
			if (getGene(genomes[i], gene) != geneValue) {
				return false;
			}
		}

		return true;
	}

	/**
	 * @notice Check if values of `gene` in `genomes` are unique.
	 * @dev Check if values of `gene` in `genomes` are unique.
	 * @param genomes Genomes to check.
	 * @param gene Gene to check.
	 * @return isUnique True if values of `gene` in `genomes` are unique, false otherwise.
	 */
	function _geneValuesAreUnique(
		uint256[] memory genomes,
		uint8 gene
	) internal pure returns (bool) {
		uint256 valuesPresentBitfield = 1 << getGene(genomes[0], gene);

		for (uint256 i = 1; i < genomes.length; i++) {
			if (valuesPresentBitfield & (1 << getGene(genomes[i], gene)) != 0) {
				return false;
			}
			valuesPresentBitfield |= 1 << getGene(genomes[i], gene);
		}

		return true;
	}
}


// File contracts/duckies/games/Seeding.sol

pragma solidity 0.8.18;

// chances are represented in per mil, thus uint32
/**
 * @title Seeding
 * @notice A contract that provides pseudo random number generation.
 * Pseudo random number generation is based on the seed created from the salt, pepper, nonce, sender address, and block timestamp.
 * Seed is divided into 32 bit slices, and each slice is used to generate a random number.
 * User of this contract must keep track of used bit slices to avoid reusing them.
 * Salt is a data based on block timestamp and msg sender, and is calculated every time a seed is generated.
 * Pepper is a random data changed periodically by external entity.
 * Nonce is incremented every time a random number is generated.
 */
contract Seeding {
	bytes32 private salt;
	bytes32 private pepper;
	uint256 private nonce;

	/**
	 * @notice Sets the pepper.
	 * @dev Pepper is a random data changed periodically by external entity.
	 * @param newPepper New pepper.
	 */
	function _setPepper(bytes32 newPepper) internal {
		pepper = newPepper;
	}

	/**
	 * @notice Creates a new seed based on the salt, pepper, nonce, sender address, and block timestamp.
	 * @dev Creates a new seed based on the salt, pepper, nonce, sender address, and block timestamp.
	 * @return New seed.
	 */
	function _randomSeed() internal returns (bytes32) {
		// use old salt to generate a new one, so that user's predictions are invalid after function that uses random is called
		salt = keccak256(abi.encode(salt, msg.sender, block.timestamp));
		unchecked {
			nonce++;
		}

		return keccak256(abi.encode(salt, pepper, nonce, msg.sender, block.timestamp));
	}
}


// File contracts/interfaces/IVoucher.sol

pragma solidity 0.8.18;

/**
 * @notice Interface describing Voucher for redeeming game items
 *
 * @dev The Voucher type must have a strict implementation on backend
 *
 * A Voucher is a document signed from the server IssuerKey and allows the execution
 * of actions on the game generally for creating game items, such as Booster Packs, Meld or reward tokens
 *
 */
interface IVoucher {
	/**
	 * @notice Custom error specifying that voucher has already been used.
	 * @param voucherCodeHash Hash of the code of the voucher that has been used.
	 */
	error VoucherAlreadyUsed(bytes32 voucherCodeHash);

	/**
	 * @notice Custom error specifying that voucher has not passed general voucher checks and is invalid.
	 * @param voucher Voucher that is invalid.
	 */
	error InvalidVoucher(Voucher voucher);

	/**
	 * @notice Custom error specifying that the message was expected to be signed by `expected` address, but was signed by `actual`.
	 * @param expected Expected address to have signed the message.
	 * @param actual Actual address that has signed the message.
	 */
	error IncorrectSigner(address expected, address actual);

	/**
	 * @dev Build and encode the Voucher from server side
	 *
	 * Voucher structure will be valid only in chainId until expire timestamp
	 * the beneficiary MUST be the same as the user redeeming the Voucher.
	 *
	 */
	struct Voucher {
		address target; // contract address which the voucher is meant for
		uint8 action; // voucher type defined by the implementation
		address beneficiary; // beneficiary account which voucher will redeem to
		address referrer; // address of the parent
		uint64 expire; // expiration time in seconds UTC
		uint32 chainId; // chain id of the voucher
		bytes32 voucherCodeHash; // hash of voucherCode
		bytes encodedParams; // voucher type specific encoded params
	}

	/**
	 * @notice Use vouchers that were issued and signed by the Back-end to receive game items.
	 * @param vouchers Vouchers issued by the Back-end.
	 * @param signature Vouchers signed by the Back-end.
	 */
	function useVouchers(Voucher[] calldata vouchers, bytes calldata signature) external;

	/**
	 * @notice Use the voucher that was signed by the Back-end to receive game items.
	 * @param voucher Voucher issued by the Back-end.
	 * @param signature Voucher signed by the Back-end.
	 */
	function useVoucher(Voucher calldata voucher, bytes calldata signature) external;

	/**
	 * @notice Event specifying that a voucher has been used.
	 * @param wallet Wallet that used a voucher.
	 * @param action The action of the voucher used.
	 * @param voucherCodeHash The code hash of the voucher used.
	 * @param chainId Id of the chain the voucher was used on.
	 */
	event VoucherUsed(address wallet, uint8 action, bytes32 voucherCodeHash, uint32 chainId);
}


// File contracts/duckies/games/Utils.sol

pragma solidity 0.8.18;

library Utils {
	using ECDSA for bytes32;

	/**
	 * @notice Invalid weights error while trying to generate a weighted random number.
	 * @param weights Empty weights array.
	 */
	error InvalidWeights(uint32[] weights);

	/**
	 * @notice Perform circular shift on the seed by 3 bytes to the left, and returns the shifted slice and the updated seed.
	 * @dev User of this contract must keep track of used bit slices to avoid reusing them.
	 * @param seed Seed to shift and extract the shifted slice from.
	 * @return bitSlice Shifted bit slice.
	 * @return updatedSeed Shifted seed.
	 */
	function _shiftSeedSlice(bytes32 seed) internal pure returns (bytes3, bytes32) {
		bytes3 slice = bytes3(seed);
		return (slice, (seed << 24) | (bytes32(slice) >> 232));
	}

	/**
	 * @notice Extracts a number from the bit slice in range [0, max).
	 * @dev Extracts a number from the bit slice in range [0, max).
	 * @param bitSlice Bit slice to extract the number from.
	 * @param max Max number to extract.
	 * @return Extracted number in range [0, max).
	 */
	function _max(bytes3 bitSlice, uint24 max) internal pure returns (uint24) {
		return uint24(bitSlice) % max;
	}

	/**
	 * @notice Generates a weighted random number given the `weights` array in range [0, weights.length).
	 * @dev Number `x` is generated with probability `weights[x] / sum(weights)`.
	 * @param weights Array of weights.
	 * @return Random number in range [0, weights.length).
	 */
	function _randomWeightedNumber(
		uint32[] memory weights,
		bytes3 bitSlice
	) internal pure returns (uint8) {
		// no sense in empty weights array
		if (weights.length == 0) revert InvalidWeights(weights);

		uint256 randomNumber = _max(bitSlice, uint24(_sum(weights)));

		uint256 segmentRightBoundary = 0;

		for (uint8 i = 0; i < weights.length; i++) {
			segmentRightBoundary += weights[i];
			if (randomNumber < segmentRightBoundary) {
				return i;
			}
		}

		// execution should never reach this
		return uint8(weights.length - 1);
	}

	/**
	 * @notice Calculates sum of all elements in array.
	 * @dev Calculates sum of all elements in array.
	 * @param numbers Array of numbers.
	 * @return sum Sum of all elements in array.
	 */
	function _sum(uint32[] memory numbers) internal pure returns (uint256 sum) {
		for (uint256 i = 0; i < numbers.length; i++) sum += numbers[i];
	}

	/**
	 * @notice Check that `signatures is `encodedData` signed by `signer`. Reverts if not.
	 * @dev Check that `signatures is `encodedData` signed by `signer`. Reverts if not.
	 * @param encodedData Data to check.
	 * @param signature Signature to check.
	 * @param signer Address of the signer.
	 */
	function _requireCorrectSigner(
		bytes memory encodedData,
		bytes memory signature,
		address signer
	) internal pure {
		address actualSigner = keccak256(encodedData).toEthSignedMessageHash().recover(signature);
		if (actualSigner != signer) revert IVoucher.IncorrectSigner(signer, actualSigner);
	}
}


// File contracts/interfaces/IDucklings.sol

pragma solidity 0.8.18;

/**
 * @title IDucklings
 * @notice This interface defines the Ducklings ERC721-compatible contract,
 * which provides basic functionality for minting, burning and querying information about the tokens.
 */
interface IDucklings is IERC721Upgradeable {
	/**
	 * @notice Token not transferable error. Is used when trying to transfer a token that is not transferable.
	 * @param tokenId Token Id that is not transferable.
	 */
	error TokenNotTransferable(uint256 tokenId);
	/**
	 * @notice Invalid magic number error. Is used when trying to mint a token with an invalid magic number.
	 * @param magicNumber Magic number that is invalid.
	 */
	error InvalidMagicNumber(uint8 magicNumber);

	struct Duckling {
		uint256 genome;
		uint64 birthdate;
	}

	// events
	/**
	 * @notice Minted event. Is emitted when a token is minted.
	 * @param to Address of the token owner.
	 * @param tokenId Id of the minted token.
	 * @param genome Genome of the minted token.
	 * @param birthdate Birthdate of the minted token.
	 * @param chainId Id of the chain where the token was minted.
	 */
	event Minted(address to, uint256 tokenId, uint256 genome, uint64 birthdate, uint256 chainId);

	/**
	 * @notice Check whether `account` is owner of `tokenId`.
	 * @dev Revert if `account` is address(0) or `tokenId` does not exist.
	 * @param account Address to check.
	 * @param tokenId Token Id to check.
	 * @return isOwnerOf True if `account` is owner of `tokenId`, false otherwise.
	 */
	function isOwnerOf(address account, uint256 tokenId) external view returns (bool);

	/**
	 * @notice Check whether `account` is owner of `tokenIds`.
	 * @dev Revert if `account` is address(0) or any of `tokenIds` do not exist.
	 * @param account Address to check.
	 * @param tokenIds Token Ids to check.
	 * @return isOwnerOfBatch True if `account` is owner of `tokenIds`, false otherwise.
	 */
	function isOwnerOfBatch(
		address account,
		uint256[] calldata tokenIds
	) external view returns (bool);

	/**
	 * @notice Get genome of `tokenId`.
	 * @dev Revert if `tokenId` does not exist.
	 * @param tokenId Token Id to get the genome of.
	 * @return genome Genome of `tokenId`.
	 */
	function getGenome(uint256 tokenId) external view returns (uint256);

	/**
	 * @notice Get genomes of `tokenIds`.
	 * @dev Revert if any of `tokenIds` do not exist.
	 * @param tokenIds Token Ids to get the genomes of.
	 * @return genomes Genomes of `tokenIds`.
	 */
	function getGenomes(uint256[] calldata tokenIds) external view returns (uint256[] memory);

	/**
	 * @notice Mint token with `genome` to `to`. Emits Minted event.
	 * @dev Revert if `to` is address(0) or `genome` has wrong magic number.
	 * @param to Address to mint token to.
	 * @param genome Genome of the token to mint.
	 * @return tokenId Id of the minted token.
	 */
	function mintTo(address to, uint256 genome) external returns (uint256);

	/**
	 * @notice Mint tokens with `genomes` to `to`. Emits Minted event for each token.
	 * @dev Revert if `to` is address(0) or any of `genomes` has wrong magic number.
	 * @param to Address to mint tokens to.
	 * @param genomes Genomes of the tokens to mint.
	 * @return tokenIds Ids of the minted tokens.
	 */
	function mintBatchTo(
		address to,
		uint256[] calldata genomes
	) external returns (uint256[] memory);

	/**
	 * @notice Burn token with `tokenId`.
	 * @dev Revert if `tokenId` does not exist.
	 * @param tokenId Id of the token to burn.
	 */
	function burn(uint256 tokenId) external;

	/**
	 * @notice Burn tokens with `tokenIds`.
	 * @dev Revert if any of `tokenIds` do not exist.
	 * @param tokenIds Ids of the tokens to burn.
	 */
	function burnBatch(uint256[] calldata tokenIds) external;
}


// File contracts/interfaces/IDuckyFamily.sol

pragma solidity 0.8.18;

interface IDuckyFamily is IVoucher {
	// Errors
	error InvalidMintParams(MintParams mintParams);
	error InvalidMeldParams(MeldParams meldParams);

	error MintingRulesViolated(uint8 collectionId, uint8 amount);
	error MeldingRulesViolated(uint256[] tokenIds);
	error IncorrectGenomesForMelding(uint256[] genomes);

	// Events
	event Melded(address owner, uint256[] meldingTokenIds, uint256 meldedTokenId, uint256 chainId);

	// Vouchers
	enum VoucherActions {
		MintPack,
		MeldFlock
	}

	struct MintParams {
		address to;
		uint8 size;
		bool isTransferable;
	}

	struct MeldParams {
		address owner;
		uint256[] tokenIds;
		bool isTransferable;
	}

	// DuckyFamily

	// for now, Solidity does not support starting value for enum
	// enum Collections {
	// 	Duckling = 0,
	// 	Zombeak,
	// 	Mythic
	// }

	enum Rarities {
		Common,
		Rare,
		Epic,
		Legendary
	}

	enum GeneDistributionTypes {
		Even,
		Uneven
	}

	enum GenerativeGenes {
		Collection,
		Rarity,
		Color,
		Family,
		Body,
		Head
	}

	enum MythicGenes {
		Collection,
		UniqId
	}

	// Config
	function getMintPrice() external view returns (uint256);

	function getMeldPrices() external view returns (uint256[4] memory);

	function getCollectionsGeneValues() external view returns (uint8[][3] memory, uint8);

	function getCollectionsGeneDistributionTypes() external view returns (uint32[3] memory);

	// Mint and Meld
	function mintPack(uint8 size) external;

	function meldFlock(uint256[] calldata meldingTokenIds) external;
}


// File contracts/duckies/games/DuckyFamily/DuckyFamilyV1.sol

// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;






/**
 * @title DuckyFamilyV1
 *
 * @notice DuckyFamily contract defines rules of Ducky Family game, which is a card game similar to Happy Families and Uno games.
 * This game also incorporates vouchers as defined in IVoucher interface.
 *
 * DuckyFamily operates on Ducklings NFT, which is defined in a corresponding contract. DuckyFamily can mint, burn and query information about NFTs
 * by calling Ducklings contract.
 *
 * Users can buy NFT (card) packs of different size. When a pack is bought, a number of cards is generated and assigned to the user.
 * The packs can be bought with Duckies token, so user should approve DuckyFamily contract to spend Duckies on his behalf.
 *
 * Each card has a unique genome, which is a 256-bit number. The genome is a combination of different genes, which describe the card and its properties.
 * There are 3 types of cards introduced in this game, which are differentiated by the 'collection' gene: Duckling, Zombeak and Mythic.
 * Duckling and Zombeak NFTs have a class system, which is defined by 'rarity' gene: Common, Rare, Epic and Legendary.
 * Mythic NFTs are not part of the class system and are considered to be the most rare and powerful cards in the game.
 *
 * All cards have a set of generative genes, which are used to describe the card, its rarity and image.
 * There are 2 types of generative genes: with even and uneven chance for each value of that gene.
 * All values of even genes are generated with equal probability, while uneven genes have a higher chance for the first values and lower for the last values.
 * Thus, while even genes can describe the card, uneven can set the rarity of the card.
 *
 * Note: don't confuse 'rarity' gene with rarity of the card. 'Rarity' gene is a part of the game logic, while rarity of the card is a value this card represents.
 * Henceforth, if a 'Common' rarity gene card has uneven generative genes with high values (which means this card has a tiny chance to being generated),
 * then this card can be more rare than some 'Rare' rarity gene cards.
 * So, when we mean 'rarity' gene, we will use quotes, while when we mean rarity of the card, we will use it without quotes.
 *
 * Duckling are the main cards in the game, as they are the only way users can get Mythic cards.
 * However, users are not obliged to use every Duckling cards to help them get Mythic, they can improve them and collect the rarest ones.
 * Users can get Duckling cards from minting packs.
 *
 * Users can improve the 'rarity' of the card by melding them. Melding is a process of combining a flock of 5 cards to create a new one.
 * The new card will have the same 'collection' gene as the first card in the flock, but the 'rarity' gene will be incremented.
 * However, users must oblige to specific rules when melding cards:
 * 1. All cards in the flock must have the same 'collection' gene.
 * 2. All cards in the flock must have the same 'rarity' gene.
 * 3a. When melding Common cards, all cards in the flock must have either the same Color or Family gene values.
 * 3b. When melding Rare and Epic cards, all cards in the flock must have both the same Color and Family gene values.
 * 3c. When melding Legendary cards, all cards in the flock must have the same Color and different Family gene values.
 * 4. Mythic cards cannot be melded.
 * 5. Legendary Zombeak cards cannot be melded.
 *
 * Other generative genes of the melded card are not random, but are calculated from the genes of the source cards.
 * This process is called 'inheritance' and is the following:
 * 1. Each generative gene is inherited separately
 * 2. A gene has a high chance of being inherited from the first card in the flock, and this chance is lower for each next card in the flock.
 * 3. A gene has a mere chance of 'positive mutation', which sets inherited gene value to be bigger than the biggest value of this gene in the flock.
 *
 * Melding is not free and has a different cost for each 'rarity' of the cards being melded.
 *
 * Zombeak are secondary cards, that you can only get when melding mutates. There is a different chance (defined in Config section below) for each 'rarity' of the Duckling cards that are being melded,
 * that the melding result card will mutate to Zombeak. If the melding mutates, then the new card will have the same 'rarity' gene as the source cards.
 * This logic makes Zombeak cards more rare than some Duckling cards, as they can only be obtained by melding mutating.
 * However, Zombeak cards cannot be melded into Mythic, which means their main value is rarity.
 *
 * Mythic are the most rare and powerful cards in the game. They can only be obtained by melding Legendary Duckling cards with special rules described above.
 * The rarity of the Mythic card is defined by the 'UniqId' gene, which corresponds to the picture of the card. The higher the 'UniqId' gene value, the rarer the card.
 * The 'UniqId' value is correlated with the 'peculiarity' of the flock that creates the Mythic: the higher the peculiarity, the higher the 'UniqId' value.
 * Peculiarity of the card is a sum of all uneven gene values of this card, and peculiarity of the flock is a sum of peculiarities of all cards in the flock.
 *
 * Mythic cards give bonuses to their owned depending on their rarity. These bonuses will be revealed in the future, but they may include
 * free Yellow tokens (with vesting claim mechanism), an ability to change existing cards, stealing / fighting other cards, etc.
 */
contract FlattenDuckyFamilyV1 is IDuckyFamily, AccessControl, Seeding {
	using Genome for uint256;

	// Roles
	bytes32 public constant MAINTAINER_ROLE = keccak256('MAINTAINER_ROLE'); // can change minting and melding price

	address public issuer; // issuer of Vouchers

	// Store the vouchers to avoid replay attacks
	mapping(bytes32 => bool) internal _usedVouchers;

	// ------- Config -------

	uint8 internal constant ducklingCollectionId = 0;
	uint8 internal constant zombeakCollectionId = 1;
	uint8 internal constant mythicCollectionId = 2;
	uint8 internal constant RARITIES_NUM = 4;

	uint8 public constant MAX_PACK_SIZE = 50;
	uint8 public constant FLOCK_SIZE = 5;

	uint8 internal constant collectionGeneIdx = Genome.COLLECTION_GENE_IDX;
	uint8 internal constant rarityGeneIdx = 1;
	uint8 internal constant flagsGeneIdx = Genome.FLAGS_GENE_IDX;
	// general genes start after Collection and Rarity
	uint8 internal constant generativeGenesOffset = 2;

	// number of values for each gene for Duckling and Zombeak collections
	uint8[][3] internal collectionsGeneValuesNum; // set in constructor

	// distribution type of each gene for Duckling and Zombeak collections (0 - even, 1 - uneven)
	uint32[3] internal collectionsGeneDistributionTypes = [
		2940, // reverse(001111101101) = 101101111100
		2940, // reverse(001111101101) = 101101111100
		107 // reverse(11010110) = 01101011
	];

	// peculiarity is a sum of uneven gene values for Ducklings
	uint16 internal maxPeculiarity;
	// mythic dispersion define the interval size in which UniqId value is generated
	uint8 internal constant MYTHIC_DISPERSION = 5;
	uint8 internal mythicAmount = 60;

	// chance of a Duckling of a certain rarity to be generated
	uint32[] internal rarityChances = [850, 120, 25, 5]; // per mil

	// chance of a Duckling of certain rarity to mutate to Zombeak while melding
	uint32[] internal collectionMutationChances = [150, 100, 50, 10]; // per mil

	uint32[] internal geneMutationChance = [955, 45]; // 4.5% to mutate gene value
	uint32[] internal geneInheritanceChances = [400, 300, 150, 100, 50]; // per mil

	// ------- Public values -------

	ERC20Burnable public duckiesContract;
	IDucklings public ducklingsContract;
	address public treasureVaultAddress;

	uint256 public mintPrice;
	uint256[RARITIES_NUM] public meldPrices; // [0] - melding Commons, [1] - melding Rares...

	// ------- Constructor -------

	/**
	 * @notice Sets Duckies, Ducklings and Treasure Vault addresses, minting and melding prices and other game config.
	 * @dev Grants DEFAULT_ADMIN_ROLE and MAINTAINER_ROLE to the deployer.
	 * @param duckiesAddress Address of Duckies ERC20 contract.
	 * @param ducklingsAddress Address of Ducklings ERC721 contract.
	 * @param treasureVaultAddress_ Address of Treasure Vault contract.
	 */
	constructor(address duckiesAddress, address ducklingsAddress, address treasureVaultAddress_) {
		_grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
		_grantRole(MAINTAINER_ROLE, msg.sender);

		duckiesContract = ERC20Burnable(duckiesAddress);
		ducklingsContract = IDucklings(ducklingsAddress);
		treasureVaultAddress = treasureVaultAddress_;

		uint256 decimalsMultiplier = 10 ** duckiesContract.decimals();

		mintPrice = 50 * decimalsMultiplier;
		meldPrices = [
			100 * decimalsMultiplier,
			200 * decimalsMultiplier,
			500 * decimalsMultiplier,
			1000 * decimalsMultiplier
		];

		// Duckling genes: (Collection, Rarity), Color, Family, Body, Head, Eyes, Beak, Wings, FirstName, Temper, Skill, Habitat, Breed
		collectionsGeneValuesNum[0] = [4, 5, 10, 25, 30, 14, 10, 36, 16, 12, 5, 28];
		// Zombeak genes: (Collection, Rarity), Color, Family, Body, Head, Eyes, Beak, Wings, FirstName, Temper, Skill, Habitat, Breed
		collectionsGeneValuesNum[1] = [2, 3, 7, 6, 9, 7, 10, 36, 16, 12, 5, 28];
		// Mythic genes: (Collection, UniqId), Temper, Skill, Habitat, Breed, Birthplace, Quirk, Favorite Food, Favorite Color
		collectionsGeneValuesNum[2] = [16, 12, 5, 28, 5, 10, 8, 4];

		maxPeculiarity = _calcConfigPeculiarity(
			collectionsGeneValuesNum[ducklingCollectionId],
			collectionsGeneDistributionTypes[ducklingCollectionId]
		);
	}

	// ------- Random -------

	/**
	 * @notice Sets the pepper for random generator.
	 * @dev Require MAINTAINER_ROLE to call. Pepper is a random data changed periodically by external entity.
	 * @param pepper New pepper.
	 */
	function setPepper(bytes32 pepper) external onlyRole(MAINTAINER_ROLE) {
		_setPepper(pepper);
	}

	// ------- Vouchers -------

	/**
	 * @notice Sets the issuer of Vouchers.
	 * @dev Require DEFAULT_ADMIN_ROLE to call.
	 * @param account Address of a new issuer.
	 */
	function setIssuer(address account) external onlyRole(DEFAULT_ADMIN_ROLE) {
		issuer = account;
	}

	/**
	 * @notice Use multiple Vouchers. Check the signature and invoke internal function for each voucher.
	 * @dev Vouchers are issued by the Back-End and signed by the issuer.
	 * @param vouchers Array of Vouchers to use.
	 * @param signature Vouchers signed by the issuer.
	 */
	function useVouchers(Voucher[] calldata vouchers, bytes calldata signature) external {
		Utils._requireCorrectSigner(abi.encode(vouchers), signature, issuer);
		for (uint8 i = 0; i < vouchers.length; i++) {
			_useVoucher(vouchers[i]);
		}
	}

	/**
	 * @notice Use a single Voucher. Check the signature and invoke internal function.
	 * @dev Vouchers are issued by the Back-End and signed by the issuer.
	 * @param voucher Voucher to use.
	 * @param signature Voucher signed by the issuer.
	 */
	function useVoucher(Voucher calldata voucher, bytes calldata signature) external {
		Utils._requireCorrectSigner(abi.encode(voucher), signature, issuer);
		_useVoucher(voucher);
	}

	/**
	 * @notice Check the validity of a voucher, decode voucher params and mint or meld tokens depending on voucher's type. Emits VoucherUsed event. Internal function.
	 * @dev Vouchers are issued by the Back-End and signed by the issuer.
	 * @param voucher Voucher to use.
	 */
	function _useVoucher(Voucher memory voucher) internal {
		_requireValidVoucher(voucher);

		_usedVouchers[voucher.voucherCodeHash] = true;

		// parse & process Voucher
		if (voucher.action == uint8(VoucherActions.MintPack)) {
			MintParams memory mintParams = abi.decode(voucher.encodedParams, (MintParams));

			// mintParams checks
			if (
				mintParams.to == address(0) ||
				mintParams.size == 0 ||
				mintParams.size > MAX_PACK_SIZE
			) revert InvalidMintParams(mintParams);

			_mintPackTo(mintParams.to, mintParams.size, mintParams.isTransferable);
		} else if (voucher.action == uint8(VoucherActions.MeldFlock)) {
			MeldParams memory meldParams = abi.decode(voucher.encodedParams, (MeldParams));

			// meldParams checks
			if (meldParams.owner == address(0) || meldParams.tokenIds.length != FLOCK_SIZE)
				revert InvalidMeldParams(meldParams);

			_meldOf(meldParams.owner, meldParams.tokenIds, meldParams.isTransferable);
		} else {
			revert InvalidVoucher(voucher);
		}

		emit VoucherUsed(
			voucher.beneficiary,
			voucher.action,
			voucher.voucherCodeHash,
			voucher.chainId
		);
	}

	/**
	 * @notice Check the validity of a voucher, reverts if invalid.
	 * @dev Voucher address must be this contract, beneficiary must be msg.sender, voucher must not be used before, voucher must not be expired.
	 * @param voucher Voucher to check.
	 */
	function _requireValidVoucher(Voucher memory voucher) internal view {
		if (_usedVouchers[voucher.voucherCodeHash])
			revert VoucherAlreadyUsed(voucher.voucherCodeHash);

		if (
			voucher.target != address(this) ||
			voucher.beneficiary != msg.sender ||
			block.timestamp > voucher.expire ||
			voucher.chainId != block.chainid
		) revert InvalidVoucher(voucher);
	}

	// -------- Config --------

	/**
	 * @notice Get the mint price in Duckies with decimals.
	 * @dev Get the mint price in Duckies with decimals.
	 * @return mintPrice Mint price in Duckies with decimals.
	 */
	function getMintPrice() external view returns (uint256) {
		return mintPrice;
	}

	/**
	 * @notice Set the mint price in Duckies without decimals.
	 * @dev Require MAINTAINER_ROLE to call.
	 * @param price Mint price in Duckies without decimals.
	 */
	function setMintPrice(uint256 price) external onlyRole(MAINTAINER_ROLE) {
		mintPrice = price * 10 ** duckiesContract.decimals();
	}

	/**
	 * @notice Get the meld price for each 'rarity' in Duckies with decimals.
	 * @dev Get the meld price for each 'rarity' in Duckies with decimals.
	 * @return meldPrices Array of meld prices in Duckies with decimals.
	 */
	function getMeldPrices() external view returns (uint256[RARITIES_NUM] memory) {
		return meldPrices;
	}

	/**
	 * @notice Set the meld price for each 'rarity' in Duckies without decimals.
	 * @dev Require MAINTAINER_ROLE to call.
	 * @param prices Array of meld prices in Duckies without decimals.
	 */
	function setMeldPrices(
		uint256[RARITIES_NUM] calldata prices
	) external onlyRole(MAINTAINER_ROLE) {
		for (uint8 i = 0; i < RARITIES_NUM; i++) {
			meldPrices[i] = prices[i] * 10 ** duckiesContract.decimals();
		}
	}

	/**
	 * @notice Get number of gene values for all collections and a number of different Mythic tokens.
	 * @dev Get number of gene values for all collections and a number of different Mythic tokens.
	 * @return collectionsGeneValuesNum Arrays of number of gene values for all collections and a mythic amount.
	 */
	function getCollectionsGeneValues() external view returns (uint8[][3] memory, uint8) {
		return (collectionsGeneValuesNum, mythicAmount);
	}

	/**
	 * @notice Get gene distribution types for all collections.
	 * @dev Get gene distribution types for all collections.
	 * @return collectionsGeneDistributionTypes Arrays of gene distribution types for all collections.
	 */
	function getCollectionsGeneDistributionTypes() external view returns (uint32[3] memory) {
		return collectionsGeneDistributionTypes;
	}

	/**
	 * @notice Set gene values number for each gene for Duckling collection.
	 * @dev Require DEFAULT_ADMIN_ROLE to call.
	 * @param duckingGeneValuesNum Array of gene values number for each gene for Duckling collection.
	 */
	function setDucklingGeneValues(
		uint8[] memory duckingGeneValuesNum
	) external onlyRole(DEFAULT_ADMIN_ROLE) {
		collectionsGeneValuesNum[0] = duckingGeneValuesNum;
		maxPeculiarity = _calcConfigPeculiarity(
			collectionsGeneValuesNum[ducklingCollectionId],
			collectionsGeneDistributionTypes[ducklingCollectionId]
		);
	}

	/**
	 * @notice Set gene distribution types for Duckling collection.
	 * @dev Require DEFAULT_ADMIN_ROLE to call.
	 * @param ducklingGeneDistrTypes Gene distribution types for Duckling collection.
	 */
	function setDucklingGeneDistributionTypes(
		uint32 ducklingGeneDistrTypes
	) external onlyRole(DEFAULT_ADMIN_ROLE) {
		collectionsGeneDistributionTypes[0] = ducklingGeneDistrTypes;
		maxPeculiarity = _calcConfigPeculiarity(
			collectionsGeneValuesNum[ducklingCollectionId],
			collectionsGeneDistributionTypes[ducklingCollectionId]
		);
	}

	/**
	 * @notice Set gene values number for each gene for Zombeak collection.
	 * @dev Require DEFAULT_ADMIN_ROLE to call.
	 * @param zombeakGeneValuesNum Array of gene values number for each gene for Duckling collection.
	 */
	function setZombeakGeneValues(
		uint8[] memory zombeakGeneValuesNum
	) external onlyRole(DEFAULT_ADMIN_ROLE) {
		collectionsGeneValuesNum[1] = zombeakGeneValuesNum;
	}

	/**
	 * @notice Set gene distribution types for Zombeak collection.
	 * @dev Require DEFAULT_ADMIN_ROLE to call.
	 * @param zombeakGeneDistrTypes Gene distribution types for Zombeak collection.
	 */
	function setZombeakGeneDistributionTypes(
		uint32 zombeakGeneDistrTypes
	) external onlyRole(DEFAULT_ADMIN_ROLE) {
		collectionsGeneDistributionTypes[1] = zombeakGeneDistrTypes;
	}

	/**
	 * @notice Set number of different Mythic tokens.
	 * @dev Require DEFAULT_ADMIN_ROLE to call.
	 * @param amount Number of different Mythic tokens.
	 */
	function setMythicAmount(uint8 amount) external onlyRole(DEFAULT_ADMIN_ROLE) {
		mythicAmount = amount;
	}

	/**
	 * @notice Set gene values number for each gene for Mythic collection.
	 * @dev Require DEFAULT_ADMIN_ROLE to call.
	 * @param mythicGeneValuesNum Array of gene values number for each gene for Mythic collection.
	 */
	function setMythicGeneValues(
		uint8[] memory mythicGeneValuesNum
	) external onlyRole(DEFAULT_ADMIN_ROLE) {
		collectionsGeneValuesNum[2] = mythicGeneValuesNum;
	}

	/**
	 * @notice Set gene distribution types for Mythic collection.
	 * @dev Require DEFAULT_ADMIN_ROLE to call.
	 * @param mythicGeneDistrTypes Gene distribution types for Mythic collection.
	 */
	function setMythicGeneDistributionTypes(
		uint32 mythicGeneDistrTypes
	) external onlyRole(DEFAULT_ADMIN_ROLE) {
		collectionsGeneDistributionTypes[2] = mythicGeneDistrTypes;
	}

	// ------- Mint -------

	/**
	 * @notice Mint a pack with `size` of Ducklings. Transfer Duckies from the sender to the TreasureVault.
	 * @dev `Size` must be less than or equal to `MAX_PACK_SIZE`.
	 * @param size Number of Ducklings in the pack.
	 */
	function mintPack(uint8 size) external {
		duckiesContract.transferFrom(msg.sender, treasureVaultAddress, mintPrice * size);
		_mintPackTo(msg.sender, size, true);
	}

	/**
	 * @notice Mint a pack with `amount` of Ducklings to `to` and set transferable flag for each token. Internal function.
	 * @dev `amount` must be less than or equal to `MAX_PACK_SIZE`.
	 * @param to Address to mint the pack to.
	 * @param amount Number of Ducklings in the pack.
	 * @param isTransferable Transferable flag for each token.
	 * @return tokenIds Array of minted token IDs.
	 */
	function _mintPackTo(
		address to,
		uint8 amount,
		bool isTransferable
	) internal returns (uint256[] memory tokenIds) {
		if (amount == 0 || amount > MAX_PACK_SIZE)
			revert MintingRulesViolated(ducklingCollectionId, amount);

		tokenIds = new uint256[](amount);
		uint256[] memory tokenGenomes = new uint256[](amount);

		for (uint256 i = 0; i < amount; i++) {
			tokenGenomes[i] = _generateGenome(ducklingCollectionId).setFlag(
				Genome.FLAG_TRANSFERABLE,
				isTransferable
			);
		}

		tokenIds = ducklingsContract.mintBatchTo(to, tokenGenomes);
	}

	/**
	 * @notice Generate genome for Duckling or Zombeak.
	 * @dev Generate and set all genes from a corresponding collection.
	 * @param collectionId Collection ID.
	 * @return genome Generated genome.
	 */
	function _generateGenome(uint8 collectionId) internal returns (uint256) {
		if (collectionId != ducklingCollectionId && collectionId != zombeakCollectionId) {
			revert MintingRulesViolated(collectionId, 1);
		}

		(bytes3 bitSlice, bytes32 seed) = Utils._shiftSeedSlice(_randomSeed());

		uint256 genome;

		genome = genome.setGene(collectionGeneIdx, collectionId);
		genome = genome.setGene(
			rarityGeneIdx,
			Utils._randomWeightedNumber(rarityChances, bitSlice)
		);
		genome = _generateAndSetGenes(
			genome,
			collectionId,
			collectionsGeneValuesNum[collectionId],
			collectionsGeneDistributionTypes[collectionId],
			seed
		);
		genome = genome.setGene(Genome.MAGIC_NUMBER_GENE_IDX, Genome.BASE_MAGIC_NUMBER);

		return genome;
	}

	/**
	 * @notice Generate and set all genes from a corresponding collection to `genome`.
	 * @dev Generate and set all genes from a corresponding collection to `genome`.
	 * @param genome Genome to set genes to.
	 * @param collectionId Collection ID.
	 * @param seed Random seed to generate genes from.
	 * @return genome Genome with set genes.
	 */
	function _generateAndSetGenes(
		uint256 genome,
		uint8 collectionId,
		uint8[] memory geneValues,
		uint32 geneDistributionTypes,
		bytes32 seed
	) internal pure returns (uint256) {
		uint8 geneValuesNum = uint8(geneValues.length);
		bytes32 newSeed;

		// generate and set each gene
		for (uint8 i = 0; i < geneValuesNum; i++) {
			GeneDistributionTypes distrType = _getDistributionType(geneDistributionTypes, i);
			bytes3 bitSlice;
			(bitSlice, newSeed) = Utils._shiftSeedSlice(seed);
			genome = _generateAndSetGene(
				genome,
				generativeGenesOffset + i,
				geneValues[i],
				distrType,
				bitSlice
			);
		}

		// set default values for Ducklings
		if (collectionId == ducklingCollectionId) {
			Rarities rarity = Rarities(genome.getGene(rarityGeneIdx));

			if (rarity == Rarities.Common) {
				genome = genome.setGene(uint8(GenerativeGenes.Body), 0);
				genome = genome.setGene(uint8(GenerativeGenes.Head), 0);
			} else if (rarity == Rarities.Rare) {
				genome = genome.setGene(uint8(GenerativeGenes.Head), 0);
			}
		}

		return genome;
	}

	/**
	 * @notice Generate and set a gene with `geneIdx` to `genome`.
	 * @dev Generate and set a gene with `geneIdx` to `genome`.
	 * @param genome Genome to set a gene to.
	 * @param geneIdx Gene index.
	 * @param geneValuesNum Number of gene values.
	 * @param distrType Gene distribution type.
	 * @param bitSlice Random bit slice to generate a gene from.
	 * @return genome Genome with set gene.
	 */
	function _generateAndSetGene(
		uint256 genome,
		uint8 geneIdx,
		uint8 geneValuesNum,
		GeneDistributionTypes distrType,
		bytes3 bitSlice
	) internal pure returns (uint256) {
		uint8 geneValue;

		if (distrType == GeneDistributionTypes.Even) {
			geneValue = uint8(Utils._max(bitSlice, geneValuesNum));
		} else {
			geneValue = uint8(_generateUnevenGeneValue(geneValuesNum, bitSlice));
		}

		// gene with value 0 means it is a default value, thus this   \/
		genome = genome.setGene(geneIdx, geneValue + 1);

		return genome;
	}

	/**
	 * @notice Generate mythic genome based on melding `genomes`.
	 * @dev Calculates flock peculiarity, and randomizes UniqId corresponding to the peculiarity.
	 * @param genomes Array of genomes to meld into Mythic.
	 * @return genome Generated Mythic genome.
	 */
	function _generateMythicGenome(
		uint256[] memory genomes,
		uint16 maxPeculiarity_,
		uint16 mythicAmount_
	) internal returns (uint256) {
		(bytes3 bitSlice, bytes32 seed) = Utils._shiftSeedSlice(_randomSeed());

		uint16 flockPeculiarity = 0;

		for (uint8 i = 0; i < genomes.length; i++) {
			flockPeculiarity += _calcPeculiarity(
				genomes[i],
				uint8(collectionsGeneValuesNum[ducklingCollectionId].length),
				collectionsGeneDistributionTypes[ducklingCollectionId]
			);
		}

		uint16 maxSumPeculiarity = maxPeculiarity_ * uint16(genomes.length);
		uint16 maxUniqId = mythicAmount_ - 1;
		uint16 pivotalUniqId = uint16((uint64(flockPeculiarity) * maxUniqId) / maxSumPeculiarity); // multiply and then divide to avoid float numbers
		(uint16 leftEndUniqId, uint16 uniqIdSegmentLength) = _calcUniqIdGenerationParams(
			pivotalUniqId,
			maxUniqId
		);

		uint16 uniqId = leftEndUniqId + uint16(Utils._max(bitSlice, uniqIdSegmentLength));

		uint256 genome;
		genome = genome.setGene(collectionGeneIdx, mythicCollectionId);
		genome = genome.setGene(uint8(MythicGenes.UniqId), uint8(uniqId));
		genome = _generateAndSetGenes(
			genome,
			mythicCollectionId,
			collectionsGeneValuesNum[mythicCollectionId],
			collectionsGeneDistributionTypes[mythicCollectionId],
			seed
		);
		genome = genome.setGene(Genome.MAGIC_NUMBER_GENE_IDX, Genome.MYTHIC_MAGIC_NUMBER);

		return genome;
	}

	// ------- Meld -------

	/**
	 * @notice Meld tokens with `meldingTokenIds` into a new token. Calls internal function.
	 * @dev Meld tokens with `meldingTokenIds` into a new token.
	 * @param meldingTokenIds Array of token IDs to meld.
	 */
	function meldFlock(uint256[] calldata meldingTokenIds) external {
		// assume all tokens have the same rarity. This is checked later.
		uint256 meldPrice = meldPrices[
			ducklingsContract.getGenome(meldingTokenIds[0]).getGene(rarityGeneIdx)
		];
		duckiesContract.transferFrom(msg.sender, treasureVaultAddress, meldPrice);

		_meldOf(msg.sender, meldingTokenIds, true);
	}

	/**
	 * @notice Meld tokens with `meldingTokenIds` into a new token. Internal function.
	 * @dev Check `owner` is indeed the owner of `meldingTokenIds`. Burn NFTs with `meldingTokenIds`. Transfers Duckies to the TreasureVault.
	 * @param meldingTokenIds Array of token IDs to meld.
	 * @param isTransferable Whether the new token is transferable.
	 * @return meldedTokenId ID of the new token.
	 */
	function _meldOf(
		address owner,
		uint256[] memory meldingTokenIds,
		bool isTransferable
	) internal returns (uint256) {
		if (meldingTokenIds.length != FLOCK_SIZE) revert MeldingRulesViolated(meldingTokenIds);
		if (!ducklingsContract.isOwnerOfBatch(owner, meldingTokenIds))
			revert MeldingRulesViolated(meldingTokenIds);

		uint256[] memory meldingGenomes = ducklingsContract.getGenomes(meldingTokenIds);
		_requireGenomesSatisfyMelding(meldingGenomes);

		ducklingsContract.burnBatch(meldingTokenIds);

		uint256 meldedGenome = _meldGenomes(meldingGenomes).setFlag(
			Genome.FLAG_TRANSFERABLE,
			isTransferable
		);
		uint256 meldedTokenId = ducklingsContract.mintTo(owner, meldedGenome);

		emit Melded(owner, meldingTokenIds, meldedTokenId, block.chainid);

		return meldedTokenId;
	}

	/**
	 * @notice Check that `genomes` satisfy melding rules. Reverts if not.
	 * @dev Check that `genomes` satisfy melding rules. Reverts if not.
	 * @param genomes Array of genomes to check.
	 */
	function _requireGenomesSatisfyMelding(uint256[] memory genomes) internal pure {
		if (
			// equal collections
			!Genome._geneValuesAreEqual(genomes, collectionGeneIdx) ||
			// Rarities must be the same
			!Genome._geneValuesAreEqual(genomes, rarityGeneIdx) ||
			// not Mythic
			genomes[0].getGene(collectionGeneIdx) == mythicCollectionId
		) revert IncorrectGenomesForMelding(genomes);

		Rarities rarity = Rarities(genomes[0].getGene(rarityGeneIdx));
		bool sameColors = Genome._geneValuesAreEqual(genomes, uint8(GenerativeGenes.Color));
		bool sameFamilies = Genome._geneValuesAreEqual(genomes, uint8(GenerativeGenes.Family));
		bool uniqueFamilies = Genome._geneValuesAreUnique(genomes, uint8(GenerativeGenes.Family));

		// specific melding rules
		if (rarity == Rarities.Common) {
			// Common
			if (
				// cards must have the same Color OR the same Family
				!sameColors && !sameFamilies
			) revert IncorrectGenomesForMelding(genomes);
		} else {
			// Rare, Epic
			if (rarity == Rarities.Rare || rarity == Rarities.Epic) {
				if (
					// cards must have the same Color AND the same Family
					!sameColors || !sameFamilies
				) revert IncorrectGenomesForMelding(genomes);
			} else {
				// Legendary
				if (
					// not Legendary Zombeak
					genomes[0].getGene(collectionGeneIdx) == zombeakCollectionId ||
					// cards must have the same Color AND be of each Family
					!sameColors ||
					!uniqueFamilies
				) revert IncorrectGenomesForMelding(genomes);
			}
		}
	}

	/**
	 * @notice Meld `genomes` into a new genome.
	 * @dev Meld `genomes` into a new genome gene by gene. Set the corresponding collection
	 * @param genomes Array of genomes to meld.
	 * @return meldedGenome Melded genome.
	 */
	function _meldGenomes(uint256[] memory genomes) internal returns (uint256) {
		uint8 collectionId = genomes[0].getGene(collectionGeneIdx);
		Rarities rarity = Rarities(genomes[0].getGene(rarityGeneIdx));

		(bytes3 bitSlice, bytes32 seed) = Utils._shiftSeedSlice(_randomSeed());

		// if melding Duckling, they can mutate or evolve into Mythic
		if (collectionId == ducklingCollectionId) {
			if (_isCollectionMutating(rarity, collectionMutationChances, bitSlice)) {
				uint256 zombeakGenome = _generateGenome(zombeakCollectionId);
				return zombeakGenome.setGene(rarityGeneIdx, uint8(rarity));
			}

			if (rarity == Rarities.Legendary) {
				return _generateMythicGenome(genomes, maxPeculiarity, mythicAmount);
			}
		}

		uint256 meldedGenome;

		// set the same collection
		meldedGenome = meldedGenome.setGene(collectionGeneIdx, collectionId);
		// increase rarity
		meldedGenome = meldedGenome.setGene(rarityGeneIdx, genomes[0].getGene(rarityGeneIdx) + 1);

		uint8[] memory geneValuesNum = collectionsGeneValuesNum[collectionId];
		uint32 geneDistTypes = collectionsGeneDistributionTypes[collectionId];

		for (uint8 i = 0; i < geneValuesNum.length; i++) {
			(bitSlice, seed) = Utils._shiftSeedSlice(seed);
			uint8 geneValue = _meldGenes(
				genomes,
				generativeGenesOffset + i,
				geneValuesNum[i],
				_getDistributionType(geneDistTypes, i),
				geneMutationChance,
				geneInheritanceChances,
				bitSlice
			);
			meldedGenome = meldedGenome.setGene(generativeGenesOffset + i, geneValue);
		}

		// randomize Body for Common and Head for Rare for Ducklings
		if (collectionId == ducklingCollectionId) {
			(bitSlice, seed) = Utils._shiftSeedSlice(seed);
			if (rarity == Rarities.Common) {
				meldedGenome = _generateAndSetGene(
					meldedGenome,
					uint8(GenerativeGenes.Body),
					geneValuesNum[uint8(GenerativeGenes.Body) - generativeGenesOffset],
					GeneDistributionTypes.Uneven,
					bitSlice
				);
			} else if (rarity == Rarities.Rare) {
				meldedGenome = _generateAndSetGene(
					meldedGenome,
					uint8(GenerativeGenes.Head),
					geneValuesNum[uint8(GenerativeGenes.Head) - generativeGenesOffset],
					GeneDistributionTypes.Uneven,
					bitSlice
				);
			}
		}

		meldedGenome = meldedGenome.setGene(Genome.MAGIC_NUMBER_GENE_IDX, Genome.BASE_MAGIC_NUMBER);

		return meldedGenome;
	}

	/**
	 * @notice Randomize if collection is mutating.
	 * @dev Randomize if collection is mutating.
	 * @param rarity Rarity of the collection.
	 * @param bitSlice Bit slice to use for randomization.
	 * @return isMutating True if mutating, false otherwise.
	 */
	function _isCollectionMutating(
		Rarities rarity,
		uint32[] memory mutationChances,
		bytes3 bitSlice
	) internal pure returns (bool) {
		// check if mutating chance for this rarity is present
		if (mutationChances.length <= uint8(rarity)) {
			return false;
		}

		uint32 mutationPercentage = mutationChances[uint8(rarity)];
		// dynamic array is needed for `_randomWeightedNumber()`
		uint32[] memory chances = new uint32[](2);
		chances[0] = mutationPercentage;
		chances[1] = 1000 - mutationPercentage; // 1000 as changes are represented in per mil
		return Utils._randomWeightedNumber(chances, bitSlice) == 0;
	}

	/**
	 * @notice Meld `gene` from `genomes` into a new gene value.
	 * @dev Meld `gene` from `genomes` into a new gene value. Gene mutation and inheritance are applied.
	 * @param genomes Array of genomes to meld.
	 * @param gene Gene to be meld.
	 * @param maxGeneValue Max gene value.
	 * @param geneDistrType Gene distribution type.
	 * @param bitSlice Bit slice to use for randomization.
	 * @return geneValue Melded gene value.
	 */
	function _meldGenes(
		uint256[] memory genomes,
		uint8 gene,
		uint8 maxGeneValue,
		GeneDistributionTypes geneDistrType,
		uint32[] memory mutationChance,
		uint32[] memory inheritanceChances,
		bytes3 bitSlice
	) internal pure returns (uint8) {
		// gene mutation
		if (
			geneDistrType == GeneDistributionTypes.Uneven &&
			Utils._randomWeightedNumber(mutationChance, bitSlice) == 1
		) {
			uint8 maxPresentGeneValue = Genome._maxGene(genomes, gene);
			return maxPresentGeneValue == maxGeneValue ? maxGeneValue : maxPresentGeneValue + 1;
		}

		// gene inheritance
		uint8 inheritanceIdx = Utils._randomWeightedNumber(inheritanceChances, bitSlice);
		return genomes[inheritanceIdx].getGene(gene);
	}

	// ------- Helpers -------

	/**
	 * @notice Get gene distribution type.
	 * @dev Get gene distribution type.
	 * @param distributionTypes Distribution types.
	 * @param idx Index of the gene.
	 * @return Gene distribution type.
	 */
	function _getDistributionType(
		uint32 distributionTypes,
		uint8 idx
	) internal pure returns (GeneDistributionTypes) {
		return
			distributionTypes & (1 << idx) == 0
				? GeneDistributionTypes.Even
				: GeneDistributionTypes.Uneven;
	}

	/**
	 * @notice Generate uneven gene value given the maximum number of values.
	 * @dev Generate uneven gene value using reciprocal distribution described below.
	 * @param valuesNum Maximum number of gene values.
	 * @param bitSlice Bit slice to use for randomization.
	 * @return geneValue Gene value.
	 */
	function _generateUnevenGeneValue(
		uint8 valuesNum,
		bytes3 bitSlice
	) internal pure returns (uint8) {
		// using reciprocal distribution
		// gene value is selected as ceil[(2N/(x+1))-N],
		// where x is random number between 0 and 1
		// Because of shape of reciprocal graph,
		// evenly distributed x values will result in unevenly distributed y values.

		// N - number of gene values
		uint256 N = uint256(valuesNum);
		// Generates number from 1 to 10^6
		uint256 x = 1 + Utils._max(bitSlice, 1_000_000);
		// Calculates uneven distributed y, value of y is between 0 and N
		uint256 y = (2 * N * 1_000) / (Math.sqrt(x) + 1_000) - N;
		return uint8(y);
	}

	/**
	 * @notice Calculate max peculiarity for a current Duckling config.
	 * @dev Sum up number of uneven gene values.
	 * @return maxPeculiarity Max peculiarity.
	 */
	function _calcConfigPeculiarity(
		uint8[] memory geneValuesNum,
		uint32 geneDistrTypes
	) internal pure returns (uint16) {
		uint16 sum = 0;

		uint8 genesNum = uint8(geneValuesNum.length);
		for (uint8 i = 0; i < genesNum; i++) {
			if (_getDistributionType(geneDistrTypes, i) == GeneDistributionTypes.Uneven) {
				// add number of values and not actual values as actual values start with 1, which means number of values and actual values are equal
				sum += geneValuesNum[i];
			}
		}

		return sum;
	}

	/**
	 * @notice Calculate peculiarity for a given genome.
	 * @dev Sum up number of uneven gene values.
	 * @param genome Genome.
	 * @return peculiarity Peculiarity.
	 */
	function _calcPeculiarity(
		uint256 genome,
		uint8 genesNum,
		uint32 geneDistrTypes
	) internal pure returns (uint16) {
		uint16 sum = 0;

		for (uint8 i = 0; i < genesNum; i++) {
			if (_getDistributionType(geneDistrTypes, i) == GeneDistributionTypes.Uneven) {
				// add number of values and not actual values as actual values start with 1, which means number of values and actual values are equal
				sum += genome.getGene(i + generativeGenesOffset);
			}
		}

		return sum;
	}

	/**
	 * @notice Calculate `leftEndUniqId` and `uniqIdSegmentLength` for UniqId generation.
	 * @dev Then UniqId is generated by adding a random number [0, `uniqIdSegmentLength`) to `leftEndUniqId`.
	 * @param pivotalUniqId Pivotal UniqId.
	 * @param maxUniqId Max UniqId.
	 * @return leftEndUniqId Left end of the UniqId segment.
	 * @return uniqIdSegmentLength Length of the UniqId segment.
	 */
	function _calcUniqIdGenerationParams(
		uint16 pivotalUniqId,
		uint16 maxUniqId
	) internal pure returns (uint16 leftEndUniqId, uint16 uniqIdSegmentLength) {
		if (pivotalUniqId < MYTHIC_DISPERSION) {
			// mythic id range overlaps with left dispersion border
			leftEndUniqId = 0;
			uniqIdSegmentLength = pivotalUniqId + MYTHIC_DISPERSION;
		} else if (maxUniqId < pivotalUniqId + MYTHIC_DISPERSION) {
			// mythic id range overlaps with right dispersion border
			leftEndUniqId = pivotalUniqId - MYTHIC_DISPERSION;
			uniqIdSegmentLength = maxUniqId - leftEndUniqId + 1; // +1 to include right border, where the last UniqId is located
		} else {
			// mythic id range does not overlap with dispersion borders
			leftEndUniqId = pivotalUniqId - MYTHIC_DISPERSION;
			uniqIdSegmentLength = 2 * MYTHIC_DISPERSION;
		}
	}
}
