package base

import "os"

type OtomatConfig struct {
}

func (b OtomatConfig) BaseUrl() string {
	return os.Getenv("OTOMAT_BASE_URL_1")
}

func (b OtomatConfig) ApiID() string {
	return os.Getenv("OTOMAT_API_ID_1")
}

func (b OtomatConfig) ApiKey() string {
	return os.Getenv("OTOMAT_API_KEY_1")
}
