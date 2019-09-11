package sgserver

type Server interface {
	Start(a ...interface{}) bool
	Stop() bool
}
