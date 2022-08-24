package dto

import "mime/multipart"

type DetachmentDTO struct {
	Id       uint                  `json:"id" form:"id" `
	Name     string                `json:"name" form:"name" binding:"required"`
	District string                `json:"district" form:"district" binding:"required"`
	Number   uint8                 `json:"number" form:"number" binding:"required"`
	Section  uint8                 `json:"section" form:"section" binding:"required"`
	Active   bool                  `json:"active" form:"active"`
	Archives string                `json:"archives" form:"archives" `
	File     *multipart.FileHeader `json:"file" form:"file" `
}

type DetachmentListDTO struct {
	Id         uint   `json:"id" `
	Name       string `json:"name"`
	District   string `json:"district"`
	Number     uint8  `json:"number" `
	Section    uint8  `json:"section" `
	Active     bool   `json:"active" `
	Archives   string `json:"archives" `
	ChurchId   uint   `json:"church_id"`
	ChurchName string `json:"church_name"`
}
