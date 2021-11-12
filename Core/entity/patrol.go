package entity

import "time"

type Patrol struct {
	Id              uint          `gorm:"primary_key:auto_increment" json:"id" `
	Name            string        `gorm:"type:varchar(255)" json:"name"`
	Url             string        `gorm:"type:text" json:"url"`
	Archives        string        `gorm:"type:text" json:"archives"`
	CreatedAt       time.Time     `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt       *time.Time    `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
	Active          bool          `gorm:"type:TINYINT" json:"active"`
	SubDetachmentId uint          `gorm:"null" json:"subdetachmentid"`
	SubDetachment   SubDetachment `gorm:"foreignkey:SubDetachmentId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"subdetachment"`
}
