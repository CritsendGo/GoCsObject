package GoCsObject

import (
	"io/ioutil"
	"net/http"
	"encoding/json"

)

type ApiResp struct {
	Error []interface{} `json:"error"`
	Info  struct {
		Query int    `json:"query"`
		Limit int    `json:"limit"`
		Page  int    `json:"page"`
		Count string `json:"count"`
	} `json:"info"`
	Result []interface {} `json:"result"`
	Debug string `json:"debug"`
}

const (
	ApiUrl = "https://api.critsend.io/"
)

var (
	ApiToken string
)

func getObjectFromApi(url string) interface{}{
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
	}
	req.Header = http.Header{"Authorization": []string{ApiToken}}
	res, err := client.Do(req)
	if err != nil {
	}
	b, err := ioutil.ReadAll(res.Body)
	resp := ApiResp{}
	json.Unmarshal(b, &resp)
	if resp.Info.Count=="1"{
		return resp.Result[0]
	}else{

	}
	return nil
}
