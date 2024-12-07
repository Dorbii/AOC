package aoc2024

import (
	"aoc/aoc2024/utils"
	"fmt"
	"strconv"
)

func Day2() {
	input, err := utils.ReadCsvFile("/aoc2024/day2_data/input.csv", false)
	res := 0
	unSafe := make(map[int][]string)
	if err != nil {
		//handle error after log configuration
	}
outerLoop:
	for i, row := range input {
		asc := false
		lasc, _ := strconv.Atoi(row[0])
		rasc, _ := strconv.Atoi(row[1])
		if lasc < rasc {
			asc = true
		}
		safe := safeCheck(row, asc)
		if !safe {
			unSafe[i] = row
			continue outerLoop
		}
		res++
	}
	fmt.Println("Day 2: part 1")
	fmt.Println(res)
	res2 := 0
	fail := 0
	fmt.Println("Day 2: part 2")
	for _, v := range unSafe {
		if tryFix(v) {
			res2++
		} else {
			fail++
		}
	}
	fmt.Println("=======DEBUG=======")
	fmt.Println("Total: ", len(input))
	fmt.Println("Safe(1st pass): ", res)
	fmt.Println("unsafe(expected): ", len(input)-res)
	fmt.Println("unSafe(actual): ", len(unSafe))
	fmt.Println("=======FIXES=======")

	fmt.Println("Fail: ", fail)
	fmt.Println("Pass: ", res2)

	fmt.Println("Answer: ", res2+res)
	fmt.Println("Answer Check: ", res2+res)
	fmt.Println("Check: ", res2+res+fail)

}

func safeCheck(slice []string, asc bool) bool {
	for i := 0; i < len(slice)-1; i++ {
		l, err := strconv.Atoi(slice[i])
		if err != nil {
			fmt.Println(err)
		}
		r, err := strconv.Atoi(slice[i+1])
		if err != nil {
			fmt.Println(err)
		}
		if (asc && l > r || !asc && l < r) || !(utils.Abs(l-r) <= 3 && utils.Abs(l-r) >= 1) {
			return false
		}
	}
	return true
}

func tryFix(slice []string) bool {
	for i := 0; i < len(slice)-1; i++ {
		mod_slice := make([]string, len(slice)-1)
		copy(mod_slice, slice)
		if i == len(slice)-1 {
			mod_slice = mod_slice[:len(slice)-1]
		} else {
			mod_slice = append(mod_slice[:i], mod_slice[i+1:]...)
		}
		mod_asc := directionCheck(mod_slice)
		if safeCheck(mod_slice, mod_asc) {
			return true
		}
	}
	return false
}

func directionCheck(slice []string) bool {
	asc := false
	lasc, _ := strconv.Atoi(slice[0])
	rasc, _ := strconv.Atoi(slice[1])
	if lasc < rasc {
		asc = true
	}
	return asc
}
