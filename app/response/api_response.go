package response

func Api() apiResponse {
	return apiResponse{}
}

type apiResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Method  string      `json:"method"`
	Data    interface{} `json:"data"`
}

func (apiResponse) Error(message string, method string, data interface{}) apiResponse {
	return apiResponse{
		Status:  false,
		Message: message,
		Method:  method,
		Data:    data,
	}
}

func (apiResponse) Success(message string, method string, data interface{}) apiResponse {
	return apiResponse{
		Status:  true,
		Message: message,
		Method:  method,
		Data:    data,
	}
}
