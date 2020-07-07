package wechat

import (
	"encoding/json"
	"errors"
	"github.com/IEatLemons/GoBase/common/helper/request"
	"github.com/IEatLemons/GoBase/common/helper/rsa"
	"github.com/mitchellh/mapstructure"
	"github.com/yuchenfw/gocrypt"
	"strconv"
)

const (
	RMPayUrl = "/open_api/third-pay/wechat"
)

var Rockminer *Rm

type Rm struct {
	Protocol string
	Host     string
}

func NewRm() *Rm {
	return Rockminer
}

func InitRm(Protocol, Host string) *Rm {
	Rockminer = &Rm{
		Protocol: Protocol,
		Host:     Host,
	}
	return Rockminer
}

type ThirdPayToRM struct {
	OrderSn     string  `json:"order_sn"`
	Token       string  `json:"token"`
	PaymentType int64   `json:"payment_type"`
	PayAmount   float64 `json:"pay_amount"`
	OrderType   int64   `json:"order_type"`
	Lang        string  `json:"lang"`
	Client      int64   `json:"client"`
	ProcessType int64   `json:"process_type"`
	Platform    string  `json:"platform"`
}

type BaseReq struct {
	RetData interface{} `json:"retData"`
	RetMsg  string      `json:"retMsg"`
	RetCode int64       `json:"retCode"`
}

type PayUrl struct {
	OrderSn   string `json:"order_sn"`
	Url       string `json:"url"`
	UrlStatus string `json:"url_status"`
	UrlCode   string `json:"url_code"`
}

func (R *Rm) GetPayUrl(Params *ThirdPayToRM) (Url string, err error) {
	Url = ""
	params := getParams(Params)
	url := R.Protocol + R.Host + RMPayUrl
	result, err := request.Request(request.POST, url, params, request.ContentTypFormUrl)
	if err != nil {
		return
	}
	var PayData PayUrl
	err = checkBase(result, &PayData)
	Url = PayData.Url
	return
}

func getParams(Params *ThirdPayToRM) request.ReqParams {
	var params = request.ReqParams{}
	params["order_sn"] = Params.OrderSn
	params["token"] = Params.Token
	params["payment_type"] = strconv.FormatInt(Params.PaymentType, 10)
	params["pay_amount"] = strconv.FormatFloat(Params.PayAmount, 'G', -1, 64)
	params["order_type"] = strconv.FormatInt(Params.OrderType, 10)
	params["lang"] = Params.Lang
	params["platform"] = Params.Platform
	params["client"] = strconv.FormatInt(Params.Client, 10)
	params["process_type"] = strconv.FormatInt(Params.ProcessType, 10)
	return params
}

func getRsaParams(params request.ReqParams, RSA *rsa.Crypt) (request.ReqParams, error) {
	var RsaParams = request.ReqParams{}
	data, _ := json.Marshal(params)

	sign, err := RSA.Sign(string(data), gocrypt.SHA256, gocrypt.Base64)
	if err != nil {
		return nil, err
	}
	RsaParams["sign"] = sign

	data2, err := RSA.Encrypt(string(data), gocrypt.Base64)
	if err != nil {
		return nil, err
	}
	RsaParams["data"] = data2
	return RsaParams, nil
}

func checkBase(Message string, result interface{}) (err error) {
	var data BaseReq
	err = json.Unmarshal([]byte(Message), &data)
	if err != nil {
		return
	}
	if data.RetCode != 0 {
		err = errors.New(data.RetMsg)
		return
	}
	err = mapstructure.Decode(data.RetData, &result)
	return
}
