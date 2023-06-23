#!/bin/bash


if [ $# -ne 2 ];
	then echo "The number of arguments must be 2. The first argument represents the home university and the second argument the destination university. Usage example: ./test_erasmus_request_chain_from_it_to_de.sh unipi tum" 
	exit 1
fi

if { [ "$1" != "unipi" ] && [ "$1" != "uniroma1" ]; } || { [ "$2" != "tum" ] && [ "$2" != "humboldt university" ]; };
then 
	echo "The first argument must be equal to unipi or uniroma1 and the second argument must be equal to tum or humboldt university. Usage example: ./test_erasmus_request_chain_from_it_to_de.sh unipi tum"
	exit 1
fi


export Karl_Schmidt=$(university_chain_ded keys show "Karl Schmidt" -a) 
echo ""
echo "Command:"
echo "university_chain_itd tx universitychainit insert-erasmus-request $1 1 6 $2 study --from $Karl_Schmidt --chain-id university_chain_it --yes"
echo ""
university_chain_itd tx universitychainit insert-erasmus-request "$1" 1 6 "$2" study --from "$Karl_Schmidt" --chain-id university_chain_it --yes


