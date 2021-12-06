package entity

import (
	"time"
)

type Detachment struct {
	Id            uint             `gorm:"primary_key:auto_increment" json:"id" `
	Name          string           `gorm:"type:varchar(255)" json:"name"`
	Url           string           `gorm:"type:text" json:"url"`
	Archives      string           `gorm:"type:text" json:"archives"`
	District      string           `gorm:"type:varchar(255)" json:"district"`
	Number        uint8            `gorm:"type:int(50)" json:"number"`
	Section       uint8            `gorm:"type:int(50)" json:"section"`
	Active        bool             `gorm:"type:TINYINT" json:"active"`
	CreatedAt     time.Time        `gorm:"<-:created_at" json:"created_at"`
	UpdatedAt     *time.Time       `gorm:"type:TIMESTAMP(6)" json:"updated_at"`
	SubDetachment *[]SubDetachment `json:"subdetachments,omitempty"`
	// ChurchId      uint             `gorm:"NULL" json:"churchid"`
	// Church        Church           `gorm:"foreignkey:ChurchId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"church"`
}
