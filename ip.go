package GoCsObject

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	ApiName = "ip"
)

type Ip struct {
	Id        int    `json:"ip_id,omitempty"`
	Value     string `json:"ip_value"`
	Reverse   string `json:"ip_reverse,omitempty"`
	Create    string `json:"ip_create,omitempty"`
	Update    string `json:"ip_update,omitempty"`
	Server    int    `json:"server_id,omitempty"`
	Removable bool   `json:"is_removable"`
	User      int    `json:"user_id,omitempty"`
	City      int    `json:"city_id,omitempty"`
	Lat       string `json:"ip_lat,omitempty"`
	Long      string `json:"ip_long,omitempty"`
	Company   int    `json:"company_id,omitempty"`
	Status    int    `json:"ip_status_id,omitempty"`
}

// NewIpByValue
// Constructor to build object by Value, if fromApi = true, try to load from API
func NewIpByValue(value string, fromApi bool) (i *Ip, err error) {
	i.Value = value
	if fromApi == true {
		err = i.Load()
	}
	return i, err
}

// NewIpById
// Constructor to build object by ID, if fromApi = true, try to load from API
func NewIpById(value int, fromApi bool) (i *Ip, err error) {
	i.Id = value
	if fromApi == true {
		err = i.Load()
	}
	return i, err
}

// Load
// Function used to load the Object to Api
func (o *Ip) Load() (err error) {
	if o.Id > 0 {
		query := ApiUrl + ApiName + "/?ip_id=" + strconv.Itoa(o.Id)
		b := getObjectFromApi(query)
		fmt.Println(b)
	} else if o.Value != "" {
		query := ApiUrl + ApiName + "/?ip_value=" + o.Value
		b := getObjectFromApi(query)
		fmt.Println(b)
	} else {
		return errors.New("please set Id or Value before querying")
	}
	return err
}

// Save
// Function used to save the Object to Api
func (o *Ip) Save() {

}

// Delete
// Function used to Delete the Object to Api
func (o *Ip) Delete() {

}
