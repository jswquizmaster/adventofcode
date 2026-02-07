package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	ranges := make(map[int]int)

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("ERROR")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		l := strings.Split(line, "-")

		start, err := strconv.Atoi(l[0])
		if err != nil {
			fmt.Println("ERROR3")
			return
		}
		end, err := strconv.Atoi(l[1])
		if err != nil {
			fmt.Println("ERROR4")
			return
		}
		val, ok := ranges[start]
		if ok {
			ranges[start] = val + 1
		} else {
			ranges[start] = 1
		}
		val, ok = ranges[end]
		if ok {
			ranges[end] = val - 1
		} else {
			ranges[end] = -1
		}
	}

	keys := make([]int, 0, len(ranges))
	for k := range ranges {
		keys = append(keys, k)
	}
	sort.Ints((keys))

	counter := 0
	answer := 0
	start := 0
	for _, k := range keys {
		if counter == 0 {
			start = k
		}
		counter = counter + ranges[k]
		if counter == 0 {
			answer = answer + (k - start) + 1
		}
	}

	fmt.Println("Answer: ", answer)
}
