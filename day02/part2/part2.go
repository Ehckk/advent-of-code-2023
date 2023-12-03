// Advent of Code 2023
// Day 02 - Part 2
// 12-02-23
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.ReadFile("../input.txt")
	text := string(file)
	lines := strings.Split(text, "\n")

	total := 0
	for _, line := range lines {
		game := strings.Split(strings.TrimSpace(line), ": ")
		values := map[string]int{"red": 0, "green": 0, "blue": 0}
		for _, subset := range strings.Split(game[1], "; ") {
			subset := strings.Split(subset, ", ")
			for _, item := range subset {
				cubes := strings.Split(item, " ")
				amount, _ := strconv.Atoi(cubes[0])
				color := cubes[1]
				values[color] = int(math.Max(float64(values[color]), float64(amount)))
			}
		}
		product := 1
		for _, value := range values {
			product *= value
		}
		total += product
	}
	fmt.Println(total)
}
