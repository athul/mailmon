package main

import (
	"bytes"
	"log"
	"text/template"
)

// Name returns the Name of the recipeint
func (r *renderCtx) Name() string {
	log.Println("Rec Name", r.Recipient)
	if r.Recipient == "" {
		return "There"
	}
	return r.Recipient
}

//Header adds the link to the header image
func (r *renderCtx) Header(link string) string {
	r.ImageURL = link
	if r.ImageURL == "" && link == "" {
		r.ImageURL = "https://i.imgur.com/4s3lF1V.png"
	}
	return ""
}

// AddButtons add buttons to the email
func (r *renderCtx) AddButtons(text, link string) string {
	var b bytes.Buffer
	newButton := Button{
		Text: text,
		Link: link,
	}
	templ, err := template.New("Button").Parse(`<a href="{{ .Link }}" class="aligncenter btn-primary" itemprop="url" style="font-family: Avenir,Arial,sans-serif; box-sizing: border-box; font-size: 14px; color: #FFF; text-decoration: none; line-height: 2em; font-weight: bold; text-align: center; cursor: pointer; display: inline-block; border-radius: 5px; text-transform: capitalize; background-color: #348eda; margin: 0; border-color: #348eda; border-style: solid; border-width: 10px 20px;">{{ .Text }}</a>`)
	if err != nil {
		log.Println(err)
	}
	templ.Execute(&b, newButton)
	return b.String()
}

func (r *renderCtx) InFooter(text, link string) string {
	var b bytes.Buffer
	footerStruct := struct {
		Link string
		Text string
	}{
		Link: link,
		Text: text,
	}

	templ, err := template.New("Footer").Parse(`Follow <a href="{{.Link}}" style="font-family: Avenir,Arial,sans-serif; box-sizing: border-box; font-size: 12px; color: #999; text-decoration: underline; margin: 0;">{{.Text}}</a> on Twitter.`)
	if err != nil {
		log.Println(err)
	}
	templ.Execute(&b, footerStruct)
	r.Footer = b.String()
	return ""
}
