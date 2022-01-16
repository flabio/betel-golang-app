package entity

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id                     uint                  `gorm:"primary_key:auto_increment" json:"id"`
	Name                   string                `gorm:"type:varchar(255)" json:"name"`
	Image                  string                `gorm:"type:text" json:"image"`
	DocumentIdentification string                `gorm:"type:text" json:"document_identification"`
	LastName               string                `gorm:"type:varchar(255)" json:"last_name"`
	Identification         string                `gorm:"type:varchar(255)" json:"identification"`
	TypeIdentification     string                `gorm:"type:varchar(2)" json:"type_identification"`
	Birthday               string                `gorm:"type:varchar(12)" json:"birthday"`
	Birthplace             string                `gorm:"type:varchar(255)" json:"birthplace"`
	Gender                 string                `gorm:"type:varchar(2)" json:"gender"`
	Rh                     string                `gorm:"type:varchar(2)" json:"rh"`
	Direction              string                `gorm:"type:varchar(255)" json:"direction"`
	PhoneNumber            string                `gorm:"type:varchar(255)" json:"phone_number"`
	CellPhone              string                `gorm:"type:varchar(255)" json:"cell_phone"`
	CivilStatus            string                `gorm:"type:varchar(255)" json:"civil_status"`
	Position               string                `gorm:"type:varchar(255)" json:"position"`
	Occupation             string                `gorm:"type:text" json:"occupation"`
	School                 string                `gorm:"type:varchar(255)" json:"school"`
	Grade                  string                `gorm:"type:varchar(255)" json:"grade"`
	HobbiesInterests       string                `gorm:"type:varchar(255)" json:"hobbies_interests"`
	Allergies              string                `gorm:"type:varchar(255)" json:"allergies"`
	BaptismWater           bool                  `gorm:"type:TINYINT" json:"baptism_water"`
	BaptismSpirit          bool                  `gorm:"type:TINYINT" json:"baptism_spirit"`
	Year                   int64                 `gorm:"type:int(4)" json:"year"`
	YearConversion         int64                 `gorm:"type:int(20)" json:"year_conversion"`
	Email                  string                `gorm:"type:varchar(255)" json:"email"`
	Password               string                `gorm:"->;<-;null" json:"-"`
	Token                  string                `gorm:"-" json:"token,omitempty"`
	Active                 bool                  `gorm:"type:TINYINT" json:"active"`
	CreatedAt              *time.Time            `gorm:"type:TIMESTAMP(6)" json:"created_at"`
	UpdatedAt              time.Time             `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
	Roles                  *Role                 `json:"rolid,omitempty"`
	UserSubdetachements    *UserSubdetachement   `json:"usersubdetachement,omitempty"`
	MinisterialAcademys    *[]MinisterialAcademy `json:"ministerialacademy,omitempty"`
	SubDetachmentId        uint                  `gorm:"NULL" json:"subdetachmentid"`
	SubDetachment          SubDetachment         `gorm:"foreignkey:SubDetachmentId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"subdetachment"`
	ChurchId               uint                  `gorm:"NULL" json:"churchid"`
	Church                 Church                `gorm:"foreignkey:ChurchId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"church"`
	CityId                 uint                  `gorm:"NULL" json:"cityid"`
	City                   City                  `gorm:"foreignkey:CityId;constraint:onUpdate:CASCADE,onDelete:SET NULL" json:"city"`
	ParentId               uint                  `gorm:"NULL" json:"parentid"`
	Parent                 Parent                `gorm:"foreignkey:ParentId;constraint:onUpdate:CASCADE,onDelete:SET NULL" json:"parent"`
	// DetachmentId uint       `gorm:"NULL" json:"detachmentid"`
	// Detachment   Detachment `gorm:"foreignkey:DetachmentId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"detachment"`

}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	if u.Email == "" {
		return errors.New("rollback invalid user")
	}
	return nil
}
