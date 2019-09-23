package sgwxopenid

import (
	"github.com/coderguang/GameEngine_go/sgtime"
)

type SWxOpenid struct {
	Code   string
	Openid string
	Time   *sgtime.DateTime
}

func (data *SWxOpenid) String() string {
	str := "\ncode:" + data.Code +
		"\nopenid:" + data.Openid +
		"\ndt:" + sgtime.NormalString(data.Time)
	return str
}
