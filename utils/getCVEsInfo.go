package utils

import (
	"strings"
)

func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func GetCVEsInfo(targetDir string, years string, severity string) []CVE {
	var cves []CVE
	yamlFilesNames := GetYamlFilesPath(targetDir, years)
	for _, yamlFileName := range yamlFilesNames {
		cve := ParseNucleiYaml(yamlFileName)
		severities := strings.Split(severity, ",")
		if contains(severities, cve.Info.Severity) {
			cves = append(cves, cve)
		}
	}
	return cves
}
