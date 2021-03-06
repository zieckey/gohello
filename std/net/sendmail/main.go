// an example to send out email without authentication with mail server.
/*package main
import (
    "fmt"
    "log"
    "net/mail"
    "net/smtp"
)

func main() {
    // the basics
    from := mail.Address{"Joe X", "zieckey@163.com"}
    to := mail.Address{"Bob Y", "zieckey@163.com"}
    body := "this is the body line1.\nthis is the body line2.\nthis is the body line3.\n"
    subject := "this is the subject line"

    // setup the remote smtpserver
    smtpserver := "mail.google.com:25"

    // setup a map for the headers
    header := make(map[string]string)
    header["From"] = from.String()
    header["To"] = to.String()
    header["Subject"] = subject

    // setup the message
    message := ""
    for k, v := range header {
        message += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    message += "\r\n" + body

    // create the smtp connection
    c, err := smtp.Dial(smtpserver)
    if err != nil {
        log.Panic(err)
    }

    // To && From
    if err = c.Mail(from.Address); err != nil {
        log.Panic(err)
    }
    if err = c.Rcpt(to.Address); err != nil {
        log.Panic(err)
    }

    // Data
    w, err := c.Data()
    if err != nil {
        log.Panic(err)
    }
    _, err = w.Write([]byte(message))
    if err != nil {
        log.Panic(err)
    }
    err = w.Close()
    if err != nil {
        log.Panic(err)
    }
    c.Quit()
}*/

//package main
//import (
//    "os/exec"
//    "fmt"
//    "log"
//)
//
//func main() {
//    message := "hello, this is the email body"
//    title := "email title ttt"
//    cmd := fmt.Sprintf("echo '%v' | mail -s ‘%v’ weizili@360.cn", message, title)
//    c := exec.Command(cmd)
//    log.Printf(c.Run())
//}


package main

import (
    "fmt"
    "github.com/mattbaird/gochimp"
)

func main() {
    //apiKey := os.Getenv("MANDRILL_KEY")
    apiKey := "111111111-1111-1111-1111-111111111"
    mandrillApi, err := gochimp.NewMandrill(apiKey)

    if err != nil {
        fmt.Println("Error instantiating client")
    }

    templateName := "welcome email"
    contentVar := gochimp.Var{"main", "<h1>Welcome aboard</h1>"}
    content := []gochimp.Var{contentVar}

    renderedTemplate, err := mandrillApi.TemplateRender(templateName, content, nil)

    if err != nil {
        fmt.Println("Error rendering template:", err.Error())
        return
    }

    recipients := []gochimp.Recipient{
        gochimp.Recipient{Email: "zieckey@163.com"},
    }

    message := gochimp.Message{
        Html:      renderedTemplate,
        Subject:   "Welcome aboard!",
        FromEmail: "bossman@example.com",
        FromName:  "Boss Man",
        To:        recipients,
    }

    _, err = mandrillApi.MessageSend(message, false)

    if err != nil {
        fmt.Println("Error sending message:", err.Error())
    }
}