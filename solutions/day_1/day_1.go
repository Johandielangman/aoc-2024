package day_1

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func countLinesInTxt(file *os.File) (int, error) {
	lineCount := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCount++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return lineCount, nil
}

type historianFindings struct {
	historianFirst  []float64
	historianSecond []float64
}

func readHistorianFile(inputPath string) (historianFindings, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return historianFindings{}, fmt.Errorf("could not read input file %q due to err %v", inputPath, err)
	}
	defer file.Close()

	lineCount, err := countLinesInTxt(file)
	if err != nil {
		return historianFindings{}, fmt.Errorf("could not count number of lines in input file %q due to err %v", inputPath, err)
	}
	file.Seek(0, 0)
	// fmt.Printf("Reading in %d lines from %q\n", lineCount, inputPath)

	historianFirst := make([]float64, lineCount)
	historianSecond := make([]float64, lineCount)

	scanner := bufio.NewScanner(file)
	for i := 0; i < lineCount; i++ {
		if scanner.Scan() {
			line := scanner.Text()
			findingParts := strings.Split(line, "   ")

			numFirst, err := strconv.Atoi(findingParts[0])
			if err != nil {
				return historianFindings{}, fmt.Errorf("could not convert %v to int for historian fist on line %d", findingParts[0], i)
			}
			numSecond, err := strconv.Atoi(findingParts[1])
			if err != nil {
				return historianFindings{}, fmt.Errorf("could not convert %v to int for historian second on line %d", findingParts[1], i)
			}

			historianFirst[i] = float64(numFirst)
			historianSecond[i] = float64(numSecond)
		} else {
			return historianFindings{}, fmt.Errorf("could not read in line %d from input file %q", i, inputPath)
		}
	}
	if err := scanner.Err(); err != nil {
		return historianFindings{}, fmt.Errorf("could not read in line from input file %q due to error %v", inputPath, err)
	}

	return historianFindings{
		historianFirst:  historianFirst,
		historianSecond: historianSecond,
	}, nil
}

func partOne(historians historianFindings) int {
	lineCount := len(historians.historianFirst)
	sort.Float64s(historians.historianFirst)
	sort.Float64s(historians.historianSecond)

	var totalDistance float64
	for i := 0; i < lineCount; i++ {
		totalDistance += math.Abs(historians.historianFirst[i] - historians.historianSecond[i])
	}

	return int(totalDistance)
}

func partTwo(historians historianFindings) int {
	historianRightFrequency := make(map[int]int)
	for _, num := range historians.historianSecond {
		historianRightFrequency[int(num)] = historianRightFrequency[int(num)] + 1
	}
	var similarityScore float64
	for _, num := range historians.historianFirst {
		if freq, ok := historianRightFrequency[int(num)]; ok {
			similarityScore += float64(num) * float64(freq)
		}

	}
	return int(similarityScore)
}

func Solution(inputPath string) (partOneSolution int, partTwoSolution int) {
	historians, err := readHistorianFile(inputPath)
	if err != nil {
		fmt.Println("could not read in historian file:", err)
		return 0, 0
	}

	partOneSolution = partOne(historians)
	partTwoSolution = partTwo(historians)

	return
}
