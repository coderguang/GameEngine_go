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
	RedisCfgFile string
)

func SetServerCfgDir(dir string) {
	serverCfgDir = dir
	MailCfgFile = serverCfgDir + "mail.json"
	MySQLCfgFile = serverCfgDir + "mysql.json"
	MongoCfgFile = serverCfgDir + "mongo.json"
	RedisCfgFile = serverCfgDir + "redis.json"
}

func GetServerCfgDir() string {
	return serverCfgDir
}
