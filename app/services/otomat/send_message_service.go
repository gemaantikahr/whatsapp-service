package otomat

import (
	"whatsapp_service/app/repositories/otomat"
	"whatsapp_service/app/repositories/otomat/models/request"
	"whatsapp_service/app/response"
)

type SendMessageService struct {
	Data request.SendMessageRequest
	Repo otomat.OtomatRepoInterface
}

func (s SendMessageService) Do() response.ServiceResponse {

	repo, err := s.Repo.SendMessage(s.Data)

	if err != nil {
		return response.Service().Error(err.Error(), nil)
	}

	if repo.Status != "success" {
		return response.Service().Error(repo.Message, nil)
	}

	return response.Service().Success(repo.Message, nil)
}
