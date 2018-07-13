package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
}

func main() {
	testRecipeScore()
	testGenerateCombinations()

	inputs, err := ioutil.ReadFile("day15.input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(inputs), "\n")

	components := []ingredient{}
	for _, line := range lines {
		var ing ingredient
		fmt.Sscanf(line, "%s capacity %d, durability %d, flavor %d, texture %d, calories %d", &ing.name, &ing.capacity, &ing.durability, &ing.flavor, &ing.texture, &ing.calories)
		components = append(components, ing)
	}

	combinations := generateCombinations(100, len(components), []int{})
	maxScore := 0
	for _, ratios := range combinations {
		maxScore = max(maxScore, recipeScore(components, ratios))
	}
	fmt.Println(maxScore)

	maxScore = 0
	for _, ratios := range combinations {
		if calculateCalories(components, ratios) == 500 {
			maxScore = max(maxScore, recipeScore(components, ratios))
		}
	}
	fmt.Println(maxScore)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sum(values []int) int {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum
}

func recipeScore(components []ingredient, ratios []int) int {
	capacity := 0
	durability := 0
	flavor := 0
	texture := 0
	for i := range ratios {
		capacity += ratios[i] * components[i].capacity
		durability += ratios[i] * components[i].durability
		flavor += ratios[i] * components[i].flavor
		texture += ratios[i] * components[i].texture
	}
	capacity = max(capacity, 0)
	durability = max(durability, 0)
	flavor = max(flavor, 0)
	texture = max(texture, 0)
	return capacity * durability * flavor * texture
}

func testRecipeScore() {
	a := ingredient{"Butterscotch", -1, -2, 6, 3, 8}
	b := ingredient{"Cinnamon", 2, 3, -2, -1, 3}
	components := []ingredient{a, b}
	fmt.Println(recipeScore(components, []int{44, 56}) == 62842880)
}

// generate all combinations of size numAddends that add up to sum
func generateCombinations(sum int, numAddends int, combination []int) [][]int {
	allCombinations := [][]int{}
	if numAddends == 1 {
		combination = append(combination, sum)
		allCombinations = append(allCombinations, combination)
	} else {
		for addend := 0; addend <= sum; addend++ {
			tempCombination := append(combination, addend)
			allCombinations = append(allCombinations, generateCombinations(sum-addend, numAddends-1, tempCombination)...)
		}
	}
	return allCombinations
}

func testGenerateCombinations() {
	fmt.Println(equals2(generateCombinations(1, 1, []int{}), [][]int{{1}}))
	fmt.Println(equals2(generateCombinations(2, 1, []int{}), [][]int{{2}}))
	fmt.Println(equals2(generateCombinations(3, 1, []int{}), [][]int{{3}}))
	fmt.Println(equals2(generateCombinations(1, 2, []int{}), [][]int{{0, 1}, {1, 0}}))
	fmt.Println(equals2(generateCombinations(2, 2, []int{}), [][]int{{0, 2}, {1, 1}, {2, 0}}))
	fmt.Println(equals2(generateCombinations(3, 2, []int{}), [][]int{{0, 3}, {1, 2}, {2, 1}, {3, 0}}))
	fmt.Println(equals2(generateCombinations(1, 3, []int{}), [][]int{{0, 0, 1}, {0, 1, 0}, {1, 0, 0}}))
	fmt.Println(equals2(generateCombinations(2, 3, []int{}), [][]int{{0, 0, 2}, {0, 1, 1}, {0, 2, 0}, {1, 0, 1}, {1, 1, 0}, {2, 0, 0}}))
	fmt.Println(equals2(generateCombinations(3, 3, []int{}), [][]int{{0, 0, 3}, {0, 1, 2}, {0, 2, 1}, {0, 3, 0}, {1, 0, 2}, {1, 1, 1}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}, {3, 0, 0}}))
}

func equals1(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func equals2(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for _, rowA := range a {
		match := false
		for _, rowB := range b {
			if equals1(rowA, rowB) {
				match = true
			}
		}
		if !match {
			return false
		}
	}
	return true
}

func calculateCalories(components []ingredient, ratios []int) int {
	total := 0
	for i := range ratios {
		total += components[i].calories * ratios[i]
	}
	return total
}
