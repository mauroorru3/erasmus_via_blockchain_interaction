#!/bin/bash


if [ $# -ne 2 ];
	then echo "The number of arguments must be 2. The first argument represents the home university and the second argument the destination university. Usage example: ./test_erasmus_request_chain_from_it_to_de.sh unipi tum" 
	exit 1
fi

if { [ "$1" != "unipi" ] && [ "$1" != "uniroma1" ]; } || { [ "$2" != "tum" ] && [ "$2" != "humboldt university" ]; };
then 
	echo "The first argument must be equal to unipi or uniroma1 and the second argument must be equal to tum or humboldt university. Usage example: ./test_erasmus_request_chain_from_it_to_de.sh unipi tum"
	exit 1
fi



Mario_Rossi=$(sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it  university_chain_itd_i keys --keyring-backend test show "Mario Rossi" --address)

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-erasmus-request $1 1 6 $2 study --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-erasmus-request $1 1 6 $2 study --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"


