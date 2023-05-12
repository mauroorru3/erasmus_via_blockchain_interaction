package utilfunc

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
	"university_chain_it/x/universitychainit/types"
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

//--------------------------------------------------
// university_info.json

const universityInfoJSON string = "university_info.json"

type Exam struct {
	ExamName         string `json:"examName"`
	Credits          uint   `json:"credits"`
	ExamType         string `json:"examType"`
	ProfessorName    string `json:"Name and Surname"`
	ProfessorAddress string `json:"Address"`
}

type Course struct {
	CourseOfStudy string `json:"courseOfStudy"`
	Exams         []Exam `json:"exams"`
}

type CourseList struct {
	Name    string   `json:"name"`
	Courses []Course `json:"courses"`
}

type CAI_struct struct {
	Department string `json:"department"`
	Name       string `json:"Name and Surname"`
	Address    string `json:"address"`
}

type DepartmentList struct {
	Name            string       `json:"name"`
	CAI             CAI_struct   `json:"CAI"`
	CoursesTypeList []CourseList `json:"courseTypeList"`
}

type Taxes_struct struct {
	First  []uint32 `json:"first"`
	Second []uint32 `json:"second"`
	Third  []uint32 `json:"third"`
	Fourth uint32   `json:"fourth"`
}

type University struct {
	Name              string           `json:"name"`
	Country           string           `json:"country"`
	Secretariat_key   string           `json:"secretariat_key"`
	University_key    string           `json:"university_key"`
	Deadline_taxes    string           `json:"deadline_taxes"`
	Deadline_erasmus  string           `json:"deadline_erasmus"`
	Max_erasmus_exams int32            `json:"max_erasmus_exams"`
	DepartmentList    []DepartmentList `json:"departmentList"`
	Taxes_brackets    Taxes_struct     `json:"taxes_brackets"`
}

type UniList struct {
	UniversityList []University `json:"universitiesList"`
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

func GetJSONFromCourseExams(NameUniversity string, departmentName string, courseType string, courseName string) (examsJSON string, numExams int, err error) {

	// Open our jsonFile
	jsonFile, err := os.OpenFile("data/"+universityInfoJSON, os.O_RDONLY, 0444)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return examsJSON, numExams, err
	}
	fmt.Println("Successfully Opened " + universityInfoJSON)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return examsJSON, numExams, err
	}

	var uList UniList
	err = json.Unmarshal([]byte(byteValue), &uList)
	fmt.Println("Successfully Unmarshalled " + universityInfoJSON)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return examsJSON, numExams, err
	}

	i := 0
	found := false
	for i < len(uList.UniversityList) && !found {
		if uList.UniversityList[i].Name == NameUniversity {
			found = true
		} else {
			i++
		}
	}
	if !found {
		fmt.Printf("University %s not found", NameUniversity)
		return examsJSON, numExams, types.ErrWrongNameUniversity
	}

	j := 0
	found = false
	for j < len(uList.UniversityList[i].DepartmentList) && !found {
		if uList.UniversityList[i].DepartmentList[j].Name == departmentName {
			found = true
		} else {
			j++
		}
	}
	if !found {
		fmt.Fprintln(os.Stderr, "Department "+departmentName+" not found")
		return examsJSON, numExams, types.ErrWrongDepartment
	}
	type_course := -1
	switch courseType {
	case "master":
		type_course = 0
	case "bachelor":
		type_course = 1
	case "single cycle":
		type_course = 2
	}
	if type_course == -1 {
		fmt.Fprintln(os.Stderr, "Course type "+courseType+" not found")
		return examsJSON, numExams, types.ErrWrongCourseType
	}

	k := 0
	found = false
	for k < len(uList.UniversityList[i].DepartmentList[j].CoursesTypeList[type_course].Courses) && !found {
		if uList.UniversityList[i].DepartmentList[j].CoursesTypeList[type_course].Courses[k].CourseOfStudy == courseName {
			found = true
		} else {
			k++
		}
	}
	if !found {
		fmt.Fprintln(os.Stderr, "Course "+courseName+" not found")
		return examsJSON, numExams, types.ErrWrongCourseOfStudy
	}

	exams := uList.UniversityList[i].DepartmentList[j].CoursesTypeList[type_course].Courses[k].Exams

	mapExams := make(map[string]ExamStruct)

	for i = 0; i < len(exams); i++ {
		examStruct := ExamStruct{
			Exam_label:      "",
			Exam_date:       "",
			Credits:         uint8(exams[i].Credits),
			Marks:           "",
			Course_year:     uint16(time.Now().Year()),
			Status:          false,
			Attendance_year: 0,
			Exam_type:       exams[i].ExamType,
		}
		mapExams[exams[i].ExamName] = examStruct
	}

	numExams = len(mapExams)

	resultByteJSON, err := json.Marshal(mapExams)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return examsJSON, numExams, err
	}

	examsJSON = string(resultByteJSON)

	return examsJSON, numExams, err

}

func ReadUniversitiesInfo() (universitiesInfo []University, err error) {

	// Open our jsonFile
	jsonFile, err := os.OpenFile("data/"+universityInfoJSON, os.O_RDONLY, 0444)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return universitiesInfo, err
	}
	fmt.Println("Successfully Opened " + universityInfoJSON)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return universitiesInfo, err
	}

	var uList UniList

	err = json.Unmarshal([]byte(byteValue), &uList)
	fmt.Println("Successfully Unmarshalled " + universityInfoJSON)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return universitiesInfo, err
	}

	universitiesInfo = uList.UniversityList

	return universitiesInfo, err

}

func IntializeTaxesStruct(incomeBracket uint32, taxesBrackets string) (taxesJSON string, err error) {

	taxesBracketsBytes := []byte(taxesBrackets)
	var taxesBracketsOriginal Taxes_struct
	err = json.Unmarshal(taxesBracketsBytes, &taxesBracketsOriginal)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return taxesJSON, err
	}

	var amount uint32

	if incomeBracket <= taxesBracketsOriginal.First[0] {
		amount = taxesBracketsOriginal.First[1]
	} else if incomeBracket < taxesBracketsOriginal.Second[0] {
		amount = taxesBracketsOriginal.Second[1]
	} else if incomeBracket < taxesBracketsOriginal.Third[0] {
		amount = taxesBracketsOriginal.Third[1]
	} else {
		amount = taxesBracketsOriginal.Fourth
	}

	year := time.Now().Year()
	baseYearTaxes := TaxesStruct{
		Year:            uint16(year),
		Amount:          amount,
		Income_bracket:  incomeBracket,
		Payment_made:    false,
		Date_of_payment: "",
	}
	var annualTaxes []TaxesStruct
	annualTaxes = append(annualTaxes, baseYearTaxes)

	resultByteJSON, err := json.Marshal(annualTaxes)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return taxesJSON, err
	}

	taxesJSON = string(resultByteJSON)

	return taxesJSON, err

}

func GetJSONFromTaxesBrackets(taxesBrackets Taxes_struct) (taxesJSON string, err error) {

	resultByteJSON, err := json.Marshal(taxesBrackets)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return taxesJSON, err
	}

	taxesJSON = string(resultByteJSON)

	return taxesJSON, err
}

func SetExamGrade(examsString string, examName string, grade string) (examsJSON string, credits uint8, err error) {

	mapExams := make(map[string]ExamStruct)

	err = json.Unmarshal([]byte(examsString), &mapExams)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return examsJSON, credits, err
	}

	val, ok := mapExams[examName]
	if ok {
		if val.Marks == "" {
			credits = val.Credits
			val.Attendance_year = uint16(time.Now().Year())
			val.Exam_date = time.Now().Format("01-02-2006")
			val.Status = true
			val.Marks = grade
			mapExams[examName] = val
		} else {
			return examsJSON, credits, types.ErrGradeAlreadyAssigned
		}
	} else {
		return examsJSON, credits, types.ErrWrongExamName
	}

	resultByteJSON, err := json.Marshal(mapExams)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error "+err.Error())
		return examsJSON, credits, err
	}

	examsJSON = string(resultByteJSON)

	return examsJSON, credits, err
}

func CheckCompleteInformation(student types.StoredStudent) (err error) {

	completeInfo := student.StudentData.CompleteInformation

	if completeInfo[0] == 0 || completeInfo[1] == 0 || completeInfo[2] == 0 {

		return types.ErrIncompleteStudentInformation
	} else {
		return err
	}
}
