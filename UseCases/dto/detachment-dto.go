package dto

type DetachmentRequest struct {
	Id       uint   `json:"id" form:"id" `
	Name     string `json:"name" form:"name" binding:"required"`
	District string `json:"district" form:"district" binding:"required"`
	Number   uint8  `json:"number" form:"number" binding:"required"`
	Section  uint8  `json:"section" form:"section" binding:"required"`
	Active   bool   `json:"active" form:"active"`
	ChurchId uint   `json:"churchid" form:"churchid" binding:"required"`
	//File     *multipart.FileHeader `json:"file" form:"file"`
	Archives string `json:"archives" form:"archive"`
}

type DetachmentResponse struct {
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
