package sgserver

import (
	"time"

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
}

func StopLogServer() {
	watiTime := 5
	sglog.Info("log server will stop after %ds", watiTime)
	time.Sleep(time.Duration(watiTime) * time.Second)
	sglog.CloseGlobalLogger()
}
