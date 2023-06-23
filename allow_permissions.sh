#!/bin/bash


checkReturnValue(){
  if [ $? -ne 0 ]; then
  echo "Command \"chmod a+rwx $1\" failed" >&2 
  fi
}


scriptFileArray=( "hub/run-hub.sh" "university_chain_it/run-university_chain_it.sh" "university_chain_de/run-university_chain_de.sh" "university_chain_it/test/test_configure_chain_it.sh" "university_chain_it/test/test_erasmus_request_chain_it.sh" "test_erasmus_request_chain_it_foreign.sh" "university_chain_it/test/test_full_chain_it.sh" "university_chain_it/test/test_insert_grade_chain_it.sh" "university_chain_de/test/test_configure_chain_de.sh" "university_chain_de/test/test_erasmus_request_chain_de.sh" "test_erasmus_request_chain_de_foreign.sh" "university_chain_de/test/test_full_chain_de.sh" "university_chain_de/test/test_insert_grade_chain_de.sh" "hub/test/test_configure_chain_hub.sh" "hermes_relayer_it_hub/run-relayer.sh" "hermes_relayer_de_hub/run-relayer.sh" "go_relayer_it_hub/run-relayer_it_hub.sh" "go_relayer_de_hub/run-relayer_de_hub.sh" "test_chain.sh" )


for i in "${scriptFileArray[@]}"
do
    chmod a+rwx "$i"
    checkReturnValue "$i"
done



