package wrappers

type ApiResponseWrapper struct {
	Code int
	Body any
}

func NewResponseWrapper(code int, response any) *ApiResponseWrapper {
	return &ApiResponseWrapper{code, response}
}
