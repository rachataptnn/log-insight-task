package logAnalyzer

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"log-analyzer/config"
	"log-analyzer/model"
	"log-analyzer/validation"
	"os"
	"sort"
)

// ResultStats represents the statistics collected from log analysis.
type ResultStats struct {
	LogLevelCnt   map[string]int
	CodeCnt       map[int]int
	RespTime      RespTimeStat
	TimeZoneCnt   map[string]int
	HostCnt       map[string]int
	SortedHostCnt []Host
}

type Host struct {
	Host  string
	Count int
}

type RespTimeStat struct {
	AllRoutes AllRoutesRespTime
	EachRoute map[string]EachRouteRespTime
}

type AllRoutesRespTime struct {
	Min int
	Avg float32
	Max int

	TotalReqCnt int

	SlowReqCnt int
	SlowRate   float32
}

type EachRouteRespTime struct {
	Min int
	Avg float32
	Max int

	TotalLatency int
	TotalReqCnt  int

	SlowReqCnt int
	SlowRate   float32
}

// InitMaps for avoid nil pointer errors.
func (r *ResultStats) InitMaps() {
	r.LogLevelCnt = make(map[string]int)
	r.CodeCnt = make(map[int]int)
	r.HostCnt = make(map[string]int)
	r.TimeZoneCnt = make(map[string]int)
	r.RespTime.EachRoute = make(map[string]EachRouteRespTime)
}

// Core function for log analysis.
func (r *ResultStats) AnalyzeLog(cfg *config.Config) error {
	r.InitMaps()

	file, err := os.Open(cfg.LogFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var logDetail model.LogDetail
		err := json.Unmarshal([]byte(line), &logDetail)
		if err != nil {
			log.Println("Error parsing line as LogDetail:", err)
			continue
		}
		r.CalculateLineStat(logDetail)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	r.CalculateAvgRespTimeAndSlowRate()
	r.SortTopURICall()

	return nil
}

// Calculates statistics for each log line.
func (r *ResultStats) CalculateLineStat(logDetail model.LogDetail) error {
	r.LogLevelCnt[logDetail.Level]++
	r.CodeCnt[logDetail.Status]++

	latency, err := validation.GetLatencyInMs(logDetail.Latency)
	if err != nil {
		return err
	}

	route, err := validation.GetRoute(logDetail.URI)
	if err != nil {
		return err
	}

	calc := r.RespTime.EachRoute[route]
	calc.TotalReqCnt++
	calc.TotalLatency += latency
	if latency > 500 {
		calc.SlowReqCnt++
	}
	r.RespTime.EachRoute[route] = calc

	summary := r.RespTime.EachRoute[route]
	if latency < summary.Min || summary.Min == 0 {
		summary.Min = latency
	}
	if latency > summary.Max {
		summary.Max = latency
	}
	r.RespTime.EachRoute[route] = summary

	r.HostCnt[logDetail.Host]++

	timeZone, err := getTimezoneKey(logDetail.Timestamp)
	if err != nil {
		return err
	}
	r.TimeZoneCnt[timeZone]++

	return nil
}

// getTimezoneKey extracts the timezone from the timestamp.
func getTimezoneKey(timestamp string) (string, error) {
	if len(timestamp) < 24 {
		return "", errors.New("timestamp is too short")
	}
	key := timestamp[23:]
	return key, nil
}

// SortTopURICall sorts the top URI calls.
func (r *ResultStats) SortTopURICall() {
	var hostSlice []Host
	for host, count := range r.HostCnt {
		hostSlice = append(hostSlice, Host{
			Host:  host,
			Count: count,
		})
	}
	sort.Slice(hostSlice, func(i, j int) bool {
		return hostSlice[i].Count > hostSlice[j].Count
	})
	if len(hostSlice) > 5 {
		r.SortedHostCnt = hostSlice[:5]
	} else {
		r.SortedHostCnt = hostSlice
	}
}

// CalculateAvgRespTimeAndSlowRate calculates average response time.
func (r *ResultStats) CalculateAvgRespTimeAndSlowRate() {
	var allRoutesStackAvg float32
	var routeCnt int
	var AllRoutesMax int
	var AllRoutesMin int
	var allRoutesTotalReqCnt int
	var allRoutesSlowReqCnt int

	eachRouteMap := r.RespTime.EachRoute
	for routeName, routeStat := range eachRouteMap {
		avgRespTime := float32(routeStat.TotalLatency) / float32(routeStat.TotalReqCnt)
		fmt.Printf("\nroute name: %v\navg resp time: %v ms\n", routeName, avgRespTime)

		eachRouteSummary := r.RespTime.EachRoute[routeName]
		avg := float32(routeStat.TotalLatency) / float32(routeStat.TotalReqCnt)
		slowRate := float32(routeStat.SlowReqCnt) / float32(routeStat.TotalReqCnt) * 100
		eachRouteSummary.SlowRate = slowRate
		eachRouteSummary.Avg = avg
		r.RespTime.EachRoute[routeName] = eachRouteSummary

		allRoutesStackAvg += avg
		routeCnt++
		allRoutesTotalReqCnt += routeStat.TotalReqCnt
		allRoutesSlowReqCnt += routeStat.SlowReqCnt
	}
	r.RespTime.AllRoutes.Avg = allRoutesStackAvg / float32(routeCnt)

	for _, v := range r.RespTime.EachRoute {
		if v.Max > AllRoutesMax {
			AllRoutesMax = v.Max
		}
		if v.Min < AllRoutesMin || AllRoutesMin == 0 {
			AllRoutesMin = v.Min
		}
	}
	r.RespTime.AllRoutes.Max = AllRoutesMax
	r.RespTime.AllRoutes.Min = AllRoutesMin

	r.RespTime.AllRoutes.TotalReqCnt = allRoutesTotalReqCnt
	r.RespTime.AllRoutes.SlowReqCnt = allRoutesSlowReqCnt
	r.RespTime.AllRoutes.SlowRate = float32(allRoutesSlowReqCnt) / float32(allRoutesTotalReqCnt) * 100
}
