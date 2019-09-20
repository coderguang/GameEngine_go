package sgmail

import (
	"strings"
)

const content_type = "Content-Type: text/plain; charset=UTF-8"

type MailCfg struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"pwd"`
	SMTP     string `json:"smtp"`
	Port     string `json:"port"`
}

//单封邮件内容
type mailData struct {
	subject    string
	toMailList []string
	body       string
}

func (data *mailData) String() string {
	str := "\nsubject: " + data.subject +
		"\nreceiver: " + strings.Join(data.toMailList, ",") +
		"\nbody: " + data.body + "\n"
	return str
}
