package services

import (
	"fmt"
	"math"
)

func Encoded(pattern string) string {
	results := []string{}

	// find 6-digit from "000000" to "999999"
	for number := 0; number <= 999999; number++ {
		numStr := fmt.Sprintf("%06d", number) // Convert to 6-digit string
		if CheckPattern(numStr, pattern) {
			results = append(results, numStr)
		}
	}

	minSum := math.MaxInt32
	minNumber := ""

	for _, num := range results {
		sum := 0
		for _, digit := range num {
			sum += int(digit - '0')
		}
		if sum < minSum {
			minSum = sum
			minNumber = num
		}
	}

	return minNumber
}

func CheckPattern(numStr string, pattern string) bool {
	for i := 0; i < len(pattern); i++ {
		switch pattern[i] {
		case 'L':
			if numStr[i] <= numStr[i+1] {
				return false
			}
		case 'R':
			if numStr[i] >= numStr[i+1] {
				return false
			}
		case '=':
			if numStr[i] != numStr[i+1] {
				return false
			}
		}
	}
	return true
}
