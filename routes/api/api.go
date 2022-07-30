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

	api.GET("/member", controllers.MemberController{}.Index)
	api.GET("/member-lite", controllers.MemberController{}.IndexLite)
	api.GET("/posts", controllers.HttpExampleController{}.GetPosts)
}

func Init() api {
	return api{}
}
