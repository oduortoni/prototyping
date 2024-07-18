package stat

import (
	"net/http"
	"text/template"
)

func StatIndex(response http.ResponseWriter, request *http.Request) {
	tmpl, _ := template.ParseFiles("ui/index.html")
	tmpl.Execute(response, nil)
}
