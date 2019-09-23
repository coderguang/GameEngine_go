package sgwxtoken

import (
	"strconv"

	"github.com/coderguang/GameEngine_go/sgtime"
)

type SWxToken struct {
	Appid       string
	Secret      string
	RequireDt   *sgtime.DateTime
	ExpiryDt    int
	AccessToken string
}

func (data *SWxToken) String() string {
	str := "\nAppid:" + data.Appid +
		"\nSecret:" + data.Secret +
		"\nRequireDt:" + sgtime.NormalString(data.RequireDt) +
		"\nExpiryDt:" + strconv.Itoa(data.ExpiryDt) +
		"\nAccessToken:" + data.AccessToken
	return str
}
