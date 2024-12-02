package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	file, err := os.Open("day-1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var slice1 []int
	var slice2 []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var fields = strings.Fields(scanner.Text())

		slice1 = append(slice1, ConvertToInteger(fields[0]))
		slice2 = append(slice2, ConvertToInteger(fields[1]))
	}

	sort.Ints(slice1)
	sort.Ints(slice2)

	var distance = 0

	for i := 0; i < len(slice1); i++ {
		distance += int(math.Abs(float64(slice2[i] - slice1[i])))
	}

	fmt.Println(distance)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
