package main

import (
	"log"

	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgthread"

	"github.com/coderguang/GameEngine_go/sgcmd"

	"github.com/coderguang/GameEngine_go/sgserver"
)

func main() {

	//sgtime.InitTimeLocation("America/Los_Angeles")

	sgserver.StartLogServer("debug", "./log/", log.LstdFlags, true)

	for {
		sglog.Info("test data")
		sgthread.SleepBySecond(3)
	}

	sgcmd.StartCmdWaitInputLoop()

	sgserver.StopLogServer()
}
