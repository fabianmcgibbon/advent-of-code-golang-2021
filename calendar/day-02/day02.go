package main

import (
	"adventofcode/utils/files"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(2, "\n")
	horPos, depth := getLocation(input)
	solution := horPos * depth
	fmt.Println(solution)
}

func getLocation(input []string) (horPos, depth int) {
	for _, line := range input {
		split := strings.Split(line, " ")
		dir := split[0]
		amount, _ := strconv.Atoi(split[1])

		switch dir {
		case "forward":
			horPos += amount
		case "up":
			depth -= amount
		case "down":
			depth += amount
		}
	}
	return horPos, depth
}
