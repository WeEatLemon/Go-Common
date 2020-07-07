package hoo

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

const (
	GET  = "GET"
	POST = "POST"
)

type Params map[string]string

type Hoo struct {
	ClientId     string
	ClientSecret string
	HOST         string
}

func NewHoo(ClientId, ClientSecret, Host string) *Hoo {
	return &Hoo{
		ClientId:     ClientId,
		ClientSecret: ClientSecret,
		HOST:         Host,
	}
}

type BaseReq struct {
	Code    string
	Message string
	Data    interface{}
}

func CheckBase(Message string, result interface{}) (err error) {
	var data BaseReq
	err = json.Unmarshal([]byte(Message), &data)
	if err != nil {
		return
	}
	if data.Code != "10000" {
		return errors.New(data.Message)
	}
	fmt.Printf("data is %+v \n", data.Data)
	err = mapstructure.Decode(data.Data, &result)
	return
}

func (H *Hoo) SortParams(params map[string]string) string {
	keys := make([]string, len(params))
	i := 0
	for k := range params {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	sorted := make([]string, len(params))
	i = 0
	for _, k := range keys {
		sorted[i] = k + "=" + url.QueryEscape(params[k])
		i++
	}
	return strings.Join(sorted, "&")
}

func (H *Hoo) getHmacHde(s string) string {
	h := hmac.New(sha256.New, []byte(H.ClientSecret))
	_, _ = io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (H *Hoo) SignEcc(params map[string]string) (DataUrl string) {
	DataUrl = H.SortParams(params)
	DataUrl = H.getHmacHde(DataUrl)
	return
}

func (H *Hoo) Request(method string, path string, params map[string]string) (result string, err error) {
	client := &http.Client{}
	params["client_id"] = H.ClientId
	params["sign"] = H.SignEcc(params)
	sorted := H.SortParams(params)
	var req *http.Request
	if method == "POST" {
		req, _ = http.NewRequest(method, H.HOST+path, strings.NewReader(sorted))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, _ = http.NewRequest(method, H.HOST+path+"?"+sorted, strings.NewReader(""))
	}

	fmt.Println("Client Do......")
	resp, err := client.Do(req)
	fmt.Println("resp is", resp)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	result = string(body)
	return
}
