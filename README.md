# Mailmon
MailMon is an Mass Emailer with a Vue Frontend and a FastAPI Backend and uses SendGrid's API to send emails to groups of People.

Mailmon is made and currently used for a sending emails to all students in My class and for educational purposes. 

## Using it
Since sending emails to many people take quite some time and server's power to compute. You can use it locally
- Clone the Project
#### Python

- Install Python Dependencies
- Get a SendGrid API key from https://sendgrid.com
- Run server
```bash
$ uvicorn main:app --reload
```
Server will start at port 8000

#### Frontend
- `$ cd frontend`
- Use Yarn or Npm to install the dependencies
```bash
$ yarn install
# OR
$ npm install
```
- Run server using 
```bash
$ yarn serve
#OR
$ npm run serve
```
Dev Server will start at port 8080

![](/frontend/mailmon-fd.png)

### Contributing
Possibly I'll rewrite the backend to Go to try the speed. Yeah and I'm accepting PRs and Issues. Feel free to open One😄
