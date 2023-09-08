#!/bin/bash

echo ""
echo "Command: sudo mkdir -p elements/val-hub-instance-1"
echo ""
sudo mkdir -p elements/val-hub-instance-1
echo ""
echo "Command: sudo mkdir -p elements/val-hub-instance-2"
echo ""
sudo mkdir -p elements/val-hub-instance-2
echo ""
echo "Command: echo -e val-hub-instance-1'\n'val-hub-instance-2 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.hub hubd_i init hub"
echo ""
echo -e val-hub-instance-1'\n'val-hub-instance-2 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.hub hubd_i init hub
echo ""
echo "Command: echo -e val-hub-instance-1'\n'val-hub-instance-2 | xargs -I {} sudo cp -r ./dir_keys/keyring-test $(pwd)/elements/{}/"
echo ""
echo -e val-hub-instance-1'\n'val-hub-instance-2 | xargs -I {} sudo cp -r ./dir_keys/keyring-test $(pwd)/elements/{}/
echo ""
#Copy manually the address and balances information in the genesis file
echo "Command: Accounts val-hub-instance-1"
jq --argjson info "$(<genesis_dir/accounts.json)" '.app_state.auth.accounts = $info' elements/val-hub-instance-1/config/genesis.json > genesis.tmp && sudo mv genesis.tmp elements/val-hub-instance-1/config/genesis.json
echo "Accounts Completed val-hub-instance-1"
echo "Command: Balances val-hub-instance-1"
jq --argjson info "$(<genesis_dir/balances.json)" '.app_state.bank.balances = $info' elements/val-hub-instance-1/config/genesis.json > genesis.tmp && sudo mv genesis.tmp elements/val-hub-instance-1/config/genesis.json
echo "Balances Completed val-hub-instance-1"
echo ""
echo "Command: Accounts val-hub-instance-2"
jq --argjson info "$(<genesis_dir/accounts.json)" '.app_state.auth.accounts = $info' elements/val-hub-instance-2/config/genesis.json > genesis.tmp && sudo mv genesis.tmp elements/val-hub-instance-2/config/genesis.json
echo "Accounts Completed val-hub-instance-2"
echo "Command: Balances val-hub-instance-2"
jq --argjson info "$(<genesis_dir/balances.json)" '.app_state.bank.balances = $info' elements/val-hub-instance-2/config/genesis.json > genesis.tmp && sudo mv genesis.tmp elements/val-hub-instance-2/config/genesis.json
echo "Balances Completed val-hub-instance-2"
echo ""
echo "Command: echo -e val-hub-instance-1'\n'val-hub-instance-2 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.hub --entrypoint sed hubd_i -Ei 's/^chain-id = .*$/chain-id = "hub"/g' /root/.hub/config/client.toml"
echo ""
echo -e val-hub-instance-1'\n'val-hub-instance-2 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.hub --entrypoint sed hubd_i -Ei 's/^chain-id = .*$/chain-id = "hub"/g' /root/.hub/config/client.toml
echo ""
echo "Command: echo -e val-hub-instance-1'\n'val-hub-instance-2 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.hub --entrypoint sed hubd_i -Ei 's/^enable-unsafe-cors = false/enable-unsafe-cors = true/g' /root/.hub/config/app.toml"
echo ""
echo -e val-hub-instance-1'\n'val-hub-instance-2 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.hub --entrypoint sed hubd_i -Ei 's/^enable-unsafe-cors = false/enable-unsafe-cors = true/g' /root/.hub/config/app.toml
echo ""
echo "Command: echo -e val-hub-instance-1'\n'val-hub-instance-2 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.hub --entrypoint sed hubd_i -Ei 's/^enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' /root/.hub/config/app.toml"
echo ""
echo -e val-hub-instance-1'\n'val-hub-instance-2 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.hub --entrypoint sed hubd_i -Ei 's/^enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' /root/.hub/config/app.toml
echo ""
echo "Command: echo -e val-hub-instance-1'\n'val-hub-instance-2 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.hub --entrypoint sed hubd_i -Ei 's/^laddr = tcp:\/\/127.0.0.1:26657/laddr = tcp:\/\/0.0.0.0:26657 /g' /root/.hub/config/config.toml"
echo ""
echo -e val-hub-instance-1'\n'val-hub-instance-2 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.hub --entrypoint sed hubd_i -Ei 's/^laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' /root/.hub/config/config.toml
echo ""
echo "Command: echo -e val-hub-instance-1'\n'val-hub-instance-2 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.hub --entrypoint sed hubd_i -Ei 's/^cors_allowed_origins = \[\]/cors_allowed_origins = \[\*\]/g' /root/.hub/config/config.toml"
echo ""
echo -e val-hub-instance-1'\n'val-hub-instance-2 | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.hub --entrypoint sed hubd_i -Ei 's/^cors_allowed_origins = \[\]/cors_allowed_origins = \["\*"\]/g' /root/.hub/config/config.toml
echo ""
echo "mv elements/val-hub-instance-1/config/genesis.json elements/val-hub-instance-2/config/"
sudo mv elements/val-hub-instance-1/config/genesis.json elements/val-hub-instance-2/config/
echo ""
echo "Command: sudo docker run --rm -i -v $(pwd)/elements/val-hub-instance-2:/root/.hub hubd_i gentx Hub-instance-2 40000000stake --keyring-backend test --account-number 0 --sequence 0 --chain-id hub --gas 1000000 --gas-prices 0.1stake"
echo ""
sudo docker run --rm -i -v $(pwd)/elements/val-hub-instance-2:/root/.hub hubd_i gentx Hub-instance-2 40000000stake --keyring-backend test --account-number 0 --sequence 0 --chain-id hub --gas 1000000 --gas-prices 0.1stake
echo ""
echo "Command: sudo mv elements/val-hub-instance-2/config/genesis.json elements/val-hub-instance-1/config/"
echo ""
sudo mv elements/val-hub-instance-2/config/genesis.json elements/val-hub-instance-1/config/
echo ""
echo "Command: sudo docker run --rm -i -v $(pwd)/elements/val-hub-instance-1:/root/.hub hubd_i gentx Hub 40000000stake --keyring-backend test --account-number 0 --sequence 0 --chain-id hub --gas 1000000 --gas-prices 0.1stake"
echo ""
sudo docker run --rm -i -v $(pwd)/elements/val-hub-instance-1:/root/.hub hubd_i gentx Hub 40000000stake --keyring-backend test --account-number 0 --sequence 0 --chain-id hub --gas 1000000 --gas-prices 0.1stake
echo ""
echo "Command: sudo cp -r elements/val-hub-instance-2/config/gentx/. elements/val-hub-instance-1/config/gentx/"
echo ""
sudo cp -r elements/val-hub-instance-2/config/gentx/. elements/val-hub-instance-1/config/gentx/
echo ""
echo "Command: sudo docker run --rm -it -v $(pwd)/elements/val-hub-instance-1:/root/.hub hubd_i collect-gentxs"
echo ""
sudo docker run --rm -it -v $(pwd)/elements/val-hub-instance-1:/root/.hub hubd_i collect-gentxs
echo ""
echo "Command: sudo docker run --rm -it -v $(pwd)/elements/val-hub-instance-1:/root/.hub hubd_i validate-genesis"
echo ""
sudo docker run --rm -it -v $(pwd)/elements/val-hub-instance-1:/root/.hub hubd_i validate-genesis 
echo ""
echo "Command: sudo cp elements/val-hub-instance-1/config/genesis.json elements/val-hub-instance-2/config"
echo ""
sudo cp elements/val-hub-instance-1/config/genesis.json elements/val-hub-instance-2/config

#sudo docker run --rm -i -v $(pwd)/elements/val-hub-instance-1:/root/.hub hubd_i tendermint show-node-id
#Enter the result in config.toml in persistent_peers: @ with the node name and : with 26656 as the port.

node_id_hub_instance_1=$(sudo docker run --rm -i -v $(pwd)/elements/val-hub-instance-1:/root/.hub hubd_i tendermint show-node-id)
persistent_peer_hub_instance_1=$node_id_hub_instance_1"@val-hub-instance-1:26656"

echo ""
echo "Command: val-hub-instance-2 persistent peer"
echo ""

sudo docker run --rm -i -v $(pwd)/elements/val-hub-instance-2:/root/.hub --entrypoint sed hubd_i -Ei 's/^persistent_peers = ""/persistent_peers = "'$persistent_peer_hub_instance_1'"/g' /root/.hub/config/config.toml

#sudo docker run --rm -i -v $(pwd)/elements/val-hub-instance-2:/root/.hub hubd_i tendermint show-node-id
#Enter the result in config.toml in persistent_peers: @ with the node name and : with 26656 as the port.


node_id_hub_instance_2=$(sudo docker run --rm -i -v $(pwd)/elements/val-hub-instance-2:/root/.hub hubd_i tendermint show-node-id)
persistent_peer_hub_instance_2=$node_id_hub_instance_2"@val-hub-instance-2:26656"

echo ""
echo "Command: val-hub-instance-1 persistent peer"
echo ""

sudo docker run --rm -i -v $(pwd)/elements/val-hub-instance-1:/root/.hub --entrypoint sed hubd_i -Ei 's/^persistent_peers = ""/persistent_peers = "'$persistent_peer_hub_instance_2'"/g' /root/.hub/config/config.toml

