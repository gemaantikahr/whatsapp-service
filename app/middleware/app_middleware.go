package middleware

import "C"
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"whatsapp_service/app/response"
)

type AppMiddleware struct {
}

func (m AppMiddleware) Do() gin.HandlerFunc {

	return func(c *gin.Context) {

		appKey := c.GetHeader("app-key")

		if os.Getenv("APP_KEY") != appKey {

			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Api().Error("you can't do this", "send_whatsapp", nil))
			return

		}

	}
}
