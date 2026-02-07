package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	code := 50
	dir := 0
	password := 0

	f, err := os.Open("input.txt")
    if err != nil {
        log.Fatal(err)
    }
    
    defer f.Close()

	scanner := bufio.NewScanner(f)

    for scanner.Scan() {
		step := scanner.Text()
		amount, err := strconv.Atoi(step[1:])
		if err != nil {
			fmt.Println("ERROR")
		}
		if step[0] == 76 {
			dir = -1
		} else {
			dir = 1
		}
		for i := 0; i < amount; i++ {
			code = code + dir
			if code == -1 {
				code = 99
			}
			if code == 100 {
				code = 0
			}
			if code == 0 {
				password = password + 1
			}
		}
		
	}

	fmt.Println("The password is ", password)
}
