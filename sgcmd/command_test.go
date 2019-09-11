package sgcmd

import (
	"log"
	"testing"

	"github.com/coderguang/GameEngine_go/sgserver"
)

func TestCmd(t *testing.T) {

	//StartCmdWaitInputLoop()

	sgserver.StartLogServer("debug", "../../log/", log.LstdFlags, true)

	printHelp()
	StartCmdWaitInput()

}
