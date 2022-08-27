package dto

type UserDTO struct {
	Id                 uint   `json:"id" form:"id"`
	Image              string `json:"imagen" form:"imagen" `
	Name               string `json:"name" form:"name" binding:"required,min=3"`
	Email              string `json:"email" form:"email" binding:"required,email"`
	LastName           string `json:"last_name" form:"last_name" binding:"required"`
	Identification     string `json:"identification" form:"identification" binding:"required"`
	TypeIdentification string `json:"type_identification" form:"type_identification" binding:"required,max=4"`
	Birthday           string `json:"birthday" form:"birthday" binding:"required"`
	Birthplace         string `json:"birthplace" form:"birthplace" binding:"required"`
	Gender             string `json:"gender" form:"gender" binding:"required,max=1"`
	Rh                 string `json:"rh" form:"rh" binding:"required,max=3"`
	Direction          string `json:"direction" form:"direction" binding:"required"`
	PhoneNumber        string `json:"phone_number" binding:"required"`
	CellPhone          string `json:"cell_phone" form:"cell_phone" `
	CivilStatus        string `json:"civil_status" form:"civil_status" binding:"required"`
	Position           string `json:"position" form:"position" binding:"required"`
	Occupation         string `json:"occupation" form:"occupation" binding:"required"`
	School             string `json:"school" form:"school" `
	Grade              string `json:"grade" form:"grade"`
	HobbiesInterests   string `json:"hobbies_interests" form:"hobbies_interests"`
	Allergies          string `json:"allergies" form:"allergies"`
	BaptismWater       bool   `json:"baptism_water" form:"baptism_water" `
	BaptismSpirit      bool   `json:"baptism_spirit" form:"baptism_spirit" `
	YearConversion     int64  `json:"year_conversion" form:"year_conversion,numeric" `
	Active             bool   `json:"active" form:"active" `
	ChurchId           uint   `json:"churchid" form:"churchid" binding:"required,numeric" `
	RolId              uint   `json:"rolid" form:"rolid" binding:"required,numeric"`
	CityId             uint   `json:"cityid" form:"cityid" binding:"required,numeric"`
	Password           string `json:"password" form:"password" binding:"required,eqfield=ConfirmPassword"`
	ConfirmPassword    string `json:"confirm_password" form:"confirm_password" binding:"required" `
}

type ScoutDTO struct {
	Id                     uint   `json:"id" form:"id"`
	Image                  string `json:"imagen" form:"imagen" `
	DocumentIdentification string `json:"document_identification" form:"document_identification" binding:"required" `
	Name                   string `json:"name" form:"name" binding:"required"`
	Email                  string `json:"email" form:"email" `
	LastName               string `json:"last_name" form:"last_name" binding:"required" `
	Identification         string `json:"identification" form:"identification" binding:"required"`
	TypeIdentification     string `json:"type_identification" form:"type_identification" binding:"required"`
	Birthday               string `json:"birthday" form:"birthday" binding:"required"`
	Birthplace             string `json:"birthplace" form:"birthplace" binding:"required"`
	Gender                 string `json:"gender" form:"gender" binding:"required"`
	Direction              string `json:"direction" form:"direction" binding:"required"`
	CellPhone              string `json:"cell_phone" form:"cell_phone" binding:"required"`
	Rh                     string `json:"rh" form:"rh" binding:"required"`
	School                 string `json:"school" form:"school" `
	Grade                  string `json:"grade" form:"grade" binding:"required"`
	HobbiesInterests       string `json:"hobbies_interests" form:"hobbies_interests"`
	Allergies              string `json:"allergies" form:"allergies"`
	Active                 bool   `json:"active" form:"active" `
	SubDetachmentId        uint   `json:"subdetachmentid" form:"subdetachmentid" binding:"required" `
	ChurchId               uint   `json:"churchid" form:"churchid" binding:"required"`
	CityId                 uint   `json:"cityid" form:"cityid" binding:"required"`
}

// LoginDTO is a model that used by client when POST from /login url
type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

type UserPasswordDTO struct {
	Id              uint   `json:"id" form:"id"`
	Password        string `json:"password" form:"password" binding:"required,eqfield=ConfirmPassword"`
	PasswordConfirm string `json:"passwordconfirm" form:"passwordconfirm" binding:"required"`
}

type UserAuthDTO struct {
	Id         uint   `json:"id" form:"id"`
	Image      string `json:"imagen" form:"imagen" `
	Name       string `json:"name" form:"name" `
	LastName   string `json:"last_name" form:"last_name" `
	Active     bool   `json:"active" form:"active" `
	ChurchId   uint   `json:"churchid" form:"churchid" `
	ChurchName string `json:"church_name" form:"church_name"`
	RolId      uint   `json:"rolid" form:"rolid"`
	RolName    string `json:"rol_name" form:"rol_name"`
	Token      string `json:"token" form:"token"`
}
