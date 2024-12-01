package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	var dimensions [][]int

	for _, line := range lines {
		if line == "" {
			continue
		}
		numbers := strings.Split(line, "x")
		arrNumbers := make([]int, len(numbers))
		for i, number := range numbers {
			arrNumbers[i], _ = strconv.Atoi(number)
		}
		dimensions = append(dimensions, arrNumbers)
	}

	solvePartOne(dimensions)
	solvePartTwo(dimensions)
}

func solvePartOne(l [][]int) {
	total := 0
	for _, dim := range l {
		side1 := dim[0] * dim[1]
		side2 := dim[0] * dim[2]
		side3 := dim[1] * dim[2]
		min := side1
		if side2 < min {
			min = side2
		}
		if side3 < min {
			min = side3
		}
		total += 2*side1 + 2*side2 + 2*side3 + min
	}
	fmt.Printf("[Part One] Total: %d\n", total)
}

func solvePartTwo(l [][]int) {
	total := 0
	for _, dim := range l {
		sort.Ints(dim)
		bow := dim[0] * dim[1] * dim[2]
		ribbon := 2 * (dim[0] + dim[1])
		total += bow + ribbon
	}
	fmt.Printf("[Part Two] Total: %d\n", total)
}
