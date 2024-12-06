package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInitialPos(arr [][]string, rows int, cols int) (int, int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if arr[i][j] == "^" {
				return i, j
			}
		}
	}

	return -1, -1
}

func main() {
	file, err := os.Open("day-6/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	guardArr := make([][]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Split(line, "")
		guardArr = append(guardArr, row)
	}

	rows := len(guardArr)
	cols := len(guardArr[0])
	currX, currY := getInitialPos(guardArr, rows, cols)
	fmt.Println("Initial Position: ", currX, currY)

	count := 1
	dy := 0
	dx := -1
	for {
		guardArr[currX][currY] = "X"

		newX := currX + dx
		newY := currY + dy

		if newX == rows || newX == -1 || newY == -1 || newY == cols {
			break
		}

		if guardArr[newX][newY] == "#" {
			if dy == 0 {
				dy = -dx
				dx = 0
			} else {
				dx = dy
				dy = 0
			}
			continue
		}

		if guardArr[newX][newY] == "." {
			count++
		}

		currX = newX
		currY = newY
	}

	fmt.Println("Distinct Positions visited by guard: ", count)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
