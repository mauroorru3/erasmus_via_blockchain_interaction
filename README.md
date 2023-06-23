# Erasmus via blockchain (Interaction between the Hub and the Italian Chain)

In this repo we focus on the interaction between the Hub and the Italian Chain.

Refers to this [link](https://github.com/mauroorru3/erasmus_via_blockchain) for the Italian Chain and to this [link](https://github.com/mauroorru3/-erasmus_via_blockchain_hub) for the hub.

## Installation

As above.

## Usage



Create the image of the Italian chain:

**cd university_chain_it**
**sudo docker build -f Dockerfile-university_chain_it . -t university_chain_it --no-cache**
**cd ..**

Create the image of the German chain:

**cd university_chain_de**
**sudo docker build -f Dockerfile-university_chain_de . -t university_chain_de --no-cache**
**cd ..**

Create the image of the Hub chain:

**cd hub**
**sudo docker build -f Dockerfile-hub . -t hub --no-cache**
**cd ..**

Create the image of the Relayer that connect the Italian chain with the hub chain:
Go relayer:

**cd go_relayer_it_hub**
**sudo docker build -f Dockerfile . -t go_relayer_it_hub --no-cache**
**cd ..**

Hermes relayer:

**cd hermes_relayer_it_hub**
**sudo docker build -f Dockerfile . -t hermes_relayer_it_hub --no-cache**
**cd ..**

Create the image of the Relayer that connect the German chain with the hub chain:
Go relayer:

**cd go_relayer_de_hub**
**sudo docker build -f Dockerfile . -t go_relayer_de_hub --no-cache** 
**cd ..**

Hermes relayer:

**cd hermes_relayer_de_hub**
**sudo docker build -f Dockerfile . -t hermes_relayer_de_hub --no-cache**
**cd ..** 

To allow scripts to be executed within containers:

**chmod +x allow_permissions.sh** 
**./allow_permissions.sh**

To start the containers with hermes relayers:

**sudo docker compose -f modular.yaml --profile hermes up**

To start the containers with go relayers:

**sudo docker compose -f modular.yaml --profile go up**

_Remember_ that the relayers must be executed when the configuration of the chains has been completed, otherwise there will be errors.
To start the hermes relayer process:
_Remember_ that the order of execution is important. 
First the relayer related to the Chain IT and the Hub must be started, and then the relayer related to the Chain DE and the Hub.
The execution must take place in a new terminal for the IT-Hub relayer and the DE-Hub relayer.

**sudo docker exec -it hermes_relayer_it_hub bash**
**./run-relayer.sh**

**sudo docker exec -it hermes_relayer_de_hub bash**
**./run-relayer.sh**

To start the go relayer process:

**sudo docker exec go_relayer_it_hub ./run-relayer_it_hub.sh**

**sudo docker exec go_relayer_de_hub ./run-relayer_de_hub.sh**


To simulate the execution:
_Remember_ that the execution of the tests must be started when the configuration of the relayers has been completed, otherwise there will be errors in communication between the chains.
Open a new terminal and run: 

**./test_chain.sh it unipi tum**

to run the tests in the italian chain, starting from the Unipi university and carrying out the Erasmus programme at TUM University
(there are also other possibilities, such as **./test_chain.sh de tum uniroma1** and so on).
Universities in the Italian chain: **unipi** and **uniroma1**.
Universities in the German chain: **tum** and **humboldt university**.

The tests that will be performed are the following (in the specific case of the **./test_chain.sh it unipi tum** command):
- Hub chain configuration;
- DE chain configuration;
- IT chain configuration;
- register new student;
- insert student's personal info;
- insert student's contact info;
- insert student's residence info;
- pay student's taxes;
- insert exam grade;
- insert Erasmus request;
- insert Erasmus exam;
- start Erasmus.

To see the updated data in the home chain:

**sudo docker exec -it university_chain_it university_chain_itd query universitychainit show-stored-student unipi_1**

To see the updated data in the destination chain:

**sudo docker exec -it university_chain_de university_chain_ded query universitychainde show-stored-student tum_1**


Once the transfer of information between the chains is complete, the following test can be performed:

**sudo docker exec university_chain_de ./test/test_insert_grade_chain_de.sh tum**

The test will check the insertion of the exam grade in the Erasmus plan.
Once the Erasmus is finished (I have set the timer temporarily to 240 seconds for the tests), the updated data on the Erasmus plan will be transferred back to the home university, where the grade conversion will be carried out.

To see the updated data in the home chain:

**sudo docker exec -it university_chain_it university_chain_itd query universitychainit show-stored-student unipi_1**

To end the execution of the containers with hermes relayers:

**sudo docker compose -f modular.yaml --profile hermes down**

To end the execution of the containers with go relayers:

**sudo docker compose -f modular.yaml --profile go down**


## Contributing


## License

[MIT](https://choosealicense.com/licenses/mit/)
