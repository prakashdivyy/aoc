package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input")
	if err != nil {
		panic(err)
	}

	var leftArray []int
	var rightArray []int

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Fields(line)

		left, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		leftArray = append(leftArray, left)

		right, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		rightArray = append(rightArray, right)
	}

	solvePartOne(leftArray, rightArray)
	solvePartTwo(leftArray, rightArray)
}

func solvePartOne(l []int, r []int) {
	leftArray := sortArray(l)
	rightArray := sortArray(r)

	total := 0

	for i := 0; i < len(leftArray); i++ {
		diff := leftArray[i] - rightArray[i]
		if diff < 0 {
			diff = -diff
		}
		total += diff
	}

	fmt.Printf("[Part One] Total: %d\n", total)
}

func solvePartTwo(l []int, r []int) {
	leftArray := sortArray(l)
	rightArray := sortArray(r)

	total := 0

	rightMap := make(map[int]int)

	for i := 0; i < len(rightArray); i++ {
		rightMap[rightArray[i]]++
	}

	for i := 0; i < len(leftArray); i++ {
		if rightMap[leftArray[i]] > 0 {
			total += leftArray[i] * rightMap[leftArray[i]]
		}
	}

	fmt.Printf("[Part Two] Total: %d\n", total)
}

func sortArray(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
