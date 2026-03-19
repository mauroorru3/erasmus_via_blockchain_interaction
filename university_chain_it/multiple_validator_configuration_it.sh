#!/bin/bash

echo ""
echo "Command: sudo mkdir -p elements/val-unipi"
echo ""
sudo mkdir -p elements/val-unipi
echo ""
echo "Command: sudo mkdir -p elements/val-uniroma1"
echo ""
sudo mkdir -p elements/val-uniroma1
echo ""
echo "Command: echo -e val-unipi'\n'val-uniroma1 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_it university_chain_itd_i init university_chain_it"
echo ""
echo -e val-unipi'\n'val-uniroma1 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_it university_chain_itd_i init university_chain_it
echo ""
echo "Command: echo -e val-unipi'\n'val-uniroma1 | xargs -I {} sudo cp -r ./dir_keys/keyring-test $(pwd)/elements/{}/"
echo ""
echo -e val-unipi'\n'val-uniroma1 | xargs -I {} sudo cp -r ./dir_keys/keyring-test $(pwd)/elements/{}/
echo ""
#Copy manually the address and balances information in the genesis file
echo "Command: Accounts val-unipi"
jq --argjson info "$(<genesis_dir/accounts.json)" '.app_state.auth.accounts = $info' elements/val-unipi/config/genesis.json > genesis.tmp && sudo mv genesis.tmp elements/val-unipi/config/genesis.json
echo "Accounts Completed val-unipi"
echo "Command: Balances val-unipi"
jq --argjson info "$(<genesis_dir/balances.json)" '.app_state.bank.balances = $info' elements/val-unipi/config/genesis.json > genesis.tmp && sudo mv genesis.tmp elements/val-unipi/config/genesis.json
echo "Balances Completed val-unipi"
echo ""
echo "Command: Accounts val-uniroma1"
jq --argjson info "$(<genesis_dir/accounts.json)" '.app_state.auth.accounts = $info' elements/val-uniroma1/config/genesis.json > genesis.tmp && sudo mv genesis.tmp elements/val-uniroma1/config/genesis.json
echo "Accounts Completed val-uniroma1"
echo "Command: Balances val-uniroma1"
jq --argjson info "$(<genesis_dir/balances.json)" '.app_state.bank.balances = $info' elements/val-uniroma1/config/genesis.json > genesis.tmp && sudo mv genesis.tmp elements/val-uniroma1/config/genesis.json
echo "Balances Completed val-uniroma1"
echo ""
echo "Command: echo -e val-unipi'\n'val-uniroma1 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_it --entrypoint sed university_chain_itd_i -Ei 's/^chain-id = .*$/chain-id = "university_chain_it"/g' /root/.university_chain_it/config/client.toml"
echo ""
echo -e val-unipi'\n'val-uniroma1 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_it --entrypoint sed university_chain_itd_i -Ei 's/^chain-id = .*$/chain-id = "university_chain_it"/g' /root/.university_chain_it/config/client.toml
echo ""
echo "Command: echo -e val-unipi'\n'val-uniroma1 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_it --entrypoint sed university_chain_itd_i -Ei 's/^enable-unsafe-cors = false/enable-unsafe-cors = true/g' /root/.university_chain_it/config/app.toml"
echo ""
echo -e val-unipi'\n'val-uniroma1 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_it --entrypoint sed university_chain_itd_i -Ei 's/^enable-unsafe-cors = false/enable-unsafe-cors = true/g' /root/.university_chain_it/config/app.toml
echo ""
echo "Command: echo -e val-unipi'\n'val-uniroma1 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_it --entrypoint sed university_chain_itd_i -Ei 's/^enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' /root/.university_chain_it/config/app.toml"
echo ""
echo -e val-unipi'\n'val-uniroma1 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_it --entrypoint sed university_chain_itd_i -Ei 's/^enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' /root/.university_chain_it/config/app.toml
echo ""
echo "Command: echo -e val-unipi'\n'val-uniroma1 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_it --entrypoint sed university_chain_itd_i -Ei 's/^laddr = "tcp:127.0.0.1:26657"/laddr = "tcp:0.0.0.0:26657"/g' /root/.university_chain_it/config/config.toml"
echo ""
echo -e val-unipi'\n'val-uniroma1 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_it --entrypoint sed university_chain_itd_i -Ei 's/^laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' /root/.university_chain_it/config/config.toml
echo ""
echo "Command: echo -e val-unipi'\n'val-uniroma1 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_it --entrypoint sed university_chain_itd_i -Ei 's/^cors_allowed_origins = \[\]/cors_allowed_origins = \["*"\]/g' /root/.university_chain_it/config/config.toml"
echo ""
echo -e val-unipi'\n'val-uniroma1 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_it --entrypoint sed university_chain_itd_i -Ei 's/^cors_allowed_origins = \[\]/cors_allowed_origins = \["\*"\]/g' /root/.university_chain_it/config/config.toml
echo ""
echo "mv elements/val-unipi/config/genesis.json elements/val-uniroma1/config/"
sudo mv elements/val-unipi/config/genesis.json elements/val-uniroma1/config/
echo ""
echo "Command: sudo docker run --rm -i -v $(pwd)/elements/val-uniroma1:/root/.university_chain_it university_chain_itd_i gentx "Humboldt University" 40000000stake --keyring-backend test --account-number 0 --sequence 0 --chain-id university_chain_it --gas 1000000 --gas-prices 0.1stake"
echo ""
sudo docker run --rm -i -v $(pwd)/elements/val-uniroma1:/root/.university_chain_it university_chain_itd_i gentx "Humboldt University" 40000000stake --keyring-backend test --account-number 0 --sequence 0 --chain-id university_chain_it --gas 1000000 --gas-prices 0.1stake
echo ""
echo "Command: sudo mv elements/val-uniroma1/config/genesis.json elements/val-unipi/config/"
echo ""
sudo mv elements/val-uniroma1/config/genesis.json elements/val-unipi/config/
echo ""
echo "Command: sudo docker run --rm -i -v $(pwd)/elements/val-unipi:/root/.university_chain_it university_chain_itd_i gentx TUM 40000000stake --keyring-backend test --account-number 0 --sequence 0 --chain-id university_chain_it --gas 1000000 --gas-prices 0.1stake"
echo ""
sudo docker run --rm -i -v $(pwd)/elements/val-unipi:/root/.university_chain_it university_chain_itd_i gentx TUM 40000000stake --keyring-backend test --account-number 0 --sequence 0 --chain-id university_chain_it --gas 1000000 --gas-prices 0.1stake
echo ""
echo "Command: sudo cp elements/val-uniroma1/config/gentx/gentx-* elements/val-unipi/config/gentx"
echo ""
# copy manually elements/val-uniroma1/config/gentx/* elements/val-unipi/config/gentx
sudo cp -r elements/val-uniroma1/config/gentx/. elements/val-unipi/config/gentx/
echo ""
echo "Command: sudo docker run --rm -it -v $(pwd)/elements/val-unipi:/root/.university_chain_it university_chain_itd_i collect-gentxs"
echo ""
sudo docker run --rm -it -v $(pwd)/elements/val-unipi:/root/.university_chain_it university_chain_itd_i collect-gentxs
echo ""
echo "Command: sudo docker run --rm -it -v $(pwd)/elements/val-unipi:/root/.university_chain_it university_chain_itd_i validate-genesis"
echo ""
sudo docker run --rm -it -v $(pwd)/elements/val-unipi:/root/.university_chain_it university_chain_itd_i validate-genesis 
echo ""
echo "Command: sudo cp elements/val-unipi/config/genesis.json elements/val-uniroma1/config"
echo ""
sudo cp elements/val-unipi/config/genesis.json elements/val-uniroma1/config

#sudo docker run --rm -i -v $(pwd)/elements/val-unipi:/root/.university_chain_it university_chain_itd_i tendermint show-node-id
#Enter the result in config.toml in persistent_peers: @ with the node name and : with 26656 as the port.

node_id_tum=$(sudo docker run --rm -i -v $(pwd)/elements/val-unipi:/root/.university_chain_it university_chain_itd_i tendermint show-node-id)
persistent_peer_tum=$node_id_tum"@val-unipi:26656"

echo ""
echo "Command: val-uniroma1 persistent peer"
echo ""

sudo docker run --rm -i -v $(pwd)/elements/val-uniroma1:/root/.university_chain_it --entrypoint sed university_chain_itd_i -Ei 's/^persistent_peers = ""/persistent_peers = "'$persistent_peer_tum'"/g' /root/.university_chain_it/config/config.toml

#sudo docker run --rm -i -v $(pwd)/elements/val-uniroma1:/root/.university_chain_it university_chain_itd_i tendermint show-node-id
#Enter the result in config.toml in persistent_peers: @ with the node name and : with 26656 as the port.


node_id_humboldt=$(sudo docker run --rm -i -v $(pwd)/elements/val-uniroma1:/root/.university_chain_it university_chain_itd_i tendermint show-node-id)
persistent_peer_humboldt=$node_id_humboldt"@val-uniroma1:26656"

echo ""
echo "Command: val-unipi persistent peer"
echo ""

sudo docker run --rm -i -v $(pwd)/elements/val-unipi:/root/.university_chain_it --entrypoint sed university_chain_itd_i -Ei 's/^persistent_peers = ""/persistent_peers = "'$persistent_peer_humboldt'"/g' /root/.university_chain_it/config/config.toml

