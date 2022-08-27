package dto

import "mime/multipart"

type SubDetachmentDTO struct {
	Id           uint                  `json:"id" form:"id" `
	Name         string                `json:"name" form:"name" binding:"required"`
	Url          string                `json:"url" form:"url"`
	Archives     string                `json:"archives" form:"archives"`
	DetachmentId uint                  `json:"detachmentid" form:"detachmentid" binding:"required"`
	Active       bool                  `json:"active" form:"active"`
	File         *multipart.FileHeader `json:"file" form:"file" `
}

type SubDetachmentListDTO struct {
	Id             uint   `json:"id" form:"id" `
	Name           string `json:"name" form:"name" `
	Archives       string `json:"archives" form:"archives"`
	DetachmentName string `json:"detachment_name" form:"detachment_name"`
	DetachmentId   uint   `json:"detachmentid" form:"detachmentid"`
	Active         bool   `json:"active" form:"active"`
}
