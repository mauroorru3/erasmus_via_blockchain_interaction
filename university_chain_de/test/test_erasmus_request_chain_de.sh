#!/bin/bash

if [ $# -ne 2 ];
	then echo "The number of arguments must be 2. The first argument represents the home university and the second argument the destination university. Usage example: ./test_erasmus_request_chain_from_de_to_it.sh tum unipi"
	exit 1
fi

if { [ "$1" != "tum" ] && [ "$1" != "humboldt university" ]; } || { [ "$2" != "unipi" ] && [ "$2" != "uniroma1" ]; };
then 
	echo "The first argument must be equal to unipi or uniroma1 and the second argument must be equal to tum or humboldt university. Usage example: ./test_erasmus_request_chain_from_de_to_it.sh tum uniroma1"
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
echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-erasmus-request "$1" 1 6 "$2" study --from "$Karl_Schmidt" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-erasmus-request "$1" 1 6 "$2" study --from "$Karl_Schmidt" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657
