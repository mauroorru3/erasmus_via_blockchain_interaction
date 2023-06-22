#!/bin/bash

export Mario_Rossi=$(university_chain_ded keys show "Mario Rossi" -a) 

echo ""
echo "Command:"
echo "university_chain_ded tx universitychainde insert-erasmus-request tum 1 6 unipi study --from $Mario_Rossi --chain-id university_chain_de --yes"
echo ""
university_chain_ded tx universitychainde insert-erasmus-request tum 1 6 unipi study --from $Mario_Rossi --chain-id university_chain_de --yes
