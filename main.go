package main

import (
	"fmt"
	"math/rand"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

// type model struct {
// 	choices  []string         // items on the to-do list
// 	cursor   int              // which to-do list item our cursor is pointing at
// 	selected map[int]struct{} // which to-do items are selected
// }

type cell struct {
	isBomb           bool
	isClosed         bool
	surroundingBombs int
}

type model struct {
	cells [][]cell
}

func generateMinefield(height int, width int) [][]cell {
	chanceOfBomb := 10
	var minefield = make([][]cell, height)
	for h := 0; h < height; h++ {
		minefield[h] = make([]cell, width)
		for w := 0; w < width; w++ {
			c := cell{isBomb: rand.Intn(100) < chanceOfBomb}
			minefield[h][w] = c
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

func fillSpacer(start string, separator string, end string, count int) string {
	runes := make([]rune, count)
	for i := 0; i < count; i++ {
		runes[i] = '─'
	}
	return fillLine(start, separator, end, runes)
}

func fillMinefieldLine(cells []cell) string {
	runes := make([]rune, len(cells))
	for i, c := range cells {
		if c.isBomb {
			runes[i] = 'B'
		} else {
			runes[i] += rune(c.surroundingBombs + '0')
		}
	}
	return fillLine("│", "│", "│", runes)
}

func fillLine(start string, separator string, end string, fill []rune) string {
	s := start
	for i, r := range fill {
		if i != 0 {
			s += separator
		}
		s += string(r)
	}
	s += end
	s += "\n"
	return s
}

func (m model) View() string {
	s := fillSpacer("┌", "┬", "┐", len(m.cells[0]))

	for i, line := range m.cells {
		if i != 0 {
			s += fillSpacer("├", "┼", "┤", len(line))
		}
		s += fillMinefieldLine(line)
	}

	s += fillSpacer("└", "┴", "┘", len(m.cells[0]))
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
