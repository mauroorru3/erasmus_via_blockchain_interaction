#!/bin/bash


if [ $# -ne 1 ];
	then echo "The number of arguments must be 1. The argument represents the university to be tested. Usage example: ./test_insert_grade_chain_it.sh.sh unipi" 
	exit 1
fi

if [ "$1" != "unipi" ] && [ "$1" != "uniroma1" ];
then 
	echo "The argument must be equal to unipi or uniroma1. Usage example: ./test_insert_grade_chain_it.sh.sh unipi"
	exit 1
fi


prof_advanced_databases=$(sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-unipi:/root/.university_chain_it  university_chain_itd_i keys --keyring-backend test show "Prof. Ermanno Naccari" --address)

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-exam-grade "$1" 1 "Advanced databases" 25 --from "$prof_advanced_databases" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-exam-grade "$1" 1 "Advanced databases" 25 --from "$prof_advanced_databases" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"

