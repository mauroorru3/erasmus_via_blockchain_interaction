#!/bin/bash

echo ""
echo "1"
echo ""
sudo mkdir -p elements/val-tum
echo ""
echo "2"
echo ""
sudo mkdir -p elements/val-humboldt
echo ""
echo "3"
echo ""
echo -e val-tum'\n'val-humboldt | xargs -I {} sudo docker run --rm -i -v $(pwd)/elements/{}:/root/.university_chain_de university_chain_ded_i init university_chain_de
echo ""
echo "4"
echo ""
sudo cp -r ./dir_keys/keyring-test $(pwd)/elements/val-humboldt/
echo ""
echo "5"
echo ""
sudo cp -r ./dir_keys/keyring-test $(pwd)/elements/val-tum/
echo ""
#Copy manually the address and balances information in the genesis file
