package tcsms

import (
	"github.com/coderguang/GameEngine_go/sgcfg"
	"github.com/zboyco/gosms"
)

var (
	globalSender *gosms.QSender
)

func InitTcSms(filename string) error {
	cfg := new(TcSmsCfg)
	err := sgcfg.ReadCfg(filename, cfg)
	if err != nil {
		return err
	}
	globalSender = &gosms.QSender{
		AppID:  cfg.Appid,
		AppKey: cfg.Secret,
	}
	return nil
}

func SingleSend(sign string, countryCode int, mobile string, id int, params ...string) (*gosms.QSingleResult, error) {
	res, err := globalSender.SingleSend(sign, countryCode, mobile, id, params...)
	return res, err
}

func MultiSendEach(sign string, countryCode int, moblies []string, id int, params ...string) (*gosms.QMultiResult, error) {
	res, err := globalSender.MultiSend(sign, countryCode, moblies, id, params...)
	return res, err
}
