#!/bin/sh
set -m

nohup geth --dev --http --http.api=eth,web3,net,personal --allow-insecure-unlock --http.addr=0.0.0.0 --http.corsdomain='*' --http.vhosts='*' &

while ! nc -z localhost 8545; do
  sleep 1
done

NODE_PATH=$(npm root -g) node /app/deploy-contracts.js

tail -f -n 300 nohup.out
