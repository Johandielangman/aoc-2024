// ~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~
//      /\_/\
//     ( o.o )
//      > ^ <
//
// Author: Johan Hanekom
// Date: December 2024
//
// ~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/johandielangman/aoc-2024/solutions/day_1"
	"github.com/johandielangman/aoc-2024/solutions/day_2"
	"github.com/johandielangman/aoc-2024/solutions/day_3"
	"github.com/johandielangman/aoc-2024/solutions/day_4"
)

func getInputDir(day string, fileName string) string {
	cwdString, err := os.Getwd()
	if err != nil {
		log.Fatal("could not get current working directory")
	}

	return filepath.Join(cwdString, "inputs", day, fileName)
}

func main() {

	fmt.Println("^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~")
	fmt.Println("	     Welcome to AOC 2024!")
	fmt.Print("^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~^~\n\n")
	fmt.Println("================ // DAY 1 // ================")
	day1_1, day1_2 := day_1.Solution(getInputDir("day_1", "input.txt"))
	fmt.Println("Solution for Day 1.1:", day1_1, "⭐")
	fmt.Println("Solution for Day 1.2:", day1_2, "⭐")
	fmt.Println("================ // DAY 2 // ================")
	day2_1, day2_2 := day_2.Solution(getInputDir("day_2", "input.txt"))
	fmt.Println("Solution for Day 2.1:", day2_1, "⭐")
	fmt.Println("Solution for Day 2.2:", day2_2, "⭐")
	fmt.Println("================ // DAY 3 // ================")
	day3_1, day3_2 := day_3.Solution(getInputDir("day_3", "input.txt"))
	fmt.Println("Solution for Day 3.1:", day3_1, "⭐")
	fmt.Println("Solution for Day 3.2:", day3_2, "⭐")
	fmt.Println("================ // DAY 4 // ================")
	day4_1, day4_2 := day_4.Solution(getInputDir("day_4", "input.txt"))
	fmt.Println("Solution for Day 4.1:", day4_1, "⭐")
	fmt.Println("Solution for Day 4.2:", day4_2, "⭐")
}
