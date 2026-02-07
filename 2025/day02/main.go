package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func compare(str string, substr string) bool {
	//fmt.Println(str)
	//fmt.Println(substr)

	i := 0
	for j := 0; j < len(str); j++ {
		if str[j] != substr[i] {
			return false
		}
		i++
		if i == len(substr) {
			i = 0
		}
	}
	if i == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	file, ferr := os.Open("input.txt")
	if ferr != nil {
		panic(ferr)
	}
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()
	ranges := strings.Split(line, ",")

	invalidID := 0
	for _, value := range ranges {
		ids := strings.Split(value, "-")
		start, err := strconv.Atoi(ids[0])
		if err != nil {
			fmt.Println("ERROR")
		}
		end, err := strconv.Atoi(ids[1])
		if err != nil {
			fmt.Println("ERROR")
		}
		fmt.Printf("\n%d - %d\n\n", start, end)
		for i := start; i <= end; i++ {
			if i > 8800 && i < 8900 {
				fmt.Println("JOE: ", i)
			}
			id := strconv.Itoa(i)
			for j := 1; j <= len(id)/2; j++ {
				if i == 8888 {
					fmt.Printf("JOEJOE %s - %s", id, id[:j])
				}
				if compare(id, id[:j]) {
					fmt.Println("Invalid ID ", id)
					invalidID = invalidID + i
					break
				}
			}
		}
	}
	fmt.Println("Invalid IDs", invalidID)
}
