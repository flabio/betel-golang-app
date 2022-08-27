package entity

import "time"

type Church struct {
	Id          uint       `gorm:"primary_key:auto_increment" json:"id"`
	Name        string     `gorm:"type:varchar(255)" json:"name"`
	Direction   string     `gorm:"type:varchar(255)" json:"direction"`
	PhoneNumber string     `gorm:"type:varchar(255)" json:"phone_number"`
	Active      bool       `gorm:"type:TINYINT" json:"active"`
	CreatedAt   time.Time  `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
	User        *[]User    `json:"users,omitempty"`
}
