package test

import (
	"github.com/joho/godotenv"
	"testing"
	"whatsapp_service/app/repositories/otomat"
	"whatsapp_service/app/repositories/otomat/models/request"
)

func TestOtomatRepository_SendMessage(t *testing.T) {

	errEnv := godotenv.Load("../../../../.env")

	if errEnv != nil {
		println("cannot load the env ", errEnv.Error())
	}

	// defined send message request
	var sendMessageRequest request.SendMessageRequest
	sendMessageRequest.Text = "okeh gema cakepp"
	sendMessageRequest.Phone = "+6289680988232"

	res, err := otomat.OtomatRepository{}.SendMessage(sendMessageRequest)

	if err != nil {
		println("something went wrong ! ", err.Error())
		return
	}

	println("the result ", res.Message)

}
