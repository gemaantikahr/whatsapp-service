package otomat

import (
	"encoding/json"
	"io"
	"whatsapp_service/app/helper"
	"whatsapp_service/app/repositories/otomat/models/request"
	"whatsapp_service/app/repositories/otomat/models/response"
	"whatsapp_service/app/services/otomat/base"
)

type OtomatRepoInterface interface {
	SendMessage(request request.SendMessageRequest) (response.SendMessageResponse, error)
}

type OtomatRepository struct {
}

func (r OtomatRepository) SendMessage(messageRequest request.SendMessageRequest) (response.SendMessageResponse, error) {

	// define send message response variable
	var sendMessageResponse response.SendMessageResponse

	// convert send message request to params as map data type
	params := helper.ConvertStructToMapString{
		Input: messageRequest,
	}.Do()

	call, err := base.OtomatApi{}.Get("").AddUrlParam(params).Call()
	if err != nil {
		return sendMessageResponse, nil
	}

	// read response body from otomat
	result, errRa := io.ReadAll(call.Body)
	if errRa != nil {
		return sendMessageResponse, nil
	}

	// mapping response body
	errMap := json.Unmarshal(result, &sendMessageResponse)
	if errMap != nil {
		return sendMessageResponse, errMap
	}

	return sendMessageResponse, nil
}
