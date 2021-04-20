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
	ErrorType   string     `json:"errorType"`
	ErrorTime   int        `json:"errorTime"`
	SuccessTime int        `json:"successTime"`
}

func (m *Model) prepareForCSV() []string {
	var wasErrorInt int8
	if m.WasError {
		wasErrorInt = 1
	}

	return []string{
		string(m.Server),
		m.RequestId,
		string(wasErrorInt),
		m.ErrorType,
		string(m.ErrorTime),
		string(m.SuccessTime),
	}
}
