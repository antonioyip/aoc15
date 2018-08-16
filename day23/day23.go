package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Registers struct {
	a int
	b int
}

type Instruction struct {
	command string
	param1  string
	param2  string
}

func main() {
	inputs, err := ioutil.ReadFile("day23.input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputs), "\n")

	instructions := make([]Instruction, len(lines))
	for i, line := range lines {
		var command, param1, param2 string
		fmt.Sscanf(line, "%s %s %s", &command, &param1, &param2)
		param1 = strings.Replace(param1, ",", "", -1)
		instructions[i] = Instruction{command, param1, param2}
	}

	registers := Registers{0, 0}
	for i := 0; i < len(instructions); {
		i += processInstruction(instructions[i], &registers)
	}
	fmt.Println(registers.b)
	registers = Registers{1, 0}
	for i := 0; i < len(instructions); {
		i += processInstruction(instructions[i], &registers)
	}
	fmt.Println(registers.b)
}

// returns next command offset
func processInstruction(ins Instruction, reg *Registers) int {
	//fmt.Println(reg, ins)
	switch ins.command {
	case "hlf":
		return half(ins.param1, reg)
	case "tpl":
		return triple(ins.param1, reg)
	case "inc":
		return increment(ins.param1, reg)
	case "jmp":
		return jump(ins.param1, reg)
	case "jie":
		return jumpIfEven(ins.param1, ins.param2, reg)
	case "jio":
		return jumpIfOne(ins.param1, ins.param2, reg)
	default:
		panic("Invalid instruction")
	}
}

func half(param string, reg *Registers) int {
	switch param {
	case "a":
		reg.a /= 2
	case "b":
		reg.b /= 2
	default:
		panic("Invalid register")
	}
	return 1
}

func triple(param string, reg *Registers) int {
	switch param {
	case "a":
		reg.a *= 3
	case "b":
		reg.b *= 3
	default:
		panic("Invalid register")
	}
	return 1
}

func increment(param string, reg *Registers) int {
	switch param {
	case "a":
		reg.a++
	case "b":
		reg.b++
	default:
		panic("Invalid register")
	}
	return 1
}

func jump(param string, reg *Registers) int {
	offset, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	return offset
}

func jumpIfEven(param1, param2 string, reg *Registers) int {
	jump := false
	switch param1 {
	case "a":
		jump = (reg.a%2 == 0)
	case "b":
		jump = (reg.b%2 == 0)
	default:
		panic("Invalid register")
	}

	if jump {
		offset, err := strconv.Atoi(param2)
		if err != nil {
			panic(err)
		}
		return offset
	}
	return 1
}

func jumpIfOne(param1, param2 string, reg *Registers) int {
	jump := false
	switch param1 {
	case "a":
		jump = reg.a == 1
	case "b":
		jump = reg.b == 1
	default:
		panic("Invalid register")
	}

	if jump {
		offset, err := strconv.Atoi(param2)
		if err != nil {
			panic(err)
		}
		return offset
	}
	return 1
}
