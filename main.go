package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	md "github.com/gomarkdown/markdown"
	"github.com/joho/godotenv"
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

	if err := json.Unmarshal(allstudents, &student); err != nil {
		log.Printf("Unmarhsall Error due to %v", err)
	}

	return student
}

func sendEmails(c *gin.Context) {
	start := time.Now()
	var wg sync.WaitGroup
	d := mail.NewDialer("smtp.gmail.com", 587, os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	d.StartTLSPolicy = mail.MandatoryStartTLS
	s, err := d.Dial()
	if err != nil {
		log.Println(err)
	}
	var mailresp = make(map[string]int)
	var emailresp []emresp
	stds := getStudents()
	subject := c.PostForm("subject")
	content := c.PostForm("content")
	htmlContent := string(md.ToHTML([]byte(content), nil, nil))
	wg.Add(1)
	go func() {
		defer wg.Done()
		m := mail.NewMessage()
		for i, r := range stds {
			log.Println(i + 1)
			m.SetHeader("Subject", subject)
			m.SetBody("text/html", htmlContent)
			m.SetAddressHeader("From", os.Getenv("USERNAME"), "TinkerHub CEK")
			m.SetAddressHeader("To", r.Email, r.Name)
			if err := mail.Send(s, m); err != nil {
				mailresp[r.Email] = 400
				log.Printf("Could not send email to %q: %v", r.Email, err)
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
	}()
	wg.Wait()
	c.JSON(200, gin.H{
		"mailresp": emailresp,
		"md":       htmlContent,
		"subject":  subject,
		"elapsed":  time.Since(start).String(),
	})
}
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

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	app := gin.Default()
	app.POST("/md", renderMD)
	app.POST("/send", sendEmails)
	app.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))
	app.Run(":8080")
}
