package day_3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func partOne(inputPath string) (int, error) {
	mulPattern := `(?m)mul\((?<digitFirst>\d{1,3}),(?<digitSecond>\d{1,3})\)`
	mulRegex := regexp.MustCompile(mulPattern)
	fileContent, err := os.ReadFile(inputPath)
	if err != nil {
		return 0, err
	}
	text := string(fileContent)
	results := mulRegex.FindAllStringSubmatch(text, -1)
	var totalMul int
	for i := range results {
		// mulMatch := results[i][0]
		digitOne, err := strconv.Atoi(results[i][1])
		if err != nil {
			return 0, err
		}
		digitTwo, err := strconv.Atoi(results[i][2])
		if err != nil {
			return 0, err
		}
		totalMul += (digitOne * digitTwo)
	}
	return totalMul, nil
}

func partTwo(inputPath string) (int, error) {
	linePattern := `mul\((?<digitFirst>\d{1,3}),(?<digitSecond>\d{1,3})\)|(?<toggle>do|don't)\(\)`
	lineRegex := regexp.MustCompile(linePattern)
	fileContent, err := os.ReadFile(inputPath)
	if err != nil {
		return 0, err
	}

	text := string(fileContent)
	totalMul := 0
	enabled := true

	results := lineRegex.FindAllStringSubmatch(text, -1)
	for i := range results {
		digitOneStr := results[i][1]
		digitTwoStr := results[i][2]
		doOrDont := results[i][3]

		digitOne, _ := strconv.Atoi(digitOneStr)
		digitTwo, _ := strconv.Atoi(digitTwoStr)

		switch doOrDont {
		case "do":
			enabled = true
		case "don't":
			enabled = false
		}
		if enabled {
			totalMul += (digitOne * digitTwo)
		}
	}

	return totalMul, nil
}

func Solution(inputPath string) (partOneSolution, partTwoSolution int) {
	partOneSolution, err := partOne(inputPath)
	if err != nil {
		fmt.Println(err)
	}

	partTwoSolution, err = partTwo(inputPath)
	if err != nil {
		fmt.Println(err)
	}

	return
}
