package routes

import (
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

	err := router.Run()
	if err != nil {
		println("something went wrong !")
		return
	}
}
