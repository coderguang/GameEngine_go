package sgmongo

import (
	"github.com/coderguang/GameEngine_go/sgcfg"
)

type MongoCfg struct {
	Host   string `json:"host"`
	Port   string `json:"port"`
	User   string `json:"user"`
	Pwd    string `json:"pwd"`
	DbName string `json:"dbName"`
}

func ReadCfg(filename string) (*MongoCfg, error) {
	cfg := new(MongoCfg)
	err := sgcfg.ReadCfg(filename, cfg)
	return cfg, err
}
