package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	line := string(input)
	solvePartOne(line)
	solvePartTwo(line)
}

func solvePartOne(l string) {
	floor := 0
	for i := 0; i < len(l); i++ {
		if l[i] == '(' {
			floor++
		} else if l[i] == ')' {
			floor--
		}
	}

	fmt.Printf("[Part One] Total: %d\n", floor)
}

func solvePartTwo(l string) {
	floor := 0
	index := 1
	for i := 0; i < len(l); i++ {
		if l[i] == '(' {
			floor++
		} else if l[i] == ')' {
			floor--
		}
		if floor == -1 {
			break
		}
		index += 1
	}

	fmt.Printf("[Part Two] Position: %d\n", index)
}
