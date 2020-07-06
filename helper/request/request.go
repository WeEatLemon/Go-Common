package request

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

const (
	POST = "POST"
	GET  = "GET"
	PUT  = "PUT"
)

const (
	ContentTypJson    = "application/json"
	ContentTypFormUrl = "application/x-www-form-urlencoded"
)

type ReqParams map[string]string

func SortParams(params ReqParams) string {
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

func Request(method string, path string, params ReqParams, ContentType string) (result string, err error) {
	client := &http.Client{}
	sorted := SortParams(params)
	fmt.Printf("sorted is %+v \n", sorted)
	fmt.Printf("path is %+v \n", path)
	var req *http.Request
	if method == GET {
		req, _ = http.NewRequest(method, path+"?"+sorted, strings.NewReader(""))
	} else {
		req, _ = http.NewRequest(method, path, strings.NewReader(sorted))
		if ContentType != "" {
			req.Header.Add("Content-Type", ContentType)
		} else {
			req.Header.Set("Content-Type", ContentTypFormUrl)
		}
	}
	fmt.Printf("req is %+v \n", req)

	fmt.Println("Client Do......")
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	result = string(body)
	return
}
