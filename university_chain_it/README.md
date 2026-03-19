# Erasmus via blockchain

Erasmus+ programme through the use of Blockchain technology.

The submitted software is a blockchain built using [Cosmos SDK](https://v1.cosmos.network/sdk) and created with [Ignite CLI](https://ignite.com/cli).

The chain enables several possibilities. The main ones are as follows:
- entry of a student with all her/his related information and examinations from her/his study plan;
- assignment of an exam grade by a professor to a student;
- Erasmus+ programme request of a student to another university in a foreign chain.

N.B. the names, surnames and addresses present are fictitious.

## Installation

First of all, you need to install [Docker](https://www.docker.com/), which is a tool that is used to automate the deployment of applications in lightweight containers so that applications can work efficiently in different environments. The installation guide for the Ubuntu platform is available [here](https://docs.docker.com/desktop/install/ubuntu/).

## Usage



```go
# Create the image:

docker build -f Dockerfile-ubuntu . -t university_chain_it_i

# To run the chain:

docker run --rm -it -v $(pwd):/university_chain_it -w /university_chain_it -p 1317:1317 -p 3000:3000 -p 4500:4500 -p 5000:5000 -p 26657:26657 --name university_chain_it university_chain_it_i ignite chain serve --reset-once

# To interact with the running chain open another shell and try: 

docker exec -it university_chain_it university_chain_itd status

# To export some accounts to simplify address insertion:

export Mario_Rossi=$(sudo docker exec university_chain_it university_chain_itd keys show "Mario Rossi" -a) 

export chain_admin=$(sudo docker exec university_chain_it university_chain_itd keys show "Admin Chain IT" -a) 

export prof_ae=$(sudo docker exec university_chain_it university_chain_itd keys show "Prof. Domenico Asprucci" -a) 

# To initialize the chain

sudo docker exec -it university_chain_it university_chain_itd tx universitychainit configure-chain --from $chain_admin --gas auto 

# To show the chain info:

sudo docker exec -it university_chain_it university_chain_itd query universitychainit show-chain-info 

# To show the universities info:

sudo docker exec -it university_chain_it university_chain_itd query universitychainit list-university-info 

# To show foreign universities info:

sudo docker exec -it university_chain_it university_chain_itd query universitychainit list-foreign-universities

# To show the professors and the exams that are held by them:

sudo docker exec -it university_chain_it university_chain_itd query universitychainit list-professors-exams

# To insert a new student:

sudo docker exec -it university_chain_it university_chain_itd tx universitychainit register-new-student unipi Mario Rossi master cs "Computer Science" --from $Mario_Rossi --gas auto  

# To insert student's personal information:

sudo docker exec -it university_chain_it university_chain_itd tx universitychainit insert-student-personal-info unipi 1 male 1994-06-06 italian italy Rome Rome 1111111111111111 20000 --from $Mario_Rossi --gas auto 

# To insert student's contact information:

sudo docker exec -it university_chain_it university_chain_itd tx universitychainit insert-student-contact-info unipi 1 "via roma" mario.rossi@example.it 0000000000 --from $Mario_Rossi --gas auto 

# To insert student's residence information:

sudo docker exec -it university_chain_it university_chain_itd tx universitychainit insert-student-residence-info unipi 1 italy PI Pisa 56100 "via roma" 3 0000000000 --from $Mario_Rossi --gas auto 

# To show the inserted student:

sudo docker exec -it university_chain_it university_chain_itd query universitychainit show-stored-student unipi_1 

# To show the accounts the addresses' budget:

sudo docker exec -it university_chain_it university_chain_itd query bank balances $Mario_Rossi 
sudo docker exec -it university_chain_it university_chain_itd query bank balances $unipi

# To pay the university taxes: 

sudo docker exec -it university_chain_it university_chain_itd tx universitychainit pay-taxes unipi 1 --from $Mario_Rossi --gas auto 

# To show the changes to the budget of the addresses after the payment of taxes:

sudo docker exec -it university_chain_it university_chain_itd query bank balances $Mario_Rossi 
sudo docker exec -it university_chain_it university_chain_itd query bank balances $unipi

# To assign a grade to an exam:

sudo docker exec -it university_chain_it university_chain_itd tx universitychainit insert-exam-grade unipi 1 "Algorithm engineering" 25 --from $prof_ae --gas auto 

# To insert an Erasmus request:

sudo docker exec -it university_chain_it university_chain_itd tx universitychainit insert-erasmus-request unipi 1 6 tum study --from $Mario_Rossi --gas auto 

# To add an Erasmus exam:

sudo docker exec -it university_chain_it university_chain_itd tx universitychainit insert-erasmus-exam unipi 1 "Advanced databases" --from $Mario_Rossi --gas auto 

# To start the Erasmus:

sudo docker exec -it university_chain_it university_chain_itd tx universitychainit start-erasmus unipi 1 --from $Mario_Rossi --gas auto 

# To perform the tests:

sudo docker run --rm -it -v $(pwd):/university_chain_it -w /university_chain_it university_chain_it_i make mock-expected-keepers 

sudo docker run --rm -it -v $(pwd):/university_chain_it -w /university_chain_it university_chain_it_i go test university_chain_it/x/universitychainit/types

sudo docker run --rm -it -v $(pwd):/university_chain_it -w /university_chain_it university_chain_it_i go test university_chain_it/x/universitychainit/keeper

```

## Contributing



## License

[MIT](https://choosealicense.com/licenses/mit/)