package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const presentSize = 3
var presentArea = [6]int {7, 7, 7, 5, 6, 7}

func main() {
answer := 0

	f, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()
	scanner := bufio.NewScanner(f)
    line := 0
	for scanner.Scan() {
		line++
		if line<31 {
			continue
		}
		el := strings.Fields(scanner.Text())
		r := strings.Split(el[0][:len(el[0])-1], "x")
		width,_ := strconv.Atoi(r[0])
		length,_ := strconv.Atoi(r[1])

		numPresents := 0
		minArea := 0
		for i, v := range el[1:] {
			present, _ := strconv.Atoi(v)
			numPresents += present
			minArea += present * presentArea[i]
		}
		if (width/presentSize) * (length/presentSize) >= numPresents {
			answer++
		} else if minArea <= width*length {
			log.Fatal("ERROR: Can not decide whether presents fit or not!")
		}
	}

	fmt.Println("Answer: ", answer)
}