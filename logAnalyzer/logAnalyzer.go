package loganalyzer

import "log-analyzer/model"

// struct resultStats
type ResultStats struct {
	LogTypeCount        map[string]int
	HTTPStatusCodeCount map[string]int
	RespTime            RespTimeStat
	CalledHostCount     map[string]int
	// TimeZoneCount       map[string]int, for now have no idea how can i know the zones
}

type RespTimeStat struct {
	Summary      AllRoutesRespTime
	EachRoute    EachRouteRespTime
	SlowReqRatio string
}

// calc in the end
type AllRoutesRespTime struct {
	Max float32
	Min float32
	Avg float32
}

type EachRouteRespTime struct {
	EachRouteRespTimes           map[string]EachRouteRespTimeSummary
	EachRouteRespTimeCalculating map[string]EachRouteRespTimeCalculating
}

// calc in the end
type EachRouteRespTimeSummary struct {
	Max float32
	Min float32
	Avg float32
}

// calc every line processing
type EachRouteRespTimeCalculating struct {
	TotalLatency int
	CalledCount  int
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
	//		route, err := getRoute(args.URI)
	//		r.RespTimeStat.EachRoute.EachRouteRespTimeCalculating[route].TotalLatency += latency //TODO: i think i should shorten this
	//      r.RespTimeStat.EachRoute.EachRouteRespTimeCalculating[route].CalledCount += latency

	//      if r.RespTimeStat.EachRoute.EachRouteRespTimeSummary[route].min > latency {
	//	      r.RespTimeStat.EachRoute.EachRouteRespTimeSummary[route].min = latency
	//      }
	// 		if r.RespTimeStat.EachRoute.EachRouteRespTimeSummary[route].max < latency {
	// 			r.RespTimeStat.EachRoute.EachRouteRespTimeSummary[route].max = latency
	// 		}
	//      // avg would be calc in calcSummaryStat(...)

	//		r.CalledHostCount[args.Host] += 1

	// 		// TODO: calc TimeZoneCount
}
