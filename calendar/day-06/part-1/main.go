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
	output := runSimulation(80, intSlice)
	fmt.Println(len(output))
}

func runSimulation(days int, input []int) []int {
	var newFish int
	for day := 0; day < days; day++ {
		newFish = 0
		for i := range input {
			switch {
			case input[i] > 0:
				input[i]--
			case input[i] == 0:
				input[i] = 6
				newFish++
			}
		}
		for f := 0; f < newFish; f++ {
			input = append(input, 8)
		}
	}

	return input
}
