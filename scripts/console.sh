#!/usr/bin/env bash

### Load env path and extra params
env_path=$1
extra_params=$2

### Load ENV from .env file
source $(pwd)/$env_path

### Now start javascript console
docker run \
  -v ${DATA_DIR}:/root/.ethereum \
  -it ethereum/client-go:stable \
  ${extra_params} \
  attach /root/.ethereum/geth.ipc