package controllers

import (
	"whatsapp_service/app/repositories"
	"whatsapp_service/app/response"
	"whatsapp_service/app/services"

	"github.com/gin-gonic/gin"
)

type MemberController struct{}

func (receiver MemberController) Index(c *gin.Context) {
	data := c.Request.URL.Query()

	if data["id"] == nil || data["id"][0] == "" {
		c.JSON(400, response.Api().Error("member id can not be null", nil))
		return
	}

	member := repositories.Member().Find(data["id"][0])

	if member.Name == "" {
		c.JSON(400, response.Api().Error("member not found", nil))
		return
	}

	c.JSON(200, response.Api().Success("loaded", member))
}

func (receiver MemberController) IndexLite(c *gin.Context) {
	data := c.Request.URL.Query()

	if data["id"] == nil || data["id"][0] == "" {
		c.JSON(400, response.Api().Error("member id can not be null", nil))
		return
	}

	service := services.SimpleMember(data["id"][0]).Do()
	if service.Status == false {
		c.JSON(400, response.Api().Error(service.Message, nil))
		return
	}

	c.JSON(200, response.Api().Success(service.Message, service.Data))
}