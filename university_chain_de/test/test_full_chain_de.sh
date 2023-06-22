#!/bin/bash


export Karl_Schmidt=$(university_chain_ded keys show "Karl Schmidt" -a) 

export tum=$(university_chain_ded keys show "TUM_University" -a) 

export chain_admin=$(university_chain_ded keys show "Admin Chain DE" -a) 

export prof_ae=$(university_chain_ded keys show "Prof. Friedrich Mayer" -a) 

echo ""
echo "Command:"
echo "university_chain_ded tx universitychainde configure-chain --from $chain_admin --chain-id university_chain_de --yes"
echo ""
university_chain_ded tx universitychainde configure-chain --from $chain_admin --chain-id university_chain_de --yes


echo ""
echo "Command:"
echo "university_chain_ded tx universitychainde register-new-student tum Karl Schmidt master cs "Computer Science" --from $Karl_Schmidt --chain-id university_chain_de --yes"
echo ""
university_chain_ded tx universitychainde register-new-student tum Karl Schmidt master cs "Computer Science" --from $Karl_Schmidt --chain-id university_chain_de --yes


echo ""
echo "Command:"
echo "university_chain_ded tx universitychainde insert-student-personal-info tum 1 male 1998-04-02 German Germany Munich Munich 1111111111111111 20000 --from $Karl_Schmidt --chain-id university_chain_de --yes"
echo ""
university_chain_ded tx universitychainde insert-student-personal-info tum 1 male 1998-04-02 German Germany Munich Munich 1111111111111111 20000 --from $Karl_Schmidt --chain-id university_chain_de --yes 

echo ""
echo "Command:"
echo "university_chain_ded tx universitychainde insert-student-contact-info tum 1 "Ludwig Street" karl.schmidt@example.it 0000000000 --from $Karl_Schmidt --chain-id university_chain_de --yes"
echo ""
university_chain_ded tx universitychainde insert-student-contact-info tum 1 "Ludwig Street" karl.schmidt@example.it 0000000000 --from $Karl_Schmidt --chain-id university_chain_de --yes

echo ""
echo "Command:"
echo "university_chain_ded tx universitychainde insert-student-residence-info tum 1 germany MU Munich 80331 "Ludwig Street" 3 0000000000 --from $Karl_Schmidt --chain-id university_chain_de --yes"
echo ""
university_chain_ded tx universitychainde insert-student-residence-info tum 1 germany MU Munich 80331 "Ludwig Street" 3 0000000000 --from $Karl_Schmidt --chain-id university_chain_de --yes

echo ""
echo "Command:"
echo "university_chain_ded tx universitychainde pay-taxes tum 1 --from $Karl_Schmidt --chain-id university_chain_de --yes"
echo ""
university_chain_ded tx universitychainde pay-taxes tum 1 --from $Karl_Schmidt --chain-id university_chain_de --yes

echo ""
echo "Command:"
echo "university_chain_ded tx universitychainde insert-exam-grade tum 1 "Algorithm engineering" 1.5 --from $prof_ae --chain-id university_chain_de --yes"
echo ""
university_chain_ded tx universitychainde insert-exam-grade tum 1 "Algorithm engineering" 1.5 --from $prof_ae --chain-id university_chain_de --yes


echo ""
echo "Command:"
echo "university_chain_ded tx universitychainde insert-erasmus-request tum 1 6 unipi study --from $Karl_Schmidt --chain-id university_chain_de --yes"
echo ""
university_chain_ded tx universitychainde insert-erasmus-request tum 1 6 unipi study --from $Karl_Schmidt --chain-id university_chain_de --yes

echo ""
echo "Command:"
echo "university_chain_ded tx universitychainde insert-erasmus-exam tum 1 "Advanced databases" --from $Karl_Schmidt --chain-id university_chain_de --yes"
echo ""
university_chain_ded tx universitychainde insert-erasmus-exam tum 1 "Advanced databases" --from $Karl_Schmidt --chain-id university_chain_de --yes

echo ""
echo "Command:"
echo "university_chain_ded tx universitychainde start-erasmus tum 1 --from $Karl_Schmidt --chain-id university_chain_de --yes"
echo ""
university_chain_ded tx universitychainde start-erasmus tum 1 --from $Karl_Schmidt --chain-id university_chain_de --yes
