package sgttopenid

import (
	"github.com/coderguang/GameEngine_go/sgtime"
)

type SByteDanceAppidCfg struct {
	Appid  string `json:"appid"`
	Secret string `json:"secret"`
}

type SByteDanceOpenid struct {
	Code   string
	Openid string
	Time   *sgtime.DateTime
}

func (data *SByteDanceOpenid) String() string {
	str := "\ncode:" + data.Code +
		"\nopenid:" + data.Openid +
		"\ndt:" + sgtime.NormalString(data.Time)
	return str
}
