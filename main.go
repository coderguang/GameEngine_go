package main

import (
	"log"

	"github.com/coderguang/GameEngine_go/sgmail"
	"github.com/coderguang/GameEngine_go/sgthread"

	"github.com/coderguang/GameEngine_go/sgcfg"

	"github.com/coderguang/GameEngine_go/sgcmd"

	"github.com/coderguang/GameEngine_go/sgserver"
)

//var wg sync.WaitGroup

func main() {

	sgcfg.SetServerCfgDir("./../globalConfig/server_config/")

	// defer func() {

	// 	if r := recover(); r != nil {
	// 		log.Println("recover from addServerToList,", r)
	// 	}

	// }()

	//sgtime.InitTimeLocation("America/Los_Angeles")

	//sgserver.StartLogServer("debug", "./log/", log.LstdFlags, true)

	sgserver.StartServer(sgserver.ServerTypeLog, "debug", "./../log/", log.LstdFlags, true)

	sgserver.StartServer(sgserver.ServerTypeMail)

	sgthread.SleepBySecond(2)

	tolist := []string{"royalchen@royalchen.com"}
	for i := 0; i < 1; i++ {
		sgmail.SendMail("GameEngine_go", tolist, "mail from GameEngine go")
		sgthread.SleepBySecond(10)
	}

	//sgserver.StartServer(sgserver.ServerTypeLog, "debug", "./../log/", log.LstdFlags, "wtwer")

	// for {
	// 	sglog.Info("test data")
	// 	sgthread.SleepBySecond(3)
	// }

	// for i := 0; i < 10; i++ {
	// 	//wg.Add(1)
	// 	go func() {
	// 		for {
	// 			sglog.Info("this is test stt")
	// 			sgthread.SleepByMillSecond(10)
	// 		}
	// 	}()
	// }

	// result, err := sgwhois.GetWhoisInfo("baidu.cn")
	// if err != nil {
	// 	sglog.Error("parse error")
	// }
	// sgwhois.ParseWhois(result)
	//sgwhois.ShowWhoisInfo(result)

	sgcmd.StartCmdWaitInputLoop()

}
