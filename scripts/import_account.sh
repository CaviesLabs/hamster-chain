#!/usr/bin/env bash

### Load env path and extra params
env_path=$1
private_key_path=$2
extra_params=$3

### Load ENV from .env file
source $(pwd)/$env_path

### Now import account
docker run \
  -v ${private_key_path}:/root/private_key.txt \
  -v ${DATA_DIR}:/root/.ethereum \
  -it ethereum/client-go:v1.10.23 \
  ${extra_params} \
  account import /root/private_key.txt
