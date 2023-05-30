#!/bin/sh

hermes keys add --chain university_chain_de --mnemonic-file "university_chain_de_keys.json"
hermes keys add --chain hub --mnemonic-file "hub_keys.json"

hermes create channel --channel-version hub-1 --a-chain university_chain_de --b-chain hub --a-port universitychainde --b-port hub --new-client-connection --yes
hermes start
