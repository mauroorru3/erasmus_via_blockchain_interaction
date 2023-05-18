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

# Create the image of the Hub:

cd hub
sudo docker build -f Dockerfile-hub . -t hub --no-cache 
cd ..

# Create the image of the Relayer:

cd relayer 
sudo docker build -f Dockerfile . -t relayer --no-cache 
cd .. 

# To allow scripts to be executed within containers:

chmod +x allow_permissions.sh 
./allow_permissions.sh 

# To start the containers

sudo docker compose -f modular.yaml up 

# To start the relayer process:

sudo docker exec relayer ./run-relayer.sh 

# To simulate the execution of the Hub and the Italian Chain, thus the chains configuration, the insertion of a student, the student's exams and so on:

./test.sh 

# To end the execution of the containers:

sudo docker compose -f modular.yaml --profile go down 



```

## Contributing



## License

[MIT](https://choosealicense.com/licenses/mit/)
