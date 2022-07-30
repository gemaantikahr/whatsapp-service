package base

import "whatsapp_service/app/response"

type ServiceInterface interface {
	Do() response.ServiceResponse
}
