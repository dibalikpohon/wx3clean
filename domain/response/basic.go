package response

type BasicResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}
