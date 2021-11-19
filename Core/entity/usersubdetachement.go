package entity

import "time"

// UserSubdetachement struct represents role table in database
type UserSubdetachement struct {
	Id              uint          `gorm:"primary_key:auto_increment" json:"id" `
	SubDetachmentId uint          `gorm:"subdetachmentid" json:"subdetachmentid"`
	SubDetachment   SubDetachment `gorm:"foreignkey:SubDetachmentId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"subdetachment"`
	UserId          uint          `gorm:"null" json:"user_id"`
	User            User          `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	Active          bool          `gorm:"type:TINYINT" json:"active"`
	CreatedAt       time.Time     `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt       *time.Time    `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
}
