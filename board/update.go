package board

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "w":
			m.Cursor.moveUp(&m)
		case "down", "s":
			m.Cursor.moveDown(&m)
		case "left", "a":
			m.Cursor.moveLeft(&m)
		case "right", "d":
			m.Cursor.moveRight(&m)
		case "f", "b":
			m.toggleFlag()
		case " ", "enter":
			m.revealCell()
		}
	}

	// If any cells can be revealed, do so and wait 100ms for visual effect
	if m.revealEmptyCellNeighbours() {
		return m, tea.Tick(time.Millisecond*100, func(t time.Time) tea.Msg {
			return msg
		})
	}
	return m, nil
}

func (board *Board) revealEmptyCellNeighbours() bool {
	var cellsToReveal []*Cell

	for y := 0; y < board.Height; y++ {
		for x := 0; x < board.Width; x++ {
			cell := board.Cells[y][x]
			if !cell.IsVisible || cell.SurroundingBombs > 0 {
				continue
			}
			for _, c := range board.getSurroundingCells(y, x) {
				if !c.IsVisible {
					cellsToReveal = append(cellsToReveal, c)
				}
			}

		}
	}

	for index := 0; index < len(cellsToReveal); index++ {
		cell := cellsToReveal[index]
		cell.IsVisible = true
		cell.IsFlagged = false
	}

	return len(cellsToReveal) > 0
}

func (board *Board) getSurroundingCells(posY int, posX int) []*Cell {
	var cellsToReveal []*Cell

	for offY := -1; offY <= 1; offY++ {
		for offX := -1; offX <= 1; offX++ {
			if offY != 0 || offX != 0 {
				Y := posY + offY
				X := posX + offX
				if Y < 0 || X < 0 || Y >= board.Height || X >= board.Width {
					continue
				}
				cellsToReveal = append(cellsToReveal, &board.Cells[posY+offY][posX+offX])
			}
		}
	}
	return cellsToReveal
}

func (board *Board) revealCell() {
	if board.CurrentCell.IsVisible || board.CurrentCell.IsFlagged {
		return
	}

	board.CurrentCell.IsVisible = true
	if board.CurrentCell.IsBomb {
		board.GameState = GameOverLoss
	}
}

func (board *Board) toggleFlag() {
	if !board.CurrentCell.IsVisible {
		board.CurrentCell.IsFlagged =
			!board.CurrentCell.IsFlagged
	}
}

func (cursor *Cursor) moveDown(board *Board) {
	if cursor.Y+1 < len(board.Cells) {
		cursor.Y++
		board.CurrentCell = &board.Cells[board.Cursor.Y][board.Cursor.X]
	}
}

func (cursor *Cursor) moveUp(board *Board) {
	if cursor.Y-1 >= 0 {
		cursor.Y--
		board.CurrentCell = &board.Cells[board.Cursor.Y][board.Cursor.X]
	}
}

func (cursor *Cursor) moveRight(board *Board) {
	if cursor.X+1 < len(board.Cells[0]) {
		cursor.X++
		board.CurrentCell = &board.Cells[board.Cursor.Y][board.Cursor.X]
	}
}

func (cursor *Cursor) moveLeft(board *Board) {
	if cursor.X-1 >= 0 {
		cursor.X--
		board.CurrentCell = &board.Cells[board.Cursor.Y][board.Cursor.X]
	}
}
