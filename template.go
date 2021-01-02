package main

import (
	"bytes"
	"log"
	"text/template"

	"github.com/gin-gonic/gin"
	md "github.com/gomarkdown/markdown"
)

func (e *renderCtx) renderEmailTemplate() string {
	t, err := template.New("Email").Parse(e.MD)
	if err != nil {
		log.Println(err)
	}
	var op bytes.Buffer
	if err := t.Execute(&op, e); err != nil {
		log.Println(err)
	}
	htmlContent := md.ToHTML([]byte(op.String()), nil, nil)
	return string(htmlContent)
}

// Render's Markdown from Request
func renderMD(c *gin.Context) {
	mdr := c.PostForm("mdb")
	if mdr != "" {
		renderedMD := md.ToHTML([]byte(mdr), nil, nil)
		log.Print("MD parsing Successfull")
		c.String(200, string(renderedMD))
	} else {
		log.Print("MD Parsing Failed - Empty String")
	}
}
