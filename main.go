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

type cell struct {
	isBomb           bool
	isChecked        bool
	isRevealed       bool
	surroundingBombs int
}

type position struct {
	x int
	y int
}
type model struct {
	cells    [][]cell
	selected position
}

type token struct {
	content    string
	isSelected bool
}

func tokenize(content string) token {
	return token{content: content}
}

func (t token) asSelected() token {
	t.isSelected = true
	return t
}

func returnOneIfEmptyAndCellExists(minefield [][]cell, y int, x int) int {
	if y < 0 || x < 0 || y >= len(minefield) || x >= len(minefield[0]) {
		return 0
	}
	if minefield[y][x].isBomb {
		return 1
	}
	return 0
}

func generateMinefield(height int, width int) [][]cell {
	chanceOfBomb := 10
	var minefield = make([][]cell, height)
	for h := range minefield {
		minefield[h] = make([]cell, width)
		for w := range minefield[h] {
			c := cell{isBomb: rand.Intn(100) < chanceOfBomb}
			minefield[h][w] = c
		}
	}
	for h := range minefield {
		for w := range minefield[h] {
			for y := -1; y <= 1; y++ {
				for x := -1; x <= 1; x++ {
					if x != 0 || y != 0 {
						minefield[h][w].surroundingBombs +=
							returnOneIfEmptyAndCellExists(minefield, h+y, w+x)
					}
				}
			}
		}
	}
	return minefield
}

func initialModel() model {
	return model{
		cells: generateMinefield(10, 40),
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func fillSpacer(start token, separator token, end token, count int) string {
	runes := make([]token, count)
	for i := 0; i < count; i++ {
		runes[i] = tokenize("───")
	}
	return fillLine(start, separator, end, runes)
}

func colour(t token) string {
	content := t.content
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

func fillLine(start token, separator token, end token, fill []token) string {
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

func (m model) View() string {
	s := fillSpacer(tokenize("┌"), tokenize("┬"), tokenize("┐"), len(m.cells[0]))

	for y, line := range m.cells {
		if y != 0 {
			s += fillSpacer(tokenize("├"), tokenize("┼"), tokenize("┤"), len(line))
		}
		tokens := make([]token, len(line))
		for x, c := range line {
			if c.isBomb {
				tokens[x] = tokenize("B")
			} else {
				tokens[x] = tokenize(fmt.Sprintf("%d", c.surroundingBombs))
				if c.surroundingBombs == 0 {
					tokens[x] = tokenize(" ")
				}
			}
		}
		s += fillLine(tokenize("│ "), tokenize(" │ "), tokenize(" │"), tokens)
	}

	s += fillSpacer(tokenize("└"), tokenize("┴"), tokenize("┘"), len(m.cells[0]))
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
