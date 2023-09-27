package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type latencyStruct struct {
	Time       time.Duration
	PacketSize string
}

type timingStruct struct {
	Details    string
	PacketSize string
	Time       time.Time
	Chain      string
}

type consumedGasStruct struct {
	ConsumedGas  string `json:"consumedGas"`
	FunctionName string `json:"functionName"`
	Id           string `json:"id"`
}

type timeStatsStruct struct {
	Details    string `json:"details"`
	PacketHash string `json:"packetHash"`
	PacketSize string `json:"packetSize"`
	Time       string `json:"time"`
}

type timeStatsStructTimeObject struct {
	Details    string
	PacketHash string
	PacketSize string
	Time       time.Time
}

var files_gas = []string{"university_chain_it/log/val-unipi/statsGasConsumed.txt", "university_chain_de/log/val-tum/statsGasConsumed.txt", "hub/log/val-hub-instance-1/statsGasConsumed.txt"}
var files_gas_res = []string{"chainItGasConsumed.txt", "chainDeGasConsumed.txt", "chainHubGasConsumed.txt"}

var filesTimingItToHub = []string{"university_chain_it/log/val-uniroma1/statsTiming.txt", "hub/log/val-hub-instance-2/statsTiming.txt"}
var filesTimingHubToDe = []string{"hub/log/val-hub-instance-2/statsTiming.txt", "university_chain_de/log/val-humboldt/statsTiming.txt"}

var statsItToHubPerPacket = "timingsITtoHubPerPacket.txt"
var statsHubToDEPerPacket = "timingsHubtoDEPerPacket.txt"
var globalFileItToHub = "globalFileTimingsITtoHub.txt"
var globalFileHubToDE = "globalFileTimingsHubtoDE.txt"

var DateLayoutMilli = "2006-01-02 15:04:05.000"

func getGasStats() (err error) {

	i := 0

	for _, fileName := range files_gas {

		file_chain_json, err := os.OpenFile(fileName, os.O_RDONLY, 0444)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error "+err.Error())
			return err
		}

		defer file_chain_json.Close()

		gasConsumptionMap := make(map[string]int64)

		scanner := bufio.NewScanner(file_chain_json)

		for scanner.Scan() {
			var consumedGasValues consumedGasStruct
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

func getGlobalTimings(fileNames []string) (globalTimings []timeStatsStructTimeObject, err error) {

	for _, fileName := range fileNames {

		file_chain_json, err := os.OpenFile(fileName, os.O_RDONLY, 0444)

		if err != nil {
			fmt.Fprintln(os.Stderr, "Error "+err.Error())
			return globalTimings, err
		}

		defer file_chain_json.Close()

		scanner := bufio.NewScanner(file_chain_json)

		for scanner.Scan() {
			var timingSpecs timeStatsStruct
			err = json.Unmarshal(scanner.Bytes(), &timingSpecs)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error "+err.Error())
				return globalTimings, err
			} else {
				date, err := time.Parse(DateLayoutMilli, timingSpecs.Time)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error "+err.Error())
					return globalTimings, err
				}
				newElem := timeStatsStructTimeObject{timingSpecs.Details, timingSpecs.PacketHash, timingSpecs.PacketSize, date}
				globalTimings = append(globalTimings, newElem)
			}
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error "+err.Error())
			return globalTimings, err
		}
	}

	sort.Slice(globalTimings, func(i, j int) bool {
		return globalTimings[i].Time.Before(globalTimings[j].Time)
	})

	return globalTimings, nil
}

func printGlobalStatsFile(globalTimings []timeStatsStructTimeObject, fileName string) (err error) {
	file_chain_middle, err := os.Create("./statistics/" + fileName)
	if err != nil {
		return err
	}

	defer file_chain_middle.Close()
	w := bufio.NewWriter(file_chain_middle)

	for _, element := range globalTimings {
		_, err = w.WriteString("PacketHash: " + element.PacketHash + "Details: " + element.Details + " PacketSize: " + element.PacketSize + " Time: " + element.Time.String() + "\n")
		if err != nil {
			return err
		}
	}
	w.Flush()
	return nil
}

func printStatsPerPacket(timingMap map[string][]timingStruct, fileName string) (err error) {

	file_chain_final, err := os.Create("./statistics/" + fileName)
	if err != nil {
		return err
	}

	defer file_chain_final.Close()
	w := bufio.NewWriter(file_chain_final)

	for key, element := range timingMap {
		_, err = w.WriteString("Packet hash: " + key + "\n")
		if err != nil {
			return err
		}
		for i := 0; i < len(element); i++ {

			_, err = w.WriteString("\tDetails: " + element[i].Details + " - PacketSize: " + element[i].PacketSize + " - Time: " + element[i].Time.String() + " - Chain: " + element[i].Chain + "\n")
			if err != nil {
				return err
			}
		}
		_, err = w.WriteString("\n")
		if err != nil {
			return err
		}

	}

	w.Flush()

	return nil
}

func getTimingMap(globalTimings []timeStatsStructTimeObject) (timingMap map[string][]timingStruct) {

	timingMap = make(map[string][]timingStruct)

	for _, timing := range globalTimings {

		val, ok := timingMap[timing.PacketHash]
		if !ok {
			var timings []timingStruct
			stringValues := strings.Split(timing.Details, " ")
			chain := stringValues[len(stringValues)-1]
			elem := timingStruct{timing.Details, timing.PacketSize, timing.Time, chain}
			timings = append(timings, elem)
			timingMap[timing.PacketHash] = timings

		} else {
			stringValues := strings.Split(timing.Details, " ")
			chain := stringValues[len(stringValues)-1]
			elem := timingStruct{timing.Details, timing.PacketSize, timing.Time, chain}
			newVal := append(val, elem)
			timingMap[timing.PacketHash] = newVal
		}
	}

	return timingMap
}

func getLatency(timingMap map[string][]timingStruct, mode string) (ltStruct []latencyStruct) {

	srcChain := ""
	dstChain := ""
	subStringSrc := ""
	subStringDst := ""

	if mode == "ITtoHub" {
		srcChain = "IT"
		dstChain = "Hub"
		subStringSrc = "Transmit"
		subStringDst = "OnRecv"

	} else if mode == "HubtoDE" {
		srcChain = "Hub"
		dstChain = "DE"
		subStringSrc = "Transmit"
		subStringDst = "OnRecv"

	} else if mode == "HubtoITAck" {
		srcChain = "Hub"
		dstChain = "IT"
		subStringSrc = "sending ack"
		subStringDst = "OnAcknowledgement"
	} else if mode == "DEtoHubAck" {
		srcChain = "DE"
		dstChain = "Hub"
		subStringSrc = "sending ack"
		subStringDst = "OnAcknowledgement"

		// reverse
	} else if mode == "DEtoHub" {
		srcChain = "DE"
		dstChain = "Hub"
		subStringSrc = "Transmit"
		subStringDst = "OnRecv"
	} else if mode == "HubtoIT" {
		srcChain = "Hub"
		dstChain = "IT"
		subStringSrc = "Transmit"
		subStringDst = "OnRecv"
	} else if mode == "HubtoDEAck" {
		srcChain = "Hub"
		dstChain = "DE"
		subStringSrc = "sending ack"
		subStringDst = "OnAcknowledgement"
	} else if mode == "ITtoHubAck" {
		srcChain = "IT"
		dstChain = "Hub"
		subStringSrc = "sending ack"
		subStringDst = "OnAcknowledgement"
	}

	i := 0
	j := i + 1
	found := false
	for _, element := range timingMap {
		found = false
		i = 0
		j = i + 1
		for j < len(element) && !found {
			found = false
			if element[i].Chain == srcChain && strings.Contains(element[i].Details, subStringSrc) {
				for j < len(element) && !found {
					if element[j].Chain == dstChain && element[i].PacketSize == element[j].PacketSize && strings.Contains(element[j].Details, subStringDst) {
						lt_element := latencyStruct{element[j].Time.Sub(element[i].Time), element[i].PacketSize}
						ltStruct = append(ltStruct, lt_element)
						found = true
					} else {
						j++
					}
				}
			} else {
				i++
				j++
			}

		}
	}
	return ltStruct
}

func printLatency(fileName string, ltStruct []latencyStruct) (err error) {

	//file_chain_final, err := os.OpenFile("./statistics/"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	file_chain_final, err := os.Create("./statistics/" + fileName)
	if err != nil {
		return err
	}

	defer file_chain_final.Close()
	w := bufio.NewWriter(file_chain_final)

	_, err = w.WriteString(time.Now().String() + "\n")
	if err != nil {
		return err
	}

	for _, element := range ltStruct {
		_, err = w.WriteString("Packet size: " + element.PacketSize + " - latency: " + element.Time.String() + "\n")
		if err != nil {
			return err
		}
	}

	w.Flush()

	return nil
}

func getTimeStats(mode string) (err error) {

	globalTimingsItToHub, err := getGlobalTimings(filesTimingItToHub)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return err
	}
	globalTimingsHubtoDe, err := getGlobalTimings(filesTimingHubToDe)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return err
	}

	//----------------------------------------

	printGlobalStatsFile(globalTimingsItToHub, globalFileItToHub)
	printGlobalStatsFile(globalTimingsHubtoDe, globalFileHubToDE)

	//----------------------------------------

	timingMapItToHub := getTimingMap(globalTimingsItToHub)
	timingMapHubToDe := getTimingMap(globalTimingsHubtoDe)

	//------------------------------------

	printStatsPerPacket(timingMapItToHub, statsItToHubPerPacket)
	printStatsPerPacket(timingMapHubToDe, statsHubToDEPerPacket)

	if mode == "normal" {

		ltITtoHub := getLatency(timingMapItToHub, "ITtoHub")
		ltHubtoDE := getLatency(timingMapHubToDe, "HubtoDE")
		ltHubtoITAckPackets := getLatency(timingMapItToHub, "HubtoITAck")
		ltDEtoHubAckPackets := getLatency(timingMapHubToDe, "DEtoHubAck")

		printLatency("LatencyITtoHub", ltITtoHub)
		printLatency("LatencyHubtoDE", ltHubtoDE)
		printLatency("LatencyHubtoIT", ltHubtoITAckPackets)
		printLatency("LatencyDEtoHub", ltDEtoHubAckPackets)

		err = getInterval(globalTimingsItToHub)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error "+err.Error())
			return err
		}

	} else {

		// reverse

		ltIDEtoHubReverse := getLatency(timingMapHubToDe, "DEtoHub")
		ltHubtoITReverse := getLatency(timingMapItToHub, "HubtoIT")
		ltHubtoDEAckPackets := getLatency(timingMapHubToDe, "HubtoDEAck")
		ltITtoHubAckPackets := getLatency(timingMapItToHub, "ITtoHubAck")

		printLatency("LatencyDEtoHub", ltIDEtoHubReverse)
		printLatency("LatencyHubtoIT", ltHubtoITReverse)
		printLatency("LatencyHubtoDE", ltHubtoDEAckPackets)
		printLatency("LatencyITtoHub", ltITtoHubAckPackets)

		err = getInterval(globalTimingsHubtoDe)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error "+err.Error())
			return err
		}

	}

	return err

}

func getInterval(globalTimings []timeStatsStructTimeObject) (err error) {

	dstChain := "Hub"
	subStringDst := "OnAcknowledgementErasmusRestictedDataPacket"
	startDetail := globalTimings[0].Details
	startTime := globalTimings[0].Time
	j := 0

	for i := 0; i < len(globalTimings); i++ {

		stringValues := strings.Split(globalTimings[i].Details, " ")
		if stringValues[len(stringValues)-1] == dstChain && strings.Contains(globalTimings[i].Details, subStringDst) && globalTimings[i].PacketSize == "28" {
			j = i
		}
	}
	endDetail := globalTimings[j].Details
	endTime := globalTimings[j].Time
	interval := endTime.Sub(startTime)

	file_chain_final, err := os.Create("./statistics/interval.txt")
	if err != nil {
		return err
	}

	defer file_chain_final.Close()
	w := bufio.NewWriter(file_chain_final)

	_, err = w.WriteString(time.Now().String() + "\n")
	if err != nil {
		return err
	}

	_, err = w.WriteString("Initial function: " + startDetail + "\n")
	if err != nil {
		return err
	}
	_, err = w.WriteString("Start time: " + startTime.String() + "\n")
	if err != nil {
		return err
	}
	_, err = w.WriteString("End function: " + endDetail + "\n")
	if err != nil {
		return err
	}
	_, err = w.WriteString("End time: " + endTime.String() + "\n")
	if err != nil {
		return err
	}
	_, err = w.WriteString("Interval: " + interval.String() + "\n")
	if err != nil {
		return err
	}

	w.Flush()

	return nil

}

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Error, the number of arguments must be equal to 1")
		return
	}

	if os.Args[1] != "normal" && os.Args[1] != "reverse" {
		fmt.Fprintln(os.Stderr, "Error, can't find normal or reverse")
		return
	}

	err := getGasStats()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return
	}
	err = getTimeStats(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return
	}

}
