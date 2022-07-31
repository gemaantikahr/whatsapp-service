package api

import (
	"net/http"
	"whatsapp_service/app/base"
	"whatsapp_service/app/controllers"
	"whatsapp_service/app/middleware"

	"github.com/gin-gonic/gin"
)

type Api struct{}

func (a Api) Do(router *base.Router) {
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

	apiWhatsapp := router.Group("/api/whatsapp", middleware.AppMiddleware{}.Do())
	apiWhatsapp.POST("/send", controllers.WhatsappController{}.SendMessage)

}

func Init() Api {
	return Api{}
}
