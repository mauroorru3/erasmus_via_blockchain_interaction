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


export Mario_Rossi=$(university_chain_itd keys show "Mario Rossi" -a) 
export university=""

if [ "$1" = "unipi" ];
then 
	university=$(university_chain_itd keys show "Unipi" -a) 
elif [ "$1" != "uniroma1" ];
then 
	university=$(university_chain_itd keys show "Uniroma1" -a) 	
fi

export chain_admin=$(university_chain_itd keys show "Admin Chain IT" -a) 

export prof_ae=$(university_chain_itd keys show "Prof. Domenico Asprucci" -a) 

echo ""
echo "Command:"
echo "university_chain_itd tx universitychainit configure-chain --from $chain_admin --gas auto --chain-id university_chain_it --yes"
echo ""
university_chain_itd tx universitychainit configure-chain --from "$chain_admin" --gas auto --chain-id university_chain_it --yes


echo ""
echo "Command:"
echo "university_chain_itd tx universitychainit register-new-student $1 Mario Rossi master cs "Computer Science" --from $Mario_Rossi --gas auto --chain-id university_chain_it --yes"
echo ""
university_chain_itd tx universitychainit register-new-student "$1" Mario Rossi master cs "Computer Science" --from "$Mario_Rossi" --gas auto --chain-id university_chain_it --yes


echo ""
echo "Command:"
echo "university_chain_itd tx universitychainit insert-student-personal-info $1 1 male 1994-06-06 italian italy Rome Rome 1111111111111111 20000 --from $Mario_Rossi --gas auto --chain-id university_chain_it --yes"
echo ""
university_chain_itd tx universitychainit insert-student-personal-info "$1" 1 male 1994-06-06 italian italy Rome Rome 1111111111111111 20000 --from "$Mario_Rossi" --gas auto --chain-id university_chain_it --yes 

echo ""
echo "Command:"
echo "university_chain_itd tx universitychainit insert-student-contact-info $1 1 "via roma" mario.rossi@example.it 0000000000 --from $Mario_Rossi --gas auto --chain-id university_chain_it --yes"
echo ""
university_chain_itd tx universitychainit insert-student-contact-info "$1" 1 "via roma" mario.rossi@example.it 0000000000 --from "$Mario_Rossi" --gas auto --chain-id university_chain_it --yes

echo ""
echo "Command:"
echo "university_chain_itd tx universitychainit insert-student-residence-info $1 1 italy PI Pisa 56100 "via roma" 3 0000000000 --from $Mario_Rossi --gas auto --chain-id university_chain_it --yes"
echo ""
university_chain_itd tx universitychainit insert-student-residence-info "$1" 1 italy PI Pisa 56100 "via roma" 3 0000000000 --from "$Mario_Rossi" --gas auto --chain-id university_chain_it --yes

echo ""
echo "Command:"
echo "university_chain_itd tx universitychainit pay-taxes $1 1 --from $Mario_Rossi --gas auto --chain-id university_chain_it --yes"
echo ""
university_chain_itd tx universitychainit pay-taxes "$1" 1 --from "$Mario_Rossi" --gas auto --chain-id university_chain_it --yes

echo ""
echo "Command:"
echo "university_chain_itd tx universitychainit insert-exam-grade $1 1 "Algorithm engineering" 25 --from $prof_ae --gas auto --chain-id university_chain_it --yes"
echo ""
university_chain_itd tx universitychainit insert-exam-grade "$1" 1 "Algorithm engineering" 25 --from "$prof_ae" --gas auto --chain-id university_chain_it --yes

echo ""
echo "Command:"
echo "university_chain_itd tx universitychainit insert-erasmus-request $1 1 6 $2 study --from $Mario_Rossi --gas auto --chain-id university_chain_it --yes"
echo ""
university_chain_itd tx universitychainit insert-erasmus-request "$1" 1 6 "$2" study --from "$Mario_Rossi" --gas auto --chain-id university_chain_it --yes

echo ""
echo "Command:"
echo "university_chain_itd tx universitychainit insert-erasmus-exam $1 1 "Advanced databases" --from $Mario_Rossi --gas auto --chain-id university_chain_it --yes"
echo ""
university_chain_itd tx universitychainit insert-erasmus-exam "$1" 1 "Advanced databases" --from "$Mario_Rossi" --gas auto --chain-id university_chain_it --yes

echo ""
echo "Command:"
echo "university_chain_itd tx universitychainit start-erasmus $1 1 --from $Mario_Rossi --gas auto --chain-id university_chain_it --yes"
echo ""
university_chain_itd tx universitychainit start-erasmus "$1" 1 --from "$Mario_Rossi" --gas auto --chain-id university_chain_it --yes
