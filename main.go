package GoCsObject

import (
	"encoding/json"
	"fmt"
	"go/types"
	"io/ioutil"
	"net/http"
)

type AutoGenerated struct {
	Error   []interface{} `json:"error"`
	Info    []interface{} `json:"info"`
	Title   string        `json:"title"`
	History []struct {
		Code string `json:"code"`
		Info string `json:"info"`
	} `json:"history"`
	Description string      `json:"description"`
	Parameters  interface{} `json:"parameters"`
}

const (
	ApiUrl   = "https://api.critsend.io/"
	ApiToken = ""
)

type ObjectField struct {
	Name string
	Type types.BasicKind
}

func CompareObject(apiName string) /*map[string]types.BasicKind*/ {
	// Build Simple Object
	client := http.Client{}
	req, err := http.NewRequest("OPTIONS", ApiUrl+apiName, nil)
	if err != nil {
		//todo:Check if object doesn't exist
	}
	res, err := client.Do(req)
	if err != nil {
		//todo:Check if query doesn't work
	}
	b, err := ioutil.ReadAll(res.Body)
	// Descode Body
	fmt.Println(string(b))
	var json2Struct AutoGenerated
	err = json.Unmarshal(b, &json2Struct)
	fmt.Println(json2Struct, err)
}
func GetObjectFromApi(url string) []byte {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
	}
	req.Header = http.Header{"Authorization": []string{ApiToken}}
	res, err := client.Do(req)
	if err != nil {
	}
	b, err := ioutil.ReadAll(res.Body)
	return b
}
