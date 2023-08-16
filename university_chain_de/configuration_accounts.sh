#!/bin/bash

#sudo docker run --rm -i -v $(pwd)/elements/val-tum:/root/.university_chain_de university_chain_ded_i tendermint show-node-id
#Enter the result in config.toml in persistent_peers: @ with the node name and : with 26656 as the port.

node_id_tum=$(sudo docker run --rm -i -v $(pwd)/elements/val-tum:/root/.university_chain_de university_chain_ded_i tendermint show-node-id)
echo $node_id_tum
persistent_peer_tum=$node_id_tum"@val-tum:26656"
echo $persistent_peer_tum

sudo docker run --rm -i -v $(pwd)/elements/val-humboldt:/root/.university_chain_de --entrypoint sed university_chain_ded_i -Ei 's/^persistent_peers = ""/persistent_peers = "'$persistent_peer_tum'"/g' /root/.university_chain_de/config/config.toml

#sudo docker run --rm -i -v $(pwd)/elements/val-humboldt:/root/.university_chain_de university_chain_ded_i tendermint show-node-id
#Enter the result in config.toml in persistent_peers: @ with the node name and : with 26656 as the port.


node_id_humboldt=$(sudo docker run --rm -i -v $(pwd)/elements/val-humboldt:/root/.university_chain_de university_chain_ded_i tendermint show-node-id)
persistent_peer_humboldt=$node_id_humboldt"@val-humboldt:26656"
echo $persistent_peer_humboldt

sudo docker run --rm -i -v $(pwd)/elements/val-tum:/root/.university_chain_de --entrypoint sed university_chain_ded_i -Ei 's/^persistent_peers = ""/persistent_peers = "'$persistent_peer_humboldt'"/g' /root/.university_chain_de/config/config.toml



