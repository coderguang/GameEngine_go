package sgmail

import (
	"errors"
	"net/smtp"
	"strings"
	"sync"

	"github.com/coderguang/GameEngine_go/sgthread"

	"github.com/coderguang/GameEngine_go/sglog"

	"github.com/coderguang/GameEngine_go/sgdef"
)

var (
	globalMailSender *mailSender
)

type mailSender struct {
	auth          smtp.Auth
	cfg           *MailCfg
	chanMailList  chan *mailData
	status        sgdef.DefServerStatus
	onceCloseFunc sync.Once
}

func NewSender(cfg *MailCfg) error {
	if globalMailSender != nil {
		return errors.New("mail sender already init")
	}
	globalMailSender = new(mailSender)
	globalMailSender.cfg = cfg
	globalMailSender.chanMailList = make(chan *mailData, 100)
	globalMailSender.status = sgdef.DefServerStatusInit
	globalMailSender.plainAuth()
	sgthread.SleepByMillSecond(500)
	go globalMailSender.loopSendMail()
	return nil
}

func CloseGlobalMailSender() {
	if globalMailSender != nil {
		globalMailSender.close()
	}
}

func IsStop() bool {
	if globalMailSender == nil {
		return true
	}
	if globalMailSender.status != sgdef.DefServerStatusRunning {
		return true
	}
	return false
}

func IsRunning() bool {
	if globalMailSender == nil {
		return false
	}
	if globalMailSender.status == sgdef.DefServerStatusRunning {
		return true
	}
	return false
}

func SendMail(subject string, tolist []string, body string) {
	data := new(mailData)
	data.subject = subject
	data.toMailList = tolist
	data.body = body
	if !IsRunning() {
		sglog.Error("send mail error,sender not running", data)
		return
	}
	globalMailSender.chanMailList <- data
}

func (sender *mailSender) close() {
	sender.status = sgdef.DefServerStatusStop
	for {
		if len(sender.chanMailList) == 0 {
			sender.onceCloseFunc.Do(func() {
				close(sender.chanMailList)
			})
			break
		}
	}
}

func (sender *mailSender) plainAuth() {
	sender.auth = smtp.PlainAuth("", sender.cfg.User, sender.cfg.Password, sender.cfg.SMTP)
}

func (sender *mailSender) loopSendMail() {
	sender.status = sgdef.DefServerStatusRunning
	for data := range sender.chanMailList {
		if err := sender.send(data); err != nil {
			sglog.Error("send mail failed,err:", err, "data", data)
		} else {
			sglog.Info("send mail ok,data", data)
		}
	}
}

func (sender *mailSender) send(data *mailData) error {
	msg := []byte("To: " + strings.Join(data.toMailList, ",") + "\r\nFrom: " + sender.cfg.Name +
		"<" + sender.cfg.User + ">\r\nSubject: " + data.subject + "\r\n" + content_type + "\r\n\r\n" + data.body)
	host := sender.cfg.SMTP + ":" + sender.cfg.Port
	err := smtp.SendMail(host, sender.auth, sender.cfg.User, data.toMailList, msg)
	return err
}
