#!/bin/bash


if [ $# -ne 1 ];
	then echo "The number of arguments must be 1. The argument represents the university to be tested. Usage example: ./test_insert_grade_chain_it.sh.sh unipi" 
	exit 1
fi

if [ "$1" != "unipi" ] && [ "$1" != "uniroma1" ];
then 
	echo "The argument must be equal to unipi or uniroma1. Usage example: ./test_insert_grade_chain_it.sh.sh unipi"
	exit 1
fi


export prof_advanced_databases=$(university_chain_itd keys show "Prof. Ermanno Naccari" -a) 

echo ""
echo "Command:"
echo "university_chain_itd tx universitychainit insert-exam-grade $1 1 "Advanced databases" 25 --from $prof_advanced_databases --gas auto --chain-id university_chain_it --yes"
echo ""
university_chain_itd tx universitychainit insert-exam-grade "$1" 1 "Advanced databases" 25 --from "$prof_advanced_databases" --gas auto --chain-id university_chain_it --yes
