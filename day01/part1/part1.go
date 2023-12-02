// Advent of Code 2023
// Day 01 - Part 1
// 12-01-23
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, _ := os.ReadFile("../input.txt")
	text := string(file)
	lines := strings.Split(text, "\n")
	values := make([]int, 0)

	for _, line := range lines {
		var chars string
		// first digit from start
		for _, ch := range line {
			if unicode.IsDigit(ch) {
				chars += string(ch)
				break
			}
		}
		// first digit from end
		for i := len(line) - 1; i >= 0; i-- {
			ch := rune(line[i])
			if unicode.IsDigit(ch) {
				chars += string(ch)
				break
			}
		}
		// cast to int
		value, _ := strconv.Atoi(chars)
		values = append(values, value)
	}
	// fmt.Println(values)

	// sum values
	var total int
	for _, val := range values {
		total += val
	}
	fmt.Print(total)
}
