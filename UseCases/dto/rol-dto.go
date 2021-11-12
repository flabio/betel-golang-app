package dto

type RolCreateDTO struct {
	Id     uint   `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	Active bool   `json:"active" form:"active"`
}

type RolUpdateDTO struct {
	Id     uint   `json:"id" form:"id" binding:"required"`
	Name   string `json:"name" form:"name" binding:"required"`
	Active bool   `json:"active" form:"active"`
}

type RolIdDTO struct {
	Id uint `json:"id" form:"id" binding:"required"`
}
type RoleDTO struct {
	RolId  uint `json:"rolId" form:"rolId" binding:"required"`
	UserId uint `json:"userId" form:"userId" binding:"required"`
}
