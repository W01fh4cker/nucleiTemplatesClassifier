package main

import (
	"flag"
	"fmt"
	"nucleiTemplatesClassifier/api"
	"nucleiTemplatesClassifier/utils"
	"os"
)

func main() {
	severity := flag.String("s", "critical,high", "Severity filter")
	years := flag.String("y", "2017,2018,2019,2020,2021,2022,2023", "Year filter")
	outputFile := flag.String("o", "output.html", "Output HTML file")
	targetDir := flag.String("d", "", "Absolute path to nuclei-templates, eg: C:\\Users\\Administrator\\nuclei-templates")
	apiPort := flag.Int("p", 8082, "http api port")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s -s critical,high -y 2017,2018,2019,2020,2021,2022,2023 -o output.html -d C:\\Users\\whoami\\nuclei-templates\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
	if *targetDir == "" {
		flag.Usage()
		os.Exit(1)
	}
	cves := utils.GetCVEsInfo(*targetDir, *years, *severity)
	err := utils.GenerateHtmlFile(*outputFile, cves)
	api.ApiStart(*targetDir, *years, *severity, *apiPort)
	if err != nil {
		fmt.Printf("生成HTML文件时发生错误：%v\n", err)
	}
}
