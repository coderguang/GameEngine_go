package wordToVoice

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"

	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/gorilla/websocket"
)

func WorldToVoice(worldstr string, host string, app_id string, apikey string, apisec string, msgTemplate *SRequireParam) ([]byte, error) {

	replaceStr := strings.Replace(worldstr, "，", ",", -1)
	rawlist := strings.Split(replaceStr, ",")

	strlist := []string{}
	tmpstr := ""
	for _, v := range rawlist {
		if len(tmpstr)+len(v) > WORD_MAX_LEN {
			strlist = append(strlist, tmpstr)
			tmpstr = v
		} else {
			tmpstr += "," + v
		}
	}
	strlist = append(strlist, tmpstr)
	sglog.Info("world to voice txt split to ", len(strlist), " part")

	authStr := assembleAuthUrl(host, apikey, apisec)

	voiceBytesList := [][]byte{}
	receiveNum := 0
	needReceiveNum := len(strlist)

	msgTemplate.Common.AppID = app_id
	msgTemplate.Data.Status = 2

	for _, v := range strlist {

		worlds := []byte(v)
		msgTemplate.Data.Text = base64.StdEncoding.EncodeToString(worlds)
		js, err := json.Marshal(msgTemplate)
		if err != nil {
			sglog.Error("param to json error,", err)
			return nil, err
		}
		//sglog.Debug("js is ", string(js))

		c, resp, err := websocket.DefaultDialer.Dial(authStr, nil)
		if err != nil {
			sglog.Error("create websocket error:", err, ",status code:", resp.StatusCode)
			return nil, err
		}
		defer c.Close()

		if resp.StatusCode != 101 {
			sglog.Error("read resp status code error,status", resp.StatusCode)
			return nil, err
		}

		receiveAllFlag := make(chan bool)

		go func() {
			defer func() {
				receiveAllFlag <- true
			}()
			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					sglog.Error("read msg from server err,", err)
					break
				}
				//sglog.Debug("recv from server:", string(message))
				result := new(SResponResult)
				err = json.Unmarshal(message, &result)
				if err != nil {
					sglog.Error("parse data error")
					break
				}
				if result.Code != 0 {
					sglog.Error("transform world error,code=", result.Code)
					break
				}
				tmpBytes, err := base64.StdEncoding.DecodeString(result.Data.Audio)
				if err != nil {
					sglog.Error("transform result to base64 error,code=", err)
					break
				}
				voiceBytesList = append(voiceBytesList, tmpBytes)
				//sglog.Info("recv from server,status:", result.Data.Status)
				if result.Data.Status == 2 {
					receiveNum++
					sglog.Debug("receive part success,current recv:", receiveNum, ",need receive:", needReceiveNum)
					break
				}
			}
		}()

		err = c.WriteMessage(websocket.TextMessage, js)
		if err != nil {
			sglog.Error("push param to server,", err)
			return nil, err
		}

		<-receiveAllFlag
	}

	if receiveNum == needReceiveNum {
		voiceBytes := bytes.Join(voiceBytesList, []byte{})
		return voiceBytes, nil
	} else {
		return nil, errors.New("receive num not match need receive num")
	}
}
