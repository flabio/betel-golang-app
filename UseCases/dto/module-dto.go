package dto

type ModuleDTO struct {
	Id     uint   `json:"id" form:"id" `
	Name   string `json:"name" form:"name" `
	Active bool   `json:"active" form:"active"`
}

type ModuleRoleDTO struct {
	ModuleId uint `json:"moduleid" form:"moduleid" binding:"required"`
	RoleId   uint `json:"roleid" form:"roleid" binding:"required"`
}
