package metrics

import "strconv"

type ServerType string

const (
	ProxyServer ServerType = "proxy"
	FinalServer ServerType = "final"
	LoadGen     ServerType = "loadgen"
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
	var (
		wasErrorInt int
		errorTime   = strconv.Itoa(m.ErrorTime)
		successTime = strconv.Itoa(m.SuccessTime)
	)
	if m.WasError {
		wasErrorInt = 1
	}
	if m.ErrorTime == 0 {
		errorTime = ""
	}
	if m.SuccessTime == 0 {
		successTime = ""
	}

	return []string{
		string(m.Server),
		m.RequestId,
		strconv.Itoa(wasErrorInt),
		m.ErrorType,
		errorTime,
		successTime,
	}
}

func getCSVHeader() []string {
	return []string{
		"Server",
		"Request id",
		"Was error",
		"Error type",
		"Error time",
		"Success time",
	}
}
