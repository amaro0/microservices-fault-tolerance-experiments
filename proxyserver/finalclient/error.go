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

func NewRequestError(errorType string, error error) *RequestError {
	return &RequestError{
		ErrorType: errorType,
		Err:       error,
	}
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("Final server returned error")
}
