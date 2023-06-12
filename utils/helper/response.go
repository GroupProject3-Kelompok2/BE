package helper

type DataResponse struct {
	Code       int         `json:"code"`
	Status     string      `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Pagination interface{} `json:"pagination,omitempty"`
}

func ResponseFormat(code int, status, message string, data interface{}, pagination interface{}) DataResponse {
	result := DataResponse{
		Code:       code,
		Status:     status,
		Message:    message,
		Data:       data,
		Pagination: pagination,
	}

	return result
}
