#!/bin/bash

sudo rm university_chain_it/data/logs.txt
sudo rm university_chain_de/data/logs.txt
sudo rm hub/data/logs.txt
sudo rm hub/data/data.txt
sudo docker exec hub ./test/test_configure_chain_hub.sh
sudo docker exec university_chain_de ./test/test_configure_chain_de.sh
sudo docker exec university_chain_it ./test/test_full_chain_it.sh

