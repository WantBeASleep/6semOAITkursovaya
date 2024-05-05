package terminal

import (
    "fmt"
    tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
    choices  []string           // items on the to-do list
    cursor   int                // which to-do list item our cursor is pointing at
    Selected map[int]struct{}   // which to-do items are Selected
}

func InitialModel(options []string) Model {
	return Model{
		choices:  options,
		Selected: make(map[int]struct{}),
	}
}

func (m Model) Init() tea.Cmd {
    return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "enter":
            return m, tea.Quit

        case "up":
            if m.cursor > 0 {
                m.cursor--
            }

        case "down":
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            }

        case "right":
            _, ok := m.Selected[m.cursor]
            if ok {
                delete(m.Selected, m.cursor)
            } else {
                m.Selected[m.cursor] = struct{}{}
            }
        }
    }
    return m, nil
}

func (m Model) View() string {
    s := "What we do?\n\n"
    for i, choice := range m.choices {
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = ">" // cursor!
        }

        checked := " " // not Selected
        if _, ok := m.Selected[i]; ok {
            checked = "x" // Selected!
        }

        s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
    }

    s += "\nPress enter to submit.\n"

    return s
}