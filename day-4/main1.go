package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("day-4/input.txt")
	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	scanner := bufio.NewScanner(file)

	arr := make([][]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		arr = append(arr, strings.Split(line, ""))
	}

	count := 0
	rows := len(arr)
	cols := len(arr[0])
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if j+3 < cols && arr[i][j] == "X" && arr[i][j+1] == "M" && arr[i][j+2] == "A" && arr[i][j+3] == "S" {
				count++
			}
			if i+3 < rows && arr[i][j] == "X" && arr[i+1][j] == "M" && arr[i+2][j] == "A" && arr[i+3][j] == "S" {
				count++
			}
			if i+3 < rows && j+3 < cols && arr[i][j] == "X" && arr[i+1][j+1] == "M" && arr[i+2][j+2] == "A" && arr[i+3][j+3] == "S" {
				count++
			}
			if i-3 >= 0 && j-3 >= 0 && arr[i][j] == "X" && arr[i-1][j-1] == "M" && arr[i-2][j-2] == "A" && arr[i-3][j-3] == "S" {
				count++
			}
			if j-3 >= 0 && arr[i][j] == "X" && arr[i][j-1] == "M" && arr[i][j-2] == "A" && arr[i][j-3] == "S" {
				count++
			}
			if i-3 >= 0 && arr[i][j] == "X" && arr[i-1][j] == "M" && arr[i-2][j] == "A" && arr[i-3][j] == "S" {
				count++
			}
			if i+3 < rows && j-3 >= 0 && arr[i][j] == "X" && arr[i+1][j-1] == "M" && arr[i+2][j-2] == "A" && arr[i+3][j-3] == "S" {
				count++
			}
			if i-3 >= 0 && j+3 < cols && arr[i][j] == "X" && arr[i-1][j+1] == "M" && arr[i-2][j+2] == "A" && arr[i-3][j+3] == "S" {
				count++
			}
		}
	}

	fmt.Println(count)
}
