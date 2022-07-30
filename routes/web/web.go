package web

import (
	"fmt"
	"whatsapp_service/app/base"

	"github.com/gin-gonic/gin"
)

type web struct{}

func (web) Do(router *base.Router) {
	router.GET("/", func(c *gin.Context) {
		fmt.Fprintf(c.Writer, "Hello world!")
	})
}

func Init() web {
	return web{}
}
