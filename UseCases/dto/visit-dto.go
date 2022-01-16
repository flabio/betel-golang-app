package dto

type VisitDTO struct {
	Id              uint   `json:"id" form:"id"`
	State           string `json:"state" form:"state"`
	Description     string `json:"description" form:"state"`
	UserId          uint   `json:"userid" form:"userid"`
	SubDetachmentId uint   `json:"subdetachmentid" form:"subdetachmentid"`
	Active          bool   `json:"active" form:"active"`
}
