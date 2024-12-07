package aoc2024

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// x mul(2,4) %&mul[3,7]!@^do_not_ mul(5,5) +mul(32,64]then (mul(11,8)mul(8,5))
func Day3() {
	data, err := os.ReadFile("aoc2024/day3_data/input_p2.txt")
	if err != nil {
		//handle error after log configuration
	}
	res := 0
	str := string(data)
	parsed := regexp.MustCompile(`(mul\(\d+,\d+\)|do\(\)|don't\(\))`)
	matched := parsed.FindAllString(str, -1)
	on_flag := true
	for _, v := range matched {
		if v == "don't()" {
			on_flag = false
			continue
		} else if v == "do()" {
			on_flag = true
		}
		if on_flag {
			parts := regexp.MustCompile(`\d+`).FindAllString(v, -1)
			if len(parts) == 2 {
				a, _ := strconv.Atoi(parts[0])
				b, _ := strconv.Atoi(parts[1])

				res += a * b
			}
		}

	}
	fmt.Print("Day 3: part 1\n")
	fmt.Println(res)
}
