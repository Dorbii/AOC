package utils

func Counts(s []int) map[int]int {
	counts := make(map[int]int)
	for _, val := range s {
		counts[val]++
	}
	return counts
}
