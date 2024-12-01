package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "iwrupvqb"
	solvePartOne(input)
	solvePartTwo(input)
}

func solvePartOne(input string) {
	number := 0
	for {
		hash := md5.Sum([]byte(input + strconv.Itoa(number)))
		hashString := fmt.Sprintf("%x", hash)
		if strings.HasPrefix(hashString, "00000") {
			fmt.Printf("[Part One] Number: %d\n", number)
			break
		}
		number++
	}
}

func solvePartTwo(input string) {
	number := 0
	for {
		hash := md5.Sum([]byte(input + strconv.Itoa(number)))
		hashString := fmt.Sprintf("%x", hash)
		if strings.HasPrefix(hashString, "000000") {
			fmt.Printf("[Part Two] Number: %d\n", number)
			break
		}
		number++
	}
}
