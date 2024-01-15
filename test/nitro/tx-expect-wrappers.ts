import { expect } from 'chai';

/**
 * Wrapper for transactions that are expected to succeed with no return values.
 */
export async function expectSucceedWithNoReturnValues(fn: () => Promise<unknown>): Promise<void> {
  const txResult = await fn();

  if (txResult === undefined) {
    expect(txResult).to.equal(undefined);
  } else {
    expect((txResult as { length: number }).length).to.equal(0);
  }
}

/**
 * Wrapper for calls to `stateIsSupported` that are expected to succeed.
 */
export async function expectSupportedState(fn: () => Promise<unknown[]>): Promise<void> {
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
export async function expectUnsupportedState(
  fn: () => Promise<unknown[]>,
  reason?: string,
): Promise<void> {
  const txResult = await fn();

  expect(txResult.length).to.equal(2);
  expect(txResult[0]).to.equal(false);
  if (reason) expect(txResult[1]).to.equal(reason);
}
