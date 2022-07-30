package test

import (
	"github.com/joho/godotenv"
	"testing"
	sendtalk2 "whatsapp_service/app/repositories/sendtalk"
	"whatsapp_service/app/repositories/sendtalk/models/request"
	"whatsapp_service/app/services/sendtalk"
)

func TestSendMessageService_Do(t *testing.T) {

	errEnv := godotenv.Load("../../../../.env")
	if errEnv != nil {
		println("error load env")
		return
	}

	// defined send message request
	var sendMessageRequest request.SendMessageRequest
	sendMessageRequest.MessageType = "text"
	sendMessageRequest.Body = "oke siap bos"
	sendMessageRequest.Phone = "+6289680988232"

	response := sendtalk.SendMessageService{
		Data: sendMessageRequest,
		Repo: sendtalk2.SendTalkRepository{},
	}.Do()

	if response.Status != true {
		println("something went wrong ! ", response.Message)
		return
	}

	println("success ", response.Message)

}
