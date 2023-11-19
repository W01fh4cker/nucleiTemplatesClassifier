package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

type Info struct {
	Name        string   `yaml:"name"`
	Severity    string   `yaml:"severity"`
	Description string   `yaml:"description"`
	Reference   []string `yaml:"reference"`
}

type Metadata struct {
	ShodanQuery string `yaml:"shodan-query"`
	FofaQuery   string `yaml:"fofa-query"`
}

type CVE struct {
	ID       string   `yaml:"id"`
	Info     Info     `yaml:"info"`
	Metadata Metadata `yaml:"metadata"`
}

func ParseNucleiYaml(yamlFileName string) CVE {
	yamlFile, err := os.ReadFile(yamlFileName)
	if err != nil {
		fmt.Printf("[-] An error occurred while reading <%s>. The error content is: %v", yamlFileName, err)
		os.Exit(1)
	}
	var cve CVE
	err = yaml.Unmarshal(yamlFile, &cve)
	if err != nil {
		fmt.Printf("[-] An error occurred while parsing <%s>. The error content is: %v", yamlFileName, err)
		os.Exit(1)
	}
	cve.Info.Description = strings.TrimRight(cve.Info.Description, "\n")
	return cve
}
