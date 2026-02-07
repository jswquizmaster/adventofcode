package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	totalVoltage := 0

	file, ferr := os.Open("input.txt")
	if ferr != nil {
		panic(ferr)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bank := scanner.Text()
		digits := [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
		lastpos := -1
		for i := 0; i < 12; i++ {
			for j := lastpos + 1; j < len(bank)-(11-i); j++ {
				d, err := strconv.Atoi(string(bank[j]))
				if err != nil {
					fmt.Println("ERROR")
					return
				}
				if d > digits[i] {
					digits[i] = d
					lastpos = j
				}
			}
		}

		voltageString := ""
		for k := 0; k < 12; k++ {
			voltageString = voltageString + strconv.Itoa(digits[k])
		}
		fmt.Printf("BANK: %s -> %s\n", bank, voltageString)
		voltage, err := strconv.Atoi(voltageString)
		if err != nil {
			fmt.Println("ERROR2")
			return
		}

		totalVoltage = totalVoltage + voltage
	}

	fmt.Println("Total voltage is ", totalVoltage)
}
