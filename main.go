package main

import (
	"fmt"
	"log"

	"github.com/charconstpointer/excry/excry"
)

func main() {
	s := `{
		"root": {
			"baz": [
				{
					"qux": 
					"quux"
				}
			],
			"foo": "bar"
		}
	}`
	exists, err := excry.Exists(s, "foo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(exists)
	existsWithVal, err := excry.ExistsWithVal(s, "foo", "bar")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(existsWithVal)
}
