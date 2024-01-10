// The last value in a result from an ethers event emission (i.e., Contract.on(<filter>, <result>))
// is an object with keys as the names of the fields emitted by the event.
// eslint-disable-next-line @typescript-eslint/no-explicit-any
export function parseEventResult(result: any[]): Record<string, any> {
  return result.at(-1).args;
}
