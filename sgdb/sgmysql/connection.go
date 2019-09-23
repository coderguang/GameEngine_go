package sgmysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func Open(cfg *MySQLCfg) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", getDbURL(cfg))
	return db, err
}

func getDbURL(cfg *MySQLCfg) string {
	dsn := cfg.User + ":" + cfg.Pwd + "@tcp(" + cfg.URL + ":" + cfg.Port + ")/" + cfg.Port + "?charset=" + cfg.Charset
	return dsn
}
