package controllers

import (
	"whatsapp_service/app/repositories"
	"whatsapp_service/app/response"
	"whatsapp_service/app/services/http_example_service"

	"github.com/gin-gonic/gin"
)

type HttpExampleController struct{}

func (HttpExampleController) GetPosts(c *gin.Context) {
	service := http_example_service.HttpExample{Repo: repositories.HttpExampleRepository{}}.Do()
	if service.Status == false {
		c.JSON(400, response.Api().Error(service.Message, nil))
		return
	}
	c.JSON(200, response.Api().Success(service.Message, service.Data))
}
