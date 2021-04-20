package metrics

type ServerType string

const (
	ProxyServer  ServerType = "proxy"
	FinalServer  ServerType = "final"
	ClientServer ServerType = "client"
)

type Model struct {
	Server      ServerType `json:"server" binding:"required"`
	RequestId   string     `json:"requestId" binding:"required"`
	WasError    bool       `json:"wasError"`
	ErrorTime   int        `json:"errorTime"`
	SuccessTime int        `json:"successTime"`
	ErrorType   string     `json:"errorType"`
}
