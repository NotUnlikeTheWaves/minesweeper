package board

import tea "github.com/charmbracelet/bubbletea"

func (m Board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "w":
			m.Cursor.moveUp(&m)
			return m, nil
		case "down", "s":
			m.Cursor.moveDown(&m)
			return m, nil
		case "left", "a":
			m.Cursor.moveLeft(&m)
			return m, nil
		case "right", "d":
			m.Cursor.moveRight(&m)
			return m, nil
		case "f", "b":
			m.toggleFlag()
			return m, nil
		case "o":
			m.revealCell()
			return m, nil
		}

	}

	return m, nil
}

func (board *Board) revealCell() {
	if board.CurrentCell.IsVisible || board.CurrentCell.IsFlagged {
		return
	}

	board.CurrentCell.IsVisible = true
}

func (board *Board) toggleFlag() {
	board.CurrentCell.IsFlagged =
		!board.CurrentCell.IsFlagged
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
