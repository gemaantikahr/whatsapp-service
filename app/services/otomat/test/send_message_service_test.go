package test

import (
	"github.com/joho/godotenv"
	"testing"
	otomat2 "whatsapp_service/app/repositories/otomat"
	"whatsapp_service/app/repositories/otomat/models/request"
	"whatsapp_service/app/services/otomat"
)

func TestSendMessageService_Do(t *testing.T) {

	errEnv := godotenv.Load("../../../../.env")
	if errEnv != nil {
		println("error load env")
		return
	}

	var sendMessageRequest request.SendMessageRequest
	sendMessageRequest.Phone = "+6289680988232"
	sendMessageRequest.Text = "from service"

	res := otomat.SendMessageService{
		Data: sendMessageRequest,
		Repo: otomat2.OtomatRepository{},
	}.Do()

	if res.Status != true {
		println("something went wrong ! ", res.Message)
		return
	}

	println("success ! ", res.Message)
}
