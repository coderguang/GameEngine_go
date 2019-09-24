package sgcfg

var serverCfgDir string

func init() {
	serverCfgDir = "./"
	SetServerCfgDir(serverCfgDir)
}

var (
	MailCfgFile  string
	MySQLCfgFile string
	MongoCfgFile string
)

func SetServerCfgDir(dir string) {
	serverCfgDir = dir
	MailCfgFile = serverCfgDir + "mail.json"
	MySQLCfgFile = serverCfgDir + "mysql.json"
	MongoCfgFile = serverCfgDir + "mongo.json"
}

func GetServerCfgDir() string {
	return serverCfgDir
}
