#!/bin/bash

# This script aims to automatically delete old images and create new, updated ones.
# It is used in the last phase of development to see the impact of small changes at code level.

sudo docker image rm university_chain_itd_i; 
cd university_chain_it; 
rm ./build/university_chain_itd; 
go build -o ./build/university_chain_itd ./cmd/university_chain_itd/main.go; 
sudo docker build -f Dockerfile-university_chain_itd . -t university_chain_itd_i; 
cd ..;
sudo docker image rm university_chain_ded_i; 
cd university_chain_de; 
rm ./build/university_chain_ded; 
go build -o ./build/university_chain_ded ./cmd/university_chain_ded/main.go; 
sudo docker build -f Dockerfile-university_chain_ded . -t university_chain_ded_i; 
cd ..;
sudo docker image rm hubd_i; 
cd hub; 
rm ./build/hubd; 
go build -o ./build/hubd ./cmd/hubd/main.go; 
sudo docker build -f Dockerfile-hubd . -t hubd_i; 
cd ..;


