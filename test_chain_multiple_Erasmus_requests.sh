#!/bin/bash

if [ $# -ne 3 ];
	then echo "The number of arguments must be 3. The first represents the chain, the second the home university and the third the destination university. Usage example: ./test_chain_multiple_Erasmus_requests.sh it unipi tum"; 
	exit 1
fi


if [ "$1" = "it" ];
then  
	if { [ "$2" = "unipi" ] || [ "$2" = "uniroma1" ]; } && { [ "$3" = "tum" ] || [ "$3" = "humboldt university" ]; };
	then
		sudo rm university_chain_it/data/logs.txt
		sudo rm university_chain_de/data/logs.txt
		sudo rm hub/data/logs.txt
		sudo rm hub/data/data.txt
		sudo docker exec hub ./test/test_configure_chain_hub.sh
		sudo docker exec university_chain_de ./test/test_configure_chain_de.sh
		sudo docker exec university_chain_it ./test/test_multiple_Erasmus_requests_chain_it.sh "$2" "$3"
	else
		echo "The second argument is not equal to unipi or uniroma1 or the third element is not equal to tum or humboldt university. Usage example: ./test_chain_multiple_Erasmus_requests.sh it unipi tum"
		exit 1
	fi
elif [ "$1" = "de" ];
then  
	if { [ "$2" = "tum" ] || [ "$2" = "humboldt university" ]; } && { [ "$3" = "unipi" ] || [ "$3" = "uniroma1" ]; };
	then
		sudo rm university_chain_it/data/logs.txt
		sudo rm university_chain_de/data/logs.txt
		sudo rm hub/data/logs.txt
		sudo rm hub/data/data.txt
		sudo docker exec hub ./test/test_configure_chain_hub.sh
		sudo docker exec university_chain_it ./test/test_configure_chain_it.sh 
		sudo docker exec university_chain_de ./test/test_multiple_Erasmus_requests_chain_de.sh "$2" "$3"
	else
		echo "The second argument is not equal to tum or humboldt university or the third element is not equal to unipi or uniroma1. Usage example: ./test_chain_multiple_Erasmus_requests.sh de tum unipi"
		exit 1
	fi
else
	echo "The first argument is not equal to it or de. Usage example: ./test_chain_multiple_Erasmus_requests.sh it unipi tum"
fi



















