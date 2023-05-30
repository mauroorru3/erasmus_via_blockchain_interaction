#!/bin/bash

sudo rm university_chain_it/data/logs.txt
sudo rm university_chain_de/data/logs.txt
sudo rm hub/data/logs.txt
sudo rm hub/data/data.txt
sudo docker exec hub ./test.sh
sudo docker exec university_chain_de ./test.sh
sudo docker exec university_chain_it ./test.sh

