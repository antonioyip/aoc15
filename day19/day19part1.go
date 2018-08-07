package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	test()

	inputs, err := ioutil.ReadFile("day19.input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputs), "\n")

	original := "CRnCaSiRnBSiRnFArTiBPTiTiBFArPBCaSiThSiRnTiBPBPMgArCaSiRnTiMgArCaSiThCaSiRnFArRnSiRnFArTiTiBFArCaCaSiRnSiThCaCaSiRnMgArFYSiRnFYCaFArSiThCaSiThPBPTiMgArCaPRnSiAlArPBCaCaSiRnFYSiThCaRnFArArCaCaSiRnPBSiRnFArMgYCaCaCaCaSiThCaCaSiAlArCaCaSiRnPBSiAlArBCaCaCaCaSiThCaPBSiThPBPBCaSiRnFYFArSiThCaSiRnFArBCaCaSiRnFYFArSiThCaPBSiThCaSiRnPMgArRnFArPTiBCaPRnFArCaCaCaCaSiRnCaCaSiRnFYFArFArBCaSiThFArThSiThSiRnTiRnPMgArFArCaSiThCaPBCaSiRnBFArCaCaPRnCaCaPMgArSiRnFYFArCaSiThRnPBPMgAr"
	molecules := map[string]bool{}
	var oldText string
	var newText string
	for _, line := range lines {
		fmt.Sscanf(line, "%s => %s", &oldText, &newText)
		// fmt.Println(oldText, newText)
		for i := 1; true; i++ {
			success, newString := replaceNth(original, oldText, newText, i)
			if !success {
				break
			}
			molecules[newString] = true
		}
	}

	fmt.Println(len(molecules))
}

// replace the nth instance of oldText with newText in input
// return true if nth oldText is found, false otherwise
// return resulting string
func replaceNth(input string, oldText string, newText string, nth int) (bool, string) {
	index := strings.Index(input, oldText)
	if index == -1 || nth < 1 {
		return false, ""
	}
	if nth == 1 {
		return true, input[:index] + newText + input[index+len(oldText):]
	}
	exists, resultStr := replaceNth(input[index+len(oldText):], oldText, newText, nth-1)
	if exists {
		return true, input[:index+len(oldText)] + resultStr
	}
	return false, ""
}

func assert(b1 bool, s1 string, b2 bool, s2 string) bool {
	if b1 == b2 {
		if b1 {
			// exists, compare strings
			return s1 == s2
		}
		// does not exist
		return true
	}
	return false
}

func test() {
	b, s := replaceNth("hohoho", "ho", "tree", 1)
	fmt.Println(assert(b, s, true, "treehoho"))

	b, s = replaceNth("hohoho", "ho", "tree", 2)
	fmt.Println(assert(b, s, true, "hotreeho"))

	b, s = replaceNth("hohoho", "ho", "tree", 3)
	fmt.Println(assert(b, s, true, "hohotree"))

	b, s = replaceNth("hohoho", "ho", "tree", 0)
	fmt.Println(assert(b, s, false, ""))

	b, s = replaceNth("hohoho", "ho", "tree", 4)
	fmt.Println(assert(b, s, false, ""))

	b, s = replaceNth("hohoho", "ae", "e", 0)
	fmt.Println(assert(b, s, false, ""))

	b, s = replaceNth("cat dog bird dog cat bird", "cat", "tree", 0)
	fmt.Println(assert(b, s, false, ""))

	b, s = replaceNth("cat dog bird dog cat bird", "cat", "tree", 1)
	fmt.Println(assert(b, s, true, "tree dog bird dog cat bird"))

	b, s = replaceNth("cat dog bird dog cat bird", "cat", "tree", 2)
	fmt.Println(assert(b, s, true, "cat dog bird dog tree bird"))

	b, s = replaceNth("cat dog bird dog cat bird", "cat", "tree", 3)
	fmt.Println(assert(b, s, false, ""))

	b, s = replaceNth("cat dog bird dog cat bird", "dog", "tree", 0)
	fmt.Println(assert(b, s, false, ""))

	b, s = replaceNth("cat dog bird dog cat bird", "dog", "tree", 1)
	fmt.Println(assert(b, s, true, "cat tree bird dog cat bird"))

	b, s = replaceNth("cat dog bird dog cat bird", "dog", "tree", 2)
	fmt.Println(assert(b, s, true, "cat dog bird tree cat bird"))

	b, s = replaceNth("cat dog bird dog cat bird", "dog", "tree", 3)
	fmt.Println(assert(b, s, false, ""))

	b, s = replaceNth("cat dog bird dog cat bird", "bird", "tree", 0)
	fmt.Println(assert(b, s, false, ""))

	b, s = replaceNth("cat dog bird dog cat bird", "bird", "tree", 1)
	fmt.Println(assert(b, s, true, "cat dog tree dog cat bird"))

	b, s = replaceNth("cat dog bird dog cat bird", "bird", "tree", 2)
	fmt.Println(assert(b, s, true, "cat dog bird dog cat tree"))

	b, s = replaceNth("cat dog bird dog cat bird", "bird", "tree", 3)
	fmt.Println(assert(b, s, false, ""))

}
