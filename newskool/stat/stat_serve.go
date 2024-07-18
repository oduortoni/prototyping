package stat

import (
	"io"
	"net/http"
	"os"
	"text/template"
	"time"
)

func StatServe(response http.ResponseWriter, r *http.Request) {
	topic := r.FormValue("content")

	if topic != "proto" {
		tmpl, _ := template.ParseFiles("ui/unknown.html")
		tmpl.Execute(response, nil)
		return
	}

	// simulate the dalay by one minute
	sleepForOneMinute()

	// Open the PDF file
	pdfFile, _ := os.Open("ui/proto.pdf")
	defer pdfFile.Close()

	// Set the appropriate HTTP headers
	response.Header().Set("Content-Type", "application/pdf")
	response.Header().Set("Content-Disposition", "inline; filename=proto.pdf")

	// Write the PDF file to the response
	io.Copy(response, pdfFile)
}

func sleepForOneMinute() {
	duration := 15 * time.Second
	time.Sleep(duration)
}
