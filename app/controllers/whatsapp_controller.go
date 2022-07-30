package controllers

import (
	"net/http"
	"whatsapp_service/app/models/whatsapp/dto/request"
	"whatsapp_service/app/response"
	"whatsapp_service/app/services"

	"github.com/gin-gonic/gin"
)

type WhatsappController struct {
}

func (c WhatsappController) SendMessage(ctx *gin.Context) {

	var sendMessageRequestDTO request.SendMessageRequestDTO
	errDTO := ctx.ShouldBind(&sendMessageRequestDTO)
	if errDTO != nil {
		ctx.JSON(http.StatusBadRequest, response.Api().Error(errDTO.Error(), "send_message", nil))
		return
	}

	service := services.ExecuteSendMessageService{Data: sendMessageRequestDTO}.Do()
	if service.Status != true {
		ctx.JSON(http.StatusBadRequest, response.Api().Error(service.Message, "send_message", nil))
		return
	}

	ctx.JSON(http.StatusOK, response.Api().Success(service.Message, "send_message", nil))
}
