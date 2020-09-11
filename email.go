package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmailPageData struct {
	Buttons []Button
	// MD      []byte
}
type Button struct {
	Text string
	Link string
}

func renderEmails(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := EmailPageData{
			Buttons: []Button{
				{Text: "Module 1 TextBooks", Link: "https://blog.athulcyriac.co"},
				{Text: "Module 2 TextBooks", Link: "https://athulcyriac.co"},
			},
		}
		tmpl.Execute(w, data)
	})
}
