#!/bin/bash

export chain_admin=$(hubd keys show "Admin Hub" -a) 

echo ""
echo "Command:"
echo "hubd tx hub configure-chain --from $chain_admin --chain-id hub --yes"
echo ""
hubd tx hub configure-chain --from $chain_admin --chain-id hub --yes

