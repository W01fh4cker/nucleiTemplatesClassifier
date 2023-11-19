package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".yaml") {
			*files = append(*files, path)
		}
		return nil
	}
}

func GetYamlFilesPath(targetDir string, years string) []string {
	var yamlFilesPath []string
	yearArray := strings.Split(years, ",")
	for _, year := range yearArray {
		targetDirc := targetDir + "\\http\\cves\\" + year
		//fmt.Printf("[*] Reading yaml file under %s\n", targetDirc)
		_ = filepath.Walk(targetDirc, visit(&yamlFilesPath))
	}

	return yamlFilesPath
}
