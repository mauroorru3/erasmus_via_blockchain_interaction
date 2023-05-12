package utilfunc

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// UniversityKeys.json

const foreignUniversityInfoJSON string = "UniversityKeys.json"

type UniversityKeys struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Address string `json:"address"`
}

type UniListKey struct {
	UniversityListKey []UniversityKeys `json:"universitiesList"`
}

func ReadForeignUniversityInfo() (universityInfo []UniversityKeys, err error) {

	// Open our jsonFile
	jsonFile, err := os.OpenFile("data/"+foreignUniversityInfoJSON, os.O_RDONLY, 0444)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return universityInfo, err
	}
	fmt.Println("Successfully Opened " + foreignUniversityInfoJSON)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return universityInfo, err
	}

	var uniK UniListKey

	err = json.Unmarshal([]byte(byteValue), &uniK)
	fmt.Println("Successfully Unmarshalled " + foreignUniversityInfoJSON)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return universityInfo, err
	}

	universityInfo = uniK.UniversityListKey

	return universityInfo, err

}
