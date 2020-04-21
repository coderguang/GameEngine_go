package sgmongo

import (
	"testing"

	"github.com/coderguang/GameEngine_go/sgcfg"
	"github.com/globalsign/mgo/bson"
)

type Person struct {
	Name  string
	Phone string
}

func Test_Mongo(t *testing.T) {
	sgcfg.SetServerCfgDir("./../../../globalConfig/server_config/")

	cfg, err := ReadCfg(sgcfg.MongoCfgFile)
	if err != nil {
		t.Error("init db config error")
	}

	session, err := NewSessionByMgo(cfg)
	if err != nil {
		t.Error("link to mongo db error")
	}

	defer session.Close()

	c := session.DB("sg").C("people")

	err = c.Insert(&Person{"royalchen", "test"}, &Person{"test two", "just test"}, &Person{"royalchen", "test111"})

	if err != nil {
		t.Error("insert data error")
	}

	result := Person{}
	err = c.Find(bson.M{"name": "royalchen"}).One(&result)
	if err != nil {
		t.Error("find data errot")
	}

	results := []Person{}
	err = c.Find(bson.M{"name": "royalchen"}).All(&results)
	if err != nil {
		t.Error("find all data error")
	}

	t.Log("test mongo ok")

}
