package dto

type RolDTO struct {
	Id     uint   `json:"id" form:"id"`
	Name   string `json:"name" form:"name" binding:"required"`
	Active bool   `json:"active" form:"active"`
}
