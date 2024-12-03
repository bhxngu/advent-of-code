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

func IsSafeReport(fields []string, alreadyCalled bool) bool {
	var seriesIncreasing bool
	seriesInitialized := false

	for i := 0; i < len(fields)-1; i++ {
		first, second := ConvertToInteger(fields[i]), ConvertToInteger(fields[i+1])

		difference := second - first

		if math.Abs(float64(difference)) >= 1 && math.Abs(float64(difference)) <= 3 {
			pairIncreasing := second > first

			if seriesInitialized {
				if pairIncreasing != seriesIncreasing {
					if alreadyCalled {
						return false
					}
					first_ans := IsSafeReport(append(fields[:i], fields[i+1:]...), true) || IsSafeReport(append(fields[:i+1], fields[i+2:]...), true)
					if i > 0 {
						second_ans := IsSafeReport(fields[1:], true)
						return first_ans || second_ans
					}
					return first_ans
				}

				if i == len(fields)-2 {
					return true
				}
			} else {
				seriesInitialized = true
				seriesIncreasing = pairIncreasing
			}
		} else {
			if alreadyCalled {
				return false
			}
			return IsSafeReport(append(fields[:i], fields[i+1:]...), true) || IsSafeReport(append(fields[:i+1], fields[i+2:]...), true)
		}
	}
	return false
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
		if IsSafeReport(fields, false) {
			count++
		}
	}

	print("Safe reports after tolerating a single bad level: ", count)

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
