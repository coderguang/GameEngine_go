package sgcfg

var serverCfgDir string

func init() {
	serverCfgDir = "./"
	SetServerCfgDir(serverCfgDir)
}

var (
	MailCfgFile string
)

func SetServerCfgDir(dir string) {
	serverCfgDir = dir
	MailCfgFile = serverCfgDir + "mail.json"
}
