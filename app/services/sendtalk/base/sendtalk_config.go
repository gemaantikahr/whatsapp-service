package base

import "os"

type SendtalkConfig struct {
}

func (c SendtalkConfig) GetBaseUrl() string {
	return os.Getenv("SENDTALK_BASE_URL")
}

func (c SendtalkConfig) GetApiKey() string {
	return os.Getenv("SENDTALK_API_KEY")

}
