package main

import (
	"fmt"

	tplValidator "github.com/jiandahao/golanger/pkg/template/validator"
)

func main() {
	templ := `type {{ .Name.ToCamel }}Model struct {
		{{ range .Fields }}
			{{- .Name.ToCamel } {{.DataType}} ` + "`gorm:\"column:{{ .Name.Source }}\"`" + ` {{if .Comment }}// {{.Comment}} {{end}}
		{{ end -}}
	}`
	_, errs := tplValidator.Validate(templ)

	if len(errs) > 0 {
		fmt.Println("sdfsdfdsf", errs)
		tplValidator.PrintErrorDetails(templ, errs)
		return
	}
}
