package excry

import (
	"encoding/json"
	"fmt"
)

type res struct {
	val string
	ok  bool
}

func Exists(input, key string) (bool, error) {
	res, err := find(input, key)
	if err != nil {
		return false, fmt.Errorf("invalid json input provided, %w", err)
	}
	if !res.ok {
		return false, nil
	}
	return true, nil
}

func ExistsWithVal(input, key, val string) (bool, error) {
	res, err := find(input, key)
	if err != nil {
		return false, fmt.Errorf("invalid json input provided, %w", err)
	}
	if !res.ok {
		return false, nil
	}
	if res.val != val {
		return false, fmt.Errorf("different value, expected %s found %s", val, res.val)
	}
	return true, nil
}

func find(input, key string) (res, error) {
	var j map[string]interface{}
	err := json.Unmarshal([]byte(input), &j)
	if err != nil {
		return res{}, fmt.Errorf("invalid json input provided, %w", err)
	}
	for k, v := range j {
		if k == key {
			return res{
				val: fmt.Sprintf("%s", v),
				ok:  true,
			}, nil
		}
		lr, err := json.Marshal(v)
		if err != nil {
			return res{}, fmt.Errorf("invalid json input provided, %w", err)
		}
		exists, _ := find(string(lr), key)
		if exists.ok {
			return res{
				val: exists.val,
				ok:  true,
			}, nil
		}
	}
	return res{}, fmt.Errorf("key not found")
}
