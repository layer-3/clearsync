import { expect } from 'chai';

/**
 * Wrapper for transactions that are expected to succeed with no return values.
 */
export async function expectSucceedWithNoReturnValues(fn: () => void) {
  const txResult = (await fn()) as any;

  expect(txResult.length).to.equal(0);
}

/**
 * Wrapper for calls to `stateIsSupported` that are expected to succeed.
 */
export async function expectSupportedState(fn: () => any): Promise<void> {
  const txResult = await fn();

  // `.stateIsSupported` returns a (bool, string) tuple
  expect(txResult.length).to.equal(2);
  expect(txResult[0]).to.equal(true);
  expect(txResult[1]).to.equal('');
}

/**
 * Wrapper for calls to `stateIsSupported` that are expected to fail.
 * Checks that the reason for failure matches the supplied `reason` string.
 */
export async function expectUnsupportedState(fn: () => void, reason?: string) {
  const txResult = (await fn()) as any;

  expect(txResult.length).to.equal(2);
  expect(txResult[0]).to.equal(false);
  if (reason) expect(txResult[1]).to.equal(reason);
}
