package main

import (
	"adventofcode/utils/files"
	"fmt"
)

func main() {
	input := files.ReadFile(3, "\n")
	o2Rating := getOxygenGeneratorRating(input)
	co2Rating := getCO2ScrubberRating(input)

	fmt.Println(o2Rating, co2Rating)
}

func getOxygenGeneratorRating(input []string) (rating string) {
	bitLength := len(input[0])
	inputLength := len(input)
	var currentLine string

	for i := 0; i < bitLength; i++ {
		ones := 0
		zeros := 0
		for j := 0; j < inputLength; j++ {
			char := fmt.Sprintf("%c", input[j][i])
			previousChar := fmt.Sprintf("%s", input[j][:i])
			if previousChar != rating {
				continue
			}

			currentLine = input[j]

			if char == "1" {
				ones++
			} else {
				zeros++
			}
		}
		if ones+zeros == 1 {
			return fmt.Sprintf("%s", currentLine)
		}
		if ones >= zeros {
			rating += "1"
		} else {
			rating += "0"
		}
	}

	return rating
}

func getCO2ScrubberRating(input []string) (rating string) {
	bitLength := len(input[0])
	inputLength := len(input)
	var currentLine string

	for i := 0; i < bitLength; i++ {
		ones := 0
		zeros := 0
		for j := 0; j < inputLength; j++ {
			char := fmt.Sprintf("%c", input[j][i])
			previousChars := fmt.Sprintf("%s", input[j][:i])
			if previousChars != rating {
				continue
			}

			currentLine = input[j]

			if char == "1" {
				ones++
			} else {
				zeros++
			}
		}
		if ones+zeros == 1 {
			return fmt.Sprintf("%s", currentLine)
		}
		if ones >= zeros {
			rating += "0"
		} else {
			rating += "1"
		}
	}

	return rating
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
