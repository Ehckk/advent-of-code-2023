package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
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

func main() {
	lines := read(os.Args[1])
	digitCheck := regexp.MustCompile(`[0-9]`)
	symbolCheck := regexp.MustCompile(`[^\d\.]`)

	grid := [][]string{}
	for _, line := range lines {
		grid = append(grid, strings.Split(strings.TrimSpace(line), ""))
	}

	hasAdjSymbol := func(i int, j int) bool {
		for r := -1; r <= 1; r++ {
			// These kinds of loops make me sad 😢
			for c := -1; c <= 1; c++ {
				// Out of bounds check
				y := i - r
				x := j - c
				if y < 0 || y > len(grid)-1 || x < 0 || x > len(grid[i])-1 {
					continue
				}
				if !symbolCheck.MatchString(grid[y][x]) {
					continue
				}
				return true
			}
		}
		return false
	}

	stack := []string{}
	total := 0
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
			hasSymbol := false
			value := 0
			k := 0
			for len(stack) > 0 {
				end := len(stack) - 1
				num, _ := strconv.Atoi(stack[end])
				stack = stack[:end]

				value += num * int(math.Pow10(k))
				k += 1

				// make sure we find an adjacent symbol around the number
				if hasSymbol {
					continue // if we haven't already :)
				}
				hasSymbol = hasAdjSymbol(i, j-k)
			}
			if hasSymbol {
				total += value
			}
			j += 1
		}
	}
	fmt.Println(total)
}
