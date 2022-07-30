package sendtalk

import (
	"encoding/json"
	"io"
	"whatsapp_service/app/repositories/sendtalk/models/request"
	"whatsapp_service/app/repositories/sendtalk/models/response"
	"whatsapp_service/app/services/sendtalk/base"
)

type SendTalkRepoInterface interface {
	SendMessage(request request.SendMessageRequest) (response.SendMessageResponse, error)
}

type SendTalkRepository struct {
}

func (r SendTalkRepository) SendMessage(messageRequest request.SendMessageRequest) (response.SendMessageResponse, error) {

	// defined
	var sendMessageResponse response.SendMessageResponse
	call, err := base.SendTalkApi{}.Post("/api/v1/message/send_whatsapp").Bodys(messageRequest).Call()

	if err != nil {
		return sendMessageResponse, err
	}

	result, errRa := io.ReadAll(call.Body)
	if errRa != nil {
		return sendMessageResponse, nil
	}

	errUm := json.Unmarshal(result, &sendMessageResponse)
	if errUm != nil {
		return sendMessageResponse, errUm
	}

	return sendMessageResponse, nil

}
