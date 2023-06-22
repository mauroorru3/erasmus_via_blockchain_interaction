package utilfunc

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
	"university_chain_it/x/universitychainit/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// grade conversion structure

const erasmusGradeConversionJSON string = "gradesConversion.json"

type GradesInfo struct {
	CountryName string   `json:"country_name"`
	Grades      []string `json:"grades"`
}

type CountryGradesList struct {
	Grades_data []GradesInfo `json:"countryList"`
}

//------------------------------------

var erasmusTypeMap = map[string]int{
	"study":       1,
	"traineeship": 1,
}

const (
	DeadlineLayout = "2006-01-02 15:04:05"
)

//-----------------------------------------
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
	Foreign_chain_name            string                    `json:"foreign_chain_name "`
	Foreign_university_student_id string                    `json:"foreign_university_student_id"`
	Status                        string                    `json:"status"`
	Contribution                  ErasmusContributionStruct `json:"contribution"`
	Exams_data                    string                    `json:"exams_data"`
}

//-----------------------------------------
// Erasmus Config

type ForeignUniversities struct {
	Name            string `json:"name"`
	AvailablePlaces int16  `json:"availablePlaces"`
}

type ErasmusTypeStruct struct {
	Name            string `json:"name"`
	Cfu_first_year  int16  `json:"cfu_first_year"`
	Cfu_other_years int16  `json:"cfu_other_years"`
	Cfu_left        int16  `json:"cfu_left"`
}

type CourseDetailsStruct struct {
	CourseOfStudy   string              `json:"courseOfStudy"`
	AvailablePlaces int16               `json:"availablePlaces"`
	ErasmusTypelist []ErasmusTypeStruct `json:"erasmusTypelist"`
}

type CourseTypeStruct struct {
	Name    string                `json:"name"`
	Courses []CourseDetailsStruct `json:"courses"`
}

type GroupStruct struct {
	Countries                []string `json:"countries"`
	Contribution_study       int16    `json:"contribution_study"`
	Contribution_traineeship int16    `json:"contribution_traineeship"`
}

type MURContributionStruct struct {
	OutOfCourse int8     `json:"outOfCourse"`
	First       []uint16 `json:"first"`
	Second      []uint16 `json:"second"`
	Third       []uint16 `json:"third"`
	Fourth      []uint16 `json:"fourth"`
	Fifth       []uint16 `json:"fifth"`
	Sixth       []uint16 `json:"sixth"`
	Seventh     uint16   `json:"seventh"`
}

type ContributionGroupsStruct struct {
	Group1 GroupStruct `json:"group1"`
	Group2 GroupStruct `json:"group2"`
	Group3 GroupStruct `json:"group3"`
}

type ErasmusConfigStruct struct {
	MinMonths                      int8                     `json:"minMonths"`
	MaxMonths                      int8                     `json:"maxMonths"`
	MinCFU                         int16                    `json:"minCFU"`
	ContributionGroups             ContributionGroupsStruct `json:"contributionGroups"`
	AdditionalContribution         int16                    `json:"additionalContribution"`
	AdditionalContributionDuration int8                     `json:"additionalContributionDuration"`
	MURContribution                MURContributionStruct    `json:"MURContribution"`
	CourseTypeList                 []CourseTypeStruct       `json:"courseTypeList"`
	UniversitiesPlacesList         []ForeignUniversities    `json:"universitiesPlacesList"`
}

const erasmusConfigJSON string = "erasmusConfig.json"

func IntializeErasmusStruct(incomeBracket uint32) (erasmusJSON string, err error) {

	erasmusInfo := ErasmusCareerStruct{
		Duration_in_months:            0,
		Start_date:                    "",
		End_date:                      "",
		Erasmus_type:                  "",
		Total_credits:                 0,
		Achieved_credits:              0,
		Total_exams:                   0,
		Exams_passed:                  0,
		Foreign_university_name:       "",
		Foreign_university_country:    "",
		Foreign_university_student_id: "",
		Foreign_chain_name:            "",
		Status:                        "",
		Contribution: ErasmusContributionStruct{
			Amount:          0,
			Income_bracket:  incomeBracket,
			Payment_made:    false,
			Date_of_payment: "",
		},
		Exams_data: "",
	}

	var erasmusCareer []ErasmusCareerStruct
	erasmusCareer = append(erasmusCareer, erasmusInfo)

	resultByteJSON, err := json.Marshal(erasmusCareer)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return erasmusJSON, err
	}

	erasmusJSON = string(resultByteJSON)

	return erasmusJSON, err

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GetChainFromUniversity(university string) (chain string, err error) {

	return chain, nil
}

func CheckErasmusParams(durationInMonths string, erasmusType string, student *types.StoredStudent, foreign_university_name string, foreign_university_country string, foreign_chain_name string) (err error) {

	// Open our jsonFile
	jsonFile, err := os.OpenFile("data/"+erasmusConfigJSON, os.O_RDONLY, 0444)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return err
	}
	fmt.Println("Successfully Opened " + erasmusConfigJSON)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return err
	}

	var erasmusConfig ErasmusConfigStruct

	err = json.Unmarshal([]byte(byteValue), &erasmusConfig)
	fmt.Println("Successfully Unmarshalled " + foreignUniversityInfoJSON)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return err
	}

	erasmusDuration, err := strconv.ParseInt(durationInMonths, 10, 8)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return err
	}

	if int8(erasmusDuration) < erasmusConfig.MinMonths || int8(erasmusDuration) > erasmusConfig.MaxMonths {
		return types.ErrWrongErasmusDuration
	}

	if student.ErasmusData.NumberMonths > 0 {
		if student.ErasmusData.NumberMonths+uint32(erasmusDuration) > uint32(erasmusConfig.MaxMonths) {
			return types.ErrLimitErasmusMonthsExceeded
		}
	}

	_, found := erasmusTypeMap[erasmusType]
	if !found {
		return types.ErrInvalidErasmusType
	}

	if student.TranscriptData.AchievedCredits < uint32(erasmusConfig.MinCFU) {
		return types.ErrTooFewCFUToAttend
	}

	found = false
	i := 0
	for !found && i < len(erasmusConfig.CourseTypeList) {
		if erasmusConfig.CourseTypeList[i].Name == student.StudentData.CourseType {
			found = true
		} else {
			i++
		}
	}
	if !found {
		return types.ErrErasmusConfFileCourseType
	}

	found = false
	j := 0
	for !found && j < len(erasmusConfig.CourseTypeList[i].Courses) {
		if erasmusConfig.CourseTypeList[i].Courses[j].CourseOfStudy == student.StudentData.CourseOfStudy {
			found = true
		} else {
			j++
		}
	}

	if !found {
		return types.ErrErasmusConfFileCourseOfStudy
	}

	found = false
	k := 0
	for !found && k < len(erasmusConfig.CourseTypeList[i].Courses[j].ErasmusTypelist) {
		if erasmusConfig.CourseTypeList[i].Courses[j].ErasmusTypelist[k].Name == erasmusType {
			found = true
		} else {
			k++
		}
	}

	if !found {
		return types.ErrErasmusConfFileErasmusType
	}

	if erasmusType == "study" {
		if student.StudentData.CurrentYearOfStudy == 1 {
			if int16(student.TranscriptData.AchievedCredits) < erasmusConfig.CourseTypeList[i].Courses[j].ErasmusTypelist[k].Cfu_first_year {
				return types.ErrTooFewCFUToAttend
			}
		} else {
			if int16(student.TranscriptData.AchievedCredits) < erasmusConfig.CourseTypeList[i].Courses[j].ErasmusTypelist[k].Cfu_other_years {
				return types.ErrTooFewCFUToAttend
			}
		}
	} else {
		if int16(student.TranscriptData.TotalCredits-student.TranscriptData.AchievedCredits) > erasmusConfig.CourseTypeList[i].Courses[j].ErasmusTypelist[k].Cfu_left {
			return types.ErrTooFewCFUToAttend
		}
	}

	student.ErasmusData.NumberMonths += uint32(erasmusDuration)

	student.ErasmusData.NumberTimes++

	var erasmusCareer []ErasmusCareerStruct

	err = json.Unmarshal([]byte(student.ErasmusData.Career), &erasmusCareer)
	if err != nil {
		return err
	}

	var contritbutionErasmus uint32

	if contains(erasmusConfig.ContributionGroups.Group1.Countries, foreign_university_country) {
		if erasmusType == "study" {
			contritbutionErasmus = uint32(erasmusConfig.ContributionGroups.Group1.Contribution_study)
		} else {
			contritbutionErasmus = uint32(erasmusConfig.ContributionGroups.Group1.Contribution_traineeship)
		}

	} else if contains(erasmusConfig.ContributionGroups.Group2.Countries, foreign_university_country) {

		if erasmusType == "study" {
			contritbutionErasmus = uint32(erasmusConfig.ContributionGroups.Group2.Contribution_study)
		} else {
			contritbutionErasmus = uint32(erasmusConfig.ContributionGroups.Group2.Contribution_traineeship)
		}

	} else {
		if erasmusType == "study" {
			contritbutionErasmus = uint32(erasmusConfig.ContributionGroups.Group3.Contribution_study)
		} else {
			contritbutionErasmus = uint32(erasmusConfig.ContributionGroups.Group3.Contribution_traineeship)
		}
	}

	lenCareer := len(erasmusCareer)

	contritbutionErasmus += uint32(erasmusConfig.AdditionalContribution) * uint32(erasmusConfig.AdditionalContributionDuration)
	if student.StudentData.NumberOfYearsOutOfCourse <= uint32(erasmusConfig.MURContribution.OutOfCourse) {
		if erasmusCareer[lenCareer-1].Contribution.Income_bracket < uint32(erasmusConfig.MURContribution.First[0]) {
			contritbutionErasmus += uint32(erasmusConfig.MURContribution.First[1])
		} else if erasmusCareer[lenCareer-1].Contribution.Income_bracket < uint32(erasmusConfig.MURContribution.Second[0]) {
			contritbutionErasmus += uint32(erasmusConfig.MURContribution.Second[1])
		} else if erasmusCareer[lenCareer-1].Contribution.Income_bracket < uint32(erasmusConfig.MURContribution.Third[0]) {
			contritbutionErasmus += uint32(erasmusConfig.MURContribution.Third[1])
		} else if erasmusCareer[lenCareer-1].Contribution.Income_bracket < uint32(erasmusConfig.MURContribution.Fourth[0]) {
			contritbutionErasmus += uint32(erasmusConfig.MURContribution.Fourth[1])
		} else if erasmusCareer[lenCareer-1].Contribution.Income_bracket < uint32(erasmusConfig.MURContribution.Fifth[0]) {
			contritbutionErasmus += uint32(erasmusConfig.MURContribution.Fifth[1])
		} else if erasmusCareer[lenCareer-1].Contribution.Income_bracket < uint32(erasmusConfig.MURContribution.Sixth[0]) {
			contritbutionErasmus += uint32(erasmusConfig.MURContribution.Sixth[1])
		}
	}

	if erasmusCareer[lenCareer-1].Duration_in_months == 0 {

		erasmusCareer[lenCareer-1].Duration_in_months = uint8(erasmusDuration)
		erasmusCareer[lenCareer-1].Erasmus_type = erasmusType
		erasmusCareer[lenCareer-1].Foreign_university_name = foreign_university_name
		erasmusCareer[lenCareer-1].Foreign_university_country = foreign_university_country
		erasmusCareer[lenCareer-1].Foreign_chain_name = foreign_chain_name
		erasmusCareer[lenCareer-1].Status = "To start" // "", To start, In progress, Waiting for updated data from the destination university or Terminated
		erasmusCareer[lenCareer-1].Contribution.Amount = contritbutionErasmus

	} else {

		erasmusInfo := ErasmusCareerStruct{
			Duration_in_months:            uint8(erasmusDuration),
			Start_date:                    "",
			End_date:                      "",
			Erasmus_type:                  erasmusType,
			Total_credits:                 0,
			Achieved_credits:              0,
			Total_exams:                   0,
			Exams_passed:                  0,
			Foreign_university_name:       foreign_university_name,
			Foreign_university_country:    foreign_university_country,
			Foreign_chain_name:            foreign_chain_name,
			Foreign_university_student_id: "",
			Status:                        "To start", // "", To start, In progress, Waiting for updated data from the destination university or Terminated
			Contribution: ErasmusContributionStruct{
				Amount:          contritbutionErasmus,
				Income_bracket:  erasmusCareer[lenCareer-1].Contribution.Income_bracket,
				Payment_made:    false,
				Date_of_payment: "",
			},
			Exams_data: "",
		}

		erasmusCareer = append(erasmusCareer, erasmusInfo)
	}

	// TO DO
	// verify also the available places in the selected foreign university

	resultByteJSON, err := json.Marshal(erasmusCareer)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return err
	}

	erasmusJSON := string(resultByteJSON)

	student.ErasmusData.Career = erasmusJSON

	return err
}

func CheckErasmusStatus(student types.StoredStudent, typeCheck string) (err error) {

	var erasmusCareer []ErasmusCareerStruct

	err = json.Unmarshal([]byte(student.ErasmusData.Career), &erasmusCareer)
	if err != nil {
		return err
	}

	lenCareer := len(erasmusCareer)

	res := erasmusCareer[lenCareer-1].Status

	switch typeCheck {
	case "start erasmus", "insert erasmus exam":

		if student.ErasmusData.ErasmusStudent == "Incoming completed" {
			return types.ErrIncomingPeriod
		} else if student.ErasmusData.ErasmusStudent == "Incoming" {
			return types.ErrIncomingPeriod
		}
		if res == "" {
			return types.ErrNoErasmusRequest
		} else if res == "In progress" {
			return types.ErrPreviousRequestInProgress
		} else if res == "Terminated" {
			return types.ErrPreviousRequestCompleted
		} else if res == "Waiting for updated data from the destination university" {
			return types.ErrPreviousRequestInProgress
		}
	case "erasmus request":
		if student.ErasmusData.ErasmusStudent == "Incoming completed" {
			return types.ErrIncomingPeriod
		} else if student.ErasmusData.ErasmusStudent == "Incoming" {
			return types.ErrIncomingPeriod
		}
		if res == "In progress" {
			return types.ErrPreviousRequestInProgress
		} else if res == "To start" {
			return types.ErrPreviousRequestCompleted
		} else if res == "Waiting for updated data from the destination university" {
			return types.ErrPreviousRequestInProgress
		}
	case "insert exam grade":

		if student.ErasmusData.ErasmusStudent == "Outgoing" {
			return types.ErrOutgoingPeriod
		} else if student.ErasmusData.ErasmusStudent == "Incoming completed" {
			return types.ErrCompletedIncomingPeriod
		}
	}

	return err

}

func InsertErasmusExam(student *types.StoredStudent, examName string, maxErasmusExams int32) (err error) {

	mapExams := make(map[string]ExamStruct)

	err = json.Unmarshal([]byte(student.TranscriptData.ExamsData), &mapExams)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return err
	}

	val, ok := mapExams[examName]
	if ok {
		if val.Marks != "" {
			return types.ErrExamAlreadyHeld
		} else {

			var erasmusCareer []ErasmusCareerStruct

			err = json.Unmarshal([]byte(student.ErasmusData.Career), &erasmusCareer)
			if err != nil {
				return err
			}

			lenCareer := len(erasmusCareer)

			var resultByteJSON []byte
			mapExamsErasmus := make(map[string]ExamStruct)

			if erasmusCareer[lenCareer-1].Exams_data == "" {

				mapExamsErasmus[examName] = val
				resultByteJSON, err = json.Marshal(mapExamsErasmus)
				if err != nil {
					return err
				}

				examsErasmusJSON := string(resultByteJSON)

				erasmusCareer[lenCareer-1].Exams_data = examsErasmusJSON

			} else {

				err = json.Unmarshal([]byte(erasmusCareer[lenCareer-1].Exams_data), &mapExamsErasmus)
				if err != nil {
					fmt.Fprintln(os.Stderr, "Error "+err.Error())
					return err
				}

				numExamsPlan := len(mapExamsErasmus)
				if numExamsPlan == int(maxErasmusExams) {
					return types.ErrMaxNumberErasmusExamReached
				}

				_, found := mapExamsErasmus[examName]
				if !found {
					mapExamsErasmus[examName] = val
					resultByteJSON, err = json.Marshal(mapExamsErasmus)
					if err != nil {
						return err
					}

					examsErasmusJSON := string(resultByteJSON)

					erasmusCareer[lenCareer-1].Exams_data = examsErasmusJSON
				} else {
					return types.ErrErasmusExamAlreadyInserted
				}
			}

			erasmusCareer[lenCareer-1].Total_exams++
			student.ErasmusData.TotalExams++
			erasmusCareer[lenCareer-1].Total_credits += val.Credits
			student.ErasmusData.TotalCredits += uint32(val.Credits)

			resultByteJSON, err = json.Marshal(erasmusCareer)
			if err != nil {
				return err
			}

			student.ErasmusData.Career = string(resultByteJSON)

		}
	} else {
		return types.ErrErasmusExamNotFound
	}

	return err

}

func CheckErasmusDeadline(ctx sdk.Context, erasmusDeadline string) error {

	deadline, err := time.Parse(DeadlineLayout, erasmusDeadline)
	if err != nil {
		return err
	}

	if deadline.Before(ctx.BlockTime()) {
		return types.ErrErasmusDeadlineExceeded
	} else {
		return nil
	}
}

func StartErasmus(ctx sdk.Context, student *types.StoredStudent, uniInfo *types.UniversityInfo) (err error) {

	var erasmusCareer []ErasmusCareerStruct

	err = json.Unmarshal([]byte(student.ErasmusData.Career), &erasmusCareer)
	if err != nil {
		return err
	}

	lenCareer := len(erasmusCareer)

	/*
		endDate := startDate.AddDate(0, int(erasmusCareer[lenCareer-1].Duration_in_months), 0)
		erasmusCareer[lenCareer-1].Start_date = startDate.Format("2006-01-02")
		erasmusCareer[lenCareer-1].End_date = endDate.Format("2006-01-02")
		erasmusCareer[lenCareer-1].Status = "In progress"
		student.ErasmusData.ErasmusStudent = "Outgoing"
	*/

	// I will enter just 30 seconds to see if the end of the Erasmus period works.

	startDate := ctx.BlockTime()

	endDate := ctx.BlockTime().Add(time.Duration(240 * time.Second))
	//endDate := ctx.BlockTime().Add(time.Duration(-1))
	erasmusCareer[lenCareer-1].Start_date = FormatDeadline(startDate)
	erasmusCareer[lenCareer-1].End_date = FormatDeadline(endDate)
	erasmusCareer[lenCareer-1].Status = "In progress"
	student.ErasmusData.ErasmusStudent = "Outgoing"

	resultByteJSON, err := json.Marshal(erasmusCareer)
	if err != nil {
		return err
	}

	student.ErasmusData.Career = string(resultByteJSON)

	return err

}

func FormatDeadline(deadline time.Time) string {
	loc, _ := time.LoadLocation("Europe/Rome")
	newTime := deadline.In(loc)
	return newTime.Format(DeadlineLayout)
	//return deadline.Format(DeadlineLayout)
}

func GetErasmusDeadline(student types.StoredStudent) (date time.Time, err error) {

	var erasmusCareer []ErasmusCareerStruct

	err = json.Unmarshal([]byte(student.ErasmusData.Career), &erasmusCareer)
	if err != nil {
		return date, err
	}

	lenCareer := len(erasmusCareer)

	finishDate, err := time.Parse(DeadlineLayout, erasmusCareer[lenCareer-1].End_date)
	if err != nil {
		return date, err
	}

	return finishDate, err
}

func CloseErasmusPeriod(student *types.StoredStudent) (err error) {

	var erasmusCareer []ErasmusCareerStruct

	err = json.Unmarshal([]byte(student.ErasmusData.Career), &erasmusCareer)
	if err != nil {
		return err
	}

	lenCareer := len(erasmusCareer)

	//student.ErasmusData.ErasmusStudent = "Outgoing completed"

	student.ErasmusData.ErasmusStudent = "Waiting for updated data from the destination university"

	//erasmusCareer[lenCareer-1].Status = "Terminated"

	erasmusCareer[lenCareer-1].Status = "Waiting for updated data from the destination university"

	resultByteJSON, err := json.Marshal(erasmusCareer)
	if err != nil {
		return err
	}

	student.ErasmusData.Career = string(resultByteJSON)

	return err
}

func SetForeignIndex(student *types.StoredStudent, foreignIndex string) (err error) {

	var erasmusCareer []ErasmusCareerStruct

	err = json.Unmarshal([]byte(student.ErasmusData.Career), &erasmusCareer)
	if err != nil {
		return err
	}

	lenCareer := len(erasmusCareer)

	erasmusCareer[lenCareer-1].Foreign_university_student_id = foreignIndex

	resultByteJSON, err := json.Marshal(erasmusCareer)
	if err != nil {
		return err
	}

	student.ErasmusData.Career = string(resultByteJSON)

	return err
}

func GetForeignIndex(student types.StoredStudent) (foreignIndex string, err error) {

	var erasmusCareer []ErasmusCareerStruct

	err = json.Unmarshal([]byte(student.ErasmusData.Career), &erasmusCareer)
	if err != nil {
		return foreignIndex, err
	}

	lenCareer := len(erasmusCareer)

	foreignIndex = erasmusCareer[lenCareer-1].Foreign_university_student_id

	return foreignIndex, err

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

func ConcludeErasmusFlag(student *types.StoredStudent) (err error) {

	var erasmusCareer []ErasmusCareerStruct

	err = json.Unmarshal([]byte(student.ErasmusData.Career), &erasmusCareer)
	if err != nil {
		return err
	}

	lenCareer := len(erasmusCareer)

	erasmusCareer[lenCareer-1].Status = "Terminated"

	resultByteJSON, err := json.Marshal(erasmusCareer)
	if err != nil {
		return err
	}

	student.ErasmusData.Career = string(resultByteJSON)

	return err
}

func UpdateErasmusData(student *types.StoredStudent, erasmusInfo *types.ErasmusInfo) (err error) {

	PrintLogs("UpdateErasmusData")

	foreignIndex, err := GetForeignIndex(*student)
	if err != nil {
		return err
	}
	student.ErasmusData = erasmusInfo

	SetForeignIndex(student, foreignIndex)
	if err != nil {
		return err
	}

	var erasmusCareer []ErasmusCareerStruct

	err = json.Unmarshal([]byte(student.ErasmusData.Career), &erasmusCareer)
	if err != nil {
		return err
	}

	lenCareer := len(erasmusCareer)

	student.ErasmusData.ErasmusStudent = "Outgoing completed"

	erasmusCareer[lenCareer-1].Status = "Terminated"

	// get the grade conversions structure

	// grade conversion taken in https://www.economia.unifi.it/upload/sub/relazioni-internazionali/TABELLA%20PER%20PAESE%20UE%20(1).pdf

	// Open our jsonFile
	jsonFile, err := os.OpenFile("data/"+erasmusGradeConversionJSON, os.O_RDONLY, 0444)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return err
	}
	fmt.Println("Successfully Opened " + erasmusGradeConversionJSON)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return err
	}

	var erasmusGradesConversion CountryGradesList

	err = json.Unmarshal([]byte(byteValue), &erasmusGradesConversion)
	fmt.Println("Successfully Unmarshalled " + foreignUniversityInfoJSON)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return err
	}

	// get the map of exams done during the Erasmus period

	mapExamsErasmus := make(map[string]ExamStruct)

	err = json.Unmarshal([]byte(erasmusCareer[lenCareer-1].Exams_data), &mapExamsErasmus)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return err
	}

	keys := make([]string, 0, len(mapExamsErasmus))
	values := make([]ExamStruct, 0, len(mapExamsErasmus))

	for k, v := range mapExamsErasmus {
		keys = append(keys, k)
		values = append(values, v)
	}

	// get the map of global exams

	mapExams := make(map[string]ExamStruct)

	err = json.Unmarshal([]byte(student.TranscriptData.ExamsData), &mapExams)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return err
	}

	// convert the grades

	found := false
	i, j, k := 0, 0, 0

	for i = 0; i < len(erasmusGradesConversion.Grades_data) && !found; {
		if erasmusGradesConversion.Grades_data[i].CountryName == erasmusCareer[lenCareer-1].Foreign_university_country {
			found = true
		} else {
			i++
		}
	}

	if !found {
		return types.ErrWrongForeignUniversity
	} else {
		for k = 0; k < len(values); k++ {
			if values[k].Marks != "" {
				found = false
				for j = 0; j < len(erasmusGradesConversion.Grades_data[i].Grades) && !found; {
					if values[k].Marks == erasmusGradesConversion.Grades_data[i].Grades[j] {
						found = true
					} else {
						j += 2
					}
				}
				if found {
					val := mapExams[keys[k]]
					val.Marks = erasmusGradesConversion.Grades_data[i].Grades[j+1]
					val.Status = values[k].Status
					val.Exam_date = values[k].Exam_date
					val.Attendance_year = values[k].Attendance_year
					mapExams[keys[k]] = val
					student.TranscriptData.AchievedCredits += uint32(val.Credits)
					student.TranscriptData.ExamsPassed += 1
				}

			}
		}

		// convert all the structure to string

		// global exams

		resultByteGlobalExamsJSON, err := json.Marshal(mapExams)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error "+err.Error())
			return err
		}

		globalExamsJSON := string(resultByteGlobalExamsJSON)

		student.TranscriptData.ExamsData = globalExamsJSON

		//-----------------

		// Erasmus exams

		resultExamsByteJSON, err := json.Marshal(mapExamsErasmus)
		if err != nil {
			return err
		}

		examsErasmusJSON := string(resultExamsByteJSON)

		erasmusCareer[lenCareer-1].Exams_data = examsErasmusJSON

		//------------------------

		// Erasmus career

		resultByteCareerJSON, err := json.Marshal(erasmusCareer)
		if err != nil {
			return err
		}

		student.ErasmusData.Career = string(resultByteCareerJSON)

		return err

	}

}
