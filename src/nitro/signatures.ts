import {Wallet, utils, Signature} from 'ethers';

import {Bytes, Uint256} from './contract/types';
import {getSignedBy, getSignersIndices, getSignersNum} from './bitfield-utils';
import {hashChallengeMessage} from './contract/challenge';
import {getChannelId} from './contract/channel';
import {Outcome} from './contract/outcome';
import {
  getFixedPart,
  hashState,
  RecoveredVariablePart,
  SignedVariablePart,
  State,
  VariablePart,
} from './contract/state';

/**
 * A {@link State} along with a {@link Signature} on it
 */
export interface SignedState {
  state: State;
  signature: Signature;
}

export function getStateSignerAddress(signedState: SignedState): string {
  const stateHash = hashState(signedState.state);
  const recoveredAddress = utils.verifyMessage(utils.arrayify(stateHash), signedState.signature);
  const participants = signedState.state.participants;

  if (participants.indexOf(recoveredAddress) < 0) {
    throw new Error(
      `Recovered address ${recoveredAddress} is not a participant in channel ${getChannelId(
        getFixedPart(signedState.state)
      )}`
    );
  }
  return recoveredAddress;
}

/**
 * Encodes, hashes and signs a State using the supplied privateKey
 * @param state a State
 * @param privateKey an ECDSA private key
 * @returns a SignedState
 */
export function signState(state: State, privateKey: string): SignedState {
  const wallet = new Wallet(privateKey);
  if (state.participants.indexOf(wallet.address) < 0) {
    throw new Error("The state must be signed with a participant's private key");
  }

  const hashedState = hashState(state);

  const signature = signData(hashedState, privateKey);
  return {state, signature};
}

export async function sign(wallet: Wallet, msgHash: string | Uint8Array): Promise<Signature> {
  // MsgHash is a hex string
  // Returns an object with v, r, and s properties.
  return utils.splitSignature(await wallet.signMessage(utils.arrayify(msgHash)));
}

// Towards #761 https://github.com/statechannels/go-nitro/issues/761
//
export type ShortenedVariablePart =
  | {
      signerIndices: number[];
      appData?: Bytes;
      outcome?: Outcome;
      isFinal?: boolean;
    }
  | number[];

// Map<turnNum => ShortenedVariablePart> | Map<turnNum => signerIndices>
export type TurnNumToShortenedVariablePart = Map<number, ShortenedVariablePart>;

/**
 * Signs a state using wallets at indices determined by signedBy bitfield.
 * @param state A state to sign.
 * @param wallets An array of wallets.
 * @param signedBy SignedBy bitfield.
 * @returns States signed by wallets at indices determined by signedBy bitfield.
 */
export async function signStateWithBitfield(
  state: State,
  wallets: Wallet[],
  signedBy: Uint256
): Promise<Signature[]> {
  if (wallets.length < getSignersNum(signedBy)) {
    throw new Error('Insufficient wallets');
  }

  const promises = getSignersIndices(signedBy).map(
    async i => await sign(wallets[i], hashState(state))
  );

  return Promise.all(promises);
}

/**
 * Convert provided ShortenedSignedBy and turnNum into RecoveredVariablePart. Missing ShortenedVariablePart fields are initialized to default values.
 * @param turnNum Turn number.
 * @param shortenedVP ShortenedVariablePart.
 * @returns RecoveredVariablePart.
 */
export function shortenedToRecoveredVariablePart(
  turnNum: number,
  shortenedVP: ShortenedVariablePart
): RecoveredVariablePart {
  // if just an array of signerIndices was supplied, convert it to ShortenedVariablePart
  if (Array.isArray(shortenedVP)) {
    shortenedVP = {
      signerIndices: shortenedVP,
    };
  }

  const outcome = shortenedVP.outcome ?? [];
  const appData = shortenedVP.appData ?? '0x';
  const isFinal = shortenedVP.isFinal ?? false;

  return {
    variablePart: {
      outcome,
      appData,
      turnNum,
      isFinal,
    },
    signedBy: getSignedBy(shortenedVP.signerIndices),
  };
}

/**
 * Convert a mapping between turnNumbers and ShortenedVariableParts to an array of RecoveredVariableParts.
 * @param turnNumToShortenedVP A mapping between turnNumbers and ShortenedVariableParts.
 * @returns An array of RecoveredVariableParts.
 */
export function shortenedToRecoveredVariableParts(
  turnNumToShortenedVP: TurnNumToShortenedVariablePart
): RecoveredVariablePart[] {
  return Array.from(turnNumToShortenedVP).map(([turnNum, shortenedVP]) => {
    return shortenedToRecoveredVariablePart(turnNum, shortenedVP);
  });
}

//
// end towards #761 https://github.com/statechannels/go-nitro/issues/761

/**
 * Maps the supplied wallets array to (a Promise of) an array of signatures by those wallets on the supplied states, using whoSignedWhat to map from wallet to state.
 */
export async function signStates(
  states: State[],
  wallets: Wallet[],
  whoSignedWhat: number[]
): Promise<Signature[]> {
  const stateHashes = states.map(s => hashState(s));
  const promises = wallets.map(async (w, i) => await sign(w, stateHashes[whoSignedWhat[i]]));
  return Promise.all(promises);
}

/**
 * Maps supplied signatures to variable parts, using whoSignedWhat.
 */
export function bindSignatures(
  variableParts: VariablePart[],
  signatures: Signature[],
  whoSignedWhat: number[]
): SignedVariablePart[] {
  const signedVariableParts = variableParts.map(
    vp =>
      ({
        variablePart: vp,
        sigs: [],
      } as SignedVariablePart)
  );

  for (let i = 0; i < signatures.length; i++) {
    signedVariableParts[whoSignedWhat[i]].sigs.push(signatures[i]);
  }

  return signedVariableParts;
}

/**
 * Maps supplied signatures to variable parts, using whoSignedWhat.
 */
export function bindSignaturesWithSignedByBitfield(
  variableParts: VariablePart[],
  signatures: Signature[],
  whoSignedWhat: number[]
): RecoveredVariablePart[] {
  const recoveredVariableParts = variableParts.map(
    vp =>
      ({
        variablePart: vp,

        signedBy: '0',
      } as RecoveredVariablePart)
  );

  for (let i = 0; i < signatures.length; i++) {
    const updatedSignedBy = Number(recoveredVariableParts[whoSignedWhat[i]].signedBy) | (2 ** i);
    recoveredVariableParts[whoSignedWhat[i]].signedBy = updatedSignedBy.toString();
  }

  return recoveredVariableParts;
}

/**
 * Signs a challenge message (necessary for submitting a challenge) using the last of the supplied signedStates and privateKey
 * @param signedStates an array of type SignedState
 * @param privateKey an ECDSA private key (must be a participant in the channel)
 * @returns a Signature
 */
export function signChallengeMessage(signedStates: SignedState[], privateKey: string): Signature {
  if (signedStates.length === 0) {
    throw new Error('At least one signed state must be provided');
  }
  const wallet = new Wallet(privateKey);
  if (signedStates[0].state.participants.indexOf(wallet.address) < 0) {
    throw new Error("The state must be signed with a participant's private key");
  }
  const challengeState = signedStates[signedStates.length - 1].state;
  const challengeHash = hashChallengeMessage(challengeState);

  return signData(challengeHash, privateKey);
}

function hashMessage(hashedData: string): string {
  return utils.hashMessage(utils.arrayify(hashedData));
}

export function signData(hashedData: string, privateKey: string): Signature {
  const signingKey = new utils.SigningKey(privateKey);

  return utils.splitSignature(signingKey.signDigest(hashMessage(hashedData)));
}
