package main

import (
	"adventofcode/utils/files"
	"fmt"
	"math"
)

func main() {
	input := files.ReadFile(3, "\n")
	alpha, gamma := getRates(input)
	solution := getSolution(alpha, gamma)
	fmt.Println(solution)
}

func getSolution(a, g []int) (solution int) {
	alphaRate := binaryToDecimal(a)
	gammaRate := binaryToDecimal(g)

	fmt.Printf("Alpha Rate: %v\n", alphaRate)
	fmt.Printf("Gamma Rate: %v\n", gammaRate)

	solution = alphaRate * gammaRate

	return solution
}

func getRates(input []string) (alphaRate, gammaRate []int) {
	bitLength := len(input[0])

	gammaRate = make([]int, bitLength)
	alphaRate = make([]int, bitLength)
	counter := make([]int, bitLength)

	for _, n := range input {
		for i := range n {
			if n[i] == '1' {
				counter[i]++
			}
		}
	}
	fmt.Println(counter)

	for i := range counter {
		counter[i] = counter[i] / (len(input) / 2)
	}

	fmt.Println(counter)

	alphaRate = counter
	gammaRate = flipBits(counter)

	fmt.Printf("Alpha Rate: %v\n", alphaRate)
	fmt.Printf("Gamma Rate: %v\n", gammaRate)

	return alphaRate, gammaRate
}

func binaryToDecimal(b []int) (d int) {
	for i := range b {
		d += b[i] * int(math.Pow(2, float64(len(b)-i-1)))
	}

	return d
}

func flipBits(binary []int) (r []int) {
	r = make([]int, len(binary))
	for i := range binary {
		r[i] = 1 - binary[i]
	}
	return r
}
