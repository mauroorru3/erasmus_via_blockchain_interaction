#!/bin/bash

if [ $# -ne 3 ];
	then echo "The number of arguments must be 3. The first represents the chain, the second the home university and the third the destination university. Usage example: ./test_chain_multiple_Erasmus_requests.sh it unipi tum"; 
	exit 1
fi


if [ "$1" = "it" ];
then  
	if { [ "$2" = "unipi" ] || [ "$2" = "uniroma1" ]; } && { [ "$3" = "tum" ] || [ "$3" = "humboldt university" ]; };
	then
		sudo rm university_chain_it/log/val-unipi/logs.txt
		sudo rm university_chain_it/log/val-uniroma1/logs.txt
		sudo rm university_chain_it/log/val-unipi/data.txt
		sudo rm university_chain_it/log/val-uniroma1/data.txt
		sudo rm university_chain_it/log/val-unipi/statsGasConsumed.txt
		sudo rm university_chain_it/log/val-unipi/statsTiming.txt
		sudo rm university_chain_it/log/val-uniroma1/statsGasConsumed.txt
		sudo rm university_chain_it/log/val-uniroma1/statsTiming.txt
		sudo rm university_chain_de/log/val-tum/logs.txt
		sudo rm university_chain_de/log/val-humboldt/logs.txt
		sudo rm university_chain_de/log/val-tum/data.txt
		sudo rm university_chain_de/log/val-humboldt/data.txt
		sudo rm university_chain_de/log/val-tum/statsGasConsumed.txt
		sudo rm university_chain_de/log/val-tum/statsTiming.txt
		sudo rm university_chain_de/log/val-humboldt/statsGasConsumed.txt
		sudo rm university_chain_de/log/val-humboldt/statsTiming.txt
		sudo rm hub/log/val-hub-instance-1/logs.txt
		sudo rm hub/log/val-hub-instance-1/data.txt
		sudo rm hub/log/val-hub-instance-1/statsGasConsumed.txt
		sudo rm hub/log/val-hub-instance-1/statsTiming.txt
		sudo rm hub/log/val-hub-instance-2/logs.txt
		sudo rm hub/log/val-hub-instance-2/data.txt
		sudo rm hub/log/val-hub-instance-2/statsGasConsumed.txt
		sudo rm hub/log/val-hub-instance-2/statsTiming.txt
		sudo ./hub/test/test_configure_chain_hub.sh
		sudo ./university_chain_de/test/test_configure_chain_de.sh "$3"
		sudo ./university_chain_it/test/test_multiple_Erasmus_requests_chain_it.sh "$2" "$3"
	else
		echo "The second argument is not equal to unipi or uniroma1 or the third element is not equal to tum or humboldt university. Usage example: ./test_chain_multiple_Erasmus_requests.sh it unipi tum"
		exit 1
	fi
elif [ "$1" = "de" ];
then  
	if { [ "$2" = "tum" ] || [ "$2" = "humboldt university" ]; } && { [ "$3" = "unipi" ] || [ "$3" = "uniroma1" ]; };
	then
		sudo rm university_chain_it/log/val-unipi/logs.txt
		sudo rm university_chain_it/log/val-uniroma1/logs.txt
		sudo rm university_chain_it/log/val-unipi/data.txt
		sudo rm university_chain_it/log/val-uniroma1/data.txt
		sudo rm university_chain_it/log/val-unipi/statsGasConsumed.txt
		sudo rm university_chain_it/log/val-unipi/statsTiming.txt
		sudo rm university_chain_it/log/val-uniroma1/statsGasConsumed.txt
		sudo rm university_chain_it/log/val-uniroma1/statsTiming.txt
		sudo rm university_chain_de/log/val-tum/logs.txt
		sudo rm university_chain_de/log/val-humboldt/logs.txt
		sudo rm university_chain_de/log/val-tum/data.txt
		sudo rm university_chain_de/log/val-humboldt/data.txt
		sudo rm university_chain_de/log/val-tum/statsGasConsumed.txt
		sudo rm university_chain_de/log/val-tum/statsTiming.txt
		sudo rm university_chain_de/log/val-humboldt/statsGasConsumed.txt
		sudo rm university_chain_de/log/val-humboldt/statsTiming.txt
		sudo rm hub/log/val-hub-instance-1/logs.txt
		sudo rm hub/log/val-hub-instance-1/data.txt
		sudo rm hub/log/val-hub-instance-1/statsGasConsumed.txt
		sudo rm hub/log/val-hub-instance-1/statsTiming.txt
		sudo rm hub/log/val-hub-instance-2/logs.txt
		sudo rm hub/log/val-hub-instance-2/data.txt
		sudo rm hub/log/val-hub-instance-2/statsGasConsumed.txt
		sudo rm hub/log/val-hub-instance-2/statsTiming.txt
		sudo ./hub/test/test_configure_chain_hub.sh
		sudo ./university_chain_it/test/test_configure_chain_it.sh "$3"
		sudo ./university_chain_de/test/test_multiple_Erasmus_requests_chain_de.sh "$2" "$3"
	else
		echo "The second argument is not equal to tum or humboldt university or the third element is not equal to unipi or uniroma1. Usage example: ./test_chain_multiple_Erasmus_requests.sh de tum unipi"
		exit 1
	fi
else
	echo "The first argument is not equal to it or de. Usage example: ./test_chain_multiple_Erasmus_requests.sh it unipi tum"
fi



















