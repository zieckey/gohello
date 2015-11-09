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

package main
import (
    "os/exec"
    "fmt"
    "log"
)

func main() {
    message := "hello, this is the email body"
    title := "email title ttt"
    cmd := fmt.Sprintf("echo '%v' | mail -s ‘%v’ weizili@360.cn", message, title)
    c := exec.Command(cmd)
    log.Printf(c.Run())
}