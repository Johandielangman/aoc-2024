# Some cool solutions I found on the internet
## Day 2

```go
package day_2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isSafe(row []int) bool {
	if len(row) <= 1 {
		return false
	}

	increments := make([]int, len(row)-1)
	for i := 0; i < len(row)-1; i++ {
		increments[i] = row[i+1] - row[i]
	}

	allPositive := true
	allNegative := true

	for _, inc := range increments {
		if inc > 3 || inc < -3 || inc == 0 {
			return false
		}
		if inc > 0 {
			allNegative = false
		}
		if inc < 0 {
			allPositive = false
		}
	}

	return allPositive || allNegative
}

func Solution(inputPath string) (partOneSolution, partTwoSolution int) {
	file, err := os.Open(inputPath)
	if err != nil {
		fmt.Println("Could not read input file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data [][]int

	for scanner.Scan() {
		rowStr := strings.Split(scanner.Text(), " ")
		row := make([]int, len(rowStr))
		for i, numStr := range rowStr {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Could not convert number:", err)
				return
			}
			row[i] = num
		}
		data = append(data, row)
	}

	// Part One: Count rows that are safe as-is
	for _, row := range data {
		if isSafe(row) {
			partOneSolution++
		}
	}

	// Part Two: Count rows that become safe by removing one number
	for _, row := range data {
		rowSafe := false
		for i := range row {
			// Create a new row with the i-th element removed
			reducedRow := make([]int, len(row)-1)
			copy(reducedRow[:i], row[:i])
			copy(reducedRow[i:], row[i+1:])

			if isSafe(reducedRow) {
				rowSafe = true
				break
			}
		}
		if rowSafe {
			partTwoSolution++
		}
	}

	return
}
```

```python
def is_safe(row):
    inc = [row[i + 1] - row[i] for i in range(len(row) - 1)]
    if set(inc) <= {1, 2, 3} or set(inc) <= {-1, -2, -3}:
        return True
    return False

data = [[int(y) for y in x.split(' ')] for x in open('input.txt').read().split('\n')]

safe_count = sum([is_safe(row) for row in data])
print(safe_count)

safe_count = sum([any([is_safe(row[:i] + row[i + 1:]) for i in range(len(row))]) for row in data])
print(safe_count)
```