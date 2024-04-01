FROM ghcr.io/foundry-rs/foundry:nightly-156cb1396b7076c6f9cb56f3719f8c90f7f52064

WORKDIR /app

RUN apk update 
RUN apk upgrade
RUN apk add npm

# Install web3 to interact with the Ethereum network with Node.js
RUN npm i -g web3

COPY contracts contracts
COPY entrypoint.sh .
COPY deploy-contracts.js .

RUN chmod +x entrypoint.sh

ENTRYPOINT ./entrypoint.sh
