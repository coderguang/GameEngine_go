package sgserver

type ServerType int

const (
	ServerTypeLog  ServerType = 1
	ServerTypeMail ServerType = 2
)

type Server interface {
	Start(startFlag chan bool, a ...interface{})
	Stop(stopFlag chan bool, a ...interface{})
	Type() ServerType
	IsStop() bool
	IsRunning() bool
}
