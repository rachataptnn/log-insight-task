package logAnalyzer

import (
	"bufio"
	"encoding/json"
	"log"
	"log-analyzer/config"
	"log-analyzer/model"
	"log-analyzer/validation"
	"os"
)

// struct resultStats
type ResultStats struct {
	LogLevelCount       map[string]int
	HTTPStatusCodeCount map[int]int
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
	EachRouteRespTimeSummary     map[string]EachRouteRespTimeSummary
	EachRouteRespTimeCalculating map[string]EachRouteRespTimeCalculating
}

// calc in the end
type EachRouteRespTimeSummary struct {
	Max int
	Min int
	Avg float32
}

// calc every line processing
type EachRouteRespTimeCalculating struct {
	TotalLatency int
	CalledCount  int
}

func (r *ResultStats) Initialize() {
	r.LogLevelCount = make(map[string]int)
	r.HTTPStatusCodeCount = make(map[int]int)
	r.CalledHostCount = make(map[string]int)
	r.RespTime.EachRoute.EachRouteRespTimeSummary = make(map[string]EachRouteRespTimeSummary)
	r.RespTime.EachRoute.EachRouteRespTimeCalculating = make(map[string]EachRouteRespTimeCalculating)
}

func (r *ResultStats) AnalyzeLog(cfg *config.Config) (*ResultStats, error) {
	r.Initialize()

	// Open the file
	file, err := os.Open(cfg.LogFilePath)
	if err != nil {
		return &ResultStats{}, err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Loop through each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		var logDetail model.LogDetail
		err := json.Unmarshal([]byte(line), &logDetail)
		if err != nil {
			log.Println("Error parsing line as LogDetail:", err)
			continue
		}

		r.CalcEachLieStat(logDetail)

	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		return &ResultStats{}, err
	}

	r.CalcSummaryStat()

	return r, nil
}

func (r *ResultStats) CalcEachLieStat(logDetail model.LogDetail) error {
	r.LogLevelCount[logDetail.Level] += 1
	r.HTTPStatusCodeCount[logDetail.Status] += 1

	latency, err := validation.GetLatencyInMs(logDetail.Latency)
	if err != nil {
		return err
	}

	route, err := validation.GetRoute(logDetail.URI)
	if err != nil {
		return err
	}

	// Retrieve the element from the map
	calc := r.RespTime.EachRoute.EachRouteRespTimeCalculating[route]

	// Update its fields
	calc.CalledCount++
	calc.TotalLatency += latency

	// Assign it back to the map
	r.RespTime.EachRoute.EachRouteRespTimeCalculating[route] = calc

	// Update summary
	summary := r.RespTime.EachRoute.EachRouteRespTimeSummary[route]
	if latency < summary.Min || summary.Min == 0 {
		summary.Min = latency
	}
	if latency > summary.Max {
		summary.Max = latency
	}
	r.RespTime.EachRoute.EachRouteRespTimeSummary[route] = summary

	hostMap := r.CalledHostCount[logDetail.Host]
	hostMap += 1

	// TODO: calc TimeZoneCount

	return nil
}

func (r *ResultStats) CalcSummaryStat() error {
	log.Println("break for see what struct look like")

	return nil
}
