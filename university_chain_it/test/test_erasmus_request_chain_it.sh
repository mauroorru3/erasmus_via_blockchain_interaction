#!/bin/bash

export Karl_Schmidt=$(university_chain_itd keys show "Karl Schmidt" -a) 

echo ""
echo "Command:"
echo "university_chain_itd tx universitychainit insert-erasmus-request unipi 1 6 tum study --from $Karl_Schmidt --chain-id university_chain_it --yes"
echo ""
university_chain_itd tx universitychainit insert-erasmus-request unipi 1 6 tum study --from $Karl_Schmidt --chain-id university_chain_it --yes
