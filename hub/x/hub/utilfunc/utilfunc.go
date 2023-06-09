package utilfunc

import (
	"encoding/json"
	"fmt"
	"hub/x/hub/types"
	"io"
	"os"
	"time"
)

// Erasmus Career

type ErasmusContributionStruct struct {
	Amount          uint32 `json:"amount"`
	Income_bracket  uint32 `json:"income_bracket"`
	Payment_made    bool   `json:"payment_made"`
	Date_of_payment string `json:"date_of_payment"`
}

type ErasmusCareerStruct struct {
	Duration_in_months            uint8                     `json:"duration_in_months"`
	Start_date                    string                    `json:"start_date"`
	End_date                      string                    `json:"end_date"`
	Erasmus_type                  string                    `json:"erasmus_type"`
	Total_credits                 uint8                     `json:"total_credits"`
	Achieved_credits              uint8                     `json:"achieved_credits"`
	Total_exams                   uint8                     `json:"total_exams"`
	Exams_passed                  uint8                     `json:"exams_passed"`
	Foreign_university_name       string                    `json:"foreign_university_name"`
	Foreign_university_country    string                    `json:"foreign_university_country"`
	Foreign_university_student_id string                    `json:"foreign_university_student_id"`
	Status                        string                    `json:"status"`
	Contribution                  ErasmusContributionStruct `json:"contribution"`
	Exams_data                    string                    `json:"exams_data"`
}

// Taxes

type TaxesStruct struct {
	Year            uint16 `json:"year"`
	Amount          uint32 `json:"amount"`
	Income_bracket  uint32 `json:"income_bracket"`
	Payment_made    bool   `json:"payment_made"`
	Date_of_payment string `json:"date_of_payment"`
}

// UniversityKeys.json

const foreignUniversityInfoJSON string = "UniversityKeys.json"

type UniversityKeys struct {
	Name       string `json:"name"`
	Country    string `json:"country"`
	Address    string `json:"address"`
	Port       string `json:"port"`
	Channel_ID string `json:"channel_id"`
	ChainName  string `json:"chain_name"`
}

type UniListKey struct {
	UniversityListKey []UniversityKeys `json:"universitiesList"`
}

const (
	DateLayout = "2006-01-02 15:04:05"
)

func FormatDate(date time.Time) string {
	loc, _ := time.LoadLocation("Europe/Rome")
	newTime := date.In(loc)
	return newTime.Format(DateLayout)
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

func CheckCompleteInformation(student types.StoredStudent) (err error) {

	completeInfo := student.StudentData.CompleteInformation

	if completeInfo[0] == 0 || completeInfo[1] == 0 || completeInfo[2] == 0 {

		return types.ErrIncompleteStudentInformation
	} else {
		return err
	}
}

func CheckTaxPayment(student types.StoredStudent) (ok bool, err error) {

	taxesDataBytes := []byte(student.TaxesData.TaxesHistory)
	var taxesData []TaxesStruct
	err = json.Unmarshal(taxesDataBytes, &taxesData)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return ok, err
	}

	return taxesData[0].Payment_made, err
}

func CheckErasmusStatus(student types.StoredStudent) (res string, err error) {

	var erasmusCareer []ErasmusCareerStruct

	err = json.Unmarshal([]byte(student.ErasmusData.Career), &erasmusCareer)
	if err != nil {
		return res, err
	}

	lenCareer := len(erasmusCareer)

	res = erasmusCareer[lenCareer-1].Status

	return res, err

}

func GetForeignUniversityName(student types.StoredStudent) (res string, err error) {

	var erasmusCareer []ErasmusCareerStruct

	err = json.Unmarshal([]byte(student.ErasmusData.Career), &erasmusCareer)
	if err != nil {
		return res, err
	}

	lenCareer := len(erasmusCareer)

	res = erasmusCareer[lenCareer-1].Foreign_university_name

	return res, err
}

func PrintLogs(text string) error {

	file, err := os.OpenFile("data/logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	dt := time.Now()

	_, err2 := file.WriteString(text + " " + FormatDate(dt) + "\n")

	if err2 != nil {
		return err2
	}

	return nil
}

func PrintData(text string) error {

	file, err := os.OpenFile("data/data.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}

	defer file.Close()

	dt := time.Now()

	_, err2 := file.WriteString(text + " " + FormatDate(dt) + "\n")

	if err2 != nil {
		return err2
	}

	return nil
}
