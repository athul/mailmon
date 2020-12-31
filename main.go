package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	md "github.com/gomarkdown/markdown"
	"github.com/joho/godotenv"
	"gopkg.in/mail.v2"
)

// Student struct handles the
// parsing of JSON from the data files
type Student struct {
	Name  string `json:"Name"`
	Roll  int    `json:"rn"`
	Email string `json:"Email"`
}

// Response once the Emails are sent
type emresp struct {
	Email string `json:"email"`
	Code  int    `json:"code"`
}

// Parses the JSON Files and returns a Student Type
func getStudents() []Student {
	var student []Student
	allstudents, _ := ioutil.ReadFile("./data/athul.json")

	if err := json.Unmarshal(allstudents, &student); err != nil {
		log.Printf("Unmarhsall Error due to %v", err)
	}
	return student
}

// Sends all the Emails
func sendEmails(c *gin.Context) {
	var (
		start = time.Now()
		wg    sync.WaitGroup
		mutex = &sync.Mutex{}
		stds  = getStudents()
	)
	// d := mail.NewDialer("smtp.yandex.com", 465, os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
	// d.StartTLSPolicy = mail.MandatoryStartTLS
	d := mail.NewDialer("localhost", 1025, "athul", "athul")
	s, err := d.Dial()
	if err != nil {
		log.Println(err)
	}
	// Map for making a Json response of Emails with a status code and Email

	log.Println(len(stds), stds)
	subject := c.PostForm("subject")
	content := c.PostForm("content")
	wg.Add(len(stds))
	var mdhtml template.HTML
	// Send Email Asynchronously using a goroutine
	for i, r := range stds {
		ccontent := renderEmailTemplate(content, r)
		htmlContent := md.ToHTML([]byte(ccontent), nil, nil)
		mdhtml = template.HTML(htmlContent)
		go func(i int, r Student) {
			defer wg.Done()
			m := mail.NewMessage()

			log.Println(i + 1)
			m.SetHeader("Subject", subject)
			m.SetBody("text/html", renderEmails(mdhtml))
			m.SetAddressHeader("From", os.Getenv("USERNAME"), "MailMon")
			m.SetAddressHeader("To", r.Email, r.Name)
			mutex.Lock()
			if err := mail.Send(s, m); err != nil {
				log.Printf("Could not send email to %q: %v", r.Email, err)
			}
			mutex.Unlock()
			m.Reset()

		}(i, r)
	}
	log.Printf("Goroutines = %d", runtime.NumGoroutine())
	wg.Wait()
	c.JSON(200, gin.H{
		"md":      renderEmails(mdhtml),
		"subject": subject,
		"elapsed": time.Since(start).String(), // Displays execution time
	})
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
func init() {
	// Loads the Env vars
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	app := gin.Default()
	app.Use(cors.Default())
	app.POST("/md", renderMD)
	app.POST("/send", sendEmails)
	app.Use(static.Serve("/", static.LocalFile("./frontend/dist", false)))
	app.Run(":8080")
}
