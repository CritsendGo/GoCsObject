package GoCsObject

import (
	"errors"
	"fmt"
	"strconv"
)

type Mx struct {
	Id        int    `json:"mx_id,omitempty"`
	Name      string `json:"mx_name"`
	Updatable string `json:"ip_updatable,omitempty"`
	Removable bool   `json:"is_removable"`
}

// NewMxByValue
// Constructor to build object by Value, if fromApi = true, try to load from API
func NewMxByValue(value string, fromApi bool) (o *Mx, err error) {
	o.Name = value
	if fromApi == true {
		err = o.Load()
	}
	return o, err
}

// Load
// Function used to load the Object to Api
func (o *Mx) Load() (err error) {
	if o.Id > 0 {
		query := ApiUrl + "ip" + "/?mx_id=" + strconv.Itoa(o.Id)
		b := getObjectFromApi(query)
		fmt.Println(b)
	} else if o.Name != "" {
		query := ApiUrl + "ip" + "/?mx_name=" + o.Name
		b := getObjectFromApi(query)
		fmt.Println(b)
	} else {
		return errors.New("please set Id or Name before querying")
	}
	return err
}

// Save
// Function used to save the Object to Api
func (o *Mx) Save() {

}

// Delete
// Function used to Delete the Object to Api
func (o *Mx) Delete() {

}
