package validation

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

// TODO:
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

// i think latency should be 'number+ms' every time but i need to check
// func getLatencyInMs(latencyStr string) (latencyF32 float32, error) {
// 		split := strings.Split(latencyStr, "ms")
//		if len(split) != 2 {
//			return 0, errors.New("not found ms in latency")
// 		}
// 		latency, err := strconv.ParseFloat(split[0], 32)
//		if err!= nil {
// 			return 0, err
// 		}
//		return float32(latency), nil
// }
