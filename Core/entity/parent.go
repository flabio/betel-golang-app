package entity

import "time"

type Parent struct {
	Id                 uint         `gorm:"primary_key:auto_increment" json:"id"`
	FullName           string       `gorm:"type:varchar(250)" json:"full_name"`
	Identification     string       `gorm:"type:varchar(255)" json:"identification"`
	TypeIdentification string       `gorm:"type:varchar(2)" json:"type_identification"`
	TypeParent         string       `gorm:"type:varchar(20)" json:"type_parent"`
	Gender             string       `gorm:"type:varchar(10)" json:"gender"`
	Direction          string       `gorm:"type:varchar(255)" json:"direction"`
	PhoneNumber        string       `gorm:"type:varchar(255)" json:"phone_number"`
	ParentScouts       *ParentScout `json:"parentScout,omitempty"`
	Active             bool         `gorm:"type:TINYINT" json:"active"`
	CreatedAt          time.Time    `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt          *time.Time   `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
}
