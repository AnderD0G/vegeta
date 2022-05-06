package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//todo:code2Session,signature
var client *http.Client = &http.Client{}

const (
	code2Session = "https://api.weixin.qq.com/sns/jscode2session"
	appId        = "wx2c4dfb4d33420782"
	appSec       = "3021cea2ab2cc3c88df373ea28c56e5d"
)

func Get(url string, q map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	for k, v := range q {
		query.Add(k, v)
	}
	req.URL.RawQuery = query.Encode()
	fmt.Println(req.URL.String())
	return client.Do(req)
}

type WxResp struct {
	SessionKey string `json:"session_key"`
	Openid     string `json:"openid"`
}

func Code2Session(code string) (WxResp, error) {
	m := map[string]string{"grant_type": "authorization_code", "appid": appId, "secret": appSec, "js_code": code}
	resp, err := Get(code2Session, m)
	wxResp := WxResp{}
	if err != nil {
		return WxResp{}, err
	}
	if all, err := ioutil.ReadAll(resp.Body); err != nil {
		return WxResp{}, err
	} else {
		err := json.Unmarshal(all, &wxResp)
		if err != nil {
			return WxResp{}, err
		}
		return wxResp, err
	}

}
