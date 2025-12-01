// Solution for Day 1
package main

import (
	"fmt"
	"strings"
	"os"
	"io"
	"strconv"
)

func Task1(str string) int {
	dial := 50
	zero_counter := 0
	for line := range strings.SplitSeq(str, "\n") {
		if len(line) < 1 {continue}
		rotation := string(line[0])
		str_num := line[1:]
		num, _ := strconv.Atoi(str_num)

		fmt.Println("Rotation", rotation, "Steps", num, "Dial", dial)
		switch rotation {
			case "R":
				dial = dial + num
			case "L":
				dial = dial - num
			default:
				fmt.Println("Undefined Behaviour!!")
		}

		dial = (dial % 100 + 100) % 100
		if dial == 0 {zero_counter++}
	}
	fmt.Printf("There are %d zeroes\n", zero_counter)
	return 0
}

func Task2(str string) int {
	dial := 50
	zero_hits := 0
	was_zero := false
	for line := range strings.SplitSeq(str, "\n") {
		if len(line) < 1 {continue}
		rotation := string(line[0])
		num, _ := strconv.Atoi(line[1:])

		switch rotation {
			case "R":
				dial = dial + num
			case "L":
				dial = dial - num
			default:
				panic("It's so over!")
		}
		loop_rounds := 0
		if dial > 99 && dial % 100 == 0 {
			loop_rounds--
		}
		for dial > 99 || dial < 0 {
			if dial > 99 {
				dial = dial - 100
				loop_rounds++
			} else if dial < 0 {
				if was_zero {
					loop_rounds--
					was_zero = false
				}
				dial = dial + 100
				loop_rounds++
			}
		}
		fmt.Println("# loops:", loop_rounds, "Dial:", dial)
		if dial == 0 {
			was_zero = true
			fmt.Println("Hit 0, dial:", dial)
			zero_hits++
		} else {
			was_zero = false
		}
		zero_hits = zero_hits + loop_rounds

	}
	fmt.Printf("There are %d zeroes passed\n", zero_hits)

	return 0
}

func main() {
	file, err := os.Open("input") // Replace "input.txt" with your file name
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close() 
	file_bytes, _ := io.ReadAll(file)
	file_string := string(file_bytes)
	
	Task1(file_string)
	Task2(file_string)
}
