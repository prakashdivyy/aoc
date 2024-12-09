package main

import (
	"fmt"
	"os"
	"regexp"
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
		muls := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`).FindAllString(line, -1)
		for _, mul := range muls {
			re := regexp.MustCompile(`\d+`)
			matches := re.FindAllString(mul, -1)
			if len(matches) == 2 {
				num1, _ := strconv.Atoi(matches[0])
				num2, _ := strconv.Atoi(matches[1])
				total += num1 * num2
			}
		}
	}
	fmt.Println(total)
}
