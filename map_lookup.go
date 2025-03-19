package main

import "fmt"

var myMap = map[string]string{
	"test":   "123",
	"wanted": "got it fast!",
	"other":  "xxx",
}

func main() {
	keys := make([]string, 0, len(myMap))
	values := make([]string, 0, len(myMap))

	for k, v := range myMap {
		keys = append(keys, k)
		values = append(values, v)
	}

	var wantedIndex int
	for i, k := range keys {
		if k == "wanted" {
			wantedIndex = i
			break
		}
	}

	fmt.Println("value:", values[wantedIndex])
}
