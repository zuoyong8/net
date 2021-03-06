package xhttp

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

//
//获取数据
func Get(url string, headers map[string]string, querys map[string]interface{}) ([]byte, error) {
	url = joinQuery(url, querys)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	transport := http.Transport{
		DisableKeepAlives: true, //禁止
	}

	client := http.Client{
		Transport: &transport,
	}

	//
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	info, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return info, nil
}

//
//post
func Post(url string, headers map[string]string, params map[string]interface{}) ([]byte, error) {
	ioReder := getBytesReader(params)
	req, err := http.NewRequest(http.MethodPost, url, ioReder)
	if err != nil {
		return nil, err
	}

	transport := http.Transport{
		DisableKeepAlives: true, //禁止
	}

	client := http.Client{
		Transport: &transport,
	}

	//
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	info, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return info, nil
}

//连接参数
func joinQuery(url string, querys map[string]interface{}) string {
	//
	if url != "" {
		if querys != nil {
			var q string
			for k, v := range querys {
				q = fmt.Sprintf("%v&%v=%v", q, k, v)
			}
			return url + "?" + q[1:]
		}
	}
	return url
}

//
//
func getBytesReader(params map[string]interface{}) *bytes.Reader {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary

	if bs, err := json.Marshal(params); err == nil && bs != nil {
		return bytes.NewReader(bs)
	}
	return nil
}
