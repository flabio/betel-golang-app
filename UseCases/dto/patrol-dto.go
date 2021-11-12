package dto

type PatrolDTO struct {
	Id              uint   `json:"id" form:"id" `
	Name            string `json:"name" form:"name"`
	Url             string `json:"url" form:"url"`
	Archives        string `json:"archives" form:"archives"`
	SubDetachmentId uint   `json:"subdetachmentid" form:"subdetachmentid"`
	Active          bool   `json:"active" form:"active"`
}
