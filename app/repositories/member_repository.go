package repositories

import (
	"whatsapp_service/app/base"
	"whatsapp_service/app/models"
)

type MemberRepoInterface interface {
	Find(id int) models.Member
}

func Member() member { return member{} }

type member struct{}

func (mr member) Find(id any) models.Member {
	m := models.Member{}
	base.OpenDB().Gorm().Where("id = ?", id).First(&m)
	return m
}
