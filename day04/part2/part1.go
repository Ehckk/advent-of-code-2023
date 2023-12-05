// Advent of Code 2023
// Day 04 - Part 1
// 12-04-23
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func read(path string) []string {
	file, err := os.Open(path + ".txt")
	if err != nil {
		panic(err)
	}
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

func main() {
	lines := read(os.Args[1])
	spaces := regexp.MustCompile(`\s+`)

	counts := make([]int, len(lines))
	for i, line := range lines {
		line = strings.Split(strings.TrimSpace(line), ":")[1]
		values := strings.Split(strings.TrimSpace(line), "|")

		winning := map[int]bool{}

		for _, value := range spaces.Split(strings.TrimSpace(values[0]), -1) {
			value, _ := strconv.Atoi(value)
			winning[value] = false
		}
		counts[i] += 1
		wins := 0
		for _, value := range spaces.Split(strings.TrimSpace(values[1]), -1) {
			value, _ := strconv.Atoi(value)
			checked, exists := winning[value]
			if exists && !checked {
				winning[value] = true
				wins += 1
			}
		}
		for c := i + 1; c < i+wins; c++ {
			if c < len(counts) {
				counts[c] += counts[i]
			}
		}
		fmt.Println(i, wins, counts)
	}

	total := 0
	for _, amount := range counts {
		total += amount
	}
	fmt.Println(counts, total)
}
