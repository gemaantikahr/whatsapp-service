package base

import "whatsapp_service/app/base"

type SendTalkApi struct {
	SendtalkConfig
}

func (a SendTalkApi) Post(endpoint string) base.NetClient {
	return base.HttpService().Post().Url(a.GetBaseUrl()+endpoint).AddHeader("API-Key", a.GetApiKey()).AddHeader("Content-Type", "application/json")
}
