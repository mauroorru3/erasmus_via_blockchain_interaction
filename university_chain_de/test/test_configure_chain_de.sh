#!/bin/bash

export chain_admin=$(university_chain_ded keys show "Admin Chain DE" -a) 

echo ""
echo "Command:"
echo "university_chain_ded tx universitychainde configure-chain --from $chain_admin --gas auto --chain-id university_chain_de --yes"
echo ""
university_chain_ded tx universitychainde configure-chain --from "$chain_admin" --gas auto --chain-id university_chain_de --yes


