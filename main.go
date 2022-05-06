package GoCsObject

import (
	"io/ioutil"
	"net/http"
)

const (
	ApiUrl = "https://api.critsend.io/"
)

var (
	ApiToken string
)

func getObjectFromApi(url string) []byte {
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
