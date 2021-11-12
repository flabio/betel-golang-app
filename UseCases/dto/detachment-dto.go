package dto

import "mime/multipart"

type DetachmentDTO struct {
	Name     string                `json:"name" form:"name" binding:"required"`
	District string                `json:"district" form:"district" binding:"required"`
	Number   uint8                 `json:"number" form:"number" binding:"required"`
	Section  uint8                 `json:"section" form:"section" binding:"required"`
	Active   bool                  `json:"active" form:"active"`
	Archives string                `json:"archives" form:"archives" `
	File     *multipart.FileHeader `json:"file" form:"file" `
}

type DetachmentUpdateDTO struct {
	Id       uint                  `json:"id" form:"id" binding:"required" `
	Name     string                `json:"name" form:"name" binding:"required"`
	District string                `json:"district" form:"district" binding:"required"`
	Number   uint8                 `json:"number" form:"number" binding:"required"`
	Section  uint8                 `json:"section" form:"section" binding:"required"`
	Active   bool                  `json:"active" form:"active" `
	Archives string                `json:"archives" form:"archives"`
	File     *multipart.FileHeader `json:"file" form:"file" `
}
