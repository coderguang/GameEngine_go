package sgmysql

import (
	"github.com/coderguang/GameEngine_go/sgthread"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Open(cfg *MySQLCfg) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", getDbURL(cfg))
	if err == nil {
		go func(dbex *gorm.DB) {
			for {
				dbex.DB().Ping()
				sgthread.SleepBySecond(60 * 60)
			}
		}(db)
	}
	return db, err
}

func getDbURL(cfg *MySQLCfg) string {
	dsn := cfg.User + ":" + cfg.Pwd + "@tcp(" + cfg.URL + ":" + cfg.Port + ")/" + cfg.DbName + "?charset=" + cfg.Charset + "&parseTime=True&loc=Local"
	return dsn
}
