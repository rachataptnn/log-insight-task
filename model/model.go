package model

// struct for LogSummary, no need?
// {"level":"INFO","ts":"2024-01-23T05:07:55.905Z","caller":"/app/domain/user/service/service.go:183","msg":"success"}

// struct for LogDetail
// {"URI":"/api/v1/payment/YZa3LRPteb1ANvDWlwb65c8GbTvPbogE","caller":"/app/cmd/api/main.go:99","header":null,"host":"https://api.example.com","latency":"190ms","level":"INFO","method":"GET","msg":"request","real_ip":"162.133.131.14","request_id":"dCnmS6BW5sqfIvL1looHUMpTWPvkWWc1","response_size":"469","status":200,"ts":"2024-01-23T05:07:55.905Z"}
type LogDetail struct {
	URI          string      `json:"URI"`
	Caller       string      `json:"caller"`
	Header       interface{} `json:"header"` // Assuming header can be of any type
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

// struct arguments from log line
type Arguments struct{}

// struct for min, max, avg response time
