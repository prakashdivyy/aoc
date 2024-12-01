package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(input), "\n")
	solvePartOne(lines)
	solvePartTwo(lines)
}

func solvePartOne(l []string) {
	total := 0

	for i := 0; i < len(l); i++ {
		number := ""
		for _, char := range l[i] {
			if char >= '0' && char <= '9' {
				number += string(char)
			}
		}
		if len(number) > 0 {
			firstDigit := string(number[0])
			lastDigit := string(number[len(number)-1])
			value := firstDigit + lastDigit
			num := 0
			fmt.Sscanf(value, "%d", &num)
			total += num
		}
	}

	fmt.Printf("[Part One] Total: %d\n", total)
}

func solvePartTwo(l []string) {
	total := 0

	for i := 0; i < len(l); i++ {
		number := ""
		line := l[i]

		for j := 0; j < len(line); j++ {
			if line[j] >= '0' && line[j] <= '9' {
				number += string(line[j])
			} else {
				if strings.HasPrefix(line[j:], "one") {
					number += "1"
				} else if strings.HasPrefix(line[j:], "two") {
					number += "2"
				} else if strings.HasPrefix(line[j:], "three") {
					number += "3"
				} else if strings.HasPrefix(line[j:], "four") {
					number += "4"
				} else if strings.HasPrefix(line[j:], "five") {
					number += "5"
				} else if strings.HasPrefix(line[j:], "six") {
					number += "6"
				} else if strings.HasPrefix(line[j:], "seven") {
					number += "7"
				} else if strings.HasPrefix(line[j:], "eight") {
					number += "8"
				} else if strings.HasPrefix(line[j:], "nine") {
					number += "9"
				}
			}
		}

		if len(number) > 0 {
			firstDigit := string(number[0])
			lastDigit := string(number[len(number)-1])
			value := firstDigit + lastDigit
			num := 0
			fmt.Sscanf(value, "%d", &num)
			total += num
		}
	}

	fmt.Printf("[Part Two] Total: %d\n", total)
}
