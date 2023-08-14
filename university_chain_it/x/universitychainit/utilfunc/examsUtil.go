package utilfunc

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
