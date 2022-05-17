package board

import "fmt"

type tokenType int

const (
	TableSpace tokenType = iota
	TableComponent
	Neighbours
	Flag
	Unkown
	Empty
)

func (m Board) View() string {
	height := len(m.Cells)*2 + 1
	width := len(m.Cells[0])*4 + 1
	viewModel := make([][]Token, height)
	for i := 0; i < len(viewModel); i++ {
		viewModel[i] = make([]Token, width)
	}

	// Plomp stuff down
	numberOfElements := len(m.Cells[0])
	addStructuralRow(viewModel[0], numberOfElements, '┌', '┬', '┐')
	for h, boardRow := range m.Cells {
		if h != 0 {
			addStructuralRow(viewModel[h*2], numberOfElements, '├', '┼', '┤')
		}
		viewRow := viewModel[h*2+1]
		for w, cell := range boardRow {
			base := w * 4
			item := createBoardPiece(cell)
			viewRow[base] = Token{Content: '│', Type: TableComponent}
			viewRow[base+1] = Token{Content: ' ', Type: TableSpace}
			viewRow[base+2] = item
			viewRow[base+3] = Token{Content: ' ', Type: TableSpace}
		}
		viewRow[len(viewRow)-1] = Token{Content: '│', Type: TableComponent}
	}
	addStructuralRow(viewModel[height-1], numberOfElements, '└', '┴', '┘')

	// Select the 'selected' cell
	vmY, vmX := translateBoardPositionToViewModelPosition(m.Cursor.Y, m.Cursor.X)
	for offsetY := -1; offsetY <= 1; offsetY++ {
		for offsetX := -2; offsetX <= 2; offsetX++ {
			viewModel[vmY+offsetY][vmX+offsetX].IsSelected = true
		}
	}

	s := ""
	for _, row := range viewModel {
		for _, cell := range row {
			s += cell.print()
		}
		s += "\n"
	}

	return s
}

func createBoardPiece(cell Cell) Token {
	if cell.IsFlagged {
		return Token{Content: '⚑', Type: Flag}
	}
	if cell.IsVisible {
		item := rune('0' + cell.SurroundingBombs)
		itemType := Neighbours
		if cell.SurroundingBombs == 0 {
			item = ' '
			itemType = Empty
		}
		return Token{Content: item, Type: itemType}
	}
	return Token{Content: '?', Type: Unkown}
}

func translateBoardPositionToViewModelPosition(y int, x int) (int, int) {
	return (1 + y*2), (2 + x*4)
}

func (t Token) print() string {
	char := t.Content

	backgroundStyle := "\033[1;40m"
	foregroundStyle := "\033[1;37m"
	if t.IsSelected {
		backgroundStyle = "\033[1;44m"
	}
	if t.Type == Neighbours {
		foregroundStyle = fmt.Sprintf("\033[1;%sm", colourNeighbour(t.Content))
	}
	if t.Type == Flag {
		foregroundStyle = "\033[1;35m"
	}
	return fmt.Sprintf("%s%s%c", backgroundStyle, foregroundStyle, char)
}

func colourNeighbour(n rune) string {
	switch n {
	case '1':
		return "32"
	case '2':
		return "33"
	case '3':
		return "31"
	case '4':
		return "35"
	case '5':
		return "36"
	case '6':
		return "37"
	case '7':
		return "37"
	case '8':
		return "37"
	}
	return "37"
}

func addStructuralRow(row []Token, numberOfElements int, start rune, separator rune, end rune) {
	row[0] = Token{Content: start, Type: TableComponent}
	for i := 0; i < numberOfElements; i++ {
		base := 1 + 4*i
		row[base] = Token{Content: '─', Type: TableComponent}
		row[base+1] = Token{Content: '─', Type: TableComponent}
		row[base+2] = Token{Content: '─', Type: TableComponent}
		row[base+3] = Token{Content: separator, Type: TableComponent}
	}
	row[len(row)-1] = Token{Content: end, Type: TableComponent}
}

type Token struct {
	Content    rune
	Type       tokenType
	IsSelected bool
}
