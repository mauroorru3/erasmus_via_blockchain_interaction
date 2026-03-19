#!/bin/bash

#This script aims to delete the data of old chain executions and restore them to their initial state.

sudo rm -rf university_chain_it/elements
sudo mkdir university_chain_it/elements
sudo cp -r university_chain_it/backup_elements/. university_chain_it/elements/
sudo rm -rf university_chain_de/elements
sudo mkdir university_chain_de/elements
sudo cp -r university_chain_de/backup_elements/. university_chain_de/elements/
sudo rm -rf hub/elements
sudo mkdir hub/elements
sudo cp -r hub/backup_elements/. hub/elements/


