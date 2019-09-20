package sgdef

type DefServerStatus int

const (
	_ DefServerStatus = iota
	DefServerStatusInit
	DefServerStatusRunning
	DefServerStatusStop
)
