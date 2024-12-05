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
			if j+1 < cols && i+1 < rows && i-1 >= 0 && j-1 >= 0 && arr[i][j] == "A" {
				// M   S
				//   A
				// M   S
				if arr[i-1][j-1] == "M" && arr[i+1][j+1] == "S" && arr[i+1][j-1] == "M" && arr[i-1][j+1] == "S" {
					count++
				}
				// M   M
				//   A
				// S   S
				if arr[i-1][j-1] == "M" && arr[i+1][j+1] == "S" && arr[i+1][j-1] == "S" && arr[i-1][j+1] == "M" {
					count++
				}
				// S   S
				//   A
				// M   M
				if arr[i-1][j-1] == "S" && arr[i+1][j+1] == "M" && arr[i+1][j-1] == "M" && arr[i-1][j+1] == "S" {
					count++
				}
				// S   M
				//   A
				// S   M
				if arr[i-1][j-1] == "S" && arr[i+1][j+1] == "M" && arr[i+1][j-1] == "S" && arr[i-1][j+1] == "M" {
					count++
				}
			}
		}
	}

	fmt.Println(count)
}
