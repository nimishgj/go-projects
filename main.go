package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	createBoard()
}

func createBoard() {
	board := make([][]int, 3)
	for i := range board {
		board[i] = make([]int, 3)
	}

	occupied := make(map[string]bool)

	for len(occupied) < 4 {
		generatedValue := rand.Intn(10)
		generatedRowIndex := rand.Intn(3)
		generatedColumnIndex := rand.Intn(3)
		position := fmt.Sprintf("%d-%d", generatedRowIndex, generatedColumnIndex)

		if !occupied[position] {
			board[generatedRowIndex][generatedColumnIndex] = generatedValue
			occupied[position] = true
		}
	}

	printFmt(board)
}

func printFmt(board [][]int) {
	fmt.Println("---------------")
	for rowIndex := 0; rowIndex < 3; rowIndex++ {
		for columnIndex := 0; columnIndex < 3; columnIndex++ {
			fmt.Printf("| %v |", board[rowIndex][columnIndex])
		}
		fmt.Println("\n---------------")
	}
}
