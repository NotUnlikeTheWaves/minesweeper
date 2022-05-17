package board

import tea "github.com/charmbracelet/bubbletea"

func (m Board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "w":
			m.Cursor.moveUp(m)
			return m, nil
		case "down", "s":
			m.Cursor.moveDown(m)
			return m, nil
		case "left", "a":
			m.Cursor.moveLeft(m)
			return m, nil
		case "right", "d":
			m.Cursor.moveRight(m)
			return m, nil
		}
	}

	return m, nil
}

func (cursor *Cursor) moveDown(board Board) {
	if cursor.Y+1 < len(board.Cells) {
		cursor.Y++
	}
}

func (cursor *Cursor) moveUp(board Board) {
	if cursor.Y-1 >= 0 {
		cursor.Y--
	}
}

func (cursor *Cursor) moveRight(board Board) {
	if cursor.X+1 < len(board.Cells[0]) {
		cursor.X++
	}
}

func (cursor *Cursor) moveLeft(board Board) {
	if cursor.X-1 >= 0 {
		cursor.X--
	}
}
