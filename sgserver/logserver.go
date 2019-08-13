package sgserver

import (
	"github.com/coderguang/GameEngine_go/sgthread"

	"github.com/coderguang/GameEngine_go/sglog"
)

func StartLogServer(level string, path string, flag int, isConsole bool) {
	logger, err := sglog.NewLogger(level, path, flag, isConsole)

	if err != nil {
		panic(err)
	}
	sglog.Swap(logger)

	go sglog.LoopLogServer()

	sglog.Info("log server init complete,path=%s,level=%s", path, level)
	sgthread.SleepBySecond(2)
}

func StopLogServer() {
	watiTime := 5
	sglog.Info("log server will stop after %ds", watiTime)
	sgthread.SleepBySecond(watiTime)
	sglog.CloseGlobalLogger()
}
