package main

import (
	"bufio"
	"math"
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
	file, err := os.Open("day-2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0
	for scanner.Scan() {
		var fields = strings.Fields(scanner.Text())

		var seriesIncreasing bool
		seriesInitialized := false
		for i := 1; i < len(fields); i++ {

			first, second := ConvertToInteger(fields[i-1]), ConvertToInteger(fields[i])

			difference := second - first

			if math.Abs(float64(difference)) >= 1 && math.Abs(float64(difference)) <= 3 {
				pairIncreasing := second > first

				if seriesInitialized {
					if pairIncreasing != seriesIncreasing {
						break
					}

					if i == len(fields)-1 {
						count++
					}
				} else {
					seriesInitialized = true
					seriesIncreasing = pairIncreasing
				}
			} else {
				break
			}
		}
	}

	print("Safe reports: ", count)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
