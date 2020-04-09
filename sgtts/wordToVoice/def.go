package wordToVoice

const (
	WORD_MAX_LEN int = 1800
)

type SRequireParam struct {
	Common struct {
		AppID string `json:"app_id"`
	} `json:"common"`
	Business struct {
		Ent    string `json:"ent"`
		Aue    string `json:"aue"`
		Sfl    int    `json:"sfl"`
		Auf    string `json:"auf"`
		Vcn    string `json:"vcn"`
		Speed  int    `json:"speed"`
		Volume int    `json:"volume"`
		Pitch  int    `json:"pitch"`
		Bgs    int    `json:"bgs"`
		Tte    string `json:"tte"`
		Reg    string `json:"reg"`
		RAM    string `json:"ram"`
		Rdn    string `json:"rdn"`
	} `json:"business"`
	Data struct {
		Status   int    `json:"status"`
		Encoding string `json:"encoding"`
		Text     string `json:"text"`
	} `json:"data"`
}

func NewParam() *SRequireParam {
	tmp := new(SRequireParam)
	tmp.Business.Ent = "intp65"
	tmp.Business.Aue = "lame"
	tmp.Business.Sfl = 1
	tmp.Business.Auf = "audio/L16;rate=16000"
	tmp.Business.Vcn = "xiaoyan"
	tmp.Business.Speed = 50
	tmp.Business.Volume = 50
	tmp.Business.Pitch = 50
	tmp.Business.Bgs = 0
	tmp.Business.Tte = "UTF8"
	tmp.Business.Reg = "2"
	tmp.Business.RAM = "0"
	tmp.Business.Rdn = "0"

	return tmp
}

type SResponResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Sid     string `json:"sid"`
	Data    struct {
		Audio  string `json:"audio"`
		Status int    `json:"status"`
		Ced    string `json:"ced"`
	} `json:"data"`
}
