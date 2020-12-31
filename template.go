package main

import (
	"bytes"
	"log"
	"text/template"
)

func renderEmailTemplate(text string, std Student) string {
	log.Println("Ivde vannu")
	defer log.Println("Poyi")
	t, err := template.New("Email").Parse(text)
	if err != nil {
		log.Println(err)
	}
	var op bytes.Buffer
	if err := t.Execute(&op, std); err != nil {
		log.Println(err)
	}
	log.Println("--", op.String(), "---")
	return op.String()
}
