package day_2

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func isIncreasingOrDecreasing(numbers []int) (bool, int) {
	if len(numbers) <= 1 {
		fmt.Println("Not enough numbers!")
		return false, 0
	}

	var jumpSigned float64
	var jumpUnSigned float64
	for i := 1; i < len(numbers); i++ {
		jump := float64(numbers[i]) - float64(numbers[i-1])
		jumpSigned += jump
		jumpUnSigned += math.Abs(jump)

		if math.Abs(jumpSigned) != jumpUnSigned {
			return false, i
		}
		if jump == 0.0 {
			return false, i
		}
		if math.Abs(jump) > 3.0 {
			return false, i
		}
	}
	return true, len(numbers)
}

func isReportSafe(reportNumbers []string) (bool, error) {
	numbers, err := convertToInts(reportNumbers)
	if err != nil {
		return false, err
	}

	isSafe, _ := isIncreasingOrDecreasing(numbers)
	return isSafe, nil
}

func remove(slice []int, s int) []int {
	// DON'T Always listen to stackoveflow!
	// https: //stackoverflow.com/questions/37334119/how-to-delete-an-element-from-a-slice-in-golang
	result := make([]int, 0, len(slice)-1)
	result = append(result, slice[:s]...)
	result = append(result, slice[s+1:]...)
	return result
}

func isReportSafeWithDampener(reportNumbers []string) (bool, error) {
	numbers, err := convertToInts(reportNumbers)
	if err != nil {
		return false, err
	}

	isSafe, _ := isIncreasingOrDecreasing(numbers)
	if isSafe {
		return true, nil
	}

	for i := 0; i < len(numbers); i++ {
		reducedNumbers := remove(numbers, i)
		isSafe, _ = isIncreasingOrDecreasing(reducedNumbers)
		if isSafe {
			return true, nil
		}
	}

	return false, nil
}

func convertToInts(reportNumbers []string) ([]int, error) {
	numbers := make([]int, len(reportNumbers))
	for i, numStr := range reportNumbers {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, fmt.Errorf("could not convert %v to an int due to err %v", numStr, err)
		}
		numbers[i] = num
	}
	return numbers, nil
}

func partOne(inputPath string) (int, error) {
	file, err := os.Open(inputPath)
	var safeCount int

	if err != nil {
		return 0, fmt.Errorf("could not read input file %q due to err %v", inputPath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		reportNumbers := strings.Split(scanner.Text(), " ")
		isSafe, err := isReportSafe(reportNumbers)
		if err != nil {
			return 0, fmt.Errorf("could not determine if the report is safe due to err %v", err)
		}
		if isSafe {
			safeCount++
		}
	}

	return safeCount, nil
}

func partTwo(inputPath string) (int, error) {
	file, err := os.Open(inputPath)
	var safeCount int

	if err != nil {
		return 0, fmt.Errorf("could not read input file %q due to err %v", inputPath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		reportNumbers := strings.Split(scanner.Text(), " ")
		isSafe, err := isReportSafeWithDampener(reportNumbers)
		if err != nil {
			return 0, fmt.Errorf("could not determine if the report is safe due to err %v", err)
		}
		if isSafe {
			safeCount++
		}
	}

	return safeCount, nil
}

func Solution(inputPath string) (partOneSolution, partTwoSolution int) {
	partOneSolution, err := partOne(inputPath)
	if err != nil {
		fmt.Println("Could not calculate solution for part one due to error:", err)
	}

	partTwoSolution, err = partTwo(inputPath)
	if err != nil {
		fmt.Println("Could not calculate solution for part two due to error:", err)
	}

	return
}
