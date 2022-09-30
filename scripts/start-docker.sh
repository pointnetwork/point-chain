#!/bin/bash

KEY="mykey"
CHAINID="point_10687-1"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t point-datadir.XXXXX)

echo "create and add new keys"
./pointd keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init Point with moniker=$MONIKER and chain-id=$CHAINID"
./pointd init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./pointd add-genesis-account \
"$(./pointd keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000apoint,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./pointd gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./pointd collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./pointd validate-genesis --home $DATA_DIR

echo "starting point node $i in background ..."
./pointd start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started point node"
tail -f /dev/null