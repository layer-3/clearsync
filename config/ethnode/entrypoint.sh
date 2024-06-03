#!/bin/sh

nohup anvil --host "0.0.0.0" & 

while ! nc -z "0.0.0.0" 8545; do
  sleep 1
done

NODE_PATH=$(npm root -g) node /app/deploy-contracts.js

while nc -z "0.0.0.0" 8545; do
  sleep 1
done
