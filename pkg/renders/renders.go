package renders

import (
	"fmt"
	"net/http"
	"text/template"
)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsetemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsetemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing in template", err)
		return
	}
}
