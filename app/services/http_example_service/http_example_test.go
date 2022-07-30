package http_example_service

import (
	"testing"
	"whatsapp_service/app/models"
	"whatsapp_service/app/repositories"
)

func TestHttpExample_Do(t *testing.T) {
	service := HttpExample{repositories.HttpExampleRepository{}}.Do()

	if service.Status == true {
		for _, value := range service.Data.(models.Posts) {
			println("title :", value.Title)
		}
	}
	println(service.Message)
}
