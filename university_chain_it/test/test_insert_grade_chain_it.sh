#!/bin/bash

export prof_advanced_databases=$(university_chain_itd keys show "Prof. Ermanno Naccari" -a) 

echo ""
echo "Command:"
echo "university_chain_itd tx universitychainit insert-exam-grade unipi 1 "Advanced databases" 25 --from $prof_advanced_databases --chain-id university_chain_it --yes"
echo ""
university_chain_itd tx universitychainit insert-exam-grade unipi 1 "Advanced databases" 25 --from $prof_advanced_databases --chain-id university_chain_it --yes
