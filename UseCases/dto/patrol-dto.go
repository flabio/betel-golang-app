package dto

type PatrolDTO struct {
	Id              uint   `json:"id" form:"id" `
	Name            string `json:"name" form:"name" binding:"required"`
	Url             string `json:"url" form:"url"`
	Archives        string `json:"archives" form:"archives"`
	SubDetachmentId uint   `json:"subdetachmentid" form:"subdetachmentid" binding:"required"`
	Active          bool   `json:"active" form:"active"`
}

type PatrolResponse struct {
	Id                uint   `json:"id"  `
	Name              string `json:"name"`
	Url               string `json:"url" `
	Archives          string `json:"archives"`
	SubDetachmentId   uint   `json:"subdetachmentid"`
	SubDetachmentName string `json:"subdetachment_name"`
	Active            bool   `json:"active"`
}
