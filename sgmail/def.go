package sgmail

type mailCfg struct {
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
