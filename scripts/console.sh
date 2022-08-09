#!/usr/bin/env bash

### Load env path
env_path=$1

### Load ENV from .env file
source $(pwd)/$env_path

### Now start javascript console
docker run \
  -v ${DATA_DIR}:/root/.ethereum \
  -it ethereum/client-go:stable \
  attach /root/.ethereum/geth.ipc