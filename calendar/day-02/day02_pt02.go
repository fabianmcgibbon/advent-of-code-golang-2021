package main

import (
	"adventofcode/utils/files"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := files.ReadFile(2, "\n")
	horPos, depth := getLocationWithAim(input)
	fmt.Println(horPos, depth)
	solution := horPos * depth
	fmt.Println(solution)
}

func getLocationWithAim(input []string) (horPos, depth int) {
	var aim int

	for _, line := range input {
		split := strings.Split(line, " ")
		move := split[0]
		amount, _ := strconv.Atoi(split[1])

		switch move {
		case "down":
			aim += amount
		case "up":
			aim -= amount
		case "forward":
			horPos += amount
			depth += aim * amount
		}
	}
	return horPos, depth
}
