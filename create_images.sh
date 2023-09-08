#!/bin/bash

# Create the image of the Italian chain:
cd university_chain_it; 
rm ./build/university_chain_itd; 
go build -o ./build/university_chain_itd ./cmd/university_chain_itd/main.go; 
sudo docker build -f Dockerfile-university_chain_itd . -t university_chain_itd_i; 
cd ..;

# Create the image of the German chain:
cd university_chain_de; 
rm ./build/university_chain_ded; 
go build -o ./build/university_chain_ded ./cmd/university_chain_ded/main.go; 
sudo docker build -f Dockerfile-university_chain_ded . -t university_chain_ded_i; 
cd ..;

# Create the image of the Hub chain:
cd hub; 
rm ./build/hubd; 
go build -o ./build/hubd ./cmd/hubd/main.go; 
sudo docker build -f Dockerfile-hubd . -t hubd_i; 
cd ..;

# Create the image of the Relayer that connect the Italian chain with the hub chain:
# Go relayer:
cd go_relayer_it_hub;
sudo docker build -f Dockerfile . -t go_relayer_it_hub --no-cache;
cd ..;

# Hermes relayer:
cd hermes_relayer_it_hub;
sudo docker build -f Dockerfile . -t hermes_relayer_it_hub --no-cache;
cd ..;

# Create the image of the Relayer that connect the German chain with the hub chain:
# Go relayer:
cd go_relayer_de_hub;
sudo docker build -f Dockerfile . -t go_relayer_de_hub --no-cache;
cd ..;

# Hermes relayer:
cd hermes_relayer_de_hub;
sudo docker build -f Dockerfile . -t hermes_relayer_de_hub --no-cache;
cd ..;

