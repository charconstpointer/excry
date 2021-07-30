package excry

import (
	"encoding/json"
	"fmt"
)

func Exists(input, key string) bool {
	var j map[string]interface{}
	err := json.Unmarshal([]byte(input), &j)
	if err != nil {
		return false
	}

	for k, v := range j {
		if k == key {
			return true
		}
		exists := Exists(fmt.Sprintf("%s", v), key)
		if exists {
			return true
		}
	}
	return false
}

func ExistsWithVal(input, key, val string) bool {
	return false
}
