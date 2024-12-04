package day_4

import (
	"fmt"
	"os"
	"strings"
)

func countDiagonals(m [][]string, w string) (cnt int) {
	// Top-left to bottom-right
	diagonalOneParts := make([]string, len(w))
	for i := 0; i < len(w); i++ {
		diagonalOneParts[i] = m[i][i]
	}
	diagonalOne := strings.Join(diagonalOneParts, "")

	if diagonalOne == w {
		cnt++
	}

	// Bottom-left to top-right
	diagonalTwoParts := make([]string, len(w))
	for i := 0; i < len(w); i++ {
		diagonalTwoParts[i] = m[len(w)-1-i][i]
	}
	diagonalTwo := strings.Join(diagonalTwoParts, "")

	if diagonalTwo == w {
		cnt++
	}

	return
}

func countXMas(m [][]string, w string) (cnt int) {
	// Top-left to bottom-right
	diagonalOneParts := make([]string, len(w))
	for i := 0; i < len(w); i++ {
		diagonalOneParts[i] = m[i][i]
	}
	diagonalOne := strings.Join(diagonalOneParts, "")

	// Bottom-left to top-right
	diagonalTwoParts := make([]string, len(w))
	for i := 0; i < len(w); i++ {
		diagonalTwoParts[i] = m[len(w)-1-i][i]
	}
	diagonalTwo := strings.Join(diagonalTwoParts, "")

	if (diagonalOne == w || diagonalOne == Reverse(w)) && (diagonalTwo == w || diagonalTwo == Reverse(w)) {
		cnt++
	}

	return
}

func countHorizontals(m [][]string, w string) (cnt int) {
	wordRows := make([]string, len(m))
	for i, array := range m {
		wordRows[i] = strings.Join(array, "")
	}
	for _, word := range wordRows {
		if word == w {
			cnt++
		}
	}
	return
}

func countVerticals(m [][]string, w string) (cnt int) {
	wordCols := make([]string, len(w))
	for j := 0; j < len(w); j++ {
		wordParts := make([]string, len(w))
		for i := 0; i < len(w); i++ {
			wordParts[i] = m[i][j]
		}
		wordCols[j] = strings.Join(wordParts, "")
	}
	for _, word := range wordCols {
		if word == w {
			cnt++
		}
	}
	return
}

func textToMatrix(fileContent string) (m [][]string, nRows int, nCols int) {
	textRows := strings.Split(fileContent, "\n")
	nRows = len(textRows)
	matrix := make([][]string, nRows)
	for i := range matrix {
		matrix[i] = make([]string, len(textRows[i]))
		textCols := strings.Split(textRows[i], "")
		nCols = len(textCols)
		copy(matrix[i], textCols)
	}
	m = matrix
	return
}

func getSubMatrix(m [][]string, startRow, endRow, startCol, endCol int) [][]string {
	subMatrix := make([][]string, endRow-startRow)
	for i := range subMatrix {
		subMatrix[i] = make([]string, endCol-startCol)
		copy(subMatrix[i], m[startRow+i][startCol:endCol])
	}
	return subMatrix
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func GrowMatrix(m [][]string, nCols, nRows, wordSize int) [][]string {
	mod := nCols % wordSize
	if mod != 0 {
		for i := range m {
			for j := nCols; j < nRows+mod; j++ {
				m[i] = append(m[i], ".")
			}
		}
		for i := nRows; i < nCols+mod; i++ {
			newRow := make([]string, nCols+mod)
			for j := range newRow {
				newRow[j] = "."
			}
			m = append(m, newRow)
		}
	}
	return m
}

func sq(start, end, step int) []int {
	var result []int
	for i := start; i <= end; i += step {
		result = append(result, i)
	}
	return result
}

func partOne(inputPath string) (int, error) {
	SearchWord := "XMAS"
	nCharWord := len(SearchWord)
	fileContent, err := os.ReadFile(inputPath)
	if err != nil {
		return 0, nil
	}
	m, nRows, nCols := textToMatrix(string(fileContent))
	if nRows != nCols {
		panic("Not the same!")
	}
	m = GrowMatrix(m, nCols, nRows, nCharWord)
	nRows = len(m)
	nCols = len(m)

	var cnt int
	for i := 0; i <= nRows-nCharWord; i++ {
		for j := 0; j <= nCols-nCharWord; j++ {
			mSub := getSubMatrix(m, i, i+nCharWord, j, j+nCharWord)
			cnt += countDiagonals(mSub, SearchWord)
			cnt += countDiagonals(mSub, Reverse(SearchWord))
		}
	}
	for i := 0; i <= nRows-nCharWord; i++ {
		for _, j := range sq(0, nCols-nCharWord, nCharWord) {
			mSub := getSubMatrix(m, i, i+nCharWord, j, j+nCharWord)
			cnt += countVerticals(mSub, SearchWord)
			cnt += countVerticals(mSub, Reverse(SearchWord))
		}
	}
	for _, i := range sq(0, nRows-nCharWord, nCharWord) {
		for j := 0; j <= nCols-nCharWord; j++ {
			mSub := getSubMatrix(m, i, i+nCharWord, j, j+nCharWord)
			cnt += countHorizontals(mSub, SearchWord)
			cnt += countHorizontals(mSub, Reverse(SearchWord))
		}
	}
	return cnt, nil
}

func partTwo(inputPath string) (int, error) {
	SearchWord := "MAS"
	nCharWord := len(SearchWord)
	fileContent, err := os.ReadFile(inputPath)
	if err != nil {
		return 0, nil
	}
	m, nRows, nCols := textToMatrix(string(fileContent))
	if nRows != nCols {
		panic("Not the same!")
	}
	m = GrowMatrix(m, nCols, nRows, nCharWord)
	nRows = len(m)
	nCols = len(m)

	var cnt int
	for i := 0; i <= nRows-nCharWord; i++ {
		for j := 0; j <= nCols-nCharWord; j++ {
			mSub := getSubMatrix(m, i, i+nCharWord, j, j+nCharWord)
			cnt += countXMas(mSub, SearchWord)
		}
	}
	return cnt, nil
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
