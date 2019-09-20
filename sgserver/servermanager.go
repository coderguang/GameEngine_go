package sgserver

import (
	"errors"
	"strconv"

	"github.com/coderguang/GameEngine_go/sglog"
)

var serverList []Server

func init() {
	serverList = []Server{}
}

func StartServer(serverType ServerType, a ...interface{}) {
	startFlag := make(chan bool)
	defer func() {
		<-startFlag
	}()
	switch serverType {
	case ServerTypeLog:
		server := new(ServerLog)
		go server.Start(startFlag, a...)
		addServerToList(server)
	}
	return
}

func StopServer(serverType ServerType, a ...interface{}) error {
	stopFlag := make(chan bool)
	defer func() {
		<-stopFlag
	}()

	serverTypeStr := strconv.Itoa(int(serverType))

	for _, v := range serverList {
		if v.Type() == serverType {
			if v.IsStop() {
				return errors.New("this server already stop,type is " + serverTypeStr)
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
		if err := StopServer(v.Type()); err != nil {
			sglog.Error("stop server error,type=", v.Type(), ",err=", err)
		}
	}

	for _, v := range serverList {
		if v.Type() == ServerTypeLog {
			if err := StopServer(v.Type()); err != nil {
				sglog.Error("stop log server error,type=", v.Type(), ",err=", err)
			}
		}
	}
}

func addServerToList(server Server) {
	serverList = append(serverList, server)
}
