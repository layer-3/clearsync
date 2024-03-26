# Pls do not increase the version of the base image, it will break the contract deployments
# FROM ethereum/solc:0.8.19 AS compiler

FROM ethereum/client-go:v1.13.14

WORKDIR /app

RUN apk update 
RUN apk upgrade
# RUN apk add git
RUN apk add npm

# Install web3 to interact with the Ethereum network with Node.js
RUN npm i -g web3

# TODO:
# COPY --from=compiler /usr/bin/solc /usr/bin/solc

# Clone the Kernel smart contract repository and checkout the specific commit
# RUN git clone https://github.com/zerodevapp/kernel.git kernel-v2.2
# RUN git -C kernel-v2.2 checkout 807b75a4da6fea6311a3573bc8b8964a34074d94

# RUN git clone --branch v2.4 https://github.com/zerodevapp/kernel.git kernel-v2.4

# # Clone the account-abstraction repository on specific tag
# RUN git clone --depth 1 --branch v0.6.0 https://github.com/eth-infinitism/account-abstraction.git

# # deps
# RUN git clone --depth 1 --branch v4.9.6 https://github.com/OpenZeppelin/openzeppelin-contracts.git deps/@openzeppelin
# RUN git clone --depth 1 --branch main https://github.com/leekt/I4337.git deps/I4337-repo
# RUN git clone --depth 1 --branch v0.0.166 https://github.com/Vectorized/solady.git deps/solady-repo

# # applying remapping to the dependencies
# RUN cp -r deps/I4337-repo/src deps/I4337
# RUN cp -r deps/solady-repo/src deps/solady

# # Compile the smart contracts
# RUN /usr/bin/solc --include-path deps --base-path account-abstraction --optimize --bin --abi ./account-abstraction/contracts/core/EntryPoint.sol -o /app/build
# RUN /usr/bin/solc --include-path deps --base-path kernel-v2.2 --optimize --bin --abi ./kernel-v2.2/src/validator/ECDSAValidator.sol -o /app/build
# RUN /usr/bin/solc --include-path deps --base-path kernel-v2.2 --optimize --bin --abi --overwrite ./kernel-v2.2/src/Kernel.sol -o /app/build
# RUN /usr/bin/solc --include-path deps --base-path kernel-v2.2 --optimize --bin --abi --overwrite ./kernel-v2.2/src/factory/KernelFactory.sol -o /app/build
# RUN /usr/bin/solc --include-path deps --base-path kernel-v2.4 --optimize --bin --abi --overwrite ./kernel-v2.4/src/validator/SessionKeyValidator.sol -o /app/build

COPY entrypoint.sh .
COPY deploy-contracts.js .

RUN chmod +x entrypoint.sh

ENTRYPOINT ./entrypoint.sh
