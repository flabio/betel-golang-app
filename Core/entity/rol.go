package entity

import "time"

type Rol struct {
	Id         uint          `gorm:"primary_key:auto_increment" json:"id"`
	Name       string        `gorm:"type:varchar(255)" json:"name"`
	Active     bool          `gorm:"type:TINYINT" json:"active"`
	RoleModule *[]RoleModule `json:"rolemodules,omitempty"`
	CreatedAt  time.Time     `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt  *time.Time    `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
}

//(HDna(sPvygCP7lrla22
