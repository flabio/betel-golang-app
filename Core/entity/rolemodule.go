package entity

import "time"

type RoleModule struct {
	Id        uint       `gorm:"primary_key:auto_increment" json:"id"`
	ModuleId  uint       `gorm:"null" json:"module"`
	Module    Module     `gorm:"foreignkey:ModuleId;constraint:onUpdate:CASCADE,onDelete:CASCADE" `
	RolId     uint       `gorm:"null" json:"rol"`
	Rol       Rol        `gorm:"foreignkey:RolId;constraint:onUpdate:CASCADE,onDelete:CASCADE" `
	CreatedAt time.Time  `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
}
