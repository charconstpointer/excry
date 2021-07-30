# excry 🕵🏻‍♂️
### let s be valid JSON string
```
s := `{
		"root": [
			{
				"qux": "quux",
				"foo": "bar"
			}
		]
	}`
```
#### you can check if fields exists
```
exists, _ := excry.Exists(s, "foo") //output: true
exists, _ := excry.Exists(s, "foo2") //output: false
```

#### you can check if fields exists with given value
```
existsWithVal, err := excry.ExistsWithVal(s, "qux", "quux") //output: true
existsWithVal, err := excry.ExistsWithVal(s, "foo", "bar2") //output: false
```