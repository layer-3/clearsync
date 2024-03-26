#!/bin/sh

nohup geth --dev --http --ws --dev.period=1 --http.api=eth,web3,net,personal --allow-insecure-unlock --http.addr=0.0.0.0 --http.corsdomain='*' --http.vhosts='*' --ws.addr=0.0.0.0 --ws.origins='*' &

while ! nc -z localhost 8545; do
  sleep 1
done

NODE_PATH=$(npm root -g) node /app/deploy-contracts.js

while nc -z localhost 8545; do
  sleep 1
done
