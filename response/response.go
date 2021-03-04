package response

import "encoding/json"

// StatusResponse struct
type StatusResponse struct {
	HTTPStatusCode       int    `json:"http_status_code"`
	Status               bool   `json:"status"`
	ResponseMessage      string `json:"response_message"`
	ResponseInternalCode int    `json:"response_internal_code"`
}

// GetJSON marshalls StatusResponse
func (s *StatusResponse) GetJSON() []byte {
	context, _ := json.Marshal(struct {
		StatusResponse *StatusResponse `json:"status_response"`
	}{
		StatusResponse: s,
	})

	return context
}

// New creates StatusResponse instance
func New() *StatusResponse {
	resp := &StatusResponse{
		HTTPStatusCode:       200,
		Status:               true,
		ResponseMessage:      "status:ok",
		ResponseInternalCode: 0,
	}
	return resp
}
