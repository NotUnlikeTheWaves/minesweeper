package board

const widthSpacing = 3

type tokenType int

const (
	TableSpace tokenType = iota
	TableComponent
	Neighbours
	Flag
)

func (m Board) View() string {
	height := len(m.Cells)*2 + 1
	width := len(m.Cells[0])*4 + 1
	viewModel := make([][]Token, height)
	for i := 0; i < len(viewModel); i++ {
		viewModel[i] = make([]Token, width)
	}

	numberOfElements := len(m.Cells[0])
	addStructuralRow(viewModel[0], numberOfElements, '┌', '┬', '┐')
	for h, boardRow := range m.Cells {
		if h != 0 {
			addStructuralRow(viewModel[h*2], numberOfElements, '├', '┼', '┤')
		}
		viewRow := viewModel[h*2+1]
		for w, cell := range boardRow {
			base := w * 4
			item := rune('0' + cell.SurroundingBombs)
			if cell.IsBomb {
				item = 'B'
			} else if cell.SurroundingBombs == 0 {
				item = ' '
			}
			viewRow[base] = Token{Content: '│', Type: TableComponent}
			viewRow[base+1] = Token{Content: ' ', Type: TableSpace}
			viewRow[base+2] = Token{Content: item, Type: Neighbours}
			viewRow[base+3] = Token{Content: ' ', Type: TableSpace}
		}
		viewRow[len(viewRow)-1] = Token{Content: '│', Type: TableComponent}
	}
	addStructuralRow(viewModel[height-1], numberOfElements, '└', '┴', '┘')

	s := ""
	for _, row := range viewModel {
		for _, cell := range row {
			s += string(cell.Content)
		}
		s += "\n"
	}
	// Send the UI for rendering
	return s
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