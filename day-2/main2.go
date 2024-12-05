package main

// Learnt that slices are just pointer to an array
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

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func IsSafeReport(fields []string, alreadyCalled bool) bool {
	var seriesIncreasing bool
	seriesInitialized := false

	for i := 0; i < len(fields)-1; i++ {
		first, second := ConvertToInteger(fields[i]), ConvertToInteger(fields[i+1])

		difference := second - first
		pairIncreasing := second > first

		if (seriesInitialized && pairIncreasing != seriesIncreasing) || math.Abs(float64(difference)) > 3 || math.Abs(float64(difference)) == 0 {
			if alreadyCalled {
				return false
			}

			a := make([]string, len(fields))
			_ = copy(a, fields)
			b := make([]string, len(fields))
			_ = copy(b, fields)
			a = RemoveIndex(a, i)
			b = RemoveIndex(b, i+1)
			firstAns := IsSafeReport(a, true) || IsSafeReport(b, true)

			if i == 1 {
				secondAns := IsSafeReport(fields[1:], true)

				return firstAns || secondAns
			}

			return firstAns
		}

		if !seriesInitialized {
			seriesInitialized = true
			seriesIncreasing = pairIncreasing
			continue
		}

		if i == len(fields)-2 {
			return true
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
