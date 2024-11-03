package services

import (
	"encoding/json"
	"log"
	"os"
)

func MaxPath(inputPath string) int {
	input := GetInput(inputPath)
	total := MaxPathSum(input)
	return total
}

func GetInput(filePath string) [][]int {
	input, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	var intSlice [][]int
	err = json.Unmarshal(input, &intSlice)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}
	return intSlice
}

func MaxPathSum(input [][]int) int {
	for i := range input[:len(input)-1] {
		// คำนวณจากชั้นล่างขึ้นบน
		level := len(input) - 2 - i
		for j := range input[level] {
			// คำนวณ max sum ของแต่ละตำแหน่ง โดยเลือกค่ามากสุดจากเส้นทางด้านล่างที่เชื่อมถึงกัน
			input[level][j] += findMax(input[level+1][j], input[level+1][j+1])
		}
	}
	return input[0][0]
}

func findMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
