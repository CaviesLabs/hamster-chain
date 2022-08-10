#!/usr/bin/env bash

### Load env path and extra params
env_path=$1
extra_params=$2

### Load ENV from .env file
source $(pwd)/$env_path

### Now import account
docker run \
  -v $(pwd)/private_key.txt:/root/private_key.txt \
  -v ${DATA_DIR}:/root/.ethereum \
  -it ethereum/client-go:v1.10.21 \
  ${extra_params} \
  account import /root/private_key.txt
