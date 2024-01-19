package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	res := trebuchet()
	fmt.Println(res)
}

func trebuchet() int {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	targets := map[string]int{
		"1":     1,
		"one":   1,
		"2":     2,
		"two":   2,
		"3":     3,
		"three": 3,
		"4":     4,
		"four":  4,
		"5":     5,
		"five":  5,
		"6":     6,
		"six":   6,
		"7":     7,
		"seven": 7,
		"8":     8,
		"eight": 8,
		"9":     9,
		"nine":  9,
	}

	total := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		maxAllowedIdx := len(line) - 1

		var first func(idx int) int
		first = func(idx int) int {
			if idx > maxAllowedIdx {
				return 0
			}

			substr := line[idx:]
			for k, v := range targets {
				if strings.HasPrefix(substr, k) {
					return v * 10
				}
			}

			return first(idx + 1)
		}

		var last func(idx int) int
		last = func(idx int) int {
			if idx > maxAllowedIdx {
				return 0
			}

			substr := line[:maxAllowedIdx-idx+1]
			for k, v := range targets {
				if strings.HasSuffix(substr, k) {
					return v
				}
			}

			return last(idx + 1)
		}

		total += first(0) + last(0)
	}

	return total
}
