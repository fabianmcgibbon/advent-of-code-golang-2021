package main

import (
	"adventofcode/utils/files"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const gridSize = 1000

type coord struct {
	x int
	y int
}

type floor struct {
	grid        [gridSize][gridSize]int
	dangerAreas int
}

func main() {
	input := files.ReadFile(5, "\n")
	f := generateFloor(input)
	fmt.Println(f.dangerAreas)
}

func generateFloor(input []string) floor {
	var f floor

	for _, entry := range input {
		start, end := parseEntry(entry)
		if start.x == end.x {
			f.logY(start, end)
		} else if start.y == end.y {
			f.logX(start, end)
		} else {
			f.logXY(start, end)
		}
	}

	return f
}

func parseEntry(entry string) (start, end coord) {
	coordinates := strings.Split(entry, " -> ")

	s := strings.Split(coordinates[0], ",")
	e := strings.Split(coordinates[1], ",")

	start.x, _ = strconv.Atoi(s[0])
	start.y, _ = strconv.Atoi(s[1])
	end.x, _ = strconv.Atoi(e[0])
	end.y, _ = strconv.Atoi(e[1])

	fmt.Println(start)
	fmt.Println(end)

	return start, end
}

func (f *floor) logX(start, end coord) {
	if start.x > end.x {
		start, end = end, start
	}
	for x := start.x; x <= end.x; x++ {
		f.grid[start.y][x]++
		if f.grid[start.y][x] == 2 {
			f.dangerAreas++
		}
	}
}

func (f *floor) logY(start, end coord) {
	if start.y > end.y {
		start, end = end, start
	}
	for y := start.y; y <= end.y; y++ {
		f.grid[y][start.x]++
		if f.grid[y][start.x] == 2 {
			f.dangerAreas++
		}
	}
}

func (f *floor) logXY(start, end coord) {
	dx, dy := 1, 1
	l := int(math.Abs(float64(end.x) - float64(start.x)))

	if start.x > end.x {
		dx = -1
	}
	if start.y > end.y {
		dy = -1
	}

	for i := 0; i <= l; i++ {
		f.grid[start.y+(dy*i)][start.x+(dx*i)]++
		if f.grid[start.y+(dy*i)][start.x+(dx*i)] == 2 {
			f.dangerAreas++
		}
	}
}
