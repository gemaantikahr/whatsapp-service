package http_example_service

import (
	"whatsapp_service/app/repositories"
	"whatsapp_service/app/response"
)

type HttpExample struct {
	Repo repositories.HttpExampleInterface
}

func (h HttpExample) Do() response.ServiceResponse {

	data := h.Repo.GetPost()

	if data == nil {
		return response.Service().Error("no data found", nil)
	}

	return response.Service().Success("loaded", data)
}
