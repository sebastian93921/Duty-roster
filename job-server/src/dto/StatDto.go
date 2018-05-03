package dto

type StatDto struct {
	Date   string 			`json:"date"`
	Time    string 			`json:"time"`
	ActivityName string `json:"activityName"`
	Location string 		`json:"location"`
	Organiser string 		`json:"organiser"`
	AnotherDept string 	`json:"anotherDept"`
	Note	string 				`json:"note"`
	NumberOfMember int 	`json:"numberOfMember"`
	NumberOfAttendance int `json:"numberOfAttendance"`
	ListOfMembers string `json:"listOfMembers"`
	Owed string 				`json:"owed"`
	DutyKey string 			`json:"dutyKey"`
	ReceivedDate string `json:"receviedDate"`
	IsFull	bool				`json:"isFull"`
}
