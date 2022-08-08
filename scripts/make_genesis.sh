#!/usr/bin/env bash

### Load env path
env_path=$1

### Load ENV from .env file
source $(pwd)/$env_path

### Run init genesis
docker run -v $(pwd)/genesis-data/:/root/data/ \
  -it ethereum/client-go:alltools-stable \
  puppeth init
