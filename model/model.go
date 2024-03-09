package model

type LogDetail struct {
	URI          string      `json:"URI"`
	Caller       string      `json:"caller"`
	Header       interface{} `json:"header"`
	Host         string      `json:"host"`
	Latency      string      `json:"latency"`
	Level        string      `json:"level"`
	Method       string      `json:"method"`
	Message      string      `json:"msg"`
	RealIP       string      `json:"real_ip"`
	RequestID    string      `json:"request_id"`
	ResponseSize string      `json:"response_size"`
	Status       int         `json:"status"`
	Timestamp    string      `json:"ts"`
}
