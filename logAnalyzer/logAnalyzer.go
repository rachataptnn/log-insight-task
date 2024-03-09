package loganalyzer

import "log-analyzer/model"

// struct resultStats
type ResultStats struct {
	LogTypeCount        map[string]int
	HTTPStatusCodeCount map[string]int
	ResponseTime        ResponseTimeStat
	CalledHostCount     map[string]int
	// TimeZoneCount       map[string]int, for now have no idea how can i know the zones
}

type ResponseTimeStat struct {
	Summary      ResponseTimeSummary
	EachRoutes   map[string]ResponseTimeEachRoutes
	SlowReqRatio string
}

type ResponseTimeSummary struct {
	Max float32
	Min float32
	Avg float32
}

type ResponseTimeEachRoutes struct {
	Max float32
	Min float32
	Avg float32
}

func AnalyzeLog() ResultStats {
	//		var result model.ResultStats
	//
	// 		for line := range reader.Lines() {
	//     		if evenLine {
	//         		logDetail, err := getLogDetail(line)
	//         		calcStat(args)
	//     		}
	//		}
	// 		calcSummaryStat(...)
	return ResultStats{}
}

func (r ResultStats) CalcStat(args model.Arguments) {
	// 		r.LogTypeCount[args.LogType] += 1
	// 		r.HTTPStatusCodeCount[args.HTTPStatusCode] += 1
	// 		latency, err := getLatencyInMs(args.Latency)
	//
	//		args.Latency
	//		args.Host
}

// func
