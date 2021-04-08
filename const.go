package toolkit

import "errors"

type ResponseType int8

const (
	RESPONSE_FLAG                  = "Response"
	RESPONSE_FLAG_SET              = "ResponseSet"
	Error                          = "error"
	Json              ResponseType = 0
	Redirect          ResponseType = 1
	Html              ResponseType = 2
)

var (
	ErrResponseAlreadySet = errors.New("response already set")
	ErrNotFoundResponse   = errors.New("not found response")
)
