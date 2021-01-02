package main

import (
	"bytes"
	"log"
	"text/template"
)

//EmailPageData is the Data set in the Email
type EmailPageData struct {
	Rec     []Student
	Buttons []Button
	MD      string
}

//Button is the type for a button in the Email
type Button struct {
	Text string
	Link string
}
type renderCtx struct {
	Recipient string
	ImageURL  string
	Buttons   []Button
	MD        string
	Footer    string
}

func (e *EmailPageData) newRenderContext(counter int) renderCtx {
	ctx := renderCtx{
		Recipient: e.Rec[counter].Name,
		Buttons:   e.Buttons,
		MD:        e.MD,
	}
	return ctx
}
func (e *renderCtx) renderEmails() string {
	var b bytes.Buffer
	e.MD = e.renderEmailTemplate()
	tmpl, err := template.New("Email").Parse(emailTemp())
	if err != nil {
		log.Println(err)
	}
	tmpl.Execute(&b, e)
	return b.String()

}
func emailTemp() string {
	return `
  <!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" style="font-family: 'Helvetica Neue', Helvetica, Arial, sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
<head>
<meta name="viewport" content="width=device-width" />
<meta http-equiv="Content-Type" content="text/html; charset=UTF-8" />
<title>Actionable emails e.g. reset password</title>


<style type="text/css">
img {
max-width: 100%;
}
body {
-webkit-font-smoothing: antialiased; -webkit-text-size-adjust: none; width: 100% !important; height: 100%; line-height: 1.6em;
}
body {
background-color: #f6f6f6;
}
button{
  border-radius: 5px;
  border-color: #348eda;
  background-color: #348eda;
  border-style: solid;
  border-width: 10px 20px;
}
button > a{
  font-family: Avenir,Arial,sans-serif;
  box-sizing: border-box;
  font-size: 14px;
  color: #FFF;
  text-decoration: none;
  line-height: 2em;
  font-weight: bold;
  text-align: center; cursor: pointer;
  display: inline-block;
  border-radius: 5px;
  background-color: #348eda;
  margin: 0; 
  border-color: #348eda;
  border-style: solid;
  border-width: 10px 20px;
}
img{
  display: block;
  margin-left: auto;
  margin-right: auto;
  width: 40%;
}
@media only screen and (max-width: 640px) {
  body {
    padding: 0 !important;
  }
  h1 {
    font-weight: 800 !important; margin: 20px 0 5px !important;
  }
  h2 {
    font-weight: 800 !important; margin: 20px 0 5px !important;
  }
  h3 {
    font-weight: 800 !important; margin: 20px 0 5px !important;
  }
  h4 {
    font-weight: 800 !important; margin: 20px 0 5px !important;
  }
  h1 {
    font-size: 22px !important;
  }
  h2 {
    font-size: 18px !important;
  }
  h3 {
    font-size: 16px !important;
  }
  .container {
    padding: 0 !important; width: 100% !important;
  }
  .content {
    padding: 0 !important;
  }
  .content-wrap {
    padding: 10px !important;
  }
  .invoice {
    width: 100% !important;
  }
  img{
    width: 100% !important;
  }
  
}
</style>
</head>

<body itemscope itemtype="http://schema.org/EmailMessage" style="font-family: Avenir,Arial,sans-serif; box-sizing: border-box; font-size: 14px; -webkit-font-smoothing: antialiased; -webkit-text-size-adjust: none; width: 100% !important; height: 100%; line-height: 1.6em; background-color: #f6f6f6; margin: 0;" bgcolor="#f6f6f6">

<table class="body-wrap" style="font-family: Avenir,Arial,sans-serif; box-sizing: border-box; font-size: 14px; width: 100%; background-color: #f6f6f6; margin: 0;" bgcolor="#f6f6f6"><tr style="font-family: Avenir,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;"><td style="font-family: Avenir,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0;" valign="top"></td>
		<td class="container" width="600" style="font-family: Avenir,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; display: block !important; max-width: 600px !important; clear: both !important; margin: 0 auto;" valign="top">
			<div class="content" style="font-family: Avenir,Arial,sans-serif; box-sizing: border-box; font-size: 14px; max-width: 600px; display: block; margin: 0 auto; padding: 20px;">
				<table class="main" width="100%" cellpadding="0" cellspacing="0" itemprop="action" itemscope itemtype="http://schema.org/ConfirmAction" style="font-family: Avenir,Arial,sans-serif; box-sizing: border-box; font-size: 14px; border-radius: 3px; background-color: #fff; margin: 0; border: 1px solid #e9e9e9;" bgcolor="#fff"><tr style="font-family: Avenir,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;"><td class="content-wrap" style="font-family: Avenir,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 20px;" valign="top">
              <meta itemprop="name" content="Confirm Email" style="font-family: Avenir,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;" />
              <table width="100%" cellpadding="0" cellspacing="0" style="font-family: Avenir,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
              <img src="{{ .ImageURL}}" align="center" />
              <tr style="font-family: Avenir,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
              <td class="content-block" style="font-family: Avenir,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0; padding: 0 0 20px;" valign="top">
										<div class="markdown">{{ .MD}}</div>
							</td>
              </tr>
              </table>
          </table>
          <div class="footer" style="font-family: Avenir,Arial,sans-serif; box-sizing: border-box; font-size: 14px; width: 100%; clear: both; color: #999; margin: 0; padding: 20px;">
          <table width="100%" style="font-family: Avenir,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
          <tr style="font-family: Avenir,Arial,sans-serif; box-sizing: border-box; font-size: 14px; margin: 0;">
          <td class="aligncenter content-block" style="font-family: Avenir,Arial,sans-serif; box-sizing: border-box; font-size: 12px; vertical-align: top; color: #999; text-align: center; margin: 0; padding: 0 0 20px;" align="center" valign="top">
          {{.Footer}}
          </td>
						</tr></table></div></div>
		</td>
		<td style="font-family: Avenir,Arial,sans-serif; box-sizing: border-box; font-size: 14px; vertical-align: top; margin: 0;" valign="top"></td>
	</tr></table></body>
</html>
  `
}
