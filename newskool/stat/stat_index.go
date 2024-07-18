package stat

import (
	"net/http"
	"text/template"
)

func StatIndex(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("ui/index.html")
	tmpl.Execute(w, nil)
}
