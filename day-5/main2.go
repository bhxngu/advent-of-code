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

func medianOfOrderedList(arr []int, notAllowed map[int]map[int]bool, recursive bool) int {
	for i := 0; i < len(arr)-1; i += 1 {
		first, second := arr[i], arr[i+1]
		if notAllowed[first] != nil {
			if notAllowed[first][second] {
				arr[i] = second
				arr[i+1] = first
				return medianOfOrderedList(arr, notAllowed, true)
			}
		}

		if i == len(arr)-2 && recursive {
			return arr[(len(arr)-1)/2]
		}
	}

	return 0
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
		numsInt := make([]int, len(nums))
		for i, num := range nums {
			numsInt[i] = ConvertToInteger(num)
		}
		medianSum += medianOfOrderedList(numsInt, notAllowed, false)
	}

	fmt.Println(medianSum)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
