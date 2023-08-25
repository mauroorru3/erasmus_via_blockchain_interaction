#!/bin/bash

chain_admin=$(sudo docker run --rm -i -v $(pwd)/hub/elements/val-hub:/root/.hub  hubd_i keys --keyring-backend test show "Admin Hub" --address)

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/hub/elements/val-hub:/root/.hub --network university_chain-prod_net-public hubd_i tx hub configure-chain --from "$chain_admin" --keyring-backend test --gas auto --chain-id hub --yes --node tcp://val-hub:26657"
echo ""
sudo docker run --rm -i -v $(pwd)/hub/elements/val-hub:/root/.hub --network university_chain-prod_net-public hubd_i tx hub configure-chain --from "$chain_admin" --keyring-backend test --gas auto --chain-id hub --yes --node "tcp://val-hub:26657"


