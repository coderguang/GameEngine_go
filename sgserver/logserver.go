package sgserver

import (
	"log"

	"github.com/coderguang/GameEngine_go/sglog"
)

type ServerLog struct {
}

func (server *ServerLog) Start(startFlag chan bool, a ...interface{}) {

	defer func() {
		startFlag <- true
	}()

	level := "debug"
	path := "./log/"
	flag := log.LstdFlags
	isConsole := true

	if len(a) >= 1 {
		if levelex, ok := a[0].(string); ok {
			level = levelex
		} else {
			log.Println("ServerLog level type error,it should be a string")
			return
		}
	}
	if len(a) >= 2 {
		if pathex, ok := a[1].(string); ok {
			path = pathex
		} else {
			log.Println("ServerLog path type error,it should be a string")
			return
		}
	}
	if len(a) >= 3 {
		if flagex, ok := a[2].(int); ok {
			flag = flagex
		} else {
			log.Println("ServerLog flag type error,it should be a int")
			return
		}
	}
	if len(a) >= 4 {
		if isConsoleEx, ok := a[3].(bool); ok {
			isConsole = isConsoleEx
		} else {

			log.Println("ServerLog isConsole type error,it should be a boolean")
			return
		}
	}

	logger, err := sglog.NewLogger(level, path, flag, isConsole)

	if err != nil {
		return
	}
	sglog.Swap(logger)

	go sglog.LoopLogServer()

	sglog.Info("log server init complete,path=", path, "level=", level)

}

func (server *ServerLog) Stop(stopFlag chan bool, a ...interface{}) {

	defer func() {
		stopFlag <- true
	}()

	sglog.Info("logger stop....")
	sglog.CloseGlobalLogger()
}

func (server *ServerLog) Type() ServerType {
	return ServerTypeLog
}

func (server *ServerLog) IsStop() bool {
	return sglog.IsStop()
}

func (server *ServerLog) IsRunning() bool {
	return true
}
