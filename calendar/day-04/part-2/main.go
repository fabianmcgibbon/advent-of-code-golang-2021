package main

import (
	"adventofcode/utils/conv"
	"adventofcode/utils/files"
	"fmt"
	"strconv"
	"strings"
)

const boardSize = 5

type game struct {
	calls  []int
	boards []bingoBoard
}
type bingoBoard struct {
	board             [boardSize][boardSize]cell
	horizontalMatches [boardSize]int
	verticalMatches   [boardSize]int
	won               bool
}

type cell struct {
	element int
	match   bool
}

func main() {
	input := files.ReadFile(4, "\n")
	solution := solvePart4(input)
	fmt.Println(solution)
}

func solvePart4(input []string) (solution int) {
	var g game

	g.calls = conv.ToIntSlice(strings.Split(input[0], ","))
	fmt.Println(g.calls)

	totalBoardInput := input[1:]

	// setup boards
	for k := 0; k < len(totalBoardInput)/(boardSize+1); k++ {
		g.boards = append(g.boards, bingoBoard{})

		increment := (boardSize + 1) * k
		boardInput := totalBoardInput[1+increment : 6+increment]
		for i := 0; i < boardSize; i++ {
			row := strings.Fields(boardInput[i])
			for j := 0; j < boardSize; j++ {
				g.boards[k].board[i][j].element, _ = strconv.Atoi(row[j])
			}
		}
	}

	// go through calls
	lastCall, winningBoardIndex := runGame(g)
	sum := sumUnmarkedNumbers(g.boards[winningBoardIndex])

	solution = lastCall * sum

	return solution
}

func sumUnmarkedNumbers(winningBoard bingoBoard) (sum int) {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if winningBoard.board[i][j].match == false {
				sum += winningBoard.board[i][j].element
			}
		}
	}
	return sum
}

func runGame(g game) (int, int) {
	// go through calls
	for _, call := range g.calls {

		for k := 0; k < len(g.boards); k++ {
			for i := 0; i < boardSize; i++ {
				for j := 0; j < boardSize; j++ {
					if g.boards[k].board[i][j].element == call {
						g.boards[k].board[i][j].match = true
						g.boards[k].horizontalMatches[i]++
						g.boards[k].verticalMatches[j]++
						if checkWin(g.boards[k]) {
							g.boards[k].won = true
							if checkLastWin(g) {
								return call, k
							}
						}
					}
				}
			}
		}
	}

	return -1, -1
}

func checkLastWin(g game) bool {
	for i := 0; i < len(g.boards); i++ {
		if g.boards[i].won == false {
			return false
		}
	}
	return true
}

func checkWin(b bingoBoard) bool {
	for i := 0; i < boardSize; i++ {
		if b.horizontalMatches[i] == boardSize || b.verticalMatches[i] == boardSize {
			return true
		}
	}
	return false
}
