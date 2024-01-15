// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {ExitFormat as Outcome} from '@statechannels/exit-format/contracts/ExitFormat.sol';

interface INitroTypes {
    struct Signature {
        uint8 v;
        bytes32 r;
        bytes32 s;
    }

    struct FixedPart {
        address[] participants;
        uint64 channelNonce;
        address appDefinition;
        uint48 challengeDuration;
    }

    struct VariablePart {
        Outcome.SingleAssetExit[] outcome;
        bytes appData;
        uint48 turnNum;
        bool isFinal;
    }

    struct SignedVariablePart {
        VariablePart variablePart;
        Signature[] sigs;
    }

    struct RecoveredVariablePart {
        VariablePart variablePart;
        uint256 signedBy; // bitmask
    }
}
