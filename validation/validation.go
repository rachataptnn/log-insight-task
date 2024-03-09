package validation

// i think latency should be 'number+ms' every time but i need to check
// func getLatencyInMs(latencyStr string) (latencyInt int, error) {
// 		split := strings.Split(latencyStr, "ms")
//		if len(split) != 2 {
//			return 0, errors.New("not found ms in latency")
// 		}
// 		latency, err := strconv.Atoi(split[0])
//		if err!= nil {
// 			return 0, err
// 		}
//		return latency, nil
// }

// wanna put this func in package logAnalyzer, anyway i nid to validate if URI is valid
// func getRoute(uri string) (route string, err error) {
//	ex URI "URI": "/api/v1/payment/YZa3LRPteb1ANvDWlwb65c8GbTvPbogE"
// 	split := strings.Split(uri, "/")
//  if len(split) < 3 {
// return "", errors.New("URI has less than 3 parts")
// }
// route = split[2] // need to check the correct index
// 	return route, nil
// }

// assume log file is not a perfect file
// maybe all LogSummary is not be placed in the odd line

// perfect log file pattern
// LogSummary
// LogDetail
// LogSummary
// LogDetail
// ...

// what if
// LogSummary
// LogDetail
// LogSummary
// LogSummary

// this func only use for check the soure file
// func is rawLogPerfect() bool {
// the first idea, LogSummary doesnt have many fields..
// i choose 'status' for validation if log detail is placed in the odd line
//     for {
//        	isOddLine {
//				logDetail,err := parseLogDetail(line)
//				if logDetail.Status != 0 {
//					fmt.Printf("\n<%+v>\n", logDetail)
// 				}
// 			}
//     }
// }

// TODO: maybe i should check if found logSummary, LogDetail 2 or 3 times in a row
