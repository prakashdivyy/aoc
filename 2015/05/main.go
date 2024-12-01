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
}

func solvePartOne(l []string) {
	total := 0
	for _, v := range l {
		if v == "" {
			continue
		}
		if isContainsVowel(v) && isDouble(v) && !isBad(v) {
			total++
		}
	}
	fmt.Printf("[Part One] Total: %d\n", total)
}

func isContainsVowel(s string) bool {
	vowels := 0
	for _, v := range s {
		if v == 'a' || v == 'e' || v == 'i' || v == 'o' || v == 'u' {
			vowels++
		}
	}
	return vowels >= 3
}

func isDouble(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func isBad(s string) bool {
	bad := []string{"ab", "cd", "pq", "xy"}
	for _, v := range bad {
		if strings.Contains(s, v) {
			return true
		}
	}
	return false
}
