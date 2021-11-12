package entity

import "time"

type RoleChurch struct {
	Id        uint       `gorm:"primary_key:auto_increment" json:"id"`
	RoleId    uint       `gorm:"foreignkey:RoleId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"role"`
	ChurchId  uint       `gorm:"foreignkey:ChurchId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"church"`
	CreatedAt time.Time  `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
}
