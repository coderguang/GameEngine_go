package sgmail

import (
	"net/smtp"
	"strings"
)

func PlainAuth(identify string, username string, passwd string, host string) (auth smtp.Auth) {
	auth = smtp.PlainAuth(identify, username, passwd, host)
	return
}

func SendMail(host string, auth smtp.Auth, toMailList []string, subject string, fromNickName string, fromEmail string, body string) error {
	const content_type = "Content-Type: text/plain; charset=UTF-8"
	msg := []byte("To: " + strings.Join(toMailList, ",") + "\r\nFrom: " + fromNickName +
		"<" + fromEmail + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail(host, auth, fromEmail, toMailList, msg)
	return err
}
