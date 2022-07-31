package request

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type SendMessageRequestDTO struct {
	Channel string `json:"channel"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

func (a SendMessageRequestDTO) ValidateSendMessageRequestDTO() error {

	return validation.ValidateStruct(&a,
		validation.Field(&a.Channel,
			validation.Required.Error("Channel cannot be empty"),
			validation.In("otomat", "sendtalk").Error("Channel not found")),

		validation.Field(&a.Phone,
			validation.Required.Error("Phone cannot be empty"),
			is.E164.Error("Phone number isn't valid"),
			validation.Length(8, 15)),

		validation.Field(&a.Message,
			validation.Required.Error("Message cannot be empty"),
			validation.Length(5, 255)),
	)
	
}
