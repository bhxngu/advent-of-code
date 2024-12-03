package main

import (
	"fmt"
	"os"
	"regexp"
)

func StringToInt(s string) int {
	var val int
	_, err := fmt.Sscanf(s, "%d", &val)

	if err != nil {
		panic(err)
	}

	return val
}

func main() {
	bytesRead, err := os.ReadFile("day-3/input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	fullString := string(bytesRead)

	firstRegex := regexp.MustCompile(`mul\(\d+\,\d+\)|don't\(\)|do\(\)`)
	matches := firstRegex.FindAllString(fullString, -1)

	total := 0
	active := true
	for _, match := range matches {
		if match == "don't()" {
			active = false
			continue
		} else if match == "do()" {
			active = true
			continue
		}

		if !active {
			continue
		}
		newRegex := regexp.MustCompile(`\d+`)
		numbers := newRegex.FindAllString(match, -1)

		total += StringToInt(numbers[0]) * StringToInt(numbers[1])
	}

	fmt.Println(total)
}
