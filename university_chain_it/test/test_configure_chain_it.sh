#!/bin/bash

export chain_admin=$(university_chain_itd keys show "Admin Chain IT" -a) 

echo ""
echo "Command:"
echo "university_chain_itd tx universitychainit configure-chain --from $chain_admin --chain-id university_chain_it --yes"
echo ""
university_chain_itd tx universitychainit configure-chain --from "$chain_admin" --chain-id university_chain_it --yes


