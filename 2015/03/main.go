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
	houses := make(map[string]bool)

	x := 0
	y := 0
	key := fmt.Sprintf("%d,%d", x, y)
	houses[key] = true

	for _, c := range l {
		switch c {
		case '^':
			y++
		case 'v':
			y--
		case '>':
			x++
		case '<':
			x--
		}
		key := fmt.Sprintf("%d,%d", x, y)
		houses[key] = true
	}

	fmt.Printf("[Part One] Total: %d\n", len(houses))
}

func solvePartTwo(l string) {
	santa := ""
	robot := ""

	for i := 0; i < len(l); i++ {
		if i%2 == 0 {
			santa += string(l[i])
		} else {
			robot += string(l[i])
		}
	}

	houses := make(map[string]bool)

	x := 0
	y := 0
	key := fmt.Sprintf("%d,%d", x, y)
	houses[key] = true

	for _, c := range santa {
		switch c {
		case '^':
			y++
		case 'v':
			y--
		case '>':
			x++
		case '<':
			x--
		}
		key := fmt.Sprintf("%d,%d", x, y)
		houses[key] = true
	}

	x = 0
	y = 0
	for _, c := range robot {
		switch c {
		case '^':
			y++
		case 'v':
			y--
		case '>':
			x++
		case '<':
			x--
		}
		key := fmt.Sprintf("%d,%d", x, y)
		houses[key] = true
	}

	fmt.Printf("[Part Two] Total: %d\n", len(houses))
}
