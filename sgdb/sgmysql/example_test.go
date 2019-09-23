package sgmysql

import (
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/coderguang/GameEngine_go/sgcfg"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Test_GORM(t *testing.T) {

	sgcfg.SetServerCfgDir("./../../../globalConfig/server_config/")

	cfg, err := ReadCfg(sgcfg.MySQLCfgFile)
	if err != nil {
		t.Error("init db config error")
	}

	url := cfg.User + ":" + cfg.Pwd + "@/test?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open("mysql", url)
	if err != nil {
		t.Error("open db error,err:", err)
	}

	defer db.Close()

	t.Log("test ok)
}
