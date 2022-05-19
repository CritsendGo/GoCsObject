package GoCsObject

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ApiUrl   = "https://api.critsend.io/"
	ApiToken = ""
)

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
