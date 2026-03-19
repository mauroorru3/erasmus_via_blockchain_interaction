#!/bin/bash

if [ $# -ne 1 ];
	then echo "The number of arguments must be 1. The argument represents the university to be tested. Usage example: ./test_insert_grade_chain_de.sh tum" 
	exit 1
fi

if [ "$1" != "tum" ] && [ "$1" != "humboldt university" ];
then 
	echo "The argument must be equal to tum or Humboldt University. Usage example: ./test_insert_grade_chain_de.sh tum"
	exit 1
fi


uni="tum"
if [ "$1" = "humboldt university" ];
then 
	uni="humboldt"
fi

export prof_advanced_databases=$(sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de  university_chain_ded_i keys --keyring-backend test show "Prof. Gustav Fischer" -a) 

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$1":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-exam-grade "$1" 1 "Advanced databases" 2.2 --from "$prof_advanced_databases" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$1":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-exam-grade "$1" 1 "Advanced databases" 2.2 --from "$prof_advanced_databases" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657

