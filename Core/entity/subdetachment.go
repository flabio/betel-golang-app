package entity

import "time"

type SubDetachment struct {
	Id           uint       `gorm:"primary_key:auto_increment" json:"id" `
	Name         string     `gorm:"type:varchar(255)" json:"name"`
	Url          string     `gorm:"type:text" json:"url"`
	Archives     string     `gorm:"type:text" json:"archives"`
	DetachmentId uint       `gorm:"NULL" json:"detachmentid"`
	Detachment   Detachment `gorm:"foreignkey:DetachmentId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"detachment"`
	User         *[]User    `json:"users,omitempty"`
	Patrol       *[]Patrol  `json:"patrols,omitempty"`
	Active       bool       `gorm:"type:TINYINT" json:"active"`
	CreatedAt    time.Time  `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt    *time.Time `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
}
