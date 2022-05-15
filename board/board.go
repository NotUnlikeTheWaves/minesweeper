package board

import tea "github.com/charmbracelet/bubbletea"

type Cell struct {
	IsBomb           bool
	IsChecked        bool
	IsRevealed       bool
	SurroundingBombs int
}

type Position struct {
	X int
	Y int
}
type Board struct {
	Cells  [][]Cell
	Cursor Position
}

func (m Board) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
