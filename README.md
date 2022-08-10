# Private Zero-fee POA Geth Network

This guide to start a private POA Geth network with 2 nodes: 1 signer (to seal new blocks) and 1 rpc node (for dapp integration).

To avoid re-org issues we should use 1 signer instead of multiple signers.

The network also accepts zero-fee transactions to be mined.

## Prerequisites

- Understand the basis of Geth private network, refer this [docs](https://geth.ethereum.org/docs/interface/private-network)
- Install docker engine (latest version)
- Install docker compose (latest version)
- Install nginx instance (latest version)
- Geth image: `ethereum/client-go:v1.10.21`

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

### Step 5: Set gasprice = 0 and get the node id

For some reasons, the option `--miner.gasprice=0` seems not to be working with `ethereum/client-go:v1.10.21`. So we have to manually adjust gas price = 0 within the ipc console.

Run the command below to access signer ipc console.

```bash
bash scripts/console.sh env.signer
```

Then type in the console to set gas price to zero

```ts
miner.setGasPrice(0)
```

Then use the command below to get the node id

```ts
admin.nodeInfo
```

Below is the sample output

```bash
> private-poa-geth git:(master) bash scripts/console.sh .env.signer

Welcome to the Geth JavaScript console!

instance: Geth/v1.10.21-stable-67109427/linux-amd64/go1.18.4
coinbase: 0xac118f16238b5aba99f3c9dddb74d3e635136fec
at block: 935 (Wed Aug 10 2022 16:20:08 GMT+0000 (UTC))
 datadir: /root/.ethereum
 modules: admin:1.0 clique:1.0 debug:1.0 engine:1.0 eth:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0

To exit, press ctrl-d or type exit
> miner.setGasPrice(0)
true
> admin.nodeInfo
{
  enode: "enode://289688c271bbf9326d5bc1ac977ac4303bd22ef0e19de00fb7122f73b03aec71f26009029deed3e2d6653f92cb1bf197a37027834775313828e27ea217bc8a6d@10.116.0.2:30303",
  <truncated-output>
}
>
```

### Step 6: Start the RPC node

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