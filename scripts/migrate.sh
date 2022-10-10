#!/usr/bin/env bash

### Load env path and extra params
env_path=$1

### Load ENV from .env file
source $(pwd)/$env_path

### Now start javascript console
docker run \
  -v ${DATA_DIR}:/root/.ethereum \
  -it ethereum/client-go:v1.10.23 \
  db freezer-migrate

