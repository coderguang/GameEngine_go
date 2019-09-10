package main

import (
	"log"

	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgthread"

	"github.com/coderguang/GameEngine_go/sgcmd"

	"github.com/coderguang/GameEngine_go/sgserver"
)

//var wg sync.WaitGroup

func main() {

	//sgtime.InitTimeLocation("America/Los_Angeles")

	sgserver.StartLogServer("debug", "./log/", log.LstdFlags, true)

	// for {
	// 	sglog.Info("test data")
	// 	sgthread.SleepBySecond(3)
	// }

	for i := 0; i < 10; i++ {
		//wg.Add(1)
		go func() {
			for {
				sglog.Info("this is test stt")
				sgthread.SleepByMillSecond(10)
			}
		}()
	}

	sgcmd.StartCmdWaitInputLoop()

	sgserver.StopLogServer()
}
