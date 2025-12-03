package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
	"io"
)

func Task1(str string) int64 {
	var sum int64 = 0 
	for bank := range strings.SplitSeq(str, "\n") {
		var max_val int64 = -1
		if len(bank) < 1 {continue}
		for idx, val_x := range bank {
			max_second := -1
			if idx + 1 == len(bank){continue}
			for _, str_y := range bank[idx+1:] {
				val_y, _ := strconv.Atoi(string(str_y))
				if val_y > max_second {
					max_second = val_y
				}
			}
			if max_second != -1 {
				app_val, _ := strconv.ParseInt(fmt.Sprintf("%s%d", string(val_x), max_second), 10, 64)
				if max_val < app_val {
					max_val = app_val
				}
			} else {
				app_val, _ := strconv.ParseInt(string(val_x), 10, 64)
				if max_val < app_val {
					max_val = app_val
				}
			}
		}
		fmt.Println("Highest value:", max_val)
		sum = sum + max_val
	}
	return sum
}

func main() {
	file, _ := os.Open("input")
	bytes_str, _ := io.ReadAll(file)
	str_string := string(bytes_str)

	fmt.Println("The optiml batteries produce:", Task1(str_string), "Jolts")
}
