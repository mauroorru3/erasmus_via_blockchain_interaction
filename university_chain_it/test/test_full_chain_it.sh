#!/bin/bash

if [ $# -ne 2 ];
	then echo "The number of arguments must be 2. The first argument represents the home university and the second argument the destination university. Usage example: ./test_full_chain_it.sh unipi tum" 
	exit 1
fi

if { [ "$1" != "unipi" ] && [ "$1" != "uniroma1" ]; } || { [ "$2" != "tum" ] && [ "$2" != "humboldt university" ]; };
then 
	echo "The first argument must be equal to unipi or uniroma1 and the second argument must be equal to tum or humboldt university. Usage example: ./test_full_chain_it.sh unipi tum"
	exit 1
fi


Mario_Rossi=$(sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-unipi:/root/.university_chain_it  university_chain_itd_i keys --keyring-backend test show "Mario Rossi" --address)
university=""

if [ "$1" = "unipi" ];
then 
	university=$(sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-unipi:/root/.university_chain_it  university_chain_itd_i keys --keyring-backend test show "Unipi" --address)
elif [ "$1" != "uniroma1" ];
then 
	university=$(sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-unipi:/root/.university_chain_it  university_chain_itd_i keys --keyring-backend test show "Uniroma1" --address)	
fi

chain_admin=$(sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-unipi:/root/.university_chain_it  university_chain_itd_i keys --keyring-backend test show "Admin Chain IT" --address)

prof_ae=$(sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-unipi:/root/.university_chain_it  university_chain_itd_i keys --keyring-backend test show "Prof. Domenico Asprucci" --address)

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit configure-chain --from "$chain_admin" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit configure-chain --from "$chain_admin" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"


while [ "$?" -ne 0 ];
do
	echo ""
	echo "Wait 10 seconds and run the command again"
	echo ""
	sleep 10
	sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit configure-chain --from "$chain_admin" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"
done

sleep 10

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit register-new-student "$1" Mario Rossi master cs "Computer Science" --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit register-new-student "$1" Mario Rossi master cs "Computer Science" --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"

while [ "$?" -ne 0 ];
do
	echo ""
	echo "Wait 10 seconds and run the command again"
	echo ""
	sleep 10
	sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit register-new-student "$1" Mario Rossi master cs "Computer Science" --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"
done

sleep 10

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-student-personal-info "$1" 1 male 1994-06-06 italian italy Rome Rome 1111111111111111 20000 --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-student-personal-info "$1" 1 male 1994-06-06 italian italy Rome Rome 1111111111111111 20000 --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"

while [ "$?" -ne 0 ];
do
	echo ""
	echo "Wait 10 seconds and run the command again"
	echo ""
	sleep 10
	sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit configure-chain --from "$chain_admin" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"
	sleep 10
	sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-student-personal-info "$1" 1 male 1994-06-06 italian italy Rome Rome 1111111111111111 20000 --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"
done

sleep 10

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-student-contact-info "$1" 1 via roma mario.rossi@example.it 0000000000 --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-student-contact-info "$1" 1 "via roma" mario.rossi@example.it 0000000000 --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"

while [ "$?" -ne 0 ];
do
	echo ""
	echo "Wait 10 seconds and run the command again"
	echo ""
	sleep 10
	sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-student-contact-info "$1" 1 "via roma" mario.rossi@example.it 0000000000 --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"
done

sleep 10

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-student-residence-info "$1" 1 italy PI Pisa 56100 "via roma" 3 0000000000 --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-student-residence-info "$1" 1 italy PI Pisa 56100 "via roma" 3 0000000000 --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"

while [ "$?" -ne 0 ];
do
	echo ""
	echo "Wait 10 seconds and run the command again"
	echo ""
	sleep 10
	sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-student-residence-info "$1" 1 italy PI Pisa 56100 "via roma" 3 0000000000 --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"
done

sleep 10

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit pay-taxes "$1" 1 --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit pay-taxes "$1" 1 --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"

while [ "$?" -ne 0 ];
do
	echo ""
	echo "Wait 10 seconds and run the command again"
	echo ""
	sleep 10
	sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit pay-taxes "$1" 1 --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"
done

sleep 10

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-exam-grade "$1" 1 "Algorithm engineering" 25 --from "$prof_ae" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-exam-grade "$1" 1 "Algorithm engineering" 25 --from "$prof_ae" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"

while [ "$?" -ne 0 ];
do
	echo ""
	echo "Wait 10 seconds and run the command again"
	echo ""
	sleep 10
	sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-exam-grade "$1" 1 "Algorithm engineering" 25 --from "$prof_ae" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"
done

sleep 10

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-erasmus-request "$1" 1 6 "$2" study --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-erasmus-request "$1" 1 6 "$2" study --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"

while [ "$?" -ne 0 ];
do
	echo ""
	echo "Wait 10 seconds and run the command again"
	echo ""
	sleep 10
	sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-erasmus-request "$1" 1 6 "$2" study --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"
done

sleep 10

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-erasmus-exam "$1" 1 "Advanced databases" --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-erasmus-exam "$1" 1 "Advanced databases" --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"

while [ "$?" -ne 0 ];
do
	echo ""
	echo "Wait 10 seconds and run the command again"
	echo ""
	sleep 10
	sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-erasmus-exam "$1" 1 "Advanced databases" --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"
done

sleep 10

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit start-erasmus "$1" 1 --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit start-erasmus "$1" 1 --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"


while [ "$?" -ne 0 ];
do
	echo ""
	echo "Wait 10 seconds and run the command again"
	echo ""
	sleep 10
	sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit start-erasmus "$1" 1 --from "$Mario_Rossi" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"
done




