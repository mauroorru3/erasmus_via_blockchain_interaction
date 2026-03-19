#!/bin/bash

if [ $# -ne 1 ];
then 
	echo ""
	echo "The number of arguments must be 1. The first argument represents the home university. Usage example: ./test_configure_chain_de.sh tum" 
	echo ""
	exit 1
fi

if { [ "$1" != "tum" ] && [ "$1" != "humboldt university" ]; };
then 
	echo ""
	echo "The first argument must be equal to tum or humboldt university. Usage example: ./test_configure_chain_de.sh tum"
	echo ""
	exit 1
fi

uni="tum"
if [ "$1" = "humboldt university" ];
then 
	uni="humboldt"
fi


chain_admin=$(sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de  university_chain_ded_i keys --keyring-backend test show "Admin Chain DE" --address)

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde configure-chain --from "$chain_admin" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde configure-chain --from "$chain_admin" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"



