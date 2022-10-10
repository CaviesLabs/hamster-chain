#!/usr/bin/env bash

### Load env path and extra params
env_path=$1
extra_params=$2

### Load ENV from .env file
source $(pwd)/$env_path

### Run init genesis
docker run -v $(pwd)/genesis-data/:/root/data/ \
  -it ethereum/client-go:alltools-v1.10.23 \
  puppeth ${extra_params} init
