package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type configuration struct {
	bold          int
	colourOutline int
	colourBomb    int
}

// https://www.tutorialspoint.com/how-to-output-colored-text-to-a-linux-terminal
var config = configuration{
	bold:          1,
	colourOutline: 30,
	colourBomb:    36,
}

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

type Token struct {
	Content    string
	IsSelected bool
}

func tokenize(content string) Token {
	return Token{Content: content}
}

func (t Token) asSelected() Token {
	t.IsSelected = true
	return t
}

func returnOneIfEmptyAndCellExists(minefield [][]Cell, y int, x int) int {
	if y < 0 || x < 0 || y >= len(minefield) || x >= len(minefield[0]) {
		return 0
	}
	if minefield[y][x].IsBomb {
		return 1
	}
	return 0
}

func generateMinefield(height int, width int) [][]Cell {
	chanceOfBomb := 10
	var minefield = make([][]Cell, height)
	for h := range minefield {
		minefield[h] = make([]Cell, width)
		for w := range minefield[h] {
			c := Cell{IsBomb: rand.Intn(100) < chanceOfBomb}
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

func initialModel() Board {
	return Board{
		Cells: generateMinefield(10, 40),
	}
}

func (m Board) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m Board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func fillSpacer(start Token, separator Token, end Token, count int) string {
	runes := make([]Token, count)
	for i := 0; i < count; i++ {
		runes[i] = tokenize("───")
	}
	return fillLine(start, separator, end, runes)
}

func colour(t Token) string {
	content := t.Content
	if strings.TrimSpace(content) == "" {
		return content
	}
	if len(content) == 1 {
		render := content
		if content[0] > '0' && content[0] < '9' {
			render = fmt.Sprintf("\033[%d;%dm%s", config.bold, 38, render)
		}
		if content == "B" {
			render = fmt.Sprintf("\033[%d;4;%dm%s\033[24m", config.bold, config.colourBomb, render)
		}
		return render
	}

	// assume structural piece otherwise
	return fmt.Sprintf("\033[%d;%dm%s", config.bold, config.colourOutline, content)
}

func fillLine(start Token, separator Token, end Token, fill []Token) string {
	s := colour(start)
	for i, r := range fill {
		if i != 0 {
			s += colour(separator)
		}
		s += colour(r)
	}
	s += colour(end)
	s += "\n"
	return s
}

func (m Board) View() string {
	s := fillSpacer(tokenize("┌"), tokenize("┬"), tokenize("┐"), len(m.Cells[0]))

	for y, line := range m.Cells {
		if y != 0 {
			s += fillSpacer(tokenize("├"), tokenize("┼"), tokenize("┤"), len(line))
		}
		tokens := make([]Token, len(line))
		for x, c := range line {
			if c.IsBomb {
				tokens[x] = tokenize("B")
			} else {
				tokens[x] = tokenize(fmt.Sprintf("%d", c.SurroundingBombs))
				if c.SurroundingBombs == 0 {
					tokens[x] = tokenize(" ")
				}
			}
		}
		s += fillLine(tokenize("│ "), tokenize(" │ "), tokenize(" │"), tokens)
	}

	s += fillSpacer(tokenize("└"), tokenize("┴"), tokenize("┘"), len(m.Cells[0]))
	// Send the UI for rendering
	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
