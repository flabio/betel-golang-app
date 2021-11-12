package dto

type SubDetachmentDTO struct {
	Id           uint   `json:"id" form:"id" `
	Name         string `json:"name" form:"name"`
	Url          string `json:"url" form:"url"`
	Archives     string `json:"archives" form:"archives"`
	DetachmentId uint   `json:"detachmentid" form:"detachmentid"`
	Active       bool   `json:"active" form:"active"`
}
