#!/bin/bash


echo "1"
echo ""
sudo docker run --rm -it -v $(pwd)/elements/val-tum:/root/.university_chain_de university_chain_ded_i collect-gentxs
echo ""
echo "2"
echo ""
sudo docker run --rm -it -v $(pwd)/elements/val-tum:/root/.university_chain_de university_chain_ded_i validate-genesis 
echo ""
echo "3"
echo ""
sudo cp elements/val-tum/config/genesis.json elements/val-humboldt/config

#sudo docker run --rm -i -v $(pwd)/elements/val-tum:/root/.university_chain_de university_chain_ded_i tendermint show-node-id
#Enter the result in config.toml in persistent_peers: @ with the node name and : with 26656 as the port.
#sudo docker run --rm -i -v $(pwd)/elements/val-humboldt:/root/.university_chain_de university_chain_ded_i tendermint show-node-id
#Enter the result in config.toml in persistent_peers: @ with the node name and : with 26656 as the port.

