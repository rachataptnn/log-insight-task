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
	err = r.AnalyzeLog(cfg)
	if err != nil {
		panic(err)
	}
	log.Println("Calc stats done")

	err = report.CreateReport(r)
	if err != nil {
		panic(err)
	}
	log.Println("write Log report done")
}
