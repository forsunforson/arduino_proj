package arduino

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	typeMap map[string]int
)

func init() {
	m := make(map[string]int, 0)
	b, err := ioutil.ReadFile("./arduino/arduino_config.json")
	if err != nil {
		fmt.Println(err)
		panic("arduino config json file missing")
	}
	err = json.Unmarshal(b, &m)
	if err != nil {
		fmt.Println(err)
		panic("arduino_config.json is broken")
	}
	typeMap = m
}

// GetCoreType int -> string arduino type
func GetCoreType(t int) string {
	for k, v := range typeMap {
		if v == t {
			return k
		}
	}
	return "unknow"
}

// GetCoreTypeInt string -> int see:./arduino_config.json
func GetCoreTypeInt(coreType string) int {
	t := strings.ToLower(coreType)
	if v, ok := typeMap[t]; ok {
		return v
	}
	return 0
}
