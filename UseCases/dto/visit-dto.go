package dto

type VisitRequest struct {
	Id              uint   `json:"id" form:"id"`
	State           string `json:"state" form:"state" binding:"required" `
	Description     string `json:"description" form:"description"`
	UserId          uint   `json:"userid" form:"userid" binding:"required" `
	SubDetachmentId uint   `json:"subdetachmentid" form:"subdetachmentid" binding:"required" `
	Active          bool   `json:"active" form:"active"`
}

type VisitResponse struct {
	Id                uint   `json:"id"`
	State             string `json:"state"`
	Description       string `json:"description" `
	UserId            uint   `json:"userid"`
	UserFullName      string `json:"user_full_name"`
	SubDetachmentId   uint   `json:"subdetachmentid"`
	SubDetachmentName string `json:"subdetachment_name"`
	Active            bool   `json:"active"`
}
