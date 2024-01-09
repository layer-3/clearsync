import chai from 'chai';
import chaiAsPromised from 'chai-as-promised';

export async function expectRevert(fn: () => Promise<any>, pattern?: string | RegExp) {
  chai.use(chaiAsPromised);

  if (pattern) {
    await chai.expect(fn()).to.be.rejectedWith(pattern);
  } else {
    await chai.expect(fn()).to.be.rejected;
  }
}
