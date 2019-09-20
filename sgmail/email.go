package sgmail

import (
	"errors"
	"net/smtp"
	"strings"

	"github.com/coderguang/GameEngine_go/sgdef"
)

var (
	globalMailSender *mailSender
)

type mailSender struct {
	auth         smtp.Auth
	cfg          *mailCfg
	chanMailList chan *mailData
	status       sgdef.DefServerStatus
}

func NewSender(cfg *mailCfg) error {
	if globalMailSender != nil {
		return errors.New("mail sender already init")
	}
	globalMailSender = new(mailSender)
	globalMailSender.cfg = cfg
	globalMailSender.chanMailList = make(chan *mailData, 100)
	globalMailSender.status = sgdef.DefServerStatusInit

	return nil
}

func (sender *mailSender) PlainAuth() {
	sender.auth = smtp.PlainAuth("", sender.cfg.User, sender.cfg.Password, sender.cfg.SMTP)
}

func (sender *mailSender) loopSendMail() {
	sender.status = sgdef.DefServerStatusRunning
	for {
		data := <-sender.chanMailList

	}
}

func SendMail(host string, auth *smtp.Auth, toMailList []string, subject string, fromNickName string, fromEmail string, body string) error {
	const content_type = "Content-Type: text/plain; charset=UTF-8"
	msg := []byte("To: " + strings.Join(toMailList, ",") + "\r\nFrom: " + fromNickName +
		"<" + fromEmail + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail(host, *auth, fromEmail, toMailList, msg)
	return err
}
