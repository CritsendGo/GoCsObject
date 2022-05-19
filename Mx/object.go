package Mx

import (
	"fmt"
	"go/types"
)

const apiName = "mx"

var isUsable = false

type Mx struct {
	Id        int    `json:"mx_id,omitempty"`
	Name      string `json:"mx_name"`
	Updatable bool   `json:"ip_updatable,omitempty"`
	Removable bool   `json:"is_removable,omitempty"`
}

var mapJson map[string]types.BasicKind
var mapName map[string]types.BasicKind

func init() {
	// Json Struct
	mapJson = make(map[string]types.BasicKind)
	mapJson["mx_id"] = types.Int
	mapJson["mx_name"] = types.String
	mapJson["is_updatable"] = types.Bool
	mapJson["is_removable"] = types.Bool
	// Name Struct
	mapName = make(map[string]types.BasicKind)
	mapName["Id"] = types.Int
	mapName["Name"] = types.String
	mapName["Updatable"] = types.Bool
	mapName["Removable"] = types.Bool
	// Check on first load if current object is the same that the production and set to usable = true
	err := checkStruct()
	if err == nil {
		isUsable = true
	}
	fmt.Println("Usable", isUsable)
}
