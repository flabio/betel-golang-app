package dto

type RolDTO struct {
	Id     uint   `json:"id" form:"id" xml:"id"`
	Name   string `json:"name" form:"name" xml:"name" binding:"required"`
	Active bool   `json:"active" form:"active" xml:"active"`
}
