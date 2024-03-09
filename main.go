package main

import (
	"log"

	"log-analyzer/config"
	"log-analyzer/logAnalyzer"
	"log-analyzer/report"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	log.Println("Load config done")

	var r logAnalyzer.ResultStats
	stats, err := r.AnalyzeLog(cfg)
	if err != nil {
		panic(err)
	}
	log.Printf("Calc stats done <%+v>", stats)

	report.CreateReport(stats)
}
