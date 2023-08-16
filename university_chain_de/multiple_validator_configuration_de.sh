#!/bin/bash

echo ""
echo "Command: sudo mkdir -p elements/val-tum"
echo ""
sudo mkdir -p elements/val-tum
echo ""
echo "Command: sudo mkdir -p elements/val-humboldt"
echo ""
sudo mkdir -p elements/val-humboldt
echo ""
echo "Command: echo -e val-tum'\n'val-humboldt | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_de university_chain_ded_i init university_chain_de"
echo ""
echo -e val-tum'\n'val-humboldt | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_de university_chain_ded_i init university_chain_de
echo ""
echo "Command: echo -e val-tum'\n'val-humboldt | xargs -I {} sudo cp -r ./dir_keys/keyring-test $(pwd)/elements/{}/"
echo ""
echo -e val-tum'\n'val-humboldt | xargs -I {} sudo cp -r ./dir_keys/keyring-test $(pwd)/elements/{}/
echo ""
#Copy manually the address and balances information in the genesis file
echo "Command: Accounts val-tum"
jq --argjson info "$(<genesis_dir/accounts.json)" '.app_state.auth.accounts = $info' elements/val-tum/config/genesis.json > genesis.tmp && sudo mv genesis.tmp elements/val-tum/config/genesis.json
echo "Accounts Completed val-tum"
echo "Command: Balances val-tum"
jq --argjson info "$(<genesis_dir/balances.json)" '.app_state.bank.balances = $info' elements/val-tum/config/genesis.json > genesis.tmp && sudo mv genesis.tmp elements/val-tum/config/genesis.json
echo "Balances Completed val-tum"
echo ""
echo "Command: Accounts val-humboldt"
jq --argjson info "$(<genesis_dir/accounts.json)" '.app_state.auth.accounts = $info' elements/val-humboldt/config/genesis.json > genesis.tmp && sudo mv genesis.tmp elements/val-humboldt/config/genesis.json
echo "Accounts Completed val-humboldt"
echo "Command: Balances val-humboldt"
jq --argjson info "$(<genesis_dir/balances.json)" '.app_state.bank.balances = $info' elements/val-humboldt/config/genesis.json > genesis.tmp && sudo mv genesis.tmp elements/val-humboldt/config/genesis.json
echo "Balances Completed val-humboldt"
echo ""
echo "Command: echo -e val-tum'\n'val-humboldt | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_de --entrypoint sed university_chain_ded_i -Ei 's/^chain-id = .*$/chain-id = "university_chain_de"/g' /root/.university_chain_de/config/client.toml"
echo ""
echo -e val-tum'\n'val-humboldt | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_de --entrypoint sed university_chain_ded_i -Ei 's/^chain-id = .*$/chain-id = "university_chain_de"/g' /root/.university_chain_de/config/client.toml
echo ""
echo "Command: echo -e val-tum'\n'val-humboldt | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_de --entrypoint sed university_chain_ded_i -Ei 's/^enable-unsafe-cors = false/enable-unsafe-cors = true/g' /root/.university_chain_de/config/app.toml"
echo ""
echo -e val-tum'\n'val-humboldt | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_de --entrypoint sed university_chain_ded_i -Ei 's/^enable-unsafe-cors = false/enable-unsafe-cors = true/g' /root/.university_chain_de/config/app.toml
echo ""
echo "Command: echo -e val-tum'\n'val-humboldt | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_de --entrypoint sed university_chain_ded_i -Ei 's/^enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' /root/.university_chain_de/config/app.toml"
echo ""
echo -e val-tum'\n'val-humboldt | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_de --entrypoint sed university_chain_ded_i -Ei 's/^enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' /root/.university_chain_de/config/app.toml
echo ""
echo "Command: echo -e val-tum'\n'val-humboldt | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_de --entrypoint sed university_chain_ded_i -Ei 's/^laddr = "tcp:127.0.0.1:26657"/laddr = "tcp:0.0.0.0:26657"/g' /root/.university_chain_de/config/config.toml"
echo ""
echo -e val-tum'\n'val-humboldt | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_de --entrypoint sed university_chain_ded_i -Ei 's/^laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' /root/.university_chain_de/config/config.toml
echo ""
echo "Command: echo -e val-tum'\n'val-humboldt | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_de --entrypoint sed university_chain_ded_i -Ei 's/^cors_allowed_origins = \[\]/cors_allowed_origins = \["*"\]/g' /root/.university_chain_de/config/config.toml"
echo ""
echo -e val-tum'\n'val-humboldt | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_de --entrypoint sed university_chain_ded_i -Ei 's/^cors_allowed_origins = \[\]/cors_allowed_origins = \["\*"\]/g' /root/.university_chain_de/config/config.toml
echo ""
echo "mv elements/val-tum/config/genesis.json elements/val-humboldt/config/"
sudo mv elements/val-tum/config/genesis.json elements/val-humboldt/config/
echo ""
echo "Command: sudo docker run --rm -i -v $(pwd)/elements/val-humboldt:/root/.university_chain_de university_chain_ded_i gentx "Humboldt University" 40000000stake --keyring-backend test --account-number 0 --sequence 0 --chain-id university_chain_de --gas 1000000 --gas-prices 0.1stake"
echo ""
sudo docker run --rm -i -v $(pwd)/elements/val-humboldt:/root/.university_chain_de university_chain_ded_i gentx "Humboldt University" 40000000stake --keyring-backend test --account-number 0 --sequence 0 --chain-id university_chain_de --gas 1000000 --gas-prices 0.1stake
echo ""
echo "Command: sudo mv elements/val-humboldt/config/genesis.json elements/val-tum/config/"
echo ""
sudo mv elements/val-humboldt/config/genesis.json elements/val-tum/config/
echo ""
echo "Command: sudo docker run --rm -i -v $(pwd)/elements/val-tum:/root/.university_chain_de university_chain_ded_i gentx TUM 40000000stake --keyring-backend test --account-number 0 --sequence 0 --chain-id university_chain_de --gas 1000000 --gas-prices 0.1stake"
echo ""
sudo docker run --rm -i -v $(pwd)/elements/val-tum:/root/.university_chain_de university_chain_ded_i gentx TUM 40000000stake --keyring-backend test --account-number 0 --sequence 0 --chain-id university_chain_de --gas 1000000 --gas-prices 0.1stake
echo ""
echo "Command: sudo cp elements/val-humboldt/config/gentx/gentx-* elements/val-tum/config/gentx"
echo ""
# copy manually elements/val-humboldt/config/gentx/* elements/val-tum/config/gentx
sudo cp -r elements/val-humboldt/config/gentx/. elements/val-tum/config/gentx/
echo ""
echo "Command: sudo docker run --rm -it -v $(pwd)/elements/val-tum:/root/.university_chain_de university_chain_ded_i collect-gentxs"
echo ""
sudo docker run --rm -it -v $(pwd)/elements/val-tum:/root/.university_chain_de university_chain_ded_i collect-gentxs
echo ""
echo "Command: sudo docker run --rm -it -v $(pwd)/elements/val-tum:/root/.university_chain_de university_chain_ded_i validate-genesis"
echo ""
sudo docker run --rm -it -v $(pwd)/elements/val-tum:/root/.university_chain_de university_chain_ded_i validate-genesis 
echo ""
echo "Command: sudo cp elements/val-tum/config/genesis.json elements/val-humboldt/config"
echo ""
sudo cp elements/val-tum/config/genesis.json elements/val-humboldt/config

#sudo docker run --rm -i -v $(pwd)/elements/val-tum:/root/.university_chain_de university_chain_ded_i tendermint show-node-id
#Enter the result in config.toml in persistent_peers: @ with the node name and : with 26656 as the port.

node_id_tum=$(sudo docker run --rm -i -v $(pwd)/elements/val-tum:/root/.university_chain_de university_chain_ded_i tendermint show-node-id)
persistent_peer_tum=$node_id_tum"@val-tum:26656"

echo ""
echo "Command: val-humboldt persistent peer"
echo ""

sudo docker run --rm -i -v $(pwd)/elements/val-humboldt:/root/.university_chain_de --entrypoint sed university_chain_ded_i -Ei 's/^persistent_peers = ""/persistent_peers = "'$persistent_peer_tum'"/g' /root/.university_chain_de/config/config.toml

#sudo docker run --rm -i -v $(pwd)/elements/val-humboldt:/root/.university_chain_de university_chain_ded_i tendermint show-node-id
#Enter the result in config.toml in persistent_peers: @ with the node name and : with 26656 as the port.


node_id_humboldt=$(sudo docker run --rm -i -v $(pwd)/elements/val-humboldt:/root/.university_chain_de university_chain_ded_i tendermint show-node-id)
persistent_peer_humboldt=$node_id_humboldt"@val-humboldt:26656"

echo ""
echo "Command: val-tum persistent peer"
echo ""

sudo docker run --rm -i -v $(pwd)/elements/val-tum:/root/.university_chain_de --entrypoint sed university_chain_ded_i -Ei 's/^persistent_peers = ""/persistent_peers = "'$persistent_peer_humboldt'"/g' /root/.university_chain_de/config/config.toml

