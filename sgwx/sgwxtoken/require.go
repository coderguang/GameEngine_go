package sgwxtoken

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/coderguang/GameEngine_go/sgwx/sgwxdef"

	"github.com/coderguang/GameEngine_go/sglog"
)

var (
	globalAccessTokenUrl []string
)

func init() {
	globalAccessTokenUrl = []string{"api.weixin.qq.com", "api2.weixin.qq.com", "sh.api.weixin.qq.com", "sz.api.weixin.qq.com", "hk.api.weixin.qq.com"}
}

func GetAccessTokenFromWx(appid string, secret string) (string, int64, error) {

	params := "/cgi-bin/token?grant_type=client_credential&appid=" + appid + "&secret=" + secret

	for _, v := range globalAccessTokenUrl {
		url := "https://" + v + params
		resp, err := http.Get(url)
		if nil != err {
			sglog.Error("get wx access token from ", url, " error,err=", err)
		} else {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if nil != err {
				continue
			}
			str := string(body)
			decoder := json.NewDecoder(bytes.NewBufferString(str))
			decoder.UseNumber()
			var result map[string]interface{}
			if err := decoder.Decode(&result); err != nil {
				sglog.Error("json parse failed,str=", str, ",err=", err)
				continue
			}

			if _, ok := result[sgwxdef.WX_ERROR_CODE_STR]; ok {
				sglog.Error("error token,code=", result[sgwxdef.WX_ERROR_CODE_STR], ",errmsg:", result[sgwxdef.WX_ERROR_MSG_STR])
				continue
			}

			access_token := result[sgwxdef.WX_ACCESS_TOKEN_STR]
			access_token_value, ok := access_token.(string)
			if !ok {
				sglog.Error("parse access_token failed,access_token=", access_token)
				continue
			}

			time_num := result[sgwxdef.WX_ACCESS_TOKEN_EXPIRY_DT_STR]
			time_num_value, err := time_num.(json.Number).Int64()
			if err != nil {
				sglog.Error("parase time_num failed,time_num=", time_num)
				continue
			}
			return access_token_value, time_num_value, nil
		}
	}
	return "", 0, errors.New("can't get access token from all wx server")
}
