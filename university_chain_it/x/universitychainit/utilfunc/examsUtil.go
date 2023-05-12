package utilfunc

import (
	"encoding/json"
)

type ExamStruct struct {
	Exam_label      string `json:"exam_label"`
	Exam_date       string `json:"exam_date"`
	Credits         uint8  `json:"credits"`
	Marks           string `json:"marks"`
	Course_year     uint16 `json:"course_year"`
	Status          bool   `json:"status"`
	Attendance_year uint16 `json:"attendance_year"`
	Exam_type       string `json:"exam_type"`
}

func (c *ExamStruct) ToMAP() (res map[string]interface{}) {
	a, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(a, &res)
	return
}

func (c *ExamStruct) LoadFromMap(m map[string]interface{}) error {
	data, err := json.Marshal(m)
	if err == nil {
		err = json.Unmarshal(data, c)
	}
	return err
}

func (c *ExamStruct) ToJSON() string {
	a, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return string(a)
}

func (c *ExamStruct) LoadFromJSON(jsonStr string) error {
	err := json.Unmarshal([]byte(jsonStr), c)
	return err
}

func (c *ExamStruct) ToByte() ([]byte, error) {
	return json.Marshal(c)
}

func (c *ExamStruct) LoadFromByte(data []byte) error {
	return json.Unmarshal(data, c)
}
