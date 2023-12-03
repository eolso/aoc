package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	r, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(r)

	var runes [][]rune
	var i int
	for scanner.Scan() {
		runes = append(runes, make([]rune, len(scanner.Text())))
		for j, r := range scanner.Text() {
			runes[i][j] = r
		}
		i++
	}

	partNumbers := make([][]int, 0, len(runes))
	for i := range runes {
		for j := range runes[i] {
			// if we find a symbol, lucy got some splainin' to do
			if isSymbol(runes[i][j]) {
				var partNumber []int

				// splain left
				if n := readNumber(runes[i], j-1); n != 0 {
					partNumber = append(partNumber, n)
				}
				// splain right
				if n := readNumber(runes[i], j+1); n != 0 {
					partNumber = append(partNumber, n)
				}

				if i > 0 {
					// splain up left
					if n := readNumber(runes[i-1], j-1); n != 0 {
						partNumber = append(partNumber, n)
					}
					// splain up
					if n := readNumber(runes[i-1], j); n != 0 {
						partNumber = append(partNumber, n)
					}
					// splain up right
					if n := readNumber(runes[i-1], j+1); n != 0 {
						partNumber = append(partNumber, n)
					}
				}

				if i < len(runes)-1 {
					// splain down left
					if n := readNumber(runes[i+1], j-1); n != 0 {
						partNumber = append(partNumber, n)
					}
					// splain down
					if n := readNumber(runes[i+1], j); n != 0 {
						partNumber = append(partNumber, n)
					}
					// splain down right
					if n := readNumber(runes[i+1], j+1); n != 0 {
						partNumber = append(partNumber, n)
					}
				}

				if len(partNumber) > 0 {
					partNumbers = append(partNumbers, partNumber)
				}
			}
		}
	}

	sum := 0
	gearRatio := 0
	gearRatioSum := 0
	for _, line := range partNumbers {
		// more lazy ass logic
		if len(line) == 2 {
			gearRatio = 1
		} else {
			gearRatio = 0
		}

		for _, num := range line {
			sum += num
			gearRatio *= num
		}

		gearRatioSum += gearRatio
	}

	fmt.Println(sum)
	fmt.Println(gearRatioSum)
}

// readNumber will return the number that touches whatever is at index and then remove it from the line. returns 0 if
// no number at index.
func readNumber(line []rune, index int) int {
	// do some bounds checking here so we can yolo more
	if index < 0 || index >= len(line) {
		return 0
	}

	if !unicode.IsDigit(line[index]) {
		return 0
	}

	// shift the index left
	for index >= 0 && unicode.IsDigit(line[index]) {
		index--
	}
	index++

	// gobble em up
	var digits []string
	for index < len(line) && unicode.IsDigit(line[index]) {
		digits = append(digits, string(line[index]))
		// overwrite what we gobbled with "." to prevent re-gobbling
		line[index] = 46
		index++
	}

	// probably better way to convert but im lazy
	n, err := strconv.Atoi(strings.Join(digits, ""))
	if err != nil {
		return 0
	}

	return n
}

func isSymbol(r rune) bool {
	return string(r) != "." && !unicode.IsDigit(r)
}
