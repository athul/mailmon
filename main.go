package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	md "github.com/gomarkdown/markdown"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	log "github.com/sirupsen/logrus"
)

// Student struct handles the JSON
type Student struct {
	Name  string `json:"Name"`
	Roll  int    `json:"rn"`
	Email string `json:"Email"`
}
type emresp struct {
	Email string
	code  int
}

func getStudents() []Student {
	var student []Student
	allstudents, _ := ioutil.ReadFile("athul-ems.json")

	json.Unmarshal(allstudents, &student)

	return student
}

func sendEmails(c *gin.Context) {
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	var mailresp = make(map[string]int)
	stds := getStudents()
	subject := c.PostForm("subject")
	content := c.PostForm("content")
	from := mail.NewEmail("Athul Cyriac Ajay", "athul8720@gmail.com")
	htmlContent := string(md.ToHTML([]byte(content), nil, nil))
	for i := 0; i < len(stds); i++ {
		log.Infof("Name:\t%s", stds[i].Name)
		to := mail.NewEmail(stds[i].Name, stds[i].Email)
		fmt.Println(to.Address)
		message := mail.NewSingleEmail(from, subject, to, "This", htmlContent)
		resp, err := client.Send(message)
		if err != nil {
			log.Errorln(err)
		} else {
			log.Infoln(resp.StatusCode, resp.Body)
			mailresp[stds[i].Email] = resp.StatusCode
			log.Println(mailresp)
		}
	}
	fmt.Println(mailresp)
	c.String(200, "Emails Send")
}
func getEmail(c *gin.Context) {
	stds := getStudents()
	for i := 0; i < len(stds); i++ {
		c.String(200, stds[i].Name+"\n"+stds[i].Email)
	}
	// c.String(200, "Emails Printed in Log")
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
	// return "", fmt.Errorf("Could not Parse Markdown since it's empty")
}
func main() {
	// stds := getStudents()
	app := gin.Default()

	// app.GET("/", func(c *gin.Context) {
	// 	c.String(200, "Hello, World!")
	// })
	app.POST("/md", renderMD)
	app.GET("/em", getEmail)
	app.POST("/send", sendEmails)
	app.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))
	app.Run(":8080")
}
