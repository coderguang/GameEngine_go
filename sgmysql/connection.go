package sgmysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Open(user string, pwd string, host string, port string, dbname string, encode string) (db *sql.DB, err error) {
	dsn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=" + encode
	db, err = sql.Open("mysql", dsn)
	return db, err
}
