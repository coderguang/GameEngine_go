package sgredis

import "github.com/coderguang/GameEngine_go/sgcfg"

type RedisCfg struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Auth string `json:"auth"`
}

func ReadCfg(filename string) (*RedisCfg, error) {
	cfg := new(RedisCfg)
	err := sgcfg.ReadCfg(filename, cfg)
	return cfg, err
}
