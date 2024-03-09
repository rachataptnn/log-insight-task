package report

import (
	"fmt"
	"log-analyzer/logAnalyzer"
)

func CreateReport(stat logAnalyzer.ResultStats) {
	fmt.Println("stat is", stat)
}
