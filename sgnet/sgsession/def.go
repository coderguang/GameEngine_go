package sgsession

type ESessionType int

const (
	Client ESessionType = 1 + iota //作为客户端时的连接
	Server                         //作为服务器时的连接
)
