#!/usr/bin/env bash

### Load env path
env_path=$1
extra_params=$2

### Load ENV from .env file
source $(pwd)/$env_path

### Now init node
docker run -v $(pwd)/genesis.json:/root/genesis.json \
  -v ${DATA_DIR}:/root/.ethereum -it ethereum/client-go:v1.10.23 \
  ${extra_params} \
  init /root/genesis.json
