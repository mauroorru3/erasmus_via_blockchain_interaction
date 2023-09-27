#!/bin/bash

if [ $# -ne 2 ];
	then echo "The number of arguments must be 2. The first argument represents the home university and the second argument the destination university. Usage example: ./test_multiple_Erasmus_requests_chain_it.sh unipi tum" 
	exit 1
fi

if { [ "$1" != "unipi" ] && [ "$1" != "uniroma1" ]; } || { [ "$2" != "tum" ] && [ "$2" != "humboldt university" ]; };
then 
	echo "The first argument must be equal to unipi or uniroma1 and the second argument must be equal to tum or humboldt university. Usage example: ./test_multiple_Erasmus_requests_chain_it.sh unipi tum"
	exit 1
fi


export university=""

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



usersArray=( $(sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-unipi:/root/.university_chain_it  university_chain_itd_i keys --keyring-backend test show "Mario Rossi" --address) $(sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-unipi:/root/.university_chain_it  university_chain_itd_i keys --keyring-backend test show "Tamara Genovesi" --address) $(sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-unipi:/root/.university_chain_it  university_chain_itd_i keys --keyring-backend test show "Franca Ferrari" --address) $(sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-unipi:/root/.university_chain_it  university_chain_itd_i keys --keyring-backend test show "Luigi Lena" --address) $(sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-unipi:/root/.university_chain_it  university_chain_itd_i keys --keyring-backend test show "Gianpiero Serafino" --address) $(sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-unipi:/root/.university_chain_it  university_chain_itd_i keys --keyring-backend test show "Antonella Marino" --address))
nameAndSurnameArray=("Mario" "Rossi" "Tamara" "Genovesi" "Franca" "Ferrari" "Luigi" "Lena" "Gianpiero" "Serafino" "Antonella" "Marino")


: << 'COMMENT'
usersArray=( $(sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-unipi:/root/.university_chain_it  university_chain_itd_i keys --keyring-backend test show "Mario Rossi" --address) $(sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-unipi:/root/.university_chain_it  university_chain_itd_i keys --keyring-backend test show "Tamara Genovesi" --address) $(sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-unipi:/root/.university_chain_it  university_chain_itd_i keys --keyring-backend test show "Franca Ferrari" --address))
nameAndSurnameArray=("Mario" "Rossi" "Tamara" "Genovesi" "Franca" "Ferrari")
COMMENT


i=0
j=1
sleep 20


for user in "${usersArray[@]}"
do

	echo ""
	echo "Command:"
	echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit register-new-student "$1" ${nameAndSurnameArray[$i]} ${nameAndSurnameArray[$i+1]} master cs "Computer Science" --from "$user" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
	echo ""
	sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit register-new-student "$1" "${nameAndSurnameArray[$i]}" "${nameAndSurnameArray[$i+1]}" master cs "Computer Science" --from "$user" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"
	
	
	i=$((i+2))
	j=$((j+1))
	
	sleep 20
done

j=1

for user in "${usersArray[@]}"
do

	echo ""
	echo "Command:"
	echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-student-personal-info "$1" "$j" male 1994-06-06 italian italy Rome Rome 1111111111111111 20000 --from "$user" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
	echo ""
	sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-student-personal-info "$1" "$j" male 1994-06-06 italian italy Rome Rome 1111111111111111 20000 --from "$user" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"
	
	j=$((j+1))
	
	sleep 20
done

j=1

for user in "${usersArray[@]}"
do

	echo ""
	echo "Command:"
	echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-student-personal-info "$1" "$j" male 1994-06-06 italian italy Rome Rome 1111111111111111 20000 --from "$user" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
	echo ""
	sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-student-personal-info "$1" "$j" male 1994-06-06 italian italy Rome Rome 1111111111111111 20000 --from "$user" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"
	
	j=$((j+1))
	
	sleep 20
done

j=1	

for user in "${usersArray[@]}"
do

	echo ""
	echo "Command:"
	echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-student-contact-info "$1" "$j" "via roma" mario.rossi@example.it 0000000000 --from "$user" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
	echo ""
	sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-student-contact-info "$1" "$j" "via roma" mario.rossi@example.it 0000000000 --from "$user" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"
	
	j=$((j+1))
	
	sleep 20
done
	
j=1
	
for user in "${usersArray[@]}"
do

	echo ""
	echo "Command:"
	echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-student-residence-info "$1" "$j" italy PI Pisa 56100 "via roma" 3 0000000000 --from "$user" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
	echo ""
	sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-student-residence-info "$1" "$j" italy PI Pisa 56100 "via roma" 3 0000000000 --from "$user" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"
	
	j=$((j+1))
	
	sleep 20
done
		
j=1

for user in "${usersArray[@]}"
do

	echo ""
	echo "Command:"
	echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit pay-taxes "$1" "$j" --from "$user" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
	echo ""
	sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit pay-taxes "$1" "$j" --from "$user" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"
	
	j=$((j+1))
	
	sleep 20
done

j=1

for user in "${usersArray[@]}"
do

	echo ""
	echo "Command:"
	echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-exam-grade "$1" "$j" "Algorithm engineering" 25 --from "$prof_ae" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
	echo ""
	sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-exam-grade "$1" "$j" "Algorithm engineering" 25 --from "$prof_ae" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"
	
	j=$((j+1))
	
	sleep 20
done

j=1	

for user in "${usersArray[@]}"
do

	echo ""
	echo "Command:"
	echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-erasmus-request "$1" "$j" 6 "$2" study --from "$user" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
	echo ""
	sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-erasmus-request "$1" "$j" 6 "$2" study --from "$user" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"
	
	j=$((j+1))
	
	sleep 20
done

j=1

for user in "${usersArray[@]}"
do

	echo ""
	echo "Command:"
	echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-erasmus-exam "$1" "$j" "Advanced databases" --from "$user" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
	echo ""
	sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit insert-erasmus-exam "$1" "$j" "Advanced databases" --from "$user" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"
	
	j=$((j+1))
	
	sleep 20
done

j=1

for user in "${usersArray[@]}"
do

	echo ""
	echo "Command:"
	echo "sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit start-erasmus "$1" "$j" --from "$user" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node tcp://val-"$1":26657"
	echo ""
	sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-"$1":/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i tx universitychainit start-erasmus "$1" "$j" --from "$user" --keyring-backend test --gas auto --chain-id university_chain_it --yes --node "tcp://val-"$1":26657"
	
	j=$((j+1))
	
	sleep 20

done


