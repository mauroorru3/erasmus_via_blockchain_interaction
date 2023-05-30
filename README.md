# Erasmus via blockchain (Interaction between the Hub and the Italian Chain)

In this repo we focus on the interaction between the Hub and the Italian Chain.

Refers to this [link](https://github.com/mauroorru3/erasmus_via_blockchain) for the Italian Chain and to this [link](https://github.com/mauroorru3/-erasmus_via_blockchain_hub) for the hub.

## Installation

As above.

## Usage



```go

# Create the image of the Italian chain:

cd university_chain_it
sudo docker build -f Dockerfile-university_chain_it . -t university_chain_it --no-cache 
cd ..

# Create the image of the German chain:

cd university_chain_de
sudo docker build -f Dockerfile-university_chain_de . -t university_chain_de --no-cache 
cd ..

# Create the image of the Hub:

cd hub
sudo docker build -f Dockerfile-hub . -t hub --no-cache 
cd ..

# Create the image of the Relayer that connect the Italian chain and the hub:
# Go relayer:

cd go_relayer_it_hub
sudo docker build -f Dockerfile . -t go_relayer_it_hub --no-cache 
cd .. 

# Hermes relayer:

cd hermes_relayer_it_hub
sudo docker build -f Dockerfile . -t hermes_relayer_it_hub --no-cache 
cd .. 

# Create the image of the Relayer that connect the German chain and the hub:
# Go relayer:

cd go_relayer_de_hub
sudo docker build -f Dockerfile . -t go_relayer_de_hub --no-cache 
cd .. 

# Hermes relayer:

cd hermes_relayer_de_hub
sudo docker build -f Dockerfile . -t hermes_relayer_de_hub --no-cache 
cd .. 

# To allow scripts to be executed within containers:

chmod +x allow_permissions.sh 
./allow_permissions.sh 

# To start the containers with hermes relayers:

sudo docker compose -f modular.yaml --profile hermes up 

# To start the containers with go relayers:

sudo docker compose -f modular.yaml --profile go up 

# To start the hermes relayer process:
# Remember that the order of execution is important. 
# First the relayer related to the Chain IT and the Hub must be started, and then the relayer related to the Chain DE and the Hub.

sudo docker exec -it hermes_relayer_it_hub bash
./run-relayer.sh

sudo docker exec -it hermes_relayer_de_hub bash
./run-relayer.sh

# To start the go relayer process:

sudo docker exec go_relayer_it_hub ./run-relayer_it_hub.sh

sudo docker exec go_relayer_de_hub ./run-relayer_de_hub.sh


# To simulate the execution of the Hub and the Italian Chain, thus the chains configuration, the insertion of a student, the student's exams and so on:

./test.sh 

# To end the execution of the containers with hermes relayers:

sudo docker compose -f modular.yaml --profile hermes down

# To end the execution of the containers with go relayers:

sudo docker compose -f modular.yaml --profile go down



```

## Contributing



## License

[MIT](https://choosealicense.com/licenses/mit/)
