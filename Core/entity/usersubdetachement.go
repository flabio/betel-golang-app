package entity

import "time"

// UserSubdetachement struct represents role table in database
type UserSubdetachement struct {
	Id     uint `gorm:"primary_key:auto_increment" json:"id" `
	RolId  uint `gorm:"foreignkey:RolId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"rols_id"`
	UserId uint `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"users_id"`
	//Rol    Rol  `gorm:"foreignkey:RolId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"rol"`
	//User      User       `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	Active    bool       `gorm:"type:TINYINT" json:"active"`
	CreatedAt time.Time  `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
}
