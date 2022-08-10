# Private POA Geth Network

This guide to start a private POA Geth network with 2 nodes: 1 signer (to seal new blocks) and 1 rpc node (for dapp integration).

To avoid re-org issues we should use 1 signer instead of multiple signers.

## Prerequisites

- Understand the basis of Geth private network, refer this [docs](https://geth.ethereum.org/docs/interface/private-network)
- Install docker engine (latest version)
- Install docker compose (latest version)
- Install nginx instance (latest version)

## Deploy

### Step 1: Prepare environment variables

Duplicate `.env.rpc.example` and `.env.signer.example` into `.env.rpc` and `.env.signer`.

Prepare `.env.rpc` and `.env.signer` with appropriate values.

For `.env.signer` leave the `SIGNER_BOOT_NODES` empty for now.

### Step 2: Initialize genesis.json

Prepare `genesis.json` by duplicating `genesis.example.json`. Replace signer address in `extradata` section.

Make sure you initialize for both rpc node and signer node.

Then run the commands below:

```bash
bash scripts/init_node.sh .env.rpc "--gcmode=archive"
bash scripts/init_node.sh .env.signer
```

### Step 3: Import signer account to signer node

Prepare a `private_key.txt` with plain-text private key of `SIGNER_ADDRESS`.

Also prepare a `password.txt` with plain-text password when you intend to secure the address once it's imported to the signer node, also fill `SIGNER_PASSWORD_PATH` with the `password.txt` file path.



Then run the command:

```bash
bash scripts/import_account.sh .env.signer
```

### Step 4: Start a signer node

Run the command below to start the signer node.

```bash
docker compose --env-file .env.signer -f compose/signer.docker-compose.yml up -d --force-recreate
```

Make sure you grab the node ID by inspecting the signer log. The node id should start with `enode://` prefix.

### Step 5: Start the RPC node

Now you grabbed the signer node id, place it in `.env.rpc` at `RPC_BOOT_NODES`

Run the command be low to start the rpc node.

```bash
docker compose --env-file .env.rpc -f compose/rpc.docker-compose.yml up -d --force-recreate
```

The rpc endpoint will be live at

- http://localhost:8545/ (http)
- http://localhost:8546/ (websocket)

## LICENSE

This repo is MIT Licensed.

## Contacts

Please feel free to contact [khang@cavies.xyz](mailto:khang@cavies.xyz) or [dev@cavies.xyz](mailto:dev@cavies.xyz) if you have any inquiries.