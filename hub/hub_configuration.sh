#!/bin/bash

echo ""
echo "Command: sudo mkdir -p elements/val-hub"
echo ""
sudo mkdir -p elements/val-hub
echo ""
echo "Command: sudo docker run --rm -i -v $(pwd)/elements/val-hub:/root/.hub hubd_i init hub"
echo ""
sudo docker run --rm -i -v $(pwd)/elements/val-hub:/root/.hub hubd_i init hub
echo ""
echo "Command: sudo cp -r ./dir_keys/keyring-test $(pwd)/elements/val-hub/"
echo ""
sudo cp -r ./dir_keys/keyring-test $(pwd)/elements/val-hub/
echo ""
#Copy manually the address and balances information in the genesis file
echo "Command: Accounts val-hub"
jq --argjson info "$(<genesis_dir/accounts.json)" '.app_state.auth.accounts = $info' elements/val-hub/config/genesis.json > genesis.tmp && sudo mv genesis.tmp elements/val-hub/config/genesis.json
echo "Accounts Completed val-hub"
echo "Command: Balances val-hub"
jq --argjson info "$(<genesis_dir/balances.json)" '.app_state.bank.balances = $info' elements/val-hub/config/genesis.json > genesis.tmp && sudo mv genesis.tmp elements/val-hub/config/genesis.json
echo "Balances Completed val-hub"
echo ""
echo "Command: sudo docker run --rm -i -v $(pwd)/elements/val-hub:/root/.hub --entrypoint sed hubd_i -Ei 's/^chain-id = .*$/chain-id = "hub"/g' /root/.hub/config/client.toml"
echo ""
sudo docker run --rm -i -v $(pwd)/elements/val-hub:/root/.hub --entrypoint sed hubd_i -Ei 's/^chain-id = .*$/chain-id = "hub"/g' /root/.hub/config/client.toml
echo ""
echo "Command: sudo docker run --rm -i -v $(pwd)/elements/val-hub:/root/.hub --entrypoint sed hubd_i -Ei 's/^enable-unsafe-cors = false/enable-unsafe-cors = true/g' /root/.hub/config/app.toml"
echo ""
sudo docker run --rm -i -v $(pwd)/elements/val-hub:/root/.hub --entrypoint sed hubd_i -Ei 's/^enable-unsafe-cors = false/enable-unsafe-cors = true/g' /root/.hub/config/app.toml
echo ""
echo "Command: sudo docker run --rm -i -v $(pwd)/elements/val-hub:/root/.hub --entrypoint sed hubd_i -Ei 's/^enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' /root/.hub/config/app.toml"
echo ""
sudo docker run --rm -i -v $(pwd)/elements/val-hub:/root/.hub --entrypoint sed hubd_i -Ei 's/^enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' /root/.hub/config/app.toml
echo ""
echo "Command: sudo docker run --rm -i -v $(pwd)/elements/val-hub:/root/.hub --entrypoint sed hubd_i -Ei 's/^laddr = "tcp:127.0.0.1:26657"/laddr = "tcp:0.0.0.0:26657"/g' /root/.hub/config/config.toml"
echo ""
sudo docker run --rm -i -v $(pwd)/elements/val-hub:/root/.hub --entrypoint sed hubd_i -Ei 's/^laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' /root/.hub/config/config.toml
echo ""
echo "Command: sudo docker run --rm -i -v $(pwd)/elements/val-hub:/root/.hub --entrypoint sed hubd_i -Ei 's/^cors_allowed_origins = \[\]/cors_allowed_origins = \["*"\]/g' /root/.hub/config/config.toml"
echo ""
sudo docker run --rm -i -v $(pwd)/elements/val-hub:/root/.hub --entrypoint sed hubd_i -Ei 's/^cors_allowed_origins = \[\]/cors_allowed_origins = \["\*"\]/g' /root/.hub/config/config.toml
echo ""
echo "Command: sudo docker run --rm -i -v $(pwd)/elements/val-hub:/root/.hub hubd_i gentx Hub 40000000stake --keyring-backend test --account-number 0 --sequence 0 --chain-id hub --gas 1000000 --gas-prices 0.1stake"
echo ""
sudo docker run --rm -i -v $(pwd)/elements/val-hub:/root/.hub hubd_i gentx Hub 40000000stake --keyring-backend test --account-number 0 --sequence 0 --chain-id hub --gas 1000000 --gas-prices 0.1stake
echo ""
echo "Command: sudo docker run --rm -it -v $(pwd)/elements/val-hub:/root/.hub hubd_i collect-gentxs"
echo ""
sudo docker run --rm -it -v $(pwd)/elements/val-hub:/root/.hub hubd_i collect-gentxs
echo ""
echo "Command: sudo docker run --rm -it -v $(pwd)/elements/val-hub:/root/.hub hubd_i validate-genesis"
echo ""
sudo docker run --rm -it -v $(pwd)/elements/val-hub:/root/.hub hubd_i validate-genesis 