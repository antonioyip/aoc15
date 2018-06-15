package main

import (
	"fmt"
)

// circuit will be a map[string]int
// wire name is the key, value is the signal

func main() {
	x := uint16(123)
	y := uint16(456)
	fmt.Println(x & y)
	fmt.Println(x | y)
	fmt.Println(x << 2)
	fmt.Println(y >> 2)
	fmt.Println(^x)
	fmt.Println(^y)
}
