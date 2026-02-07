package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var aufgaben []string
	var operands []string
	result := 1
	finalResult := 0

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for row := 0; row < 4; row++ {
		scanner.Scan()
		aufgaben = append(aufgaben, scanner.Text())
	}
	scanner.Scan()
	operands = strings.Fields(scanner.Text())
	index := 0
	for i := 0; i < len(aufgaben[0]); i++ {
		n := strings.TrimSpace(string(aufgaben[0][i]) + string(aufgaben[1][i]) + string(aufgaben[2][i]) + string(aufgaben[3][i]))
		num, err := strconv.Atoi(n)
		if err != nil {
			if operands[index] == "+" {
				result--
			}
			fmt.Println("Result: ", result)
			finalResult = finalResult + result
			result = 1
			index++
			continue
		}
		if operands[index] == "+" {
			result = result + num
		} else {
			result = result * num
		}
	}

	fmt.Println("Final result: ", finalResult)
}
