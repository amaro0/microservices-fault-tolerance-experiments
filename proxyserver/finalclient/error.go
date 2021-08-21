package finalclient

import "fmt"

const (
	TimeoutError    = "timeout"
	UnexpectedError = "unexpected"
	ClientError     = "client"
	UnknownError    = "unknown"
)

type RequestError struct {
	ErrorType string
	Err       error
}

func NewRequestError(errorType string) *RequestError {
	return &RequestError{
		ErrorType: errorType,
	}
}

func (r *RequestError) AttachError(e error) {
	r.Err = e
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("Final server returned error")
}
