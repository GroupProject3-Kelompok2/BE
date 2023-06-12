package helper

type DataResponse struct {
	Code       int         `json:"code,omitempty"`
	Status     string      `json:"status,omitempty"`
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
