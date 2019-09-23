package sgmail

import (
	"strings"

	"github.com/coderguang/GameEngine_go/sgcfg"
)

const content_type = "Content-Type: text/plain; charset=UTF-8"

type MailCfg struct {
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"pwd"`
	SMTP     string `json:"smtp"`
	Port     string `json:"port"`
	UseTLS   bool   `json:"useTLS"`
}

//单封邮件内容
type mailData struct {
	subject    string
	toMailList []string
	body       string
}

func ReadCfg(filename string) (*MailCfg, error) {
	cfg := new(MailCfg)
	err := sgcfg.ReadCfg(filename, cfg)
	return cfg, err
}

func (data *mailData) String() string {
	str := "\nsubject: " + data.subject +
		"\nreceiver: " + strings.Join(data.toMailList, ",") +
		"\nbody: " + data.body + "\n"
	return str
}
