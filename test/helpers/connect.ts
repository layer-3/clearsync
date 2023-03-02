import type { SignerWithAddress } from '@nomiclabs/hardhat-ethers/signers';
import type { Contract } from 'ethers';

export function connect<ContractT extends Contract>(
  contract: ContractT,
  caller: SignerWithAddress,
): ContractT {
  return contract.connect(caller) as ContractT;
}

export function connectGroup<ContractT extends Contract>(
  contract: ContractT,
  callers: SignerWithAddress[],
): ContractT[] {
  const connectedContracts: ContractT[] = [];

  for (const caller of callers) {
    // no need to deep clone contract, as `connect` returns a copy of original contract
    connectedContracts.push(connect(contract, caller));
  }

  return connectedContracts;
}
