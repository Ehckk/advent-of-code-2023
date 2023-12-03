// Advent of Code 2023
// Day 02 - Part 1
// 12-02-23
package main

import (
	"fmt"
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
		id, _ := strconv.Atoi(strings.Split(game[0], " ")[1])

		var possible bool = true
		for _, subset := range strings.Split(game[1], "; ") {
			if !possible {
				break
			}

			limits := map[string]int{"red": 12, "green": 13, "blue": 14}
			values := make(map[string]int)

			subset := strings.Split(subset, ", ")
			for _, item := range subset {
				cubes := strings.Split(item, " ")
				amount, _ := strconv.Atoi(cubes[0])
				color := cubes[1]
				values[color] = amount
			}

			for key, value := range values {
				if value > limits[key] {
					possible = false
				}
			}
		}

		if possible {
			total += id
		}

	}
	fmt.Println(total)
}
