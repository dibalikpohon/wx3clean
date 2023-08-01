package response

import "net/http"

type What struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type DataNotFoundErrorResponse struct {
	Status  string `json:"status"`
	Code    uint16 `json:"code"`
	Message string `json:"message"`
	What    What   `json:"what"`
}

func NewDataNotFoundErrorResponse(field, value string) DataNotFoundErrorResponse {
	return DataNotFoundErrorResponse{
		Status:  "Error",
		Code:    http.StatusNotFound,
		Message: http.StatusText(http.StatusNotFound),
		What: What{
			Field: field,
			Value: value,
		},
	}
}
