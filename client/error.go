package client

type ErrorDetail struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Field   string `json:"field"`
}

type ErrorResponse struct {
	Errors []ErrorDetail `json:"errors"`
}

func (a *ErrorResponse) Error() string {
	return a.Errors[0].Message
}
