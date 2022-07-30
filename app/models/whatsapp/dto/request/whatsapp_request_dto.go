package request

type SendMessageRequestDTO struct {
	Channel string `json:"channel" binding:"required"`
	Phone   string `json:"phone" binding:"required"`
	Message string `json:"message" binding:"required"`
}
