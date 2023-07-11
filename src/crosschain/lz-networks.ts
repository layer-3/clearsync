interface LZNetworkInfo {
  lzChainId: number;
  endpointAddress: string;
}

export const mainnetChainIdToLZInfo = new Map<number, LZNetworkInfo>([
  [1, { lzChainId: 101, endpointAddress: '0x66A71Dcef29A0fFBDBE3c6a460a3B5BC225Cd675' }],
  [137, { lzChainId: 109, endpointAddress: '0x3c2269811836af69497E5F486A85D7316753cf62' }],
  [56, { lzChainId: 102, endpointAddress: '0x3c2269811836af69497E5F486A85D7316753cf62' }],
  [43_114, { lzChainId: 106, endpointAddress: '0x3c2269811836af69497E5F486A85D7316753cf62' }],
  [42_161, { lzChainId: 110, endpointAddress: '0x3c2269811836af69497E5F486A85D7316753cf62' }],
  [10, { lzChainId: 111, endpointAddress: '0x3c2269811836af69497E5F486A85D7316753cf62' }],
  [250, { lzChainId: 112, endpointAddress: '0xb6319cC6c8c27A8F5dAF0dD3DF91EA35C4720dd7' }],
  [1101, { lzChainId: 158, endpointAddress: '0x9740FF91F1985D8d2B71494aE1A2f723bb3Ed9E4' }],
  [324, { lzChainId: 165, endpointAddress: '0x9b896c0e23220469C7AE69cb4BbAE391eAa4C8da' }],
]);

export const testnetChainIdToLZInfo = new Map<number, LZNetworkInfo>([
  [5, { lzChainId: 10_121, endpointAddress: '0xbfD2135BFfbb0B5378b56643c2Df8a87552Bfa23' }],
  [80_001, { lzChainId: 10_109, endpointAddress: '0xf69186dfBa60DdB133E91E9A4B5673624293d8F8' }],
  [97, { lzChainId: 10_102, endpointAddress: '0x6Fcb97553D41516Cb228ac03FdC8B9a0a9df04A1' }],
  [43_113, { lzChainId: 10_106, endpointAddress: '0x93f54D755A063cE7bB9e6Ac47Eccc8e33411d706' }],
  [421_613, { lzChainId: 10_143, endpointAddress: '0x6aB5Ae6822647046626e83ee6dB8187151E1d5ab' }],
  [420, { lzChainId: 10_132, endpointAddress: '0xae92d5aD7583AD66E49A0c67BAd18F6ba52dDDc1' }],
  [4002, { lzChainId: 10_112, endpointAddress: '0x7dcAD72640F835B0FA36EFD3D6d3ec902C7E5acf' }],
  [1442, { lzChainId: 10_158, endpointAddress: '0x6aB5Ae6822647046626e83ee6dB8187151E1d5ab' }],
  [280, { lzChainId: 10_165, endpointAddress: '0x093D2CF57f764f09C3c2Ac58a42A2601B8C79281' }],
]);
