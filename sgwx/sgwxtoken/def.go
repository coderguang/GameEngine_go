package sgwxtoken

import (
	"github.com/coderguang/GameEngine_go/sgtime"
)

type SWxToken struct {
	Appid       string
	Secret      string
	RequireDt   *sgtime.DateTime
	ExpiryDt    int
	AccessToken string
}

func (data *SWxToken) String() {
	str := "\nAppid:" + data.Appid +
		"\nSecret:" + data.Secret +
		"\nRequireDt:" + sgtime.NormalString(data.RequireDt) +
		"\nExpiryDt:" + sgtime.NormalString(data.ExpiryDt) +
		"\nAccessToken:" + data.AccessToken
}
