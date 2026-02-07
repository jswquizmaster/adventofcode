package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const depth = 70
const width = 141

var field [depth][width]int

func trackBeam(r int, c int) int {
	for i := r; i < depth; i++ {
		if field[i][c] > 0 {
			return field[i][c]
		}
		if field[i][c] == 0 {
			sum := trackBeam(i+1, c-1) + trackBeam(i+1, c+1)
			field[i][c] = sum
			return sum
		}
	}
	return 1
}

func main() {

	// Load puzzle input into field array
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	start := strings.IndexRune(scanner.Text(), 'S')
	scanner.Scan()
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			if line[i] == '^' {
				field[row][i] = 0
			} else {
				field[row][i] = -1
			}
		}
		scanner.Scan()
		row++
	}

	// Solve puzzle
	answer := trackBeam(0, start)
	fmt.Println("Answer: ", answer)

}
