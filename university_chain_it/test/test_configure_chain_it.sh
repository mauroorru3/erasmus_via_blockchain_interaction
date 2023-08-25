#!/bin/bash


if [ $# -ne 1 ];
	then echo "The number of arguments must be 1. The first argument represents the home university. Usage example: ./test_configure_chain_it.sh unipi" 
	exit 1
fi

if { [ "$1" != "unipi" ] && [ "$1" != "uniroma1" ]; };
then 
	echo "The first argument must be equal to unipi or uniroma1. Usage example: ./test_configure_chain_it.sh unipi"
	exit 1
fi



chain_admin=$(sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-unipi:/root/.university_chain_it  university_chain_itd_i keys --keyring-backend test show "Admin Chain IT" --address)

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit configure-chain --from "$chain_admin" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit configure-chain --from "$chain_admin" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"


