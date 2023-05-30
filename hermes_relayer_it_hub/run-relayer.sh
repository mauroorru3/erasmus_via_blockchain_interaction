#!/bin/sh

hermes keys add --chain university_chain_it --mnemonic-file "university_chain_it_keys.json"
hermes keys add --chain hub --mnemonic-file "hub_keys.json"

hermes create channel --channel-version hub-1 --a-chain university_chain_it --b-chain hub --a-port universitychainit --b-port hub --new-client-connection --yes
hermes start
