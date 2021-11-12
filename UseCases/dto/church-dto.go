package dto

type ChurchDTO struct {
	Id          uint   `json:"id" form:"id" `
	Name        string `json:"name" form:"name" binding:"required"`
	Direction   string `json:"direction" form:"direction" binding:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" binding:"required"`
	Active      bool   `json:"active" form:"active"`
}
