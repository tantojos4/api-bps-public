package response

type StatusOk struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func NewStatusOk(message string) *StatusOk {
	return &StatusOk{
		Message: message,
		Status:  "ok",
	}
}
