# Erasmus via blockchain (Hub version)

In this repo we focus on the Hub.

Refers to this [link](https://github.com/mauroorru3/erasmus_via_blockchain).

## Installation

Refers to this [link](https://github.com/mauroorru3/erasmus_via_blockchain).

## Usage



```go
# Create the image:

sudo docker build -f Dockerfile-ubuntu . -t hub_i

# To run the chain:

sudo docker run --rm -it -v $(pwd):/hub -w /hub -p 1317:1317 -p 3000:3000 -p 4500:4500 -p 5000:5000 -p 26657:26657 --name hub hub_i ignite chain serve --reset-once

# To interact with the running chain open another shell and try: 

sudo docker exec -it hub hubd status

# To export the admin to simplify address insertion:

export chain_admin=$(sudo docker exec hub hubd keys show "Admin Hub" -a) 

# To initialize the chain

sudo docker exec -it hub hubd tx hub configure-chain --from $chain_admin --gas auto 

# To show the chain info:

sudo docker exec -it hub hubd query hub show-chain-info 

# To show the universities info:

sudo docker exec -it hub hubd query hub list-universities 

# To perform the tests:

sudo docker run --rm -it -v $(pwd):/hub -w /hub hub_i go test hub/x/hub/types

sudo docker run --rm -it -v $(pwd):/hub -w /hub hub_i go test hub/x/hub/keeper

```

## Contributing



## License

[MIT](https://choosealicense.com/licenses/mit/)