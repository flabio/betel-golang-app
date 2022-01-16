package entity

import "time"

type Attendance struct {
	Id uint `gorm:"primary_key:auto_increment" json:"id"`

	Notebook          bool          `gorm:"type:TINYINT" json:"notebook"`
	Bible             bool          `gorm:"type:TINYINT" json:"bible"`
	Attendance        bool          `gorm:"type:TINYINT" json:"attendance"`
	Uniform           bool          `gorm:"type:TINYINT" json:"uniform"`
	Offering          bool          `gorm:"type:TINYINT" json:"offering"`
	Active            bool          `gorm:"type:TINYINT" json:"active"`
	UserId            uint          `gorm:"null" json:"user_id"`
	User              User          `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
	SubDetachmentId   uint          `gorm:"null" json:"subdetachment_id"`
	SubDetachment     SubDetachment `gorm:"foreignkey:SubDetachmentId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"subdetachment"`
	WeekSubDetachment string        `gorm:"type:varchar(100)" json:"week_sub_detachment"`
	CreatedAt         time.Time     `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt         *time.Time    `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
}
