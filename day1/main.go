package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	res := trebuchet()
	fmt.Println(res)
}

func first(line string, comparator []string, currIdx int) int {
	if currIdx > len(line)-1 {
		return 0
	}

	target := string(line[currIdx])

	if slices.Contains(comparator, target) {
		if d, err := strconv.Atoi(target); err == nil {
			return d * 10
		}
	}

	return first(line, comparator, currIdx+1)
}

func last(line string, comparator []string, currIdx int) int {
	if currIdx < 0 {
		return 0
	}

	target := string(line[currIdx])
	if slices.Contains(comparator, target) {
		if d, err := strconv.Atoi(target); err == nil {
			return d
		}
	}

	return last(line, comparator, currIdx-1)
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
		total += first(line, targets, 0) + last(line, targets, len(line)-1)
	}

	return total
}
