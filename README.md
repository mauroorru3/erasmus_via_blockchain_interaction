# Erasmus via blockchain (Interaction between the Hub and the Italian Chain)

In this repo we focus on the interaction between the Hub and the Italian Chain.

Refers to this [link](https://github.com/mauroorru3/erasmus_via_blockchain) for the Italian Chain and to this [link](https://github.com/mauroorru3/-erasmus_via_blockchain_hub) for the hub.

## Installation

As above.

## Usage

To allow scripts to be executed within containers:

```bash
chmod +x allow_permissions.sh
```
```bash
./allow_permissions.sh
```

To create all the images we need:

- Italian chain;
- German chain;
- Hub chain;
- Go relayer that connect the Italian chain with the Hub chain;
- Hermes relayer that connect the Italian chain with the Hub chain;
- Go relayer that connect the German chain with the Hub chain;
- Hermes relayer that connect the German chain with the Hub chain;

we execute the following script:
```bash
./create_images.sh
```

To run the initial configuration (it must be executed before running all the containers):

```bash
./start_configuration.sh
```

To start the containers with hermes relayers:

```bash
sudo docker compose --file docker-compose.yml --project-name university_chain-prod --profile hermes up
```

To start the containers with go relayers:

```bash
sudo docker compose --file docker-compose.yml --project-name university_chain-prod --profile go up
```

_Remember_ that the relayers must be executed when the configuration of the chains has been completed, otherwise there will be errors.
To start the hermes relayer process:
_Remember_ that the order of execution is important. 
First the relayer related to the Chain IT and the Hub must be started, and then the relayer related to the Chain DE and the Hub.
The execution must take place in a new terminal for the IT-Hub relayer and the DE-Hub relayer.

```bash
sudo docker exec -it hermes_relayer_it_hub bash
```
```bash
./run-relayer.sh
```

```bash
sudo docker exec -it hermes_relayer_de_hub bash
```
```bash
./run-relayer.sh
```


To start the go relayer process:

```bash
sudo docker exec go_relayer_it_hub ./run-relayer_it_hub.sh
```
```bash
sudo docker exec go_relayer_de_hub ./run-relayer_de_hub.sh
```


To simulate the execution:
_Remember_ that the execution of the tests must be started when the configuration of the relayers has been completed, otherwise there will be errors in communication between the chains.
Open a new terminal and run: 

```bash
./test_chain.sh it unipi tum
```

to run the tests in the italian chain, starting from the Unipi university and carrying out the Erasmus programme at TUM University
(there are also other possibilities, such as `./test_chain.sh de tum uniroma1` and so on).
Universities in the Italian chain: `unipi` and `uniroma1`.
Universities in the German chain: `tum` and `humboldt university`.

The tests that will be performed are the following (in the specific case of the `./test_chain.sh it unipi tum` command):
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

```bash
sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-unipi:/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i query universitychainit show-stored-student unipi_1  --chain-id university_chain_it --node "tcp://val-unipi:26657"
```

To see the updated data in the destination chain:

```bash
sudo docker run --rm -i -v $(pwd)/university_chain_de/elements/val-tum:/root/.university_chain_de --network university_chain-prod_net-public university_chain_ded_i query universitychainde show-stored-student tum_1  --chain-id university_chain_de --node "tcp://val-tum:26657"
```


Once the transfer of information between the chains is complete, the following test can be performed:

```bash
./university_chain_de/test/test_insert_grade_chain_de.sh tum
```

The test will check the insertion of the exam grade in the Erasmus plan.
Once the Erasmus is finished (I have set the timer temporarily to 120 seconds for the tests), the updated data on the Erasmus plan will be transferred back to the home university, where the grade conversion will be carried out.

To see the updated data in the home chain:

```bash
sudo docker run --rm -i -v $(pwd)/university_chain_it/elements/val-unipi:/root/.university_chain_it --network university_chain-prod_net-public university_chain_itd_i query universitychainit show-stored-student unipi_1  --chain-id university_chain_it --node "tcp://val-unipi:26657"
```

Besides waiting for the normal Erasmus conclusion, it is also possible to run the Erasmus period extension test and the Erasmus early termination test. Respectively:

```bash
./university_chain_it/test/test_extend_erasmus_request_chain_it.sh unipi
```
```bash
./university_chain_it/test/test_end_erasmus_request_chain_it.sh unipi
```


To end the execution of the containers with hermes relayers:

```bash
sudo docker compose --file docker-compose.yml --project-name university_chain-prod --profile hermes down
```

To end the execution of the containers with go relayers:

```bash
sudo docker compose --file docker-compose.yml --project-name university_chain-prod --profile go down
```

To calculate statistics, first run a test, then e.g. `./test_chain.sh it unipi tum`, and, from the main directory, run `./statistics/getStatistics.go normal`. Or if you ran the test from the German chain to the Italian chain, then for example `./test_chain.sh de tum unipi`, run `go run ./statistics/getStatistics.go reverse`. You will find the test results inside the statistics directory.


## Contributing


## License

[MIT](https://choosealicense.com/licenses/mit/)
