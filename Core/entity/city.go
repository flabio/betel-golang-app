package entity

import "time"

type City struct {
	Id         uint       `gorm:"primary_key:auto_increment" json:"id"`
	Name       string     `gorm:"type:varchar(255)" json:"name"`
	Estate     string     `gorm:"type:varchar(255)" json:"estate"`
	PostalCode string     `gorm:"type:varchar(255)" json:"postal_code"`
	Active     bool       `gorm:"type:TINYINT" json:"active"`
	CreatedAt  time.Time  `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt  *time.Time `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
}
