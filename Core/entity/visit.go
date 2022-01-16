package entity

import "time"

type Visit struct {
	Id              uint          `gorm:"primary_key:auto_increment" json:"id"`
	State           string        `gorm:"type:varchar(50)" json:"state"`
	Description     string        `gorm:"type:text" json:"description"`
	UserId          uint          `gorm:"null" json:"userid"`
	User            User          `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	SubDetachmentId uint          `gorm:"null" json:"subdetachmentid"`
	SubDetachment   SubDetachment `gorm:"foreignkey:SubDetachmentId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"subdetachment"`
	CreatedAt       time.Time     `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt       *time.Time    `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
	Active          bool          `gorm:"type:TINYINT" json:"active"`
}
