package main

import (
	"adventofcode/utils/conv"
	"adventofcode/utils/files"
	"fmt"
	"math"
	"strings"
)

func main() {
	input := files.ReadFile(7, "\n")
	intSlice := conv.ToIntSlice(strings.Split(input[0], ","))
	output := calculateFuel(intSlice)
	fmt.Println(output)
}

func calculateFuel(input []int) (fuel int) {
	c := math.Ceil(mean(input))
	f := math.Floor(mean(input))

	var fuels [2]int
	for i, m := range []float64{c, f} {
		for _, v := range input {
			diff := int(math.Abs(float64(v) - m))
			for j := 1; j <= diff; j++ {
				fuels[i] += j
			}
		}
	}

	fuel = int(math.Min(float64(fuels[0]), float64(fuels[1])))

	return fuel
}

func mean(input []int) (res float64) {
	return float64(sum(input)) / float64(len(input))
}

func sum(s []int) (res int) {
	for _, v := range s {
		res += v
	}

	return res
}
