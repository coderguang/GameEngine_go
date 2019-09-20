package sgserver

type ServerMail struct {
}

func (server *ServerMail) Start(startFlag chan bool, a ...interface{}) {

}

func (server *ServerMail) Stop(stopFlag chan bool, a ...interface{}) {

}

func (server *ServerMail) Type() ServerType {
	return ServerTypeMail
}

func (server *ServerMail) IsStop() bool {
	return true
}

func (server *ServerMail) IsRunning() bool {
	return true
}
