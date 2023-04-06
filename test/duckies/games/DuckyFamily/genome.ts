const BITS_PER_GENE = 8;

export class Genome {
  genome = 0n;

  constructor(from?: number | bigint) {
    switch (typeof from) {
      case 'bigint': {
        this.genome = from;
        break;
      }

      case 'number': {
        this.genome = BigInt(from);
        break;
      }

      default: {
        this.genome = 0n;
        break;
      }
    }
  }

  static from(from: number[]): Genome {
    const genome = new Genome();
    for (const [i, geneValue] of from.entries()) {
      genome.setGene(i, geneValue);
    }
    return genome;
  }

  setGene(gene: number, value: number): this {
    // number of bytes from genome's rightmost and geneBlock's rightmost
    const shiftingBy = BigInt(gene * BITS_PER_GENE);

    let genome = this.genome;

    // remember genes we will shift off
    const shiftedPart = genome & ((1n << shiftingBy) - 1n);

    // shift right so that genome's rightmost bit is the geneBlock's rightmost
    genome >>= shiftingBy;

    // clear previous gene value by shifting it off
    genome >>= BigInt(BITS_PER_GENE);
    genome <<= BigInt(BITS_PER_GENE);

    // update gene's value
    genome += BigInt(value);

    // reserve space for restoring previously shifted off values
    genome <<= shiftingBy;

    // restore previously shifted off values
    genome += shiftedPart;

    this.genome = genome;

    return this;
  }

  randomizeGene(gene: number, maxValue: number): this {
    this.setGene(gene, Math.round(Math.random() * maxValue));
    return this;
  }

  getGene(gene: number): number {
    // number of bytes from genome's rightmost and geneBlock's rightmost
    const shiftingBy = BigInt(gene * BITS_PER_GENE);

    const temp = this.genome >> shiftingBy;
    return Number.parseInt((temp & ((1n << BigInt(BITS_PER_GENE)) - 1n)).toString());
  }
}
