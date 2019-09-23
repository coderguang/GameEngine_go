package sgmongo

import (
	"github.com/globalsign/mgo"
)

func NewSesssion(cfg *MongoCfg) (*mgo.Session, error) {
	dsn := getMongoURL(cfg)
	session, err := mgo.Dial(dsn)
	return session, err
}

func getMongoURL(cfg *MongoCfg) string {
	// [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	dsn := "mongodb://"
	if cfg.User != "" {
		dsn += cfg.User + ":" + cfg.Pwd + "@"
	}
	dsn += cfg.Host + ":" + cfg.Port
	return dsn
}
