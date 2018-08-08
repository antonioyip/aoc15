package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	inputs, err := ioutil.ReadFile("day19.input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputs), "\n")

	machine := map[string]string{}
	var smaller string
	var larger string
	for _, line := range lines {
		fmt.Sscanf(line, "%s => %s", &smaller, &larger)
		machine[larger] = smaller
	}

	testReduce()

	startMolecule := "CRnCaSiRnBSiRnFArTiBPTiTiBFArPBCaSiThSiRnTiBPBPMgArCaSiRnTiMgArCaSiThCaSiRnFArRnSiRnFArTiTiBFArCaCaSiRnSiThCaCaSiRnMgArFYSiRnFYCaFArSiThCaSiThPBPTiMgArCaPRnSiAlArPBCaCaSiRnFYSiThCaRnFArArCaCaSiRnPBSiRnFArMgYCaCaCaCaSiThCaCaSiAlArCaCaSiRnPBSiAlArBCaCaCaCaSiThCaPBSiThPBPBCaSiRnFYFArSiThCaSiRnFArBCaCaSiRnFYFArSiThCaPBSiThCaSiRnPMgArRnFArPTiBCaPRnFArCaCaCaCaSiRnCaCaSiRnFYFArFArBCaSiThFArThSiThSiRnTiRnPMgArFArCaSiThCaPBCaSiRnBFArCaCaPRnCaCaPMgArSiRnFYFArCaSiThRnPBPMgAr"
	fmt.Println(depthFirstSearch(machine, startMolecule, "e"))
}

func testReduce() {
	machine := map[string]string{"H": "e", "O": "e", "HO": "H", "OH": "H", "HH": "O"}
	fmt.Println(depthFirstSearch(machine, "HOHOHO", "e"))
}

func depthFirstSearch(machine map[string]string, startMolecule string, finalMolecule string) int {
	sortedKeys := make(StringSlice, len(machine))
	i := 0
	for key := range machine {
		sortedKeys[i] = key
		i++
	}
	sort.Sort(sortedKeys)
	return depthFirstSearchCore(machine, sortedKeys, startMolecule, finalMolecule)
}

func depthFirstSearchCore(machine map[string]string, sortedMachineKeys StringSlice, startMolecule string, finalMolecule string) int {
	fmt.Println(startMolecule)
	for _, pattern := range sortedMachineKeys {
		count := strings.Count(startMolecule, pattern)
		if count == 0 {
			continue
		}
		tempMolecule := strings.Replace(startMolecule, pattern, machine[pattern], -1)
		if tempMolecule == finalMolecule {
			// match found
			return count
		}
		recursiveCount := depthFirstSearchCore(machine, sortedMachineKeys, tempMolecule, finalMolecule)
		if -1 != recursiveCount {
			return count + recursiveCount
		}
	}
	// finalMolecule not found
	return -1
}

type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	if len(p[i]) == len(p[j]) {
		return p[i] > p[j]
	}
	return len(p[i]) > len(p[j])
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
