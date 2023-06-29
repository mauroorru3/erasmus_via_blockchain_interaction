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


usersArray=( $(university_chain_itd keys show "Mario Rossi" -a) $(university_chain_itd keys show "Tamara Genovesi" -a) 
	$(university_chain_itd keys show "Franca Ferrari" -a) $(university_chain_itd keys show "Luigi Lena" -a) 
	$(university_chain_itd keys show "Gianpiero Serafino" -a) $(university_chain_itd keys show "Antonella Marino" -a))
nameAndSurnameArray=("Mario" "Rossi" "Tamara" "Genovesi" "Franca" "Ferrari" "Luigi" "Lena" "Gianpiero" "Serafino" "Antonella" "Marino")
i=0
j=1


for user in "${usersArray[@]}"
do
	echo ""
	echo "Command:"
	echo "university_chain_itd tx universitychainit register-new-student $1 ${nameAndSurnameArray[$i]} ${nameAndSurnameArray[$i+1]} master cs "Computer Science" --from $user --gas auto --chain-id university_chain_it --yes"
	echo ""
	university_chain_itd tx universitychainit register-new-student "$1" "${nameAndSurnameArray[$i]}" "${nameAndSurnameArray[$i+1]}" master cs "Computer Science" --from "$user" --gas auto --chain-id university_chain_it --yes


	echo ""
	echo "Command:"
	echo "university_chain_itd tx universitychainit insert-student-personal-info $1 $j male 1994-06-06 italian italy Rome Rome 1111111111111111 20000 --from $user --gas auto --chain-id university_chain_it --yes"
	echo ""
	university_chain_itd tx universitychainit insert-student-personal-info "$1" "$j" male 1994-06-06 italian italy Rome Rome 1111111111111111 20000 --from "$user" --gas auto --chain-id university_chain_it --yes 

	echo ""
	echo "Command:"
	echo "university_chain_itd tx universitychainit insert-student-contact-info $1 $j "via roma" mario.rossi@example.it 0000000000 --from $user --gas auto --chain-id university_chain_it --yes"
	echo ""
	university_chain_itd tx universitychainit insert-student-contact-info "$1" "$j" "via roma" mario.rossi@example.it 0000000000 --from "$user" --gas auto --chain-id university_chain_it --yes

	echo ""
	echo "Command:"
	echo "university_chain_itd tx universitychainit insert-student-residence-info $1 $j italy PI Pisa 56100 "via roma" 3 0000000000 --from $user --gas auto --chain-id university_chain_it --yes"
	echo ""
	university_chain_itd tx universitychainit insert-student-residence-info "$1" "$j" italy PI Pisa 56100 "via roma" 3 0000000000 --from "$user" --gas auto --chain-id university_chain_it --yes

	echo ""
	echo "Command:"
	echo "university_chain_itd tx universitychainit pay-taxes $1 $j --from $user --gas auto --chain-id university_chain_it --yes"
	echo ""
	university_chain_itd tx universitychainit pay-taxes "$1" "$j" --from "$user" --gas auto --chain-id university_chain_it --yes

	echo ""
	echo "Command:"
	echo "university_chain_itd tx universitychainit insert-exam-grade $1 $j "Algorithm engineering" 25 --from $prof_ae --gas auto --chain-id university_chain_it --yes"
	echo ""
	university_chain_itd tx universitychainit insert-exam-grade "$1" "$j" "Algorithm engineering" 25 --from "$prof_ae" --gas auto --chain-id university_chain_it --yes

	echo ""
	echo "Command:"
	echo "university_chain_itd tx universitychainit insert-erasmus-request $1 $j 6 $2 study --from $user --gas auto --chain-id university_chain_it --yes"
	echo ""
	university_chain_itd tx universitychainit insert-erasmus-request "$1" "$j" 6 "$2" study --from "$user" --gas auto --chain-id university_chain_it --yes

	echo ""
	echo "Command:"
	echo "university_chain_itd tx universitychainit insert-erasmus-exam $1 $j "Advanced databases" --from $user --gas auto --chain-id university_chain_it --yes"
	echo ""
	university_chain_itd tx universitychainit insert-erasmus-exam "$1" "$j" "Advanced databases" --from "$user" --gas auto --chain-id university_chain_it --yes

	echo ""
	echo "Command:"
	echo "university_chain_itd tx universitychainit start-erasmus $1 $j --from $user --gas auto --chain-id university_chain_it --yes"
	echo ""
	university_chain_itd tx universitychainit start-erasmus "$1" "$j" --from "$user" --gas auto --chain-id university_chain_it --yes

	i=$((i+2))
	j=$((j+1))

done

echo ""
echo "Command:"
echo "university_chain_itd tx universitychainit end-erasmus-before-deadline $1 3 --from $(university_chain_itd keys show "Franca Ferrari" -a) --gas auto --chain-id university_chain_it --yes"
echo ""
university_chain_itd tx universitychainit end-erasmus-before-deadline "$1" 3 --from "$(university_chain_itd keys show "Franca Ferrari" -a)"  --gas auto --chain-id university_chain_it --yes





