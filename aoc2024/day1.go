package aoc2024

import (
	"aoc/aoc2024/utils"
	"fmt"
	"sort"
	"strconv"
)

type Input struct {
	Left  []int
	Right []int
}

func Day1() {
	res_part1 := 0
	res_part2 := 0
	input, err := utils.ReadCsvFile("/aoc2024/day1_data/input.csv", true)
	if err != nil {
		//handle error after log configuration
	}
	data := Input{}
	for _, row := range input {
		l, err := strconv.Atoi(row[0])
		if err != nil {
			//handle error after log configuration
		}
		r, err := strconv.Atoi(row[1])
		if err != nil {
			//handle error after log configuration
		}
		data.Left = append(data.Left, l)
		data.Right = append(data.Right, r)
	}
	sort.Ints(data.Left)
	sort.Ints(data.Right)
	right_counts := utils.Counts(data.Right)
	for i := 0; i < len(data.Left); i++ {
		res_part1 += utils.Abs(data.Left[i] - data.Right[i])
		res_part2 += data.Left[i] * right_counts[data.Left[i]]
	}
	fmt.Println("Day 1: part 1")
	fmt.Println(res_part1)
	fmt.Println("Day 1: part 2")
	fmt.Println(res_part2)

}
