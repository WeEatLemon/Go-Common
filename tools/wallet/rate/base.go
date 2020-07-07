package rate

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
)

func Request(method string, path string, params map[string]string) (result string, err error) {
	sorted := SortParams(params)
	c := &http.Client{}
	var req *http.Request
	if method == POST {
		req, _ = http.NewRequest(method, path, strings.NewReader(sorted))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, _ = http.NewRequest(method, path+"?"+sorted, strings.NewReader(""))
	}

	fmt.Println("Client Do......")
	resp, err := c.Do(req)
	fmt.Println("resp is", resp)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	result = string(body)
	return
}

func SortParams(params map[string]string) string {
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
