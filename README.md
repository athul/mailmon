# Mailmon

MailMon is an Mass Emailer with a Vue Frontend and a Go Backend(Gin) and uses Gmail's SMTP(with ) to send emails to groups of People.

Mailmon is made and currently used for a sending emails to all students in My class and for educational purposes.

## Using it

Since sending emails to many people take quite some time and server's power to compute. It's better that you use it locally.

- Clone the Project

### Frontend

- `$ cd frontend`

- Use Yarn or Npm to install the dependencies

```bash
$ yarn install
# OR
$ npm install
```

- Build the Files

```bash
$ yarn build
#OR
$ npm run build
```

This command will generate static(HTML,CSS and JS) files which will be used by the server to show the webpage like below

### Server

- Install Go and do the setup
- run `go mod download` to download the dependencies
- Run the server

```bash
$ go run main.go # to run the server
# OR
$ go build main.go # generates static binary
$ ./main # Execute the binary
```

A server will start at port [8080](http://localhost:8080). The `/` endpoint will show our frontend.

### SMTP

- Create a `.env` file with details like this

```text
USERNAME=<your_email>
PASSWORD=<your_email_password>
```

![Frontend Image](/frontend/mailmon-fd.png)

### Contributing

Possibly I'll rewrite the backend to Go to try the speed. Yeah and I'm accepting PRs and Issues. Feel free to open One😄
