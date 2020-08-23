package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	md "github.com/gomarkdown/markdown"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mail.v2"
)

// Student struct handles the JSON
type Student struct {
	Name  string `json:"Name"`
	Roll  int    `json:"rn"`
	Email string `json:"Email"`
}
type emresp struct {
	Email string `json:"email"`
	Code  int    `json:"code"`
}

func getStudents() []Student {
	var student []Student
	allstudents, _ := ioutil.ReadFile("./data/athul.json")

	json.Unmarshal(allstudents, &student)

	return student
}

func sendEmails(c *gin.Context) {
	d := mail.NewDialer("smtp.gmail.com", 587, os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	d.StartTLSPolicy = mail.MandatoryStartTLS
	s, err := d.Dial()
	if err != nil {
		log.Errorln(err)
	}
	var mailresp = make(map[string]int)
	var emailresp []emresp
	stds := getStudents()
	subject := c.PostForm("subject")
	content := c.PostForm("content")
	htmlContent := string(md.ToHTML([]byte(content), nil, nil))
	m := mail.NewMessage()
	for _, r := range stds {
		m.SetHeader("Subject", subject)
		m.SetBody("text/html", htmlContent)
		m.SetAddressHeader("From", os.Getenv("USERNAME"), "TinkerHub CEK")
		m.SetAddressHeader("To", r.Email, r.Name)
		err := mail.Send(s, m)
		if err != nil {
			mailresp[r.Email] = 400
			log.Errorln("Could not send email to %q: %v", r.Email, err)
		} else {
			mailresp[r.Email] = 200
		}
		m.Reset()
	}
	log.Println(mailresp)
	for k, v := range mailresp {
		emailresp = append(emailresp, emresp{
			Email: k,
			Code:  v,
		})
	}
	c.JSON(200, gin.H{
		"mailresp": emailresp,
		"md":       htmlContent,
		"subject":  subject,
	})
}
func renderMD(c *gin.Context) {
	mdr := c.PostForm("mdb")
	if mdr != "" {
		renderedMD := md.ToHTML([]byte(mdr), nil, nil)
		log.Info("MD parsing Successfull")
		c.String(200, string(renderedMD))
	} else {
		log.Info("MD Parsing Failed - Empty String")
	}
}
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := gin.Default()
	app.POST("/md", renderMD)
	app.POST("/send", sendEmails)
	app.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))
	app.Run(":8080")
}
