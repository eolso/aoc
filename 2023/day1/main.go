package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	r, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(r)

	sum := 0
	for scanner.Scan() {
		switch digits := convertWords(scanner.Text()); len(digits) {
		case 0:
			continue
		case 1:
			sum += digits[0]*10 + digits[0]
			fmt.Printf("%s -> %v -> %d%d\n", scanner.Text(), digits, digits[0], digits[0])
		default:
			sum += digits[0]*10 + digits[len(digits)-1]
			fmt.Printf("%s -> %v -> %d%d\n", scanner.Text(), digits, digits[0], digits[len(digits)-1])
		}
	}

	fmt.Printf("Sum: %d\n", sum)
}

func convertWords(s string) []int {
	digitMapping := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	type tuple struct {
		index int
		value int
	}

	s = strings.ToLower(s)

	// Lazy ass searching, make sure to sort later
	var tuples []tuple

	// Grab the digits
	for i := 0; i < 10; i++ {
		buffer := s
		a := strconv.Itoa(i)

		for {
			if n := strings.Index(buffer, a); n != -1 {
				tuples = append(tuples, tuple{n, i})
				buffer = strings.Replace(buffer, a, "-", 1)
			} else {
				break
			}
		}
	}

	// Grab the words
	for k, v := range digitMapping {
		buffer := s

		for {
			if n := strings.Index(buffer, k); n != -1 {
				tuples = append(tuples, tuple{n, v})
				buffer = strings.Replace(buffer, k, strings.Repeat("-", len(k)), 1)
			} else {
				break
			}
		}
	}

	sort.SliceStable(tuples, func(i, j int) bool {
		return tuples[i].index < tuples[j].index
	})

	values := make([]int, len(tuples))
	for i, tuple := range tuples {
		values[i] = tuple.value
	}

	return values
}
