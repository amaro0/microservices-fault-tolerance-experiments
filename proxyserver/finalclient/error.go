package finalclient

import "fmt"

const (
	TimeoutError = iota
	UnexpectedError
	ClientError
	UnknownError
)

type RequestError struct {
	ErrorCode int
	Err       error
}

func NewRequestError(errorCode int, error error) *RequestError {
	return &RequestError{
		ErrorCode: errorCode,
		Err:       error,
	}
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("Final server returned error")
}
