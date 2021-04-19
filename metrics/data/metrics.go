package data

type ServerType string

const (
	ProxyServer  ServerType = "proxy"
	FinalServer  ServerType = "final"
	ClientServer ServerType = "client"
)

type Metrics struct {
	Server      ServerType `json:"server" binding:"required"`
	RequestId   string     `json:"requestId" binding:"required"`
	WasError    string     `json:"wasError" binding:"required"`
	ErrorTime   int        `json:"errorTime" validate:"min=0"`
	SuccessTime int        `json:"successTime" validate:"min=0"`
	ErrorType   string     `json:"errorType"`
}
