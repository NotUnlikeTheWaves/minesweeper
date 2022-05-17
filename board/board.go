package board

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Cell struct {
	IsBomb           bool
	IsFlagged        bool
	IsVisible        bool
	SurroundingBombs int
}

type Cursor struct {
	X int
	Y int
}

type Board struct {
	Cells       [][]Cell
	Cursor      Cursor
	CurrentCell *Cell
}

func (m Board) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
