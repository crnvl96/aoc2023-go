package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	defer file.Close()

	scanner := bufio.NewScanner(file)
	targets := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		limit := len(line) - 1
	first:
		for _, c := range line {
			for _, t := range targets {
				if t == string(c) {
					if d, err := strconv.Atoi(t); err == nil {
						total += 10 * d
						break first
					}
				}
			}
		}
	last:
		for c := limit; c >= 0; c-- {
			for _, t := range targets {
				if t == string(line[c]) {
					if d, err := strconv.Atoi(t); err == nil {
						total += d
						break last
					}
				}
			}
		}
	}
	return total
}
