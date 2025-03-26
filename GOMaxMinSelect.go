package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func maxMinSelect(array []int) (int, int) {
	if len(array) == 0 {
		fmt.Println("Error: Array is empty")
		os.Exit(1)
	}
	return maxMinHelper(array, 0, len(array)-1)
}

func maxMinHelper(array []int, left, right int) (int, int) {
	// Base case: Single element
	if left == right {
		return array[left], array[left]
	}

	// Base case: Two elements
	if left == right-1 {
		if array[left] > array[right] {
			return array[left], array[right]
		}
		return array[right], array[left]
	}

	// Recursively split the array
	mid := (left + right) / 2
	maxLeft, minLeft := maxMinHelper(array, left, mid)
	maxRight, minRight := maxMinHelper(array, mid+1, right)

	// Combine results
	return max(maxLeft, maxRight), min(minLeft, minRight)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Insert the numbers, with a space between each: ")
	scanner.Scan()
	input := scanner.Text()

	// Split the input and convert to integers
	parts := strings.Fields(input)
	numbers := make([]int, 0, len(parts))

	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Println("Error: Invalid input. Please enter valid integers.")
			return
		}
		numbers = append(numbers, num)
	}

	// Find and print max and min
	maxValue, minValue := maxMinSelect(numbers)

	fmt.Println("Maximum value:", maxValue)
	fmt.Println("Minimum value:", minValue)
}