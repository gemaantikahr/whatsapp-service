package http_example_service

import (
	"strconv"
	"testing"
	"whatsapp_service/app/models"
	"whatsapp_service/mocks"

	"github.com/golang/mock/gomock"
)

func TestHttpExample_Do(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	expected := models.Posts{}
	for i := 0; i < 3; i++ {
		expected = append(expected, struct {
			Body   string `json:"body"`
			ID     int64  `json:"id"`
			Title  string `json:"title"`
			UserID int64  `json:"userId"`
		}(struct {
			Body   string
			ID     int64
			Title  string
			UserID int64
		}{Body: "Body Mocked " + strconv.Itoa(i+1), ID: int64(i + 1), Title: "Title Mocked " + strconv.Itoa(i+1), UserID: 3}))
	}

	repoMock := mocks.NewMockHttpExampleInterface(ctrl)
	repoMock.EXPECT().GetPost().Return(expected).AnyTimes()

	service := HttpExample{repoMock}.Do()

	if service.Status == true {
		for _, value := range service.Data.(models.Posts) {
			println("title :", value.Title)
		}
	}
	println(service.Message)
}
