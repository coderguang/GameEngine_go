package sgredis

import (
	"strconv"
	"testing"

	"github.com/coderguang/GameEngine_go/sgcfg"
)

func Test_Redis(t *testing.T) {

	sgcfg.SetServerCfgDir("./../../../globalConfig/server_config/")

	cfg, err := ReadCfg(sgcfg.RedisCfgFile)
	if err != nil {
		t.Error("init redis cfg error")
		return
	}

	conn, err := NewConnection(cfg.Host, strconv.Itoa(cfg.Port))
	if err != nil {
		t.Error("connection error,", err)
		return
	}

	_, err = conn.Do("Auth", cfg.Auth)
	if err != nil {
		t.Error("auth error,", err)
		return
	}

	t.Log("test ok")
}
