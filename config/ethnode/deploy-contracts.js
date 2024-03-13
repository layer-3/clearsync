//web.archive.org/web/20180401231221/https://tokenmarket.net/blog/creating-ethereum-smart-contract-transactions-in-client-side-javascript/

let fs = require('fs');
var { Web3 } = require('web3');

var web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:8545'));
let eth = web3.eth;

// TODO: move to the env
const deployer = '0x25ba87CA70739Bc8448D018Ad4A11F35Ea5a2DF9';
const deployerPk = '6aa35771f25b5098020350399171952bdaafd8b381eb777577befd5ee995a122';

async function main() {
  await eth.personal.importRawKey(deployerPk, 'password');
  await web3.eth.personal.unlockAccount(deployer, 'password', 600);

  let coinbase = await eth.getCoinbase();

  await eth.sendTransaction({
    from: coinbase,
    to: deployer,
    value: web3.utils.toWei(50, 'ether'),
  });

  let entryPointAddress = await DeployEntryPoint();
  await DeployKernelECDSAValidator();
  let kernelAddress = await DeployKernel(entryPointAddress);
  await DeployKernelFactory(deployer, entryPointAddress, kernelAddress);
  await DeploySessionKeyValidator();
}

main();

// 0x07bd68335Ff013481b0fED98c190EaeB36e52b3D
async function DeployEntryPoint() {
  let abi = JSON.parse(fs.readFileSync('/app/build/EntryPoint.abi'));
  let bytecode = '0x' + fs.readFileSync('/app/build/EntryPoint.bin', 'utf8');

  let contract = new eth.Contract(abi);
  let gas = await eth.estimateGas({ data: bytecode });

  let receipt = await contract.deploy({ data: bytecode }).send({
    from: deployer,
    gas,
    gasPrice: '30000000000',
  });

  return receipt.options.address;
}

// 0x0E3c0cb9F2Ae0053f2b236b698C2028112b333a7
async function DeployKernelECDSAValidator() {
  let abi = JSON.parse(fs.readFileSync('/app/build/ECDSAValidator.abi'));
  let bytecode = '0x' + fs.readFileSync('/app/build/ECDSAValidator.bin', 'utf8');

  let contract = new eth.Contract(abi);
  let gas = await eth.estimateGas({ data: bytecode });

  await contract.deploy({ data: bytecode }).send({
    from: deployer,
    gas,
    gasPrice: '30000000000',
  });
}

// 0x8Bdf2ceE549101447fA141fFfc9f6e3B2BE8BBF2
async function DeployKernel(entryPointAddress) {
  let abi = JSON.parse(fs.readFileSync('/app/build/Kernel.abi'));
  let bytecode = '0x' + fs.readFileSync('/app/build/Kernel.bin', 'utf8');

  let contract = new eth.Contract(abi);
  let gas = await contract.deploy({ data: bytecode, arguments: [entryPointAddress] }).estimateGas();

  let receipt = await contract.deploy({ data: bytecode, arguments: [entryPointAddress] }).send({
    from: deployer,
    gas,
    gasPrice: '30000000000',
  });

  return receipt.options.address;
}

// 0x9CBDd0D809f3490d52E3609044D4cf78f4df3a5f
async function DeployKernelFactory(deployerAddress, entryPointAddress, kernelAddress) {
  let abi = JSON.parse(fs.readFileSync('/app/build/KernelFactory.abi'));
  let bytecode = '0x' + fs.readFileSync('/app/build/KernelFactory.bin', 'utf8');

  let contract = new eth.Contract(abi);
  let gas = await contract
    .deploy({ data: bytecode, arguments: [deployerAddress, entryPointAddress] })
    .estimateGas();

  let receipt = await contract
    .deploy({ data: bytecode, arguments: [deployerAddress, entryPointAddress] })
    .send({
      from: deployer,
      gas,
      gasPrice: '30000000000',
    });

  let data = await contract.methods.setImplementation(kernelAddress, true).encodeABI();

  gas = await web3.eth.estimateGas({
    from: deployer,
    to: receipt.options.address,
    data,
  });

  await web3.eth.sendTransaction({
    from: deployer,
    to: receipt.options.address,
    data,
    gas,
  });
}

// 0x18D865C12377cf6d106953b83eE1b5bA7c3073Ac
async function DeploySessionKeyValidator() {
  let abi = JSON.parse(fs.readFileSync('/app/build/SessionKeyValidator.abi'));
  let bytecode = '0x' + fs.readFileSync('/app/build/SessionKeyValidator.bin', 'utf8');

  let gas = await eth.estimateGas({ data: bytecode });
  let contract = new eth.Contract(abi);

  await contract.deploy({ data: bytecode }).send({
    from: deployer,
    gas,
    gasPrice: '30000000000',
  });
}
