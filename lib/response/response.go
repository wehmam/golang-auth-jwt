package response

type (
	Response struct {
		ResponseCode string `json:"responceCode"`
		ResponseMessage string `json:"responseMessage"`
		Data interface{} `json:"data,omitempty"`
	}
)