#!/usr/bin/env bash

### Load env path
env_path=$1

### Load ENV from .env file
source $(pwd)/$env_path

### Now init node
docker run -v $(pwd)/genesis.json:/root/genesis.json \
  -v ${SIGNER_DATA_DIR}:/root/.ethereum -it ethereum/client-go:stable \
  init /root/genesis.json
