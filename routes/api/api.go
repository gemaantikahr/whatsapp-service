package api

import (
	"net/http"
	"whatsapp_service/app/base"
	"whatsapp_service/app/controllers"

	"github.com/gin-gonic/gin"
)

type api struct{}

func (a api) Do(router *base.Router) {
	api := router.Group("/api")

	// Just for example
	api.GET("/hello_world", func(c *gin.Context) {
		c.JSON(http.StatusOK, struct {
			Status  bool   `json:"status"`
			Message string `json:"message"`
		}{
			Status:  true,
			Message: "api hello world!",
		})
	})

	apiWhatsapp := router.Group("/api/whatsapp")
	apiWhatsapp.POST("/send", controllers.WhatsappController{}.SendMessage)

}

func Init() api {
	return api{}
}
