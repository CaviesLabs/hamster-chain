#!/usr/bin/env bash

### Load env path
env_path=$1

### Load ENV from .env file
source $(pwd)/$env_path

### Now import account
docker run \
  -v $(pwd)/private_key.txt:/root/private_key.txt \
  -v ${DATA_DIR}:/root/.ethereum \
  -it ethereum/client-go:stable \
  account import /root/private_key.txt
