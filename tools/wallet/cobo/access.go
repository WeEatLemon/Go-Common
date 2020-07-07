package cobo

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcd/btcec"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

const (
	GET  = "GET"
	POST = "POST"
)

type BaseReq struct {
	Success bool
	Result  interface{}
}

func CheckBase(Message string, result interface{}) (err error) {
	var data BaseReq
	err = json.Unmarshal([]byte(Message), &data)

	if err != nil {
		return
	}
	err = mapstructure.Decode(data.Result, &result)
	return
}

type Params map[string]string

type Key struct {
	PriKey    *btcec.PrivateKey
	ApiKey    string
	ApiSecret []byte
}

func GenerateRandomKeyPair() (Key Key) {
	Key.ApiSecret = make([]byte, 32)
	if _, err := rand.Read(Key.ApiSecret); err != nil {
		panic(err)
	}

	Key.PriKey, _ = btcec.PrivKeyFromBytes(btcec.S256(), Key.ApiSecret)
	Key.ApiKey = fmt.Sprintf("%x", Key.PriKey.PubKey().SerializeCompressed())
	return
}

func Hash256(s string) (hashString string) {
	hashResult := sha256.Sum256([]byte(s))
	hashString = string(hashResult[:])
	return
}
func Hash256x2(s string) string {
	return Hash256(Hash256(s))
}

type Cobo struct {
	ApiKey    string
	ApiSecret string
	HOST      string
	ApiPub    string
}

func NewCobo(ApiKey, ApiSecret, HOST, ApiPub string) *Cobo {
	return &Cobo{
		ApiKey:    ApiKey,
		ApiSecret: ApiSecret,
		HOST:      HOST,
		ApiPub:    ApiPub,
	}
}

func (Co *Cobo) SortParams(params map[string]string) string {
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

func (Co *Cobo) SignEcc(message string) string {
	apiSecret, _ := hex.DecodeString(Co.ApiSecret)
	privateKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), apiSecret)
	sig, _ := privateKey.Sign([]byte(Hash256x2(message)))
	return fmt.Sprintf("%x", sig.Serialize())
}

func (Co *Cobo) VerifyEcc(message string, signature string) bool {
	pkey, _ := hex.DecodeString(Co.ApiPub)
	pubKey, _ := btcec.ParsePubKey(pkey, btcec.S256())

	sigBytes, _ := hex.DecodeString(signature)
	sigObj, _ := btcec.ParseSignature(sigBytes, btcec.S256())

	verified := sigObj.Verify([]byte(Hash256x2(message)), pubKey)
	return verified
}

func (Co *Cobo) Request(method string, path string, params map[string]string) string {
	client := &http.Client{}
	nonce := fmt.Sprintf("%d", time.Now().Unix()*1000)
	sorted := Co.SortParams(params)
	var req *http.Request
	if method == "POST" {
		req, _ = http.NewRequest(method, Co.HOST+path, strings.NewReader(sorted))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, _ = http.NewRequest(method, Co.HOST+path+"?"+sorted, strings.NewReader(""))
	}
	content := strings.Join([]string{method, path, nonce, sorted}, "|")

	req.Header.Set("Biz-Api-Key", Co.ApiKey)
	req.Header.Set("Biz-Api-Nonce", nonce)
	req.Header.Set("Biz-Api-Signature", Co.SignEcc(content))

	//fmt.Printf("%+v \n", req)

	resp, _ := client.Do(req)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	timestamp := resp.Header["Biz-Timestamp"][0]
	signature := resp.Header["Biz-Resp-Signature"][0]
	success := Co.VerifyEcc(string(body)+"|"+timestamp, signature)
	fmt.Println("verify success?", success)
	return string(body)
}
