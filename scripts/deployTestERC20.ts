const artifactFile = require('../artifacts/contracts/test/TestERC20.sol/TestERC20.json');
const API_KEY = process.env['DEFENDER_API_KEY'];
const API_SECRET = process.env['DEFENDER_API_SECRET'];
const { Defender } = require('@openzeppelin/defender-sdk');
const client = new Defender({ apiKey: API_KEY, apiSecret: API_SECRET });


const main = async (): Promise<any> => {
  return client.deploy.deployContract({
    contractName: 'TestERC20',
    contractPath: 'contracts/test/TestERC20.sol',
    artifactPayload: JSON.stringify(artifactFile),
    network: 'goerli',
    licenseType: 'MIT',
    constructorInputs: ['TEST', 'TEST', '1000000000000000000'],
    verifySourceCode: false
  });
}

main().then((deployment) => {
  console.log(deployment);
}).catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
