package sgmysql

import "github.com/coderguang/GameEngine_go/sgcfg"

type MySQLCfg struct {
	URL  string `json:"url"`
	Port string `json:"port"`
	User string `json:"user"`
	Pwd  string `json:"pwd"`
}

func ReadCfg(filename string) (*MySQLCfg, error) {
	cfg := new(MySQLCfg)
	err := sgcfg.ReadCfg(filename, cfg)
	return cfg, err
}
