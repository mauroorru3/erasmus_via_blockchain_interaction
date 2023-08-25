#!/bin/bash

if [ $# -ne 2 ];
	then echo "The number of arguments must be 2. The first argument represents the home university and the second argument the destination university. Usage example: ./test_multiple_Erasmus_requests_chain_de.sh tum unipi" 
	exit 1
fi

if { [ "$1" != "tum" ] && [ "$1" != "humboldt university" ]; } || { [ "$2" != "unipi" ] && [ "$2" != "uniroma1" ]; };
then 
	echo "The first argument must be equal to tum or humboldt university and the second argument must be equal to unipi or uniroma1. Usage example: ./test_multiple_Erasmus_requests_chain_de.sh tum unipi"
	exit 1
fi

uni="tum"
if [ "$1" = "humboldt university" ];
then 
	uni="humboldt"
fi


export university=""

if [ "$1" = "tum" ];
then 
	university=$(sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de  university_chain_ded_i keys --keyring-backend test show "TUM" -a) 
elif [ "$1" != "humboldt university" ];
then 
	university=$(sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de  university_chain_ded_i keys --keyring-backend test show "Humboldt University" -a) 
fi

export chain_admin=$(sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de  university_chain_ded_i keys --keyring-backend test show "Admin Chain DE" -a) 

export prof_ae=$(sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de  university_chain_ded_i keys --keyring-backend test show "Prof. Friedrich Mayer" -a) 


echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde configure-chain --from $chain_admin --keyring-backend test --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde configure-chain --from "$chain_admin" --keyring-backend test --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"



usersArray=( $(sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de  university_chain_ded_i keys --keyring-backend test show "Karl Schmidt" -a) $(sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de  university_chain_ded_i keys --keyring-backend test show "Paul Becker" -a) 
	$(sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de  university_chain_ded_i keys --keyring-backend test show "Mike Eisenhauer" -a) $(sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de  university_chain_ded_i keys --keyring-backend test show "Daniela Gloeckner" -a) 
	$(sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de  university_chain_ded_i keys --keyring-backend test show "Nadine Baier" -a) $(sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de  university_chain_ded_i keys --keyring-backend test show "Jennifer Diederich" -a))
nameAndSurnameArray=("Karl" "Schmidt" "Paul" "Becker" "Mike" "Eisenhauer" "Daniela" "Gloeckner" "Nadine" "Baier" "Jennifer" "Diederich")
i=0
j=1


for user in "${usersArray[@]}"
do
	echo ""
	echo "Command:"
	echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde register-new-student $1 ${nameAndSurnameArray[$i]} ${nameAndSurnameArray[$i+1]} master cs "Computer Science" --from "$user" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
	echo ""
	sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde register-new-student "$1" "${nameAndSurnameArray[$i]}" "${nameAndSurnameArray[$i+1]}" master cs "Computer Science" --from "$user" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"
	
	while [ "$?" -ne 0 ];
	do
		echo ""
		echo "Wait 10 seconds and run the command again"
		echo ""
		sleep 10
		sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde register-new-student "$1" "${nameAndSurnameArray[$i]}" "${nameAndSurnameArray[$i+1]}" master cs "Computer Science" --from "$user" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"
	done
	
	sleep 10

	echo ""
	echo "Command:"
	echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-student-personal-info $1 $j male 1994-06-06 italian italy Rome Rome 1111111111111111 20000 --from $user --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
	echo ""
	sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-student-personal-info "$1" "$j" male 1994-06-06 italian italy Rome Rome 1111111111111111 20000 --from "$user" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"
	
	
	while [ "$?" -ne 0 ];
	do
		echo ""
		echo "Wait 10 seconds and run the command again"
		echo ""
		sleep 10
		sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde register-new-student "$1" "${nameAndSurnameArray[$i]}" "${nameAndSurnameArray[$i+1]}" master cs "Computer Science" --from "$user" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"
		sleep 10
		sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-student-personal-info "$1" "$j" male 1994-06-06 italian italy Rome Rome 1111111111111111 20000 --from "$user" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"
	done
	
	sleep 10

	echo ""
	echo "Command:"
	echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-student-contact-info $1 $j "via roma" mario.rossi@example.it 0000000000 --from $user --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
	echo ""
	sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-student-contact-info "$1" "$j" "via roma" mario.rossi@example.it 0000000000 --from "$user" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"


	while [ "$?" -ne 0 ];
	do
		echo ""
		echo "Wait 10 seconds and run the command again"
		echo ""
		sleep 10
		sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-student-contact-info "$1" "$j" "via roma" mario.rossi@example.it 0000000000 --from "$user" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"
	done
	
	sleep 10

	echo ""
	echo "Command:"
	echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-student-residence-info $1 $j italy PI Pisa 56100 "via roma" 3 0000000000 --from $user --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
	echo ""
	sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-student-residence-info "$1" "$j" italy PI Pisa 56100 "via roma" 3 0000000000 --from "$user" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"


	while [ "$?" -ne 0 ];
	do
		echo ""
		echo "Wait 10 seconds and run the command again"
		echo ""
		sleep 10
		sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-student-residence-info "$1" "$j" italy PI Pisa 56100 "via roma" 3 0000000000 --from "$user" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"
	done
	
	sleep 10

	echo ""
	echo "Command:"
	echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde pay-taxes $1 $j --from $user --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
	echo ""
	sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde pay-taxes "$1" "$j" --from "$user" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"


	while [ "$?" -ne 0 ];
	do
		echo ""
		echo "Wait 10 seconds and run the command again"
		echo ""
		sleep 10
		sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde pay-taxes "$1" "$j" --from "$user" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"
	done
	
	sleep 10

	echo ""
	echo "Command:"
	echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-exam-grade $1 $j "Algorithm engineering" 1.5 --from $prof_ae --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
	echo ""
	sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-exam-grade "$1" "$j" "Algorithm engineering" 1.5 --from "$prof_ae" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"

	while [ "$?" -ne 0 ];
	do
		echo ""
		echo "Wait 10 seconds and run the command again"
		echo ""
		sleep 10
		sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-exam-grade "$1" "$j" "Algorithm engineering" 1.5 --from "$prof_ae" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"
	done
	
	sleep 10

	echo ""
	echo "Command:"
	echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-erasmus-request $1 $j 6 $2 study --from $user --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
	echo ""
	sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-erasmus-request "$1" "$j" 6 "$2" study --from "$user" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"

	while [ "$?" -ne 0 ];
	do
		echo ""
		echo "Wait 10 seconds and run the command again"
		echo ""
		sleep 10
		sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-erasmus-request "$1" "$j" 6 "$2" study --from "$user" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"
	done
	
	sleep 10

	echo ""
	echo "Command:"
	echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-erasmus-exam $1 $j "Advanced databases" --from $user --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
	echo ""
	sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-erasmus-exam "$1" "$j" "Advanced databases" --from "$user" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"


	while [ "$?" -ne 0 ];
	do
		echo ""
		echo "Wait 10 seconds and run the command again"
		echo ""
		sleep 10
		sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-erasmus-exam "$1" "$j" "Advanced databases" --from "$user" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"
	done
	
	sleep 10

	echo ""
	echo "Command:"
	echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde start-erasmus $1 $j --from $user --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
	echo ""
	sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde start-erasmus "$1" "$j" --from "$user" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"
	
	
	while [ "$?" -ne 0 ];
	do
		echo ""
		echo "Wait 10 seconds and run the command again"
		echo ""
		sleep 10
		sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde start-erasmus "$1" "$j" --from "$user" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"
	done
	
	sleep 10

	i=$((i+2))
	j=$((j+1))
	

done

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde end-erasmus-before-deadline $1 3 --from $(sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de  university_chain_ded_i keys --keyring-backend test show "Mike Eisenhauer" -a) --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde end-erasmus-before-deadline "$1" 3 --from "$(sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de  university_chain_ded_i keys --keyring-backend test show "Mike Eisenhauer" -a)"  --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"

while [ "$?" -ne 0 ];
	do
		echo ""
		echo "Wait 10 seconds and run the command again"
		echo ""
		sleep 10
		sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde end-erasmus-before-deadline "$1" 3 --from "$(sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de  university_chain_ded_i keys --keyring-backend test show "Mike Eisenhauer" -a)"  --keyring-backend test --gas auto --chain-id university_chain_de --yes --node "tcp://val-"$uni":26657"
	done



