package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/NotUnlikeTheWaves/minesweeper/board"
	tea "github.com/charmbracelet/bubbletea"
)

func returnOneIfEmptyAndCellExists(minefield [][]board.Cell, y int, x int) int {
	if y < 0 || x < 0 || y >= len(minefield) || x >= len(minefield[0]) {
		return 0
	}
	if minefield[y][x].IsBomb {
		return 1
	}
	return 0
}

// TODO: Generate a deterministic mine field
func generateMinefield(height int, width int) [][]board.Cell {
	chanceOfBomb := 10
	var minefield = make([][]board.Cell, height)
	for h := range minefield {
		minefield[h] = make([]board.Cell, width)
		for w := range minefield[h] {
			c := board.Cell{IsBomb: rand.Intn(100) < chanceOfBomb}
			minefield[h][w] = c
		}
	}
	for h := range minefield {
		for w := range minefield[h] {
			for y := -1; y <= 1; y++ {
				for x := -1; x <= 1; x++ {
					if x != 0 || y != 0 {
						minefield[h][w].SurroundingBombs +=
							returnOneIfEmptyAndCellExists(minefield, h+y, w+x)
					}
				}
			}
		}
	}
	return minefield
}

func initialBoard() board.Board {
	return board.Board{
		Cells: generateMinefield(10, 40),
	}
}

func main() {
	p := tea.NewProgram(initialBoard())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
