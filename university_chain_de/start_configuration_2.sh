#!/bin/bash

echo "1"
echo ""
echo -e val-tum'\n'val-humboldt | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_de --entrypoint sed university_chain_ded_i -Ei 's/^chain-id = .*$/chain-id = "university_chain_de"/g' /root/.university_chain_de/config/client.toml
echo ""
echo "2"
echo ""
echo -e val-tum'\n'val-humboldt | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_de --entrypoint sed university_chain_ded_i -Ei 's/^enable-unsafe-cors = false/enable-unsafe-cors = true/g' /root/.university_chain_de/config/app.toml
echo ""
echo "3"
echo ""
echo -e val-tum'\n'val-humboldt | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_de --entrypoint sed university_chain_ded_i -Ei 's/^enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' /root/.university_chain_de/config/app.toml
echo ""
echo "4"
echo ""
echo -e val-tum'\n'val-humboldt | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_de --entrypoint sed university_chain_ded_i -Ei 's/^laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' /root/.university_chain_de/config/config.toml
echo ""
echo "5"
echo ""
echo -e val-tum'\n'val-humboldt | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_de --entrypoint sed university_chain_ded_i -Ei 's/^cors_allowed_origins = \[\]/cors_allowed_origins = \["\*"\]/g' /root/.university_chain_de/config/config.toml
echo ""
echo "6"
echo ""
sudo docker run --rm -i -v $(pwd)/elements/val-humboldt:/root/.university_chain_de university_chain_ded_i gentx "Humboldt University" 40000000stake --keyring-backend test --account-number 0 --sequence 0 --chain-id university_chain_de --gas 1000000 --gas-prices 0.1stake
echo ""
echo "7"
echo ""
sudo mv elements/val-humboldt/config/genesis.json elements/val-tum/config/
echo ""
echo "8"
echo ""
sudo docker run --rm -i -v $(pwd)/elements/val-tum:/root/.university_chain_de university_chain_ded_i gentx TUM 40000000stake --keyring-backend test --account-number 0 --sequence 0 --chain-id university_chain_de --gas 1000000 --gas-prices 0.1stake
echo ""
echo "9"
echo ""
# copy manually elements/val-humboldt/config/gentx/* elements/val-tum/config/gentx
sudo cp elements/val-humboldt/config/gentx/gentx-* elements/val-tum/config/gentx 
echo ""
echo "10"
echo ""
sudo docker run --rm -it -v $(pwd)/elements/val-tum:/root/.university_chain_de university_chain_ded_i collect-gentxs
echo ""
echo "11"
echo ""
sudo docker run --rm -it -v $(pwd)/elements/val-tum:/root/.university_chain_de university_chain_ded_i validate-genesis 
echo ""
echo "12"
echo ""
sudo cp elements/val-tum/config/genesis.json elements/val-humboldt/config

#sudo docker run --rm -i -v $(pwd)/elements/val-tum:/root/.university_chain_de university_chain_ded_i tendermint show-node-id
#Enter the result in config.toml in persistent_peers: @ with the node name and : with 26656 as the port.
#sudo docker run --rm -i -v $(pwd)/elements/val-humboldt:/root/.university_chain_de university_chain_ded_i tendermint show-node-id
#Enter the result in config.toml in persistent_peers: @ with the node name and : with 26656 as the port.

