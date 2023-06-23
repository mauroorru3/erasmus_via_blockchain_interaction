#!/bin/bash

if [ $# -ne 1 ];
	then echo "The number of arguments must be 1. The argument represents the university to be tested. Usage example: ./test_insert_grade_chain_de.sh tum" 
	exit 1
fi

if [ "$1" != "tum" ] && [ "$1" != "humboldt university" ];
then 
	echo "The argument must be equal to tum or Humboldt University. Usage example: ./test_insert_grade_chain_de.sh tum"
	exit 1
fi


export prof_advanced_databases=$(university_chain_ded keys show "Prof. Gustav Fischer" -a) 

echo ""
echo "Command:"
echo "university_chain_ded tx universitychainde insert-exam-grade $1 1 "Advanced databases" 2.2 --from $prof_advanced_databases --chain-id university_chain_de --yes"
echo ""
university_chain_ded tx universitychainde insert-exam-grade "$1" 1 "Advanced databases" 2.2 --from "$prof_advanced_databases" --chain-id university_chain_de --yes

