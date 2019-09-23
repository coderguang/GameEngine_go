package sgmysql

import (
	"testing"

	"github.com/jinzhu/gorm"

	"github.com/coderguang/GameEngine_go/sgcfg"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Animal struct {
	ID   int64
	Name string `gorm:"default:'galeone'"`
	Age  int64
}

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

	an := Animal{Age: 99, Name: "tt"}

	if !db.HasTable(Animal{}) {
		db.CreateTable(Animal{})
		db.Create(&an)
	}

	ans := []Animal{}

	db.Where("age<?", 40).Find(&ans)

	t.Log("test ok")
}
