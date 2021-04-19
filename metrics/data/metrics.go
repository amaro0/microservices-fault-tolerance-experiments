package data

type Metrics struct {
	Server      string `json:"server" binding:"required"`
	RequestId   string `json:"requestId" binding:"required"`
	WasError    string `json:"wasError" binding:"required"`
	ErrorTime   int    `json:"errorTime" validate:"min=0"`
	SuccessTime int    `json:"successTime" validate:"min=0"`
	ErrorType   string `json:"errorType"`
}
