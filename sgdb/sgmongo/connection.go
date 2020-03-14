package sgmongo

import (
	"context"
	"time"

	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/globalsign/mgo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewSessionByMgo(cfg *MongoCfg) (*mgo.Session, error) {
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

func NewSessionByOffice(cfg *MongoCfg) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(getMongoURL(cfg)))
	//connection
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		sglog.Error("connect mongo error,", getMongoURL(cfg))
		return nil, err
	}

	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		sglog.Error("ping mongo error,", err)
		return nil, err
	}
	return client, nil
}
