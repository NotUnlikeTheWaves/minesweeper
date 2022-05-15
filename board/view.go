package board

import (
	"fmt"
	"strings"

	"github.com/NotUnlikeTheWaves/minesweeper/config"
)

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

func (t Token) asSelected() Token {
	t.IsSelected = true
	return t
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
			render = fmt.Sprintf("\033[%d;%dm%s", config.Config.Bold, 38, render)
		}
		if content == "B" {
			render = fmt.Sprintf("\033[%d;4;%dm%s\033[24m", config.Config.Bold, config.Config.ColourBomb, render)
		}
		return render
	}

	// assume structural piece otherwise
	return fmt.Sprintf("\033[%d;%dm%s", config.Config.Bold, config.Config.ColourOutline, content)
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

type Token struct {
	Content    string
	IsSelected bool
}

func tokenize(content string) Token {
	return Token{Content: content}
}
