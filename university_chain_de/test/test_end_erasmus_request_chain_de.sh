#!/bin/bash


if [ $# -ne 1 ];
	then echo "The number of arguments must be 1. The first argument represents the home university. Usage example: ./test_end_erasmus_request_chain_de.sh unipi" 
	exit 1
fi

if { [ "$1" != "unipi" ] && [ "$1" != "uniroma1" ]; };
then 
	echo "The first argument must be equal to unipi or uniroma1. Usage example: ./test_end_erasmus_request_chain_de.sh unipi"
	exit 1
fi


uni="tum"
if [ "$1" = "humboldt university" ];
then 
	uni="humboldt"
fi


Karl_Schmidt=$(sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de  university_chain_ded_i keys --keyring-backend test show "Karl Schmidt" -a) 

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde end-erasmus-before-deadline "$1" 1 --from $Karl_Schmidt --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde end-erasmus-before-deadline "$1" 1 --from $Karl_Schmidt --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"

