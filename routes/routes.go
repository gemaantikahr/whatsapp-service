package routes

import (
	"os"
	"whatsapp_service/app/base"
	"whatsapp_service/routes/api"
	"whatsapp_service/routes/web"
)

type RouteInterface interface {
	Do(router *base.Router)
}

func Init() {
	router := base.NewRouter()
	api.Init().Do(router)
	web.Init().Do(router)

	router.Run(":" + os.Getenv("APP_PORT"))
}
