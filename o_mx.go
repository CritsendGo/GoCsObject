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
func NewMxByValue(value string, fromApi bool) (*Mx, error) {
	o := &Mx{Name: value}
	var err error
	if fromApi == true {
		err = o.Load()
	}
	return o, err
}

// Load
// Function used to load the Object to Api
func (o *Mx) Load() (err error) {
	if o.Id > 0 {
		b := getObjectFromApi(ApiUrl + "mx" + "/?mx_id=" + strconv.Itoa(o.Id))
		fmt.Println(b)
	} else if o.Name != "" {
		b := getObjectFromApi(ApiUrl + "mx" + "/?mx_name=" + o.Name)
		v,_:=b.(map[string]interface{})
		o.Id,_=strconv.Atoi(fmt.Sprint(v["mx_id"]))

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
