package main

import (
	"adventofcode/utils/conv"
	"adventofcode/utils/files"
	"fmt"
)

func main() {
	inputSliceAsString := files.ReadFile(1, "\n")
	input := conv.ToIntSlice(inputSliceAsString)

	solution := getIncreases(input)

	fmt.Println(solution)
}

func getIncreases(integers []int) (increases int) {
	for i := range integers {
		if i == 0 {
			continue
		}
		if integers[i] > integers[i-1] {
			increases += 1
		}
	}
	return increases
}
