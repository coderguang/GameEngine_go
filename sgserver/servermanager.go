package sgserver

import (
	"errors"
	"log"
	"strconv"

	"github.com/coderguang/GameEngine_go/sgthread"

	"github.com/coderguang/GameEngine_go/sglog"
)

var serverList []Server

func init() {
	serverList = []Server{}
}

func getServerInstance(serverType ServerType) Server {
	switch serverType {
	case ServerTypeLog:
		return new(ServerLog)
	case ServerTypeMail:
		return new(ServerMail)
	}
	return nil
}

func StartServer(serverType ServerType, a ...interface{}) {

	for _, v := range serverList {
		if v.Type() == serverType && v.IsRunning() {
			sglog.Error("server had already running,type is", v.Type())
			return
		}
	}

	startFlag := make(chan bool)
	server := getServerInstance(serverType)
	if server == nil {
		sglog.Error("unknow server type,type is ", serverType)
		return
	}
	go server.Start(startFlag, a...)
	addServerToList(startFlag, server)
	return
}

func StopServer(serverType ServerType, a ...interface{}) error {
	stopFlag := make(chan bool)
	defer func() {
		<-stopFlag
		sglog.Info("server stop ok,type=", serverType)
	}()

	serverTypeStr := strconv.Itoa(int(serverType))

	for _, v := range serverList {
		if v.Type() == serverType {
			if v.IsStop() || !v.IsRunning() {
				return errors.New("this server already stop or not running,type is " + serverTypeStr)
			} else {
				go v.Stop(stopFlag, a...)
				return nil
			}
		}
	}
	return errors.New("no this server,type is " + serverTypeStr)
}

func StopAllServer() {
	for _, v := range serverList {
		if v.Type() == ServerTypeLog {
			continue
		}
		if v.IsStop() || !v.IsRunning() {
			sglog.Error("stop server error not running or already stop,type=", v.Type())
			continue
		}
		if err := StopServer(v.Type()); err != nil {
			sglog.Error("stop server error,type=", v.Type(), ",err=", err)
		}
	}

	sgthread.SleepBySecond(2)

	for _, v := range serverList {
		if v.Type() == ServerTypeLog {
			if err := StopServer(v.Type()); err != nil {
				sglog.Error("stop log server error,type=", v.Type(), ",err=", err)
			}
		}
	}
}

func addServerToList(startFlag chan bool, server Server) {

	startResult := <-startFlag
	if startResult {
		serverList = append(serverList, server)
	} else {
		if server.Type() != ServerTypeLog {
			sglog.Error("server start error,please check,type=", server.Type())
		} else {
			log.Println("log server start error,please check")
		}
		sgthread.SleepBySecond(2)
		panic(errors.New("server start error type is " + strconv.Itoa(int(server.Type()))))
	}
}
