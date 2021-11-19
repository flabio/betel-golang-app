package entity

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id                  uint                  `gorm:"primary_key:auto_increment" json:"id"`
	Name                string                `gorm:"type:varchar(255)" json:"name"`
	Image               string                `gorm:"type:text" json:"image"`
	LastName            string                `gorm:"type:varchar(255)" json:"last_name"`
	Identification      string                `gorm:"type:varchar(255)" json:"identification"`
	TypeIdentification  string                `gorm:"type:varchar(2)" json:"type_identification"`
	Birthday            string                `gorm:"type:varchar(12)" json:"birthday"`
	Birthplace          string                `gorm:"type:varchar(255)" json:"birthplace"`
	Gender              string                `gorm:"type:varchar(2)" json:"gender"`
	Direction           string                `gorm:"type:varchar(255)" json:"direction"`
	PhoneNumber         string                `gorm:"type:varchar(255)" json:"phone_number"`
	CellPhone           string                `gorm:"type:varchar(255)" json:"cell_phone"`
	CivilStatus         string                `gorm:"type:varchar(255)" json:"civil_status"`
	Position            string                `gorm:"type:varchar(255)" json:"position"`
	Occupation          string                `gorm:"type:text" json:"occupation"`
	BaptismWater        bool                  `gorm:"type:TINYINT" json:"baptism_water"`
	BaptismSpirit       bool                  `gorm:"type:TINYINT" json:"baptism_spirit"`
	YearConversion      int64                 `gorm:"type:int(20)" json:"year_conversion"`
	Email               string                `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password            string                `gorm:"->;<-;not null" json:"-"`
	Token               string                `gorm:"-" json:"token,omitempty"`
	Active              bool                  `gorm:"type:TINYINT" json:"active"`
	CreatedAt           time.Time             `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt           *time.Time            `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
	Roles               *Role                 `json:"rolid,omitempty"`
	UserSubdetachements *UserSubdetachement   `json:"usersubdetachement,omitempty"`
	MinisterialAcademys *[]MinisterialAcademy `json:"ministerialacademy,omitempty"`

	DetachmentId uint       `gorm:"NULL" json:"detachmentid"`
	Detachment   Detachment `gorm:"foreignkey:DetachmentId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"detachment"`
	ChurchId     uint       `gorm:"NULL" json:"churchid"`
	Church       Church     `gorm:"foreignkey:ChurchId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"church"`
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	if u.Email == "" {
		return errors.New("rollback invalid user")
	}
	return nil
}
