// Advent of Code 2023
// Day 01 - Part 2
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
	nums := map[string]rune{
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	file, _ := os.ReadFile("../input.txt")
	text := string(file)
	lines := strings.Split(text, "\n")
	values := make([]int, 0)

	startsWithName := func(target string) (bool, rune) {
		// Loop over num names
		for key, num := range nums {
			// Does this part end with the num name
			if strings.HasSuffix(target, key) {
				return true, num
			}
		}
		return false, ' '
	}

	endsWithName := func(target string) (bool, rune) {
		// Loop over num names
		for key, num := range nums {
			// Does this part start with the num name
			if strings.HasPrefix(target, key) {
				return true, num
			}
		}
		return false, ' '
	}

	for _, line := range lines {
		var chars string
		// first digit from start
		for i := 0; i < len(line); i++ {
			ch := rune(line[i])
			if unicode.IsDigit(ch) {
				chars += string(ch)
				break
			}
			// Check for num name
			result, ch := startsWithName(line[:i+1])
			if result {
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
			// Check for num name
			result, ch := endsWithName(line[i:])
			if result {
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
