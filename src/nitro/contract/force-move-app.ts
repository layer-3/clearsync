import { utils, Contract } from 'ethers';

import ForceMoveAppArtifact from '../../../artifacts/contracts/nitro/interfaces/IForceMoveApp.sol/IForceMoveApp.json';
import { State, getVariablePart } from './state';

//  https://github.com/ethers-io/ethers.js/issues/602#issuecomment-574671078
export const ForceMoveAppContractInterface = new utils.Interface(ForceMoveAppArtifact.abi);

/**
 * Calls the valiTransition method on the supplied ForceMoveApp using eth_call
 * @param fromState a State
 * @param toState a State
 * @param appContract a ForceMoveApp contract address
 * @returns a Promise that resolves to a boolean
 */
export async function validTransition(
  fromState: State,
  toState: State,
  appContract: Contract,
): Promise<boolean> {
  const numberOfParticipants = toState.participants.length;
  const fromVariablePart = getVariablePart(fromState);
  const toVariablePart = getVariablePart(toState);

  return await appContract.validTransition(fromVariablePart, toVariablePart, numberOfParticipants);
}

/**
 * Encodes a validTransition method call as the data on an ethereum transaction. Useful for testing gas consumption of a ForceMoveApp.
 */
export function createValidTransitionTransaction(
  fromState: State,
  toState: State,
): { data: string } {
  const numberOfParticipants = toState.participants.length;
  const fromVariablePart = getVariablePart(fromState);
  const toVariablePart = getVariablePart(toState);
  const data = ForceMoveAppContractInterface.encodeFunctionData('validTransition', [
    fromVariablePart,
    toVariablePart,
    numberOfParticipants,
  ]);
  return { data };
}
