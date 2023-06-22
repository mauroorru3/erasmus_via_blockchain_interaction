#!/bin/bash

export prof_advanced_databases=$(university_chain_ded keys show "Prof. Gustav Fischer" -a) 

echo ""
echo "Command:"
echo "university_chain_ded tx universitychainde insert-exam-grade tum 1 "Advanced databases" 25 --from $prof_advanced_databases --chain-id university_chain_de --yes"
echo ""
university_chain_ded tx universitychainde insert-exam-grade tum 1 "Advanced databases" 2.2 --from $prof_advanced_databases --chain-id university_chain_de --yes

