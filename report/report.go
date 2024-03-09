package report

import (
	"fmt"
	"log-analyzer/logAnalyzer"
	"os"
)

var (
	LogReportTemplateFinal = `## Log Report

%s

%s

%s
	
	%s

%s

%s`

	ErrorLogNumber = `#### 1. Error log number and ratio analysis
    Total: %s
    Error: %s
    Error Ratio: %s
    
    Overall:
%s
	
`

	LogAnalysisByStatusCode = `#### 2. Log analysis by HTTP status code
	HTTP STATUS CODES:
%s
	
`

	ResponseTimeAnalysis = `#### 3. Response time analysis
	All Routes:
		Min: %s ms
		Avg: %s ms
		Max: %s ms
		
		Total Req: %s req 
		Slow Req:  %s req
		SlowRate:  %s`

	ResponseTimeAnalysisEachRoute = `Each Routes
%s
	
`

	ParseTheReqURI = `#### 4. Parse the request URI
%s
`

	AnalysisByTimePeriod = `#### 5. Analysis by time period 
%s

`
)

func CreateReport(stat logAnalyzer.ResultStats) error {
	textErrorLogNumber := putStatToErrorLogNumber(stat)
	textLogAnalysisByStatusCode := putStatToLogAnalysisByStatusCode(stat)
	textResponseTimeAnalysis := putStatToResponseTimeAnalysis(stat)
	textResponseTimeAnalysisEachRoute := putStatToResponseTimeAnalysisEachRoute(stat)
	textParseTheReqURI := putStatToParseTheReqURI(stat)
	textAnalysisByTimePeriod := putStatToAnalysisByTimePeriod(stat)

	finalText := fmt.Sprintf(LogReportTemplateFinal,
		textErrorLogNumber,
		textLogAnalysisByStatusCode,
		textResponseTimeAnalysis,
		textResponseTimeAnalysisEachRoute,
		textParseTheReqURI,
		textAnalysisByTimePeriod,
	)

	err := os.WriteFile("log-report.md", []byte(finalText), 0644)
	if err != nil {
		return err
	}

	return nil
}

func putStatToErrorLogNumber(stat logAnalyzer.ResultStats) string {
	totalCnt := stat.RespTime.AllRoutes.TotalReqCnt
	errorCnt := 0
	var errorRatio float32

	_, ok := stat.LogLevelCnt["ERROR"]
	if ok {
		errorCnt = stat.LogLevelCnt["ERROR"]
		errorRatio = float32(stat.LogLevelCnt["ERROR"]) / float32(totalCnt) * 100
	}

	overAll := ""
	for key, val := range stat.LogLevelCnt {
		overAll += fmt.Sprintf("        %v: %v\n", key, val)
	}

	totalCntStr := fmt.Sprintf("%v", totalCnt)
	errorCntStr := fmt.Sprintf("%v", errorCnt)
	errorRatioStr := fmt.Sprintf("%.2f percent", errorRatio)

	return fmt.Sprintf(ErrorLogNumber, totalCntStr, errorCntStr, errorRatioStr, overAll)
}

func putStatToLogAnalysisByStatusCode(stat logAnalyzer.ResultStats) string {
	overAll := ""
	for key, val := range stat.CodeCnt {
		keyStr := fmt.Sprintf("%v", key)
		valStr := fmt.Sprintf("%v", val)
		overAll += fmt.Sprintf("        %v: %v\n", keyStr, valStr)
	}

	return fmt.Sprintf(LogAnalysisByStatusCode, overAll)
}

func putStatToResponseTimeAnalysis(stat logAnalyzer.ResultStats) string {
	min := fmt.Sprintf("%v", stat.RespTime.AllRoutes.Min)
	avg := fmt.Sprintf("%.2f ms", stat.RespTime.AllRoutes.Avg)
	max := fmt.Sprintf("%v", stat.RespTime.AllRoutes.Max)

	totalReqCnt := fmt.Sprintf("%v", stat.RespTime.AllRoutes.TotalReqCnt)
	slowReqCnt := fmt.Sprintf("%v", stat.RespTime.AllRoutes.SlowReqCnt)
	slowRate := fmt.Sprintf("%.2f percent", stat.RespTime.AllRoutes.SlowRate)

	return fmt.Sprintf(ResponseTimeAnalysis,
		min,
		avg,
		max,
		totalReqCnt,
		slowReqCnt,
		slowRate,
	)
}

func putStatToResponseTimeAnalysisEachRoute(stat logAnalyzer.ResultStats) string {
	overAll := ""
	eachRoute := `		%s:
				Min: %s ms
				Avg: %s ms
				Max: %s ms

				Total Req: %s req 
				Slow Req:  %s req
				SlowRate:  %s percent
`
	for route, val := range stat.RespTime.EachRoute {
		totalReqCnt := fmt.Sprintf("%v", val.TotalReqCnt)
		slowReqCnt := fmt.Sprintf("%v", val.SlowReqCnt)
		slowRate := fmt.Sprintf("%v", val.SlowRate)
		minStr := fmt.Sprintf("%v", val.Min)
		avgStr := fmt.Sprintf("%v", val.Avg)
		maxStr := fmt.Sprintf("%v", val.Max)

		eachRouteFilled := fmt.Sprintf(eachRoute,
			"/"+route,
			minStr,
			avgStr,
			maxStr,
			totalReqCnt,
			slowReqCnt,
			slowRate,
		)

		overAll += eachRouteFilled + "\n"
	}

	return fmt.Sprintf(ResponseTimeAnalysisEachRoute, overAll)
}

func putStatToParseTheReqURI(stat logAnalyzer.ResultStats) string {
	overAll := ""
	for _, val := range stat.SortedHostCnt {
		cntStr := fmt.Sprintf("%v", val.Count)
		overAll += fmt.Sprintf("        %v: %v\n", val.Host, cntStr)
	}

	return fmt.Sprintf(ParseTheReqURI, overAll)
}

func putStatToAnalysisByTimePeriod(stat logAnalyzer.ResultStats) string {
	overAll := ""
	for key, val := range stat.TimeZoneCnt {
		cntStr := fmt.Sprintf("%v", val)
		overAll += fmt.Sprintf("        %v: %v\n", key, cntStr)
	}

	return fmt.Sprintf(AnalysisByTimePeriod, overAll)
}
