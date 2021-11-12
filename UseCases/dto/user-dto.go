package dto

type UserDTO struct {
	Id                 uint   `json:"id" form:"id"`
	Image              string `json:"imagen" form:"imagen" `
	Name               string `json:"name" form:"name" `
	Email              string `json:"email" form:"email" `
	Password           string `json:"password" form:"password"`
	ConfirmPassword    string `json:"confirm_password" form:"confirm_password" `
	LastName           string `json:"last_name" form:"last_name" `
	Identification     string `json:"identification" form:"identification"`
	TypeIdentification string `json:"type_identification" form:"type_identification" `
	Birthday           string `json:"birthday" form:"birthday"`
	Birthplace         string `json:"birthplace" form:"birthplace" `
	Gender             string `json:"gender" form:"gender" `
	Direction          string `json:"direction" form:"direction" `
	PhoneNumber        string `json:"phone_number" `
	CellPhone          string `json:"cell_phone" form:"cell_phone" `
	CivilStatus        string `json:"civil_status" form:"civil_status" `
	Position           string `json:"position" form:"position" `
	Occupation         string `json:"occupation" form:"occupation" `
	BaptismWater       bool   `json:"baptism_water" form:"baptism_water" `
	BaptismSpirit      bool   `json:"baptism_spirit" form:"baptism_spirit" `
	YearConversion     int64  `json:"year_conversion" form:"year_conversion" `
	Active             bool   `json:"active" form:"active" `
	DetachmentId       uint   `json:"detachmentid" form:"detachmentid" `
	ChurchId           uint   `json:"churchid" form:"churchid" `
	RolId              uint   `json:"rolid" form:"rolid"`
}

//LoginDTO is a model that used by client when POST from /login url
type LoginDTO struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password" `
}

type UserPasswordDTO struct {
	Id              uint   `json:"id" form:"id"`
	Password        string `json:"password" form:"password" `
	PasswordConfirm string `json:"passwordconfirm" form:"passwordconfirm" `
}
