package dto

type StudyCarriedDTO struct {
	Id              uint   `json:"id" form:"id" `
	Description     string `json:"description" form:"description" binding:"required"`
	NumberCompleted uint8  `json:"number_completed" form:"number_completed" binding:"required"`
	Active          bool   `json:"active" form:"active"`
	RoleId          uint   `json:"roleid" form:"roleid" binding:"required"`
}
