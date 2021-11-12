package entity

import "time"

type StudyCarried struct {
	Id              uint       `gorm:"primary_key:auto_increment" json:"id"`
	Description     string     `gorm:"type:text" json:"title"`
	NumberCompleted uint8      `gorm:"type:int(4)" json:"number_completed"`
	Active          bool       `gorm:"type:TINYINT" json:"active"`
	CreatedAt       time.Time  `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
	RoleId          uint       `gorm:"NULL" json:"-"`
	Role            Role       `gorm:"foreignkey:RoleId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"role"`
}
