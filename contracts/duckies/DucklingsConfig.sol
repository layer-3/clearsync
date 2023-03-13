// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity 0.8.18;

abstract contract DucklingsConfig {
	enum BaseTraits {
		CollectionId,
		Class // may be absent if `GenerativeCollectionConfig.segregation` is set to false
	}

	// ===================================
	// THIS CAN BE CONFIGURED EVEN FURTHER

	enum Classes {
		Common,
		Rare,
		Epic,
		Legendary
	}

	// weights of generating NFT of a specified class
	uint8[] internal classWeights; // = [70, 20, 5, 1];

	// weights of a melding NFT inheriting a certain trait from some parent or mutating this trait
	uint8[] internal inheritanceWeights; // = [6, 5, 4, 3, 2, 1];

	// ===================================

	enum DistributionTypes {
		Equal,
		Unequal
	}

	enum AcquireType {
		ByPaymentOrVoucher, // NFTs can be acquired both by payment and by voucher
		ByPayment, // NFTs can be acquired only by payment
		ByVoucher, // NFTs can be acquired only by voucher
		ByEvolutionOrMutation // NFTs can not be acquired manually, instead, they can be minted as a result of evolution or mutation.
	}

	enum GenerationType {
		Pregenerated, // NFTs are represented as a whole unique pictures
		Generative // NFTs are represented as a set of distinct assets (like head, body, etc.)
	}

	struct CollectionConfig {
		uint8 id; // collection id
		uint64 availableAfter; // UNIX timestamp collection minting is available after
		uint64 availableBefore; // UNIX timestamp collection minting is not available after
		//
		AcquireType acquireType; // minting type
		//
		GenerationType genType; // collection generation type
		PregenCollectionConfig pregenColConfig; // pre-generated collection config. Makes sense only when `genType` is `Pregenerated`
		GenerativeCollectionConfig generColConfig; // generative collection config. Makes sense only when `genType` is `Generative`
	}

	enum PregenType {
		Unique, // there are no same NFTs in circulation
		UserUnique, // there are no same NFTs owned by the same account, but several accounts can own the same NFT
		Repeatable // the same NFT can be minted several times
	}

	struct PregenCollectionConfig {
		uint8 size; // Number of unique NFT assets available. For PregenType.Unique also specifies max number of NFTs.
		PregenType pregenType; // collection pre-generation type
	}

	uint8 internal constant _NIL_COLLECTION_ID = 0;

	struct GenerativeCollectionConfig {
		uint8[] traitValuesNum; // number of values for each trait
		uint32 traitDistributionType; // bit field: 0 - equal distribution, 1 - unequal distribution, defined by a special function
		//
		// segregation
		bool segregation; // whether NFTs from this collection are divided into classes
		uint8[][] defaultedTraits; // [class][trait indices] - indices of traits for each class that are going to be set to a default value
		// This is true when `segregation` is true. Otherwise, it does not make sense.abi
		// bool meldable; // whether several NFTs can be melded into another one
		uint8[][2] meldingRules; // rules applied to melding into class in the same collection. [0] - uniqueValuesOfTraitWithIndex. [1] - equalValuesOfTraitWithIndex.
		//
		// evolution
		uint8 evolutionCollectionId; // id of the collection NFTs of the highest class are melded into
		uint8[][2] evolutionRules; // rules applied to melding into NFT of different collection (evolution collection). [0] - uniqueValuesOfTraitWithIndex. [1] - equalValuesOfTraitWithIndex.
		//
		// mutation
		int8 traitMutationAmount; // value a trait mutates by. Is represented by `int`, so that trait can both upgrade and degrade
		uint8 mutationCollectionId; // id of the collection of a mutated NFT
		uint8[] mutationWeights; // chances of collection mutating when melding NFTs of each class
	}
}
