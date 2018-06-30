package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	inputs, err := ioutil.ReadFile("day12.input")
	if err != nil {
		panic(err)
	}

	var v interface{}
	err = json.Unmarshal(inputs, &v)
	if err != nil {
		panic(err)
	}
	fmt.Println(tally(v))
}

func tally(v interface{}) float64 {
	sum := 0.0
	if num, ok := v.(float64); ok {
		sum += num
	} else if array, ok := v.([]interface{}); ok {
		for _, entry := range array {
			sum += tally(entry)
		}
	} else if object, ok := v.(map[string]interface{}); ok {
		for _, entry := range object {
			sum += tally(entry)
			if entry == "red" {
				return 0.0
			}
		}
	}
	return sum
}
