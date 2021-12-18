package main

import (
	"adventofcode/utils/conv"
	"adventofcode/utils/files"
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	input := files.ReadFile(7, "\n")
	intSlice := conv.ToIntSlice(strings.Split(input[0], ","))
	output := calculateFuel(intSlice)
	fmt.Println(output)
}

func calculateFuel(input []int) (fuel int) {
	med := float64(median(input))

	for _, v := range input {
		fuel += int(math.Abs(float64(v) - med))
	}

	return fuel
}

func median(input []int) (res int) {
	sort.Ints(input)
	n := len(input)

	// even
	if len(input)%2 == 0 {
		res = (input[n/2] + input[n/2-1]) / 2

		// odd
	} else {
		res = input[(n-1)/2]
	}

	return res
}
