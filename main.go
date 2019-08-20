package main

import (
	"log"

	"github.com/coderguang/GameEngine_go/sgcmd"

	"github.com/coderguang/GameEngine_go/sgserver"
)

func main() {
	sgserver.StartLogServer("debug", "./log/", log.LstdFlags, true)

	sgcmd.StartCmdWaitInputLoop()

	sgserver.StopLogServer()
}
