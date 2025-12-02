package main

import (
	"fmt"
	"strings"
	"strconv"
	"os"
	"io"
)

func Map[T any, R any](items []T, fx func(T) R) []R {
    var result []R
    for _, v := range items {
        result = append(result, fx(v))
    }
    return result
}

func Task1(input string) int64 {
    var sum int64 = 0

    for rng := range strings.SplitSeq(input, ",") {
        rng = strings.TrimSpace(rng)
        if rng == "" {
            continue
        }
        parts := strings.Split(rng, "-")
        if len(parts) != 2 {
            continue
        }

        start, err1 := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
        end,   err2 := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
        if err1 != nil || err2 != nil {
            continue
        }

        for idx := start; idx <= end; idx++ {
            s := strconv.FormatInt(idx, 10)
            n := len(s)
            if n%2 != 0 {
                continue
            }
            half := n / 2
            if s[:half] == s[half:] {
                // fmt.Println("Found Match:", idx) // optional
                sum += idx
            }
        }
    }
    return sum
}

func check_slices(str string) int64 {
	var val int64 = 0
	is_repeat := true
	if len(str) < 2 {return 0}
	for i := 1; i <= len(str) / 2; i++ {
		is_repeat = true
		// fmt.Println("Processing string number:", str)
		if len(str) % i != 0 {continue}
		part := string(str[:i])
		for l := i; l < len(str); l = l + i {
			// fmt.Println("Loop part:", str[l:l+i])
			if str[l:l+i] != part {
				is_repeat = false
				break
			}
		}
		if is_repeat {
			val, _ = strconv.ParseInt(str, 10, 64)
			return val
		}

	} 
	return val
}

func Task2(str string) int64 {
    var sum int64 = 0
	var numbers_checked int64 = 0

    for rng := range strings.SplitSeq(str, ",") {
        rng = strings.TrimSpace(rng)
        if rng == "" {
            continue
        }
        parts := strings.Split(rng, "-")
        if len(parts) != 2 {
            continue
        }

        start, err1 := strconv.ParseInt(strings.TrimSpace(parts[0]), 10, 64)
        end,   err2 := strconv.ParseInt(strings.TrimSpace(parts[1]), 10, 64)
        if err1 != nil || err2 != nil {
            continue
        }
		numbers_checked = numbers_checked + (end - start)
        for idx := start; idx <= end; idx++ {
            s := strconv.FormatInt(idx, 10)
            sum = sum + check_slices(s)
        }
    }

	fmt.Println("Checked # numbers:", numbers_checked)
	return sum
}

func main() {
	fh, err := os.Open("input")
	if err != nil {
		panic("Could not read file!")
	}
	byte_input, _ := io.ReadAll(fh)
	str_input := string(byte_input)
	sol1 := Task1(str_input)
	sol2 := Task2(str_input)
	fmt.Println("Solution Task1:", sol1)
	fmt.Println("Solution Task2:", sol2)
	fmt.Println("Checking 'check_slices'(12121213):", check_slices(string("12121213")))
	fmt.Println("Checking 'check_slices'(12121212):", check_slices(string("12121212")))
	fmt.Println("Checking 'check_slices'(111):", check_slices(string("111")))
}
