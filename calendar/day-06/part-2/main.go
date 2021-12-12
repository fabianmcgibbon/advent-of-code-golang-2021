package main

import (
	"adventofcode/utils/conv"
	"adventofcode/utils/files"
	"fmt"
	"strings"
)

func main() {
	input := files.ReadFile(6, "\n")
	intSlice := conv.ToIntSlice(strings.Split(input[0], ","))
	output := runSimulation(256, intSlice)
	fmt.Println(sum(output))
}

func runSimulation(days int, input []int) []int {
	fish := make([]int, 9)

	// initial fish
	for _, i := range input {
		fish[i]++
	}

	for day := 0; day < days; day++ {
		fish = passDay(fish)
	}

	return fish
}

func passDay(s []int) (res []int) {
	res = s[1:]
	res = append(res, s[0])
	res[6] += s[0]

	return res
}

func sum(s []int) (res int) {
	for _, v := range s {
		res += v
	}

	return res
}
