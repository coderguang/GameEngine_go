package sgcfg

var serverCfgDir string

func init() {
	serverCfgDir = "./"
	SetServerCfgDir(serverCfgDir)
}

var (
	MailCfgFile  string
	MySQLCfgFile string
)

func SetServerCfgDir(dir string) {
	serverCfgDir = dir
	MailCfgFile = serverCfgDir + "mail.json"
	MySQLCfgFile = serverCfgDir + "mysql.json"
}
