package utils

import (
	"fmt"
	"html/template"
	"os"
)

func GenerateHtmlFile(outputFile string, cves []CVE) error {
	tmpl := template.Must(template.New("table").Parse(`
		<!DOCTYPE html>
		<html>
		<head>
			<title>CVE Data</title>
		</head>
		<body>
			<table border="1">
				<tr>
					<th>ID</th>
					<th>Name</th>
					<th>Severity</th>
					<th>Description</th>
					<th>Reference</th>
					<th>Shodan Query</th>
					<th>Fofa Query</th>
				</tr>
				{{range .}}
				<tr>
					<td>{{.ID}}</td>
					<td>{{.Info.Name}}</td>
					<td>{{.Info.Severity}}</td>
					<td>{{.Info.Description}}</td>
					<td>{{range .Info.Reference}}<a href="{{.}}">{{.}}</a><br>{{end}}</td>
					<td>{{.Metadata.ShodanQuery}}</td>
					<td>{{.Metadata.FofaQuery}}</td>
				</tr>
				{{end}}
			</table>
		</body>
		</html>
	`))

	htmlFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer htmlFile.Close()

	err = tmpl.Execute(htmlFile, cves)
	if err != nil {
		return err
	}

	fmt.Printf("[+] Results have been exported to: %s\n", outputFile)
	return nil
}
