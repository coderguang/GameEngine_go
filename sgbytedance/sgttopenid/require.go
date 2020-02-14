package sgttopenid

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgwx/sgwxdef"
)

func GetOpenIdFromWx(appid string, secret string, code string) (string, error) {
	//GET https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
	url := "https://developer.toutiao.com/api/apps/jscode2session?appid=" + appid + "&secret=" + secret + "&js_code=" + code + "&grant_type=authorization_code"
	resp, err := http.Get(url)

	if nil != err {
		sglog.Error("get bytedance openid from", url, " error,err=", err)
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		sglog.Error("get bytedance openid error,read resp body error,err:", err)
		return "", err
	}
	str := string(body)
	decoder := json.NewDecoder(bytes.NewBufferString(str))
	decoder.UseNumber()
	var result map[string]interface{}
	if err := decoder.Decode(&result); err != nil {
		sglog.Error("json parse failed,str:", str, ",err:", err)
		return "", err
	}
	if _, ok := result[sgwxdef.WX_ERROR_CODE_STR]; ok {
		sglog.Error("bytedance error openid,code=", result[sgwxdef.WX_ERROR_CODE_STR], ",errmsg=", result[sgwxdef.WX_ERROR_MSG_STR])
		errmsgV, _ := result[sgwxdef.WX_ERROR_MSG_STR].(string)
		return "", errors.New(errmsgV)
	}

	tmpOpenid := result[sgwxdef.WX_OPEN_ID_STR]
	tmpOpenidValue, ok := tmpOpenid.(string)
	if !ok {
		sglog.Error("parse tmp_openid failed,tmp_openid=", tmpOpenid)
		return "", errors.New("parse tmp_openid failed")
	}
	return tmpOpenidValue, nil
}
