package services

import (
	"whatsapp_service/app/repositories"
	"whatsapp_service/app/response"
)

func SimpleMember(memberId any) simpleMember {
	return simpleMember{memberId: memberId}
}

type simpleMember struct{ memberId any }

func (s simpleMember) Do() response.ServiceResponse {
	member := repositories.Member().Find(s.memberId)
	if member.Name == "" {
		return response.Service().Error("member not found", nil)
	}

	data := make(map[string]interface{})
	data["name"] = member.Name
	data["email"] = member.Email

	return response.Service().Success("loaded", data)
}
