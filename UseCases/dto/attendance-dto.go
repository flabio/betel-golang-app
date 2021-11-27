package dto

type AttendanceDTO struct {
	Id         uint `json:"id" form:"id" `
	Notebook   bool `form:"notebook" json:"notebook"`
	Bible      bool `form:"bible   " json:"bible"`
	Attendance bool `form:"attendan" json:"attendance"`
	Uniform    bool `form:"uniform " json:"uniform"`
	Offering   bool `form:"offering" json:"offering"`
	Active     bool `form:"active  " json:"active"`
	UserId     uint `form:"user_id  " json:"user_id"`
}
