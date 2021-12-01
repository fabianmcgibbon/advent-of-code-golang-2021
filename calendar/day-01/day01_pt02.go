package main

import (
	"adventofcode/utils/conv"
	"adventofcode/utils/files"
	"fmt"
)

func main() {
	inputSliceAsString := files.ReadFile(1, "\n")
	input := conv.ToIntSlice(inputSliceAsString)

	solution := getWindowIncreases(input)

	fmt.Println(solution)
}

func getWindowIncreases(integers []int) (increases int) {
	for i := range integers {
		if i < 3 {
			continue
		}
		if integers[i] > integers[i-3] {
			increases += 1
		}
	}
	return increases
}
