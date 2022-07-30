package sendtalk

import (
	"whatsapp_service/app/repositories/sendtalk"
	"whatsapp_service/app/repositories/sendtalk/models/request"
	"whatsapp_service/app/response"
)

type SendMessageService struct {
	Data request.SendMessageRequest
	Repo sendtalk.SendTalkRepoInterface
}

func (s SendMessageService) Do() response.ServiceResponse {

	repo, err := s.Repo.SendMessage(s.Data)

	if err != nil {
		return response.Service().Error(err.Error(), nil)
	}

	if repo.Error.Message != "" {
		return response.Service().Error(repo.Error.Message, nil)
	}

	return response.Service().Success(repo.Data.Message, repo)

}
