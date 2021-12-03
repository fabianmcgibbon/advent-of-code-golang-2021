package main

import (
	"adventofcode/utils/files"
	"fmt"
	"math"
)

func main() {
	input := files.ReadFile(3, "\n")
	solution := getSolution1(input)
	fmt.Println(solution)
}

func getSolution1(input []string) (solution int) {
	bitLength := len(input[0])
	inputLength := len(input)

	gammaRate := make([]string, bitLength)
	alphaRate := make([]string, bitLength)

	for i := 0; i < bitLength; i++ {
		ones := 0
		for j := 0; j < inputLength; j++ {
			char := fmt.Sprintf("%c", input[j][i])
			if char == "1" {
				ones++
			}
		}
		if ones > (inputLength / 2) {
			gammaRate[i] = "1"
			alphaRate[i] = "0"
		} else {
			gammaRate[i] = "0"
			alphaRate[i] = "1"
		}
	}

	fmt.Println(alphaRate)
	fmt.Println(gammaRate)

	return 1
}

func binaryToDecimal(num int) int {
	var remainder int
	index := 0
	decimalNum := 0
	for num != 0 {
		remainder = num % 10
		num = num / 10
		decimalNum = decimalNum + remainder*int(math.Pow(2, float64(index)))
		index++
	}
	return decimalNum
}
