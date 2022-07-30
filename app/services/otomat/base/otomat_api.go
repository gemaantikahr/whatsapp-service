package base

import (
	"whatsapp_service/app/base"
)

type OtomatApi struct {
	OtomatConfig
}

func (b OtomatApi) Get(endpoint string) base.NetClient {

	data := map[string]string{
		"api_id":  b.ApiID(),
		"api_key": b.ApiKey(),
	}

	return base.HttpService().Get().Url(b.BaseUrl() + endpoint).AddUrlParam(data)
}
