// Advent of Code 2023
// Day 03 - Part 2
// 12-03-23
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func read(path string) []string {
	file, _ := os.Open(path + ".txt")
	fileScnr := bufio.NewScanner(file)
	fileScnr.Split(bufio.ScanLines)

	var lines []string
	for fileScnr.Scan() {
		text := strings.TrimSpace(fileScnr.Text())
		if len(text) == 0 {
			continue
		}
		lines = append(lines, text)
	}
	file.Close()
	return lines
}

type Key struct{ y, x int }

func main() {
	lines := read(os.Args[1])
	digitCheck := regexp.MustCompile(`[0-9]`)
	gearCheck := regexp.MustCompile(`\*`)

	grid := [][]string{}
	for _, line := range lines {
		grid = append(grid, strings.Split(strings.TrimSpace(line), ""))
	}

	// map of sets ;-;
	gears := map[Key][]int{}

	hasAdjGears := func(i int, j int) []Key {
		adjGears := []Key{}
		for r := -1; r <= 1; r++ {
			// These kinds of loops make me sad 😢
			for c := -1; c <= 1; c++ {
				// Out of bounds check
				y := i - r
				x := j - c
				if y < 0 || y > len(grid)-1 || x < 0 || x > len(grid[i])-1 {
					continue
				}
				if !gearCheck.MatchString(grid[y][x]) {
					continue
				}
				adjGears = append(adjGears, Key{y, x})
			}
		}
		return adjGears
	}

	stack := []string{}
	for i, chars := range grid {
		j := 0
		for j < len(chars) {
			for digitCheck.MatchString(chars[j]) {
				stack = append(stack, chars[j])
				j += 1
				if j == len(chars) {
					break
				}
			}
			// hey we found some digits
			if len(stack) == 0 {
				j += 1
				continue
			}

			// now lets find a symbol and sum the values :)
			value := 0
			k := 0
			// and the adjacent gears to this number
			adjGears := []Key{}
			for len(stack) > 0 {
				end := len(stack) - 1
				num, _ := strconv.Atoi(stack[end])
				stack = stack[:end]
				value += num * int(math.Pow10(k))
				k += 1
				adjGears = append(adjGears, hasAdjGears(i, j-k)...)
			}
			for _, gear := range adjGears {
				if slices.Contains(gears[gear], value) {
					continue
				}
				gears[gear] = append(gears[gear], value)
			}
		}
	}
	// sum the gear numbers but only if the gear has 2 adjacent numbers
	total := 0
	for _, values := range gears {
		if len(values) != 2 {
			continue
		}
		factor := 1
		for _, value := range values {
			factor *= value
		}
		total += factor
	}
	fmt.Println(total)
}
