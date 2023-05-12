#!/bin/bash


export Mario_Rossi=$(university_chain_itd keys show "Mario Rossi" -a) 

export unipi=$(university_chain_itd keys show Unipi -a) 

export chain_admin=$(university_chain_itd keys show "Admin Chain IT" -a) 

export prof_ae=$(university_chain_itd keys show "Prof. Domenico Asprucci" -a) 

university_chain_itd tx universitychainit configure-chain --from $chain_admin --chain-id university_chain_it --yes

university_chain_itd tx universitychainit register-new-student unipi Mario Rossi master cs "Computer Science" --from $Mario_Rossi --chain-id university_chain_it --yes

university_chain_itd tx universitychainit insert-student-personal-info unipi 1 male 1994-06-06 italian italy Rome Rome 1111111111111111 20000 --from $Mario_Rossi --chain-id university_chain_it --yes 

university_chain_itd tx universitychainit insert-student-contact-info unipi 1 "via roma" mario.rossi@example.it 0000000000 --from $Mario_Rossi --chain-id university_chain_it --yes

university_chain_itd tx universitychainit insert-student-residence-info unipi 1 italy PI Pisa 56100 "via roma" 3 0000000000 --from $Mario_Rossi --chain-id university_chain_it --yes

university_chain_itd tx universitychainit pay-taxes unipi 1 --from $Mario_Rossi --chain-id university_chain_it --yes

university_chain_itd tx universitychainit insert-exam-grade unipi 1 "Algorithm engineering" 25 --from $prof_ae --chain-id university_chain_it --yes

university_chain_itd tx universitychainit insert-erasmus-request unipi 1 6 tum study --from $Mario_Rossi --chain-id university_chain_it --yes

university_chain_itd tx universitychainit insert-erasmus-exam unipi 1 "Advanced databases" --from $Mario_Rossi --chain-id university_chain_it --yes

#university_chain_itd tx universitychainit start-erasmus unipi 1 --from $Mario_Rossi --gas auto 
