// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

import './IVoucher.sol';

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
