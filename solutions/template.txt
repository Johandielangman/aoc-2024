package day_4

import "fmt"

func partOne(inputPath string) (int, error) {
	fmt.Println(inputPath)
	return 0, nil
}

func partTwo(inputPath string) (int, error) {
	fmt.Println(inputPath)
	return 0, nil
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
