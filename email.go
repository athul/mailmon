package main

import (
	"bytes"
	"html/template"
	"log"
)

//EmailPageData is the Data set in the Email
type EmailPageData struct {
	Buttons []Button
	MD      template.HTML
}

//Button is the type for a button in the Email
type Button struct {
	Text string
	Link string
}

func renderEmails(md template.HTML) string {
	var b bytes.Buffer
	tmpl, err := template.New("Email").Parse(getEmailTemplate())
	if err != nil {
		log.Println(err)
	}
	data := EmailPageData{
		Buttons: []Button{
			{
				Text: "Jiofi CLI",
				Link: "https://blog.athulcyriac.co/jiofi/",
			},
		},
		MD: md,
	}

	tmpl.Execute(&b, data)
	return b.String()

}

func getEmailTemplate() string {
	return `
	<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
  <head>
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
    <style type="text/css" rel="stylesheet" media="all">
      html,
      body {
        margin: 0 auto !important;
        padding: 0 !important;
        height: 100% !important;
        width: 100% !important;
        font-family:Avenir;
      }
      .body-action {
        width: 100%;
        margin: 30px auto;
        padding: 0;
        text-align: center;
      }
      .button {
        background-color: #f07d6e;
        display: inline-block;
        cursor: pointer;
        color: white;
        font-family: Avenir;
        font-size: 17px;
        padding: 16px 31px;
        text-decoration: none;
        border-radius:50px;
        text-shadow: 0px 1px 0px #2f6627;
        -webkit-text-size-adjust: none;
        mso-hide: all;
      }
      .email-logo {
        max-height: 50px;
      }
      .email-header {
        padding: 25px 0;
        text-align: center;
      }
	  }
    @media only screen and (max-width: 500px) {
      .button {
        width: 100% !important;
      }
    }
    </style>
  </head>
  <body>
    <div class="email-header">
      <img src="https://i.imgur.com/4s3lF1V.png" class="email-logo" />
	</div>
	{{if .Buttons}}
    <table class="body-action">
      <tr>
        <td align="center">
          <div class="button">
            {{ range .Buttons }}
            <a href="{{ .Link }}" target="_blank">
              {{ .Text }}
            </a>
            {{ end }}
          </div>
        </td>
      </tr>
	</table>
	{{ end }}
    <div class="markdown">
      {{.MD}}
    </div>
  </body>
</html>`
}
