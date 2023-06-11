package helper

type DataResponse struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
}

func ResponseFormat(code int, message string, data interface{}, pagination interface{}) DataResponse {
	result := DataResponse{
		Code:       code,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	}

	return result
}
