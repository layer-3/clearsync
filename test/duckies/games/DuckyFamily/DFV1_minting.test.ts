describe('DuckyFamilyV1 minting', () => {
  describe('generateGenome', () => {
    it('duckling has correct structure');

    it('zombeak has correct structure');

    it('mythic has correct structure');
  });

  describe('generateAndSetGenes', () => {
    it('has correct numbers of genes');

    it('does not exceed max gene values');
  });

  describe('mintPack', () => {
    it('duckies are paid for mint');

    it('correct amount of tokens is minted');

    it('revert on amount == 0');

    it('revert on amount > MAX_PACK_SIZE');

    it('event is emitted');
  });
});
