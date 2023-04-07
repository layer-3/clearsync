describe('DuckyFamilyV1 vouchers', () => {
  describe('issuer', () => {
    it('admin can set issuer');

    it('revert on not admin set issuer');
  });

  describe('useVoucher', () => {
    describe('general revert', () => {
      it('revert on incorrect signer');

      it('revert on using same voucher for second time');

      it('revert on target address != contract address');

      it('revert on beneficiary != sender');

      it('revert on expired voucher');

      it('revert on wrong chainId');

      it('revert on invalid voucher action');
    });

    describe('mint voucher', () => {
      it('successfuly use mint voucher');

      it('duckies are not paid for mint');

      it('revert on to == address(0)');

      it('revert on size == 0');

      it('revert on size > MAX_PACK_SIZE');

      it('event emitted');
    });

    describe('meld voucher', () => {
      it('successfuly use meld voucher');

      it('duckies are not paid for meld');

      it('revert on owner == address(0)');

      it('revert on number of tokens != FLOCK_SIZE');

      it('event is emitted');
    });
  });

  describe('useVouchers', () => {
    it('can use several mint vouchers');

    it('can use several meld vouchers');

    it('revert on incorrect signer');
  });
});
