package main

import (
	"fmt"
	"log"

	"github.com/charconstpointer/excry/excry"
)

func main() {
	s := `{
		"root": [
			{
				"qux": "quux",
				"foo": "bar"
			}
		]
	}`

	exists, _ := excry.Exists(s, "foo")
	fmt.Println(exists)
	existsWithVal, err := excry.ExistsWithVal(s, "qux", "quux")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(existsWithVal)
}
