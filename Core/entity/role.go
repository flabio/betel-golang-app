package entity

// Role struct represents role table in database
type Role struct {
	Id           uint            `gorm:"primary_key:auto_increment" json:"id" `
	RolId        uint            `gorm:"rolid" json:"rolid"`
	UserId       uint            `gorm:"null" json:"user_id"`
	Rol          Rol             `gorm:"foreignkey:RolId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"rol"`
	RoleModule   *[]RoleModule   `json:"rolemodules,omitempty"`
	StudyCarried *[]StudyCarried `json:"studycarried,omitempty"`
	RoleChurch   *[]RoleChurch   `json:"rolechurchs,omitempty"`
	//User       User          `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
