package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(r)

	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	possibleSum := 0
	powerSum := 0
	for scanner.Scan() {
		var gameId int
		_, _ = fmt.Sscanf(scanner.Text(), "Game %d:", &gameId)
		gameData := strings.Split(scanner.Text(), ":")[1]

		if isPossible(limits, gameData) {
			possibleSum += gameId
		}

		powerSum += minimumPower(gameData)
	}

	fmt.Printf("1: %d\n", possibleSum)
	fmt.Printf("2: %d\n", powerSum)
}

func isPossible(limits map[string]int, gameData string) bool {
	for _, game := range strings.Split(gameData, ";") {
		for _, set := range strings.Split(strings.TrimSpace(game), ",") {
			var color string
			var amount int

			_, _ = fmt.Sscanf(strings.TrimSpace(set), "%d %s", &amount, &color)

			if limits[color] < amount {
				return false
			}
		}
	}

	return true
}

func minimumPower(gameData string) int {
	maxColors := map[string]int{
		"red":   1,
		"green": 1,
		"blue":  1,
	}

	for _, game := range strings.Split(gameData, ";") {
		for _, set := range strings.Split(strings.TrimSpace(game), ",") {
			var color string
			var amount int

			_, _ = fmt.Sscanf(strings.TrimSpace(set), "%d %s", &amount, &color)

			if amount > maxColors[color] {
				maxColors[color] = amount
			}
		}
	}

	power := 1
	for _, v := range maxColors {
		power *= v
	}

	return power
}
