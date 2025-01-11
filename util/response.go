package util

type BaseResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

func NewSuccessResponse(message string, data any) BaseResponse {
	return BaseResponse{
		Success: true,
		Data:    data,
		Message: message,
	}
}

func NewFailedResponse(message string) BaseResponse {
	return BaseResponse{
		Message: message,
	}
}
