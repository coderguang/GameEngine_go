package sgserver

import (
	"encoding/json"
	"io/ioutil"

	"github.com/coderguang/GameEngine_go/sgcfg"

	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgmail"
)

type ServerMail struct {
}

func (server *ServerMail) InitCfg() (*sgmail.MailCfg, error) {
	config, err := ioutil.ReadFile(sgcfg.MailCfgFile)
	if err != nil {
		sglog.Error("read config error,err:", err)
		return nil, err
	}
	cfg := new(sgmail.MailCfg)
	p := &cfg
	err = json.Unmarshal([]byte(config), p)
	if err != nil {
		sglog.Error("parse config error,err=", err)
		return nil, err
	}
	return cfg, nil
}

func (server *ServerMail) Start(startFlag chan bool, a ...interface{}) {
	startResult := false
	defer func() {
		startFlag <- startResult
	}()

	cfg, err := server.InitCfg()
	if err != nil {
		return
	}
	sgmail.NewSender(cfg)
	sglog.Info("mail server init complete")
	startResult = true
	return
}

func (server *ServerMail) Stop(stopFlag chan bool, a ...interface{}) {

	stopResult := false
	defer func() {
		stopFlag <- stopResult
	}()
	sglog.Info("mail stop....")
	sgmail.CloseGlobalMailSender()
	stopResult = true

}

func (server *ServerMail) Type() ServerType {
	return ServerTypeMail
}

func (server *ServerMail) IsStop() bool {
	return sgmail.IsStop()
}

func (server *ServerMail) IsRunning() bool {
	return sgmail.IsRunning()
}
