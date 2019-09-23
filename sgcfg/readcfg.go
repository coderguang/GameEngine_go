package sgcfg

import (
	"encoding/json"
	"io/ioutil"

	"github.com/coderguang/GameEngine_go/sglog"
)

func ReadCfg(filename string, cfgType interface{}) error {
	config, err := ioutil.ReadFile(filename)
	if err != nil {
		sglog.Error("read ", filename, " config error,err:", err)
		return err
	}
	err = json.Unmarshal([]byte(config), &cfgType)
	if err != nil {
		sglog.Error("parse ", filename, " config error,err=", err)
		return err
	}
	return nil
}
