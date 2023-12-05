package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	r, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(r)

	winners := make([]int, 0, 188)
	for scanner.Scan() {
		winners = append(winners, len(winningNumbers(scanner.Text())))
	}

	var pointSum, cardSum int
	for i := range winners {
		pointSum += calculatePoints(winners[i])
		// this card + this card's winners + this card's winners' winners
		cardSum += 1 + winners[i] + countCards(winners[i+1:], winners[i])
	}

	fmt.Println(pointSum)
	fmt.Println(cardSum)
}

// winningNumbers uselessly tracks the winning numbers matched of a card. in my defense I really thought we'd use em
// for part 2.
func winningNumbers(card string) []int {
	sets := strings.Split(strings.Split(card, ":")[1], "|")
	winners := sets[0]
	hand := sets[1]

	var nums []int

	for _, w := range strings.Split(strings.TrimSpace(winners), " ") {
		for _, h := range strings.Split(strings.TrimSpace(hand), " ") {
			if w == h {
				if n, err := strconv.Atoi(h); err == nil {
					nums = append(nums, n)
				}
			}
		}
	}

	return nums
}

// countCards recursively counts the next n winning cards and its baby's mamas' mamas. mamas, mamas, baby mamas' mamas
func countCards(winners []int, n int) int {
	if len(winners) == 0 || n == 0 {
		return 0
	}

	var cards int
	for i := 0; i < n; i++ {
		cards += winners[i] + countCards(winners[i+1:], winners[i])
	}

	return cards
}

func calculatePoints(winners int) int {
	return int(math.Pow(2.0, float64(winners-1)))
}
