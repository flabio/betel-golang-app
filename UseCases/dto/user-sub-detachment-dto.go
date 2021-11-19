package dto

type UserSubDetachmentDTO struct {
	UserId          uint `json:"userid" form:"userid"`
	SubDetachmentId uint `json:"subdetachmentid" form:"subdetachmentid"`
}
