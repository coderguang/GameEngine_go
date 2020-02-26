package sgqqopenid

import (
	"github.com/coderguang/GameEngine_go/sgtime"
)

type SQQAppidCfg struct {
	Appid  string `json:"appid"`
	Secret string `json:"secret"`
}

type SQQDanceOpenid struct {
	Code   string
	Openid string
	Time   *sgtime.DateTime
}

func (data *SQQDanceOpenid) String() string {
	str := "\ncode:" + data.Code +
		"\nopenid:" + data.Openid +
		"\ndt:" + sgtime.NormalString(data.Time)
	return str
}
