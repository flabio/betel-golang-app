package entity

import "time"

type ParentScout struct {
	Id        uint       `gorm:"primary_key:auto_increment" json:"id" `
	UserId    uint       `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"userid"`
	RoleId    uint       `gorm:"foreignkey:RoleId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"roleid"`
	CreatedAt time.Time  `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
	Active    bool       `gorm:"type:TINYINT" json:"active"`
}
