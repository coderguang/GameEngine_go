package sgwhois

import (
	"bytes"
	"errors"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/coderguang/GameEngine_go/sgtime"

	"github.com/coderguang/GameEngine_go/sglog"
)

type Whois struct {
	Raw            string //域名查询返回的整个字符串信息
	Name           string
	Zone           string
	Domain         string
	CreateDtStr    string
	UpdateDtStr    string
	ExpiryDtStr    string
	OldExpiryDtStr string
	IsRegist       int //  1：已注册   0：未注册  -1：查询失败 2:该域名为后缀域名，无法注册
	Name_length    int
	CreateDt       *sgtime.DateTime
	UpdateDt       *sgtime.DateTime
	ExpiryDt       *sgtime.DateTime
}

func (data *Whois) IsEqual(other *Whois) bool {
	if sgtime.GetTotalSecond(data.ExpiryDt) == sgtime.GetTotalSecond(other.ExpiryDt) {
		return true
	} else {
		return false
	}
}

func ShowWhoisInfo(data *Whois) {
	sglog.Info("==================start show=============")
	sglog.Debug("name:", data.Domain, ",status:", data.IsRegist, "\n")
	sglog.Debug("create dt:", sgtime.NormalString(data.CreateDt))
	sglog.Debug("update dt", sgtime.NormalString(data.UpdateDt))
	sglog.Debug("expiry dt:", sgtime.NormalString(data.ExpiryDt))
	sglog.Debug("raw create dt:", data.CreateDtStr)
	sglog.Debug("raw update dt:", data.UpdateDtStr)
	sglog.Debug("raw expiry dt:", data.ExpiryDtStr)
	sglog.Info("==================end show=============")
}

func GetIpByDomain(domain string) (addrs []string, err error) {
	return net.LookupHost(domain)
}

func GetWhoisInfo(domain string) (*Whois, error) {
	var (
		parts      []string
		zone       string
		name       string
		connection net.Conn
	)

	result := new(Whois)
	result.CreateDt = sgtime.New()
	result.UpdateDt = sgtime.New()
	result.ExpiryDt = sgtime.New()

	result.Domain = domain
	parts = strings.Split(domain, ".")
	if len(parts) < 2 {
		result.IsRegist = SG_WHOIS_STATUS_CHECK_FAILD
		sglog.Error("Domain(", domain, ") name is wrong!")
		return result, errors.New("domain is wrong")
	}
	name = parts[len(parts)-2]
	zone = parts[len(parts)-1]
	server, ok := servers[zone]
	if !ok {
		result.IsRegist = SG_WHOIS_STATUS_CHECK_FAILD
		sglog.Error("no such zone server,zone:", zone)
		return result, errors.New("no such zone server")
	}
	connection, err := net.DialTimeout("tcp", net.JoinHostPort(server, "43"), time.Second*5)
	if connection != nil {
		defer connection.Close()
	}
	if err != nil {
		result.IsRegist = SG_WHOIS_STATUS_CHECK_FAILD
		sglog.Error("connect to dns server error:", err)
		return result, errors.New("connect to dns server error")
	}
	connection.Write([]byte("" + domain + "\r\n"))
	sglog.Info("domain:", domain, ",wait for get domainInfo")
	//超过30s即超时
	connection.SetReadDeadline(time.Now().Add(time.Second * 30))

	buf := new(bytes.Buffer)
	readNum, err := buf.ReadFrom(connection)

	//buffer, err = ioutil.ReadAll(connection)
	if err != nil {
		result.IsRegist = SG_WHOIS_STATUS_CHECK_FAILD
		sglog.Error(domain, " connection readAll error", err)
		return result, errors.New(" connection readAll error")
	}
	if readNum <= 0 {
		result.IsRegist = SG_WHOIS_STATUS_CHECK_FAILD
		sglog.Error(domain, " connection readAll error,buffer size =0")
		return result, errors.New("connection readAll error,buffer size =0")
	}

	sglog.Info("domain:", domain, ",read respond success")
	result.Raw = buf.String()
	result.Domain = domain
	result.Zone = zone
	result.Name = name
	result.Name_length = len(result.Name)
	buf = nil
	return result, nil
}

func ParseWhois(info *Whois) {
	if SG_WHOIS_STATUS_CHECK_FAILD == info.IsRegist {
		sglog.Info("fail to check domain:", info.Domain)
		return
	}
	switch info.Zone {
	case "com", "net":
		{
			ParseWhoisCom(info)
		}
	case "cn":
		{
			ParseWhoisCn(info)
		}
	}
}

func ParseWhoisCom(info *Whois) {
	str_list := strings.Split(info.Raw, "\r\n")
	if strings.Contains(info.Raw, "No match for domain") {
		info.IsRegist = SG_WHOIS_STATUS_CAN_REGIST_NOW
		return
	}

	if !strings.Contains(info.Raw, "Registry Expiry Date") {
		info.IsRegist = SG_WHOIS_STATUS_CHECK_FAILD
		return
	}

	info.IsRegist = SG_WHOIS_STATUS_HAD_REGIST
	time_str_len := 19
	start_index := 17
	for _, line := range str_list {
		if info.UpdateDtStr != "" && info.CreateDtStr != "" && info.ExpiryDtStr != "" {
			break
		}
		if strings.Contains(line, "Updated Date") {
			start_index = 17
			info.UpdateDtStr = string([]byte(line)[start_index:(start_index + time_str_len)])
			continue
		}
		if strings.Contains(line, "Creation Date") {
			start_index = 18
			info.CreateDtStr = string([]byte(line)[start_index:(start_index + time_str_len)])
			continue
		}
		if strings.Contains(line, "Registry Expiry Date") {
			start_index = 25
			info.ExpiryDtStr = string([]byte(line)[start_index:(start_index + time_str_len)])
			continue
		}
	}

	info.OldExpiryDtStr = info.ExpiryDtStr
	createDt, _ := time.Parse(sgtime.FORMAT_TIME_RFC_3339_SIMPLE, info.CreateDtStr)
	updateDt, _ := time.Parse(sgtime.FORMAT_TIME_RFC_3339_SIMPLE, info.UpdateDtStr)
	expiryDt, _ := time.Parse(sgtime.FORMAT_TIME_RFC_3339_SIMPLE, info.ExpiryDtStr)

	createDt = createDt.In(sgtime.GetTimeZone())
	updateDt = updateDt.In(sgtime.GetTimeZone())
	expiryDt = expiryDt.In(sgtime.GetTimeZone())

	info.CreateDt = &createDt
	info.UpdateDt = &updateDt
	info.ExpiryDt = &expiryDt

	info.CreateDt.In(sgtime.GetTimeZone())
}

func ParseWhoisCn(info *Whois) {
	str_list := strings.Split(info.Raw, "\n")
	if strings.Contains(info.Raw, "No matching record") {
		info.IsRegist = SG_WHOIS_STATUS_CAN_REGIST_NOW
		return
	}

	if strings.Contains(info.Raw, "the Domain Name you apply can not be registered online") {
		sglog.Error("info raw ", info.Raw, ",name=", info.Domain)
		info.IsRegist = SG_WHOIS_STATUS_LIMIT_BY_GOVERNMENT
		return
	}

	if !strings.Contains(info.Raw, "Expiration Time") {
		info.IsRegist = SG_WHOIS_STATUS_CHECK_FAILD
		return
	}

	info.IsRegist = SG_WHOIS_STATUS_HAD_REGIST

	time_str_len := 19
	start_index := 17
	for _, line := range str_list {
		if info.CreateDtStr != "" && info.ExpiryDtStr != "" {
			break
		}
		if strings.Contains(line, "Registration Time") {
			start_index = 19
			info.CreateDtStr = string([]byte(line)[start_index:(start_index + time_str_len)])
			continue
		}
		if strings.Contains(line, "Expiration Time") {
			start_index = 17
			info.ExpiryDtStr = string([]byte(line)[start_index:(start_index + time_str_len)])
			continue
		}
	}

	createDt, err := time.Parse(sgtime.FORMAT_TIME_NORMAL, info.CreateDtStr)
	if err != nil {
		info.CreateDt = &createDt
		info.UpdateDt = &createDt
	}

	expiryDt, err := time.Parse(sgtime.FORMAT_TIME_NORMAL, info.ExpiryDtStr)
	if err != nil {
		info.ExpiryDt = &expiryDt
	}
}

func IsHightValueDomainByName(domain string) bool {
	strlist := strings.Split(domain, ".")
	if len(strlist) == 2 {
		name := strlist[0]
		zone := strlist[1]
		return IsHightValueDomain(name, zone)
	}
	return false
}

func IsHightValueDomain(name string, zone string) bool {

	name_len := len(name)

	if name_len <= 2 {
		return true
	}

	if name_len >= 6 {
		return false
	}

	if GetCharacterNum(name) <= 2 {
		return true
	}

	if 3 == name_len {
		if isAllNumber(name) {
			return true
		}
		if isIncludeNumber(name) {
			return false
		}
		return true
	}

	if 4 == name_len {
		if isAllNumber(name) {
			return true
		}
		if !isIncludeNumber(name) {
			return false
		}
		if zone == "com" {
			return true
		}
		return false
	}

	if 5 == name_len {
		if isAllNumber(name) && zone == "com" {
			return true
		}
		return false
	}

	return false
}

func isAllNumber(name string) bool {
	_, err := strconv.Atoi(name)

	if err == nil {
		return true
	}
	return false
}

func isIncludeNumber(name string) bool {
	for i := 0; i < 10; i++ {
		if strings.Contains(name, strconv.Itoa(i)) {
			return true
		}
	}
	return false
}

func GetCharacterNum(name string) int {
	m := make(map[byte]int)
	n := len(name)
	for i := 0; i < n; i++ {
		ch := name[i]
		m[ch] = 1
	}
	return len(m)
}
