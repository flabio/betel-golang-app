package entity

import "time"

type WeeksDetachment struct {
	Id              uint       `gorm:"primary_key:auto_increment" json:"id" `
	Name            string     `gorm:"type:varchar(255)" json:"name"`
	CreatedAt       time.Time  `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt       *time.Time `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
	Active          bool       `gorm:"type:TINYINT" json:"active"`
	SubDetachmentId uint       `gorm:"foreignkey:SubDetachmentId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"subdetachmentid"`
}
