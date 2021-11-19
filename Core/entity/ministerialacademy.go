package entity

import "time"

type MinisterialAcademy struct {
	Id          uint       `gorm:"primary_key:auto_increment" json:"id"`
	Name        string     `gorm:"type:varchar(255)" json:"name"`
	DateAcademy string     `gorm:"type:Date" json:"dateacademy"`
	Place       string     `gorm:"type:text" json:"place"`
	Active      bool       `gorm:"type:TINYINT" json:"active"`
	CreatedAt   time.Time  `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
	UserId      uint       `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user_id"`
}
