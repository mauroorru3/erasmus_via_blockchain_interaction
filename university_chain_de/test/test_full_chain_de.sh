#!/bin/bash


if [ $# -ne 2 ];
	then echo "The number of arguments must be 2. The first argument represents the home university and the second argument the destination university. Usage example: ./test_full_chain_de.sh tum unipi" 
	exit 1
fi

if { [ "$1" != "tum" ] && [ "$1" != "humboldt university" ]; } || { [ "$2" != "unipi" ] && [ "$2" != "uniroma1" ]; };
then 
	echo "The first argument must be equal to unipi or uniroma1 and the second argument must be equal to tum or humboldt university. Usage example: ./test_full_chain_de.sh tum unipi"
	exit 1
fi


uni="tum"
if [ "$1" = "humboldt university" ];
then 
	uni="humboldt"
fi



Karl_Schmidt=$(sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de  university_chain_ded_i keys --keyring-backend test show "Karl Schmidt" -a) 

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
echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde configure-chain --from $chain_admin --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde configure-chain --from "$chain_admin" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657

sleep 20

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde register-new-student $1 Karl Schmidt master cs "Computer Science" --from $Karl_Schmidt --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde register-new-student "$1" Karl Schmidt master cs "Computer Science" --from "$Karl_Schmidt" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657

sleep 20

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-student-personal-info $1 1 male 1998-04-02 German Germany Munich Munich 1111111111111111 20000 --from $Karl_Schmidt --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-student-personal-info "$1" 1 male 1998-04-02 German Germany Munich Munich 1111111111111111 20000 --from "$Karl_Schmidt" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657 

sleep 20

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-student-contact-info $1 1 "Ludwig Street" karl.schmidt@example.it 0000000000 --from $Karl_Schmidt --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-student-contact-info "$1" 1 "Ludwig Street" karl.schmidt@example.it 0000000000 --from "$Karl_Schmidt" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657

sleep 20

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-student-residence-info $1 1 germany MU Munich 80331 "Ludwig Street" 3 0000000000 --from $Karl_Schmidt --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-student-residence-info "$1" 1 germany MU Munich 80331 "Ludwig Street" 3 0000000000 --from "$Karl_Schmidt" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657

sleep 20

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde pay-taxes $1 1 --from $Karl_Schmidt --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde pay-taxes "$1" 1 --from "$Karl_Schmidt" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657

sleep 20

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-exam-grade $1 1 "Algorithm engineering" 1.5 --from $prof_ae --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-exam-grade "$1" 1 "Algorithm engineering" 1.5 --from "$prof_ae" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657

sleep 20

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-erasmus-request $1 1 6 $2 study --from $Karl_Schmidt --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-erasmus-request "$1" 1 6 "$2" study --from "$Karl_Schmidt" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657

sleep 20

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-erasmus-exam $1 1 "Advanced databases" --from $Karl_Schmidt --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde insert-erasmus-exam "$1" 1 "Advanced databases" --from "$Karl_Schmidt" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657

sleep 20

echo ""
echo "Command:"
echo "sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde start-erasmus $1 1 --from $Karl_Schmidt --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657"
echo ""
sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-"$uni":/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i tx universitychainde start-erasmus "$1" 1 --from "$Karl_Schmidt" --keyring-backend test --gas auto --chain-id university_chain_de --yes --node tcp://val-"$uni":26657


