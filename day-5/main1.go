package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ConvertToInteger(s string) int {
	val, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return val
}

func main() {
	file, err := os.Open("day-5/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	notAllowed := make(map[int]map[int]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		pair := strings.Split(scanner.Text(), "|")
		first, second := ConvertToInteger(pair[0]), ConvertToInteger(pair[1])
		if notAllowed[second] == nil {
			notAllowed[second] = make(map[int]bool)
		}
		notAllowed[second][first] = true
	}

	medianSum := 0
	for scanner.Scan() {
		nums := strings.Split(scanner.Text(), ",")

		length := len(nums)
		for i := 0; i < length-1; i += 1 {
			first, second := ConvertToInteger(nums[i]), ConvertToInteger(nums[i+1])
			if notAllowed[first] != nil {
				if notAllowed[first][second] {
					break
				}
			}
			if i == length-2 {
				medianSum += ConvertToInteger(nums[(length-1)/2])
			}
		}
	}

	fmt.Println(medianSum)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
