import { describe } from 'mocha';

describe('EscrowApp', () => {
  describe('succeed', () => {
    it('success when all rules obeyed');
  });

  describe('revert', () => {
    it('revert when turn num > 3');

    it('revert when more that 2 participants');

    describe('settlement', () => {
      it('revert when asset allocated to more then 1 destination');

      it('revert when another asset added');

      it('revert when asset allocation amount has changed');

      it('revert when 3 distinct destinations');

      it('revert if destination has not changed');
    });
  });
});
