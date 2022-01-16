package dto

type AttendanceDTO struct {
	Id                uint   `json:"id" form:"id"`
	WeekSubDetachment string `json:"week_sub_detachment" form:"week_sub_detachment"`
	Notebook          bool   `json:"notebook" form:"notebook"`
	Bible             bool   `json:"bible" form:"bible"`
	Attendance        bool   `json:"attendance" form:"attendance"`
	Uniform           bool   `json:"uniform" form:"uniform"`
	Offering          bool   `json:"offering" form:"offering"`
	Active            bool   `json:"active" form:"active"`
	UserId            uint   `json:"user_id" form:"user_id"`
	SubDetachmentId   uint   `json:"sub_detachment_id" form:"sub_detachment_id"`
}
