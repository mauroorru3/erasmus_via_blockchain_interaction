package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type ConsumedGasStruct struct {
	ConsumedGas  string `json:"consumedGas"`
	FunctionName string `json:"functionName"`
}

type TimeStatsStruct struct {
	FunctionName string
	PacketSize   int
	Time         string
}

var files_gas = []string{"university_chain_it/log/val-unipi/statsGasConsumed.txt", "university_chain_de/log/val-tum/statsGasConsumed.txt", "hub/log/val-hub-instance-1/statsGasConsumed.txt"}
var files_gas_res = []string{"chainItGasConsumed.txt", "chainDeGasConsumed.txt", "chainHubGasConsumed.txt"}

const file_it_chain_timing string = "university_chain_it/log/val-unipi/statsTiming.txt"
const file_de_chain_timing string = "university_chain_de/log/val-tum/statsTiming.txt"
const file_hub_chain_timing string = "hub/log/val-hub-instance-1/statsTiming.txt"

func getGasStats() (err error) {

	i := 0

	for _, fileName := range files_gas {

		file_chain_json, err := os.OpenFile(fileName, os.O_RDONLY, 0444)
		// if we os.Open returns an error then handle it
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error "+err.Error())
			return err
		}

		// defer the closing of our jsonFile so that we can parse it later on
		defer file_chain_json.Close()

		gasConsumptionMap := make(map[string]int64)
		//timeStatsMap := make(map[string][]TimeStatsStruct)

		scanner := bufio.NewScanner(file_chain_json)
		// optionally, resize scanner's capacity for lines over 64K, see next example
		for scanner.Scan() {
			var consumedGasValues ConsumedGasStruct
			err = json.Unmarshal(scanner.Bytes(), &consumedGasValues)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error "+err.Error())
				return err
			} else {
				GasInt, err := strconv.ParseInt(consumedGasValues.ConsumedGas, 10, 64)
				if err != nil {
					return nil
				} else {
					gasConsumptionMap[consumedGasValues.FunctionName] = GasInt
				}
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error "+err.Error())
			return err
		}

		file_chain_final, err := os.Create("./statistics/" + files_gas_res[i])
		if err != nil {
			return err
		}

		defer file_chain_final.Close()
		w := bufio.NewWriter(file_chain_final)

		for key, element := range gasConsumptionMap {

			elementString := strconv.FormatInt(element, 10)
			_, err = w.WriteString("Function: " + key + " - Gas consumed: " + elementString + "\n")
			if err != nil {
				return err
			}
		}

		w.Flush()

		i++

	}

	return nil

}

func main() {

	err := getGasStats()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return
	} else {
	}

}
