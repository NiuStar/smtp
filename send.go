package smtp

import (
	"fmt"
	"strings"
)

func SendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	fmt.Println(send_to)
	err := SendMail(host, auth, user, send_to, msg)
	return err
}

type MailAuth struct {
	user string
	password string
	host string
	auth Auth
}

func NewMailAuth(user, password, host string) *MailAuth {
	hp := strings.Split(host, ":")
	auth := PlainAuth("", user, password, hp[0])
	return &MailAuth{user:user,password:password,host:host,auth:auth}
}

func (this *MailAuth)SendMail(to,subject,body,mailtype string) error {

	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " +  this.user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	fmt.Println(send_to)
	err := SendMail(this.host, this.auth, this.user, send_to, msg)
	return err
}