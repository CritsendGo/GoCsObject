package Mx

import (
	"errors"
	"fmt"
	"github.com/CritsendGo/GoCsObject"
	"strconv"
)

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
		query := GoCsObject.ApiUrl + "ip" + "/?mx_id=" + strconv.Itoa(o.Id)
		b := GoCsObject.GetObjectFromApi(query)
		fmt.Println(b)
	} else if o.Name != "" {
		query := GoCsObject.ApiUrl + "ip" + "/?mx_name=" + o.Name
		b := GoCsObject.GetObjectFromApi(query)
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
