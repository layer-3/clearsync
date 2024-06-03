let fs = require('fs');
var { Web3 } = require('web3');

var web3 = new Web3(new Web3.providers.HttpProvider('http://0.0.0.0:8545'));
let eth = web3.eth;

let deployerAccount;

const deployer = process.env.DEPLOYER_ADDRESS;
const deployerPk = process.env.DEPLOYER_PK;
const expectedEntryPointAddress = process.env.ENTRY_POINT_ADDRESS;
const expectedKernelECDSAValidatorAddress = process.env.KERNEL_ECDSA_VALIDATOR_ADDRESS;
const expectedKernelAddress = process.env.KERNEL_ADDRESS;
const expectedKernelFactoryAddress = process.env.KERNEL_FACTORY_ADDRESS;
const expectedSessionKeyValidatorAddress = process.env.SESSION_KEY_VALIDATOR_ADDRESS;

async function main() {
  deployerAccount = web3.eth.accounts.privateKeyToAccount('0x' + deployerPk);

  let entryPointAddress = await DeployEntryPoint();
  await DeployKernelECDSAValidator();
  let kernelAddress = await DeployKernel(entryPointAddress);
  await DeployKernelFactory(deployer, entryPointAddress, kernelAddress);
  await DeploySessionKeyValidator();
}

main();

async function DeployEntryPoint() {
  let { abi, bin } = JSON.parse(fs.readFileSync('/app/contracts/EntryPoint.json'));

  let contractAddress = await deployContract(abi, bin);

  assert(
    cmpAddresses(contractAddress, expectedEntryPointAddress),
    `Get unexpected EntryPoint address, expected: ${expectedEntryPointAddress}, got: ${contractAddress}`,
  );

  return contractAddress;
}

async function DeployKernelECDSAValidator() {
  let { abi, bin } = JSON.parse(fs.readFileSync('/app/contracts/KernelECDSAValidator.json'));

  let contractAddress = await deployContract(abi, bin);

  assert(
    cmpAddresses(contractAddress, expectedKernelECDSAValidatorAddress),
    `Get unexpected ECDSAValidator address, expected: ${expectedKernelECDSAValidatorAddress}, got: ${contractAddress}`,
  );
}

async function DeployKernel(entryPointAddress) {
  let { abi, bin } = JSON.parse(fs.readFileSync('/app/contracts/Kernel.json'));

  let contractAddress = await deployContract(abi, bin, [entryPointAddress]);

  assert(
    cmpAddresses(contractAddress, expectedKernelAddress),
    `Get unexpected Kernel address, expected: ${expectedKernelAddress}, got: ${contractAddress}`,
  );

  return contractAddress;
}

async function DeployKernelFactory(deployerAddress, entryPointAddress, kernelAddress) {
  let { abi, bin } = JSON.parse(fs.readFileSync('/app/contracts/KernelFactory.json'));

  let contractAddress = await deployContract(abi, bin, [deployerAddress, entryPointAddress]);

  assert(
    cmpAddresses(contractAddress, expectedKernelFactoryAddress),
    `Get unexpected KernelFactory address, expected: ${expectedKernelFactoryAddress}, got: ${contractAddress}`,
  );

  let contract = new eth.Contract(abi);
  let data = await contract.methods.setImplementation(kernelAddress, true).encodeABI();

  let gas = await eth.estimateGas({ from: deployer, to: contractAddress, data });

  let signedTx = await deployerAccount.signTransaction({
    from: deployer,
    to: contractAddress,
    gas,
    gasPrice: '30000000000',
    data,
  });

  await eth.sendSignedTransaction(signedTx.rawTransaction);
}

async function DeploySessionKeyValidator() {
  let { abi, bin } = JSON.parse(fs.readFileSync('/app/contracts/SessionKeyValidator.json'));
  
  let contractAddress = await deployContract(abi, bin);

  assert(
    cmpAddresses(contractAddress, expectedSessionKeyValidatorAddress),
    `Get unexpected SessionKeyValidator address, expected: ${expectedSessionKeyValidatorAddress}, got: ${contractAddress}`,
  );
}

function cmpAddresses(adr1, adr2) {
  return adr1.toLowerCase() === adr2.toLowerCase();
}

function assert(value, msg) {
  if (!value) {
    throw new Error(msg);
  }
}

async function deployContract(abi, bin, arguments = undefined) {
  let contract = new eth.Contract(abi);
  let deployData = contract
    .deploy({
      data: bin,
      arguments,
    })
    .encodeABI();

  let gas = await eth.estimateGas({ data: deployData });

  let signedTx = await deployerAccount.signTransaction({
    from: deployer,
    gas,
    gasPrice: '30000000000',
    data: deployData,
  });

  let receipt = await eth.sendSignedTransaction(signedTx.rawTransaction);

  return receipt.contractAddress;
}
