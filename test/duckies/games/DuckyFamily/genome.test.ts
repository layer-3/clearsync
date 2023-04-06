import { expect } from 'chai';

import { Genome } from './genome';

describe('genome', () => {
  it('from', () => {
    const from = [1, 2, 3, 4];
    const genome = Genome.from(from);

    expect(genome.getGene(0)).to.equal(1);
    expect(genome.getGene(1)).to.equal(2);
    expect(genome.getGene(2)).to.equal(3);
    expect(genome.getGene(3)).to.equal(4);
  });

  it('getGene', () => {
    const genome = new Genome(0b100_0000_0010);

    expect(new Genome(0b00).getGene(0)).to.equal(0);
    expect(new Genome(0b01).getGene(0)).to.equal(1);
    expect(new Genome(0b10).getGene(0)).to.equal(2);
    expect(new Genome(0b1_0000_0000).getGene(1)).to.equal(1);

    expect(genome.getGene(0)).to.equal(2);
    expect(genome.getGene(1)).to.equal(4);
  });

  it('setGene', () => {
    const genome = new Genome();

    expect(genome.setGene(0, 0).genome).to.equal(0b0);
    expect(genome.setGene(0, 1).genome).to.equal(0b1);
    expect(genome.setGene(0, 5).genome).to.equal(0b101);

    expect(genome.setGene(1, 1).genome).to.equal(0b1_0000_0101);
    expect(genome.setGene(1, 5).genome).to.equal(0b101_0000_0101);
  });

  it('randomizeGene', () => {
    const maxNum = 5;

    for (let i = 0; i < maxNum * 3; i++) {
      const genome = new Genome();
      genome.randomizeGene(0, maxNum);
      expect(genome.getGene(0)).to.be.within(0, maxNum);
    }

    for (let i = 0; i < maxNum * 3; i++) {
      const genome = new Genome();
      genome.randomizeGene(1, maxNum);
      expect(genome.getGene(1)).to.be.within(0, maxNum);
    }
  });
});
