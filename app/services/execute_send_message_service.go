package services

import (
	"whatsapp_service/app/models/whatsapp/dto/request"
	otomat2 "whatsapp_service/app/repositories/otomat"
	request2 "whatsapp_service/app/repositories/otomat/models/request"
	sendtalk2 "whatsapp_service/app/repositories/sendtalk"
	request3 "whatsapp_service/app/repositories/sendtalk/models/request"
	"whatsapp_service/app/response"
	"whatsapp_service/app/services/otomat"
	"whatsapp_service/app/services/sendtalk"
)

type ExecuteSendMessageService struct {
	Data request.SendMessageRequestDTO
}

func (s ExecuteSendMessageService) Do() response.ServiceResponse {

	switch s.Data.Channel {

	case "otomat":

		var sendMessageRequest request2.SendMessageRequest
		sendMessageRequest.Phone = s.Data.Phone
		sendMessageRequest.Text = s.Data.Message

		service := otomat.SendMessageService{
			Data: sendMessageRequest,
			Repo: otomat2.OtomatRepository{},
		}.Do()

		if service.Status != true {
			return response.Service().Error(service.Message, nil)
		}

		return response.Service().Success("message sent from otomat !", nil)
	case "sendtalk":

		var sendMessageRequest request3.SendMessageRequest
		sendMessageRequest.MessageType = "text"
		sendMessageRequest.Body = s.Data.Message
		sendMessageRequest.Phone = s.Data.Phone

		service := sendtalk.SendMessageService{
			Data: sendMessageRequest,
			Repo: sendtalk2.SendTalkRepository{},
		}.Do()

		if service.Status != true {
			return response.Service().Error(service.Message, nil)
		}

		return response.Service().Success("message sent from sendtalk !", nil)
	}

	return response.Service().Error("channel not found !", nil)
}
