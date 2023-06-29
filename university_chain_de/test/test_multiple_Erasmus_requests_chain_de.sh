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


export university=""

if [ "$1" = "tum" ];
then 
	university=$(university_chain_ded keys show "TUM" -a) 
elif [ "$1" != "humboldt university" ];
then 
	university=$(university_chain_ded keys show "Humboldt University" -a) 	
fi

export chain_admin=$(university_chain_ded keys show "Admin Chain DE" -a) 

export prof_ae=$(university_chain_ded keys show "Prof. Friedrich Mayer" -a) 


echo ""
echo "Command:"
echo "university_chain_ded tx universitychainde configure-chain --from $chain_admin --gas auto --chain-id university_chain_de --yes"
echo ""
university_chain_ded tx universitychainde configure-chain --from "$chain_admin" --gas auto --chain-id university_chain_de --yes


usersArray=( $(university_chain_ded keys show "Karl Schmidt" -a) $(university_chain_ded keys show "Paul Becker" -a) 
	$(university_chain_ded keys show "Mike Eisenhauer" -a) $(university_chain_ded keys show "Daniela Gloeckner" -a) 
	$(university_chain_ded keys show "Nadine Baier" -a) $(university_chain_ded keys show "Jennifer Diederich" -a))
nameAndSurnameArray=("Karl" "Schmidt" "Paul" "Becker" "Mike" "Eisenhauer" "Daniela" "Gloeckner" "Nadine" "Baier" "Jennifer" "Diederich")
i=0
j=1


for user in "${usersArray[@]}"
do
	echo ""
	echo "Command:"
	echo "university_chain_ded tx universitychainde register-new-student $1 ${nameAndSurnameArray[$i]} ${nameAndSurnameArray[$i+1]} master cs "Computer Science" --from $user --gas auto --chain-id university_chain_de --yes"
	echo ""
	university_chain_ded tx universitychainde register-new-student "$1" "${nameAndSurnameArray[$i]}" "${nameAndSurnameArray[$i+1]}" master cs "Computer Science" --from "$user" --gas auto --chain-id university_chain_de --yes


	echo ""
	echo "Command:"
	echo "university_chain_ded tx universitychainde insert-student-personal-info $1 $j male 1994-06-06 italian italy Rome Rome 1111111111111111 20000 --from $user --gas auto --chain-id university_chain_de --yes"
	echo ""
	university_chain_ded tx universitychainde insert-student-personal-info "$1" "$j" male 1994-06-06 italian italy Rome Rome 1111111111111111 20000 --from "$user" --gas auto --chain-id university_chain_de --yes 

	echo ""
	echo "Command:"
	echo "university_chain_ded tx universitychainde insert-student-contact-info $1 $j "via roma" mario.rossi@example.it 0000000000 --from $user --gas auto --chain-id university_chain_de --yes"
	echo ""
	university_chain_ded tx universitychainde insert-student-contact-info "$1" "$j" "via roma" mario.rossi@example.it 0000000000 --from "$user" --gas auto --chain-id university_chain_de --yes

	echo ""
	echo "Command:"
	echo "university_chain_ded tx universitychainde insert-student-residence-info $1 $j italy PI Pisa 56100 "via roma" 3 0000000000 --from $user --gas auto --chain-id university_chain_de --yes"
	echo ""
	university_chain_ded tx universitychainde insert-student-residence-info "$1" "$j" italy PI Pisa 56100 "via roma" 3 0000000000 --from "$user" --gas auto --chain-id university_chain_de --yes

	echo ""
	echo "Command:"
	echo "university_chain_ded tx universitychainde pay-taxes $1 $j --from $user --gas auto --chain-id university_chain_de --yes"
	echo ""
	university_chain_ded tx universitychainde pay-taxes "$1" "$j" --from "$user" --gas auto --chain-id university_chain_de --yes

	echo ""
	echo "Command:"
	echo "university_chain_ded tx universitychainde insert-exam-grade $1 $j "Algorithm engineering" 1.5 --from $prof_ae --gas auto --chain-id university_chain_de --yes"
	echo ""
	university_chain_ded tx universitychainde insert-exam-grade "$1" "$j" "Algorithm engineering" 1.5 --from "$prof_ae" --gas auto --chain-id university_chain_de --yes

	echo ""
	echo "Command:"
	echo "university_chain_ded tx universitychainde insert-erasmus-request $1 $j 6 $2 study --from $user --gas auto --chain-id university_chain_de --yes"
	echo ""
	university_chain_ded tx universitychainde insert-erasmus-request "$1" "$j" 6 "$2" study --from "$user" --gas auto --chain-id university_chain_de --yes

	echo ""
	echo "Command:"
	echo "university_chain_ded tx universitychainde insert-erasmus-exam $1 $j "Advanced databases" --from $user --gas auto --chain-id university_chain_de --yes"
	echo ""
	university_chain_ded tx universitychainde insert-erasmus-exam "$1" "$j" "Advanced databases" --from "$user" --gas auto --chain-id university_chain_de --yes

	echo ""
	echo "Command:"
	echo "university_chain_ded tx universitychainde start-erasmus $1 $j --from $user --gas auto --chain-id university_chain_de --yes"
	echo ""
	university_chain_ded tx universitychainde start-erasmus "$1" "$j" --from "$user" --gas auto --chain-id university_chain_de --yes

	i=$((i+2))
	j=$((j+1))

done

echo ""
echo "Command:"
echo "university_chain_ded tx universitychainde end-erasmus-before-deadline $1 3 --from $(university_chain_ded keys show "Mike Eisenhauer" -a) --gas auto --chain-id university_chain_de --yes"
echo ""
university_chain_ded tx universitychainde end-erasmus-before-deadline "$1" 3 --from "$(university_chain_ded keys show "Mike Eisenhauer" -a)"  --gas auto --chain-id university_chain_de --yes





