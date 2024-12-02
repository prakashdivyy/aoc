package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")

	solvePartOne(lines)
}

func solvePartOne(lines []string) {
	total := 0
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		inc := false
		dec := false
		invalid := false

		for i := 0; i < len(parts); i++ {
			if i < len(parts)-1 {
				num1, _ := strconv.Atoi(parts[i])
				num2, _ := strconv.Atoi(parts[i+1])
				if i == 0 {
					inc = num1 < num2
					dec = num1 > num2
				} else {
					if inc && num1 > num2 {
						invalid = true
						break
					} else if dec && num1 < num2 {
						invalid = true
						break
					}
				}
				if math.Abs(float64(num1-num2)) > 3 || math.Abs(float64(num1-num2)) == 0 {
					invalid = true
					break
				}
			}
		}

		if !invalid {
			total++
		}
	}

	fmt.Printf("[Part One] Total: %d\n", total)
}
