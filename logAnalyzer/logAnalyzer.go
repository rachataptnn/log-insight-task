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

// struct resultStats
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
	EachRoute EachRouteRespTime
}

// calc in the end
type AllRoutesRespTime struct {
	Max      int
	Min      int
	Avg      float32
	SlowRate float32
}

type EachRouteRespTime struct {
	EachRouteRespTimeSummary map[string]EachRouteRespTimeSummary
	EachRouteRespTimeForCalc map[string]EachRouteRespTimeForCalcAvg
}

// calc in the end
type EachRouteRespTimeSummary struct {
	Max      int
	Min      int
	Avg      float32
	SlowRate float32
}

// calc every line processing
type EachRouteRespTimeForCalcAvg struct {
	TotalLatency int
	TotalReqCnt  int
	SlowReqCnt   int
}

func (r *ResultStats) Initialize() {
	r.LogLevelCnt = make(map[string]int)
	r.CodeCnt = make(map[int]int)
	r.HostCnt = make(map[string]int)
	r.SortedHostCnt = make([]Host, 0)
	r.TimeZoneCnt = make(map[string]int)
	r.RespTime.EachRoute.EachRouteRespTimeSummary = make(map[string]EachRouteRespTimeSummary)
	r.RespTime.EachRoute.EachRouteRespTimeForCalc = make(map[string]EachRouteRespTimeForCalcAvg)
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
	r.LogLevelCnt[logDetail.Level] += 1
	r.CodeCnt[logDetail.Status] += 1

	latency, err := validation.GetLatencyInMs(logDetail.Latency)
	if err != nil {
		return err
	}

	route, err := validation.GetRoute(logDetail.URI)
	if err != nil {
		return err
	}

	// Retrieve the element from the map
	calc := r.RespTime.EachRoute.EachRouteRespTimeForCalc[route]
	// Update its fields
	calc.TotalReqCnt++
	calc.TotalLatency += latency
	if latency > 500 {
		calc.SlowReqCnt++
	}

	// Assign it back to the map
	r.RespTime.EachRoute.EachRouteRespTimeForCalc[route] = calc

	// Update summary
	summary := r.RespTime.EachRoute.EachRouteRespTimeSummary[route]
	if latency < summary.Min || summary.Min == 0 {
		summary.Min = latency
	}
	if latency > summary.Max {
		summary.Max = latency
	}
	r.RespTime.EachRoute.EachRouteRespTimeSummary[route] = summary

	r.HostCnt[logDetail.Host] += 1

	timeZone, err := getTimezoneKey(logDetail.Timestamp)
	if err != nil {
		return err
	}
	r.TimeZoneCnt[timeZone] += 1

	return nil
}

func (r *ResultStats) CalcSummaryStat() error {
	log.Println("break for see what struct look like")

	r.calcAvgRespTime()

	// time out rate response time,
	// all routes
	// each route

	// sort top URI call
	r.sortTopURICall()

	return nil
}

// for collecting data that telling timezone, i would like to assume that 'another timezone' that not present in the log file would look like this:
// Timestamp                               Key
// "ts": "2024-01-22T21:07:55.905-08:00",  "-08:00"     represent Pacific Standard Time (PST)
// "ts": "2024-01-23T00:07:55.905-05:00"   "-05:00"     represent Eastern Standard Time (EST)
// "ts": "2024-01-23T05:07:55.905Z",       "Z"          represent Greenwich Mean Time (GMT)
// "ts": "2024-01-23T12:07:55.905+07:00"   "+07:00"     represent Indochina Time (ICT)
// some more...
func getTimezoneKey(timestamp string) (string, error) {
	if len(timestamp) < 24 {
		return "", errors.New("timestamp is too short")
	}

	key := timestamp[23:]
	return key, nil
}

func (r *ResultStats) sortTopURICall() {
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

func (r *ResultStats) calcAvgRespTime() {
	// var allRoutes
	var allRoutesStackAvg float32
	var routeCnt int
	var AllRoutesMax int
	var AllRoutesMin int

	var allRoutesTotalReqCnt int
	var allRoutesSlowReqCnt int

	eachRouteMap := r.RespTime.EachRoute.EachRouteRespTimeForCalc
	for routeName, routeStat := range eachRouteMap {
		avgRespTime := float32(routeStat.TotalLatency) / float32(routeStat.TotalReqCnt)
		fmt.Printf("\nroute name: %v\navg resp time: %v ms\n", routeName, avgRespTime)

		eachRouteSummary := r.RespTime.EachRoute.EachRouteRespTimeSummary[routeName]
		avg := float32(routeStat.TotalLatency) / float32(routeStat.TotalReqCnt)
		slowRate := float32(routeStat.SlowReqCnt) / float32(routeStat.TotalReqCnt) * 100
		eachRouteSummary.SlowRate = slowRate
		eachRouteSummary.Avg = avg
		r.RespTime.EachRoute.EachRouteRespTimeSummary[routeName] = eachRouteSummary

		allRoutesStackAvg += avg
		routeCnt++
		allRoutesTotalReqCnt += routeStat.TotalReqCnt
		allRoutesSlowReqCnt += routeStat.SlowReqCnt
	}
	r.RespTime.AllRoutes.Avg = allRoutesStackAvg / float32(routeCnt)

	for _, v := range r.RespTime.EachRoute.EachRouteRespTimeSummary {
		if v.Max > AllRoutesMax {
			AllRoutesMax = v.Max
		}
		if v.Min < AllRoutesMin || AllRoutesMin == 0 {
			AllRoutesMin = v.Min
		}
	}
	r.RespTime.AllRoutes.Max = AllRoutesMax
	r.RespTime.AllRoutes.Min = AllRoutesMin

	r.RespTime.AllRoutes.SlowRate = float32(allRoutesSlowReqCnt) / float32(allRoutesTotalReqCnt) * 100
}
