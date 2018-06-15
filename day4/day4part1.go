package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
)

func md5hash(input string) string {
	sum := md5.Sum([]byte(input))
	return hex.EncodeToString(sum[:])
}

func test(input string, expected bool) {
	hash := md5hash(input)
	//fmt.Println(hash)
	result := hash[0:5] == "00000"
	if result == expected {
		fmt.Println("Success")
	} else {
		fmt.Println("Failure")
	}
}

func main() {
	test("abcdef609043", true)
	test("pqrstuv1048970", true)
	test("12345", false)

	input := "iwrupvqb"
	for i := 0; i < 10000000000; i++ {
		hash := md5hash(input + strconv.Itoa(i))
		if hash[0:5] == "00000" {
			fmt.Println(i)
			return
		}
	}
}
