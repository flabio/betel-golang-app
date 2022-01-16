package dto

type ParentDTO struct {
	Id                 uint   `json:"id" form:"id"`
	FullName           string `json:"full_name" form:"full_name"`
	Identification     string `json:"identification" form:"identification"`
	TypeIdentification string `json:"type_identification" form:"type_identification"`
	TypeParent         string `json:"type_parent" form:"type_parent"`
	Gender             string `json:"gender" form:"gender"`
	Direction          string `json:"direction" form:"direction"`
	PhoneNumber        string `json:"phone_number" form:"phone_number"`
}
