package main

import (
	"bufio"
	"fmt"
	"os"
)

var board [140][]rune

func check(row int, col int) int {
	if row < 0 || row >= len(board) {
		return 0
	}
	if col < 0 || col >= len(board[row]) {
		return 0
	}
	if board[row][col] != '@' {
		return 0
	}
	return 1
}

func main() {
	numPapers := 0
	file, ferr := os.Open("input.txt")
	if ferr != nil {
		panic(ferr)
	}
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		board[i] = []rune(scanner.Text())
		i++
	}

start:
	for i := range board {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '@' {
				paper := check(i-1, j-1) + check(i-1, j) + check(i-1, j+1) +
					check(i, j-1) + +check(i, j+1) +
					check(i+1, j-1) + check(i+1, j) + check(i+1, j+1)
				if paper < 4 {
					numPapers++
					board[i][j] = '.'
					goto start
				}
			}
		}
	}

	fmt.Println("# papers: ", numPapers)
}
