package cli

import (
    "fmt"
    tea "github.com/charmbracelet/bubbletea"
    color "kra/cli/color"
)

var c = color.GetColorSprintFuncs()

type Model struct {
    choices  []string           // items on the to-do list
    cursor   int                // which to-do list item our cursor is pointing at
    Selected []bool   // which to-do items are Selected
}

func InitialModel(options []string) Model {
	return Model{
		choices:  options,
		Selected: make([]bool, len(options)),
	}
}

func (m Model) Init() tea.Cmd {
    return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c":
            panic("ouq")
        
        case "ctrl+d":
            return m, tea.Quit

        case "enter":
            m.Selected[m.cursor] = !m.Selected[m.cursor]

        case "up":
            if m.cursor > 0 {
                m.cursor--
            }

        case "down":
            if m.cursor < len(m.choices)-1 {
                m.cursor++
            }

        case "right":
            for i := range m.Selected {
                m.Selected[i] = !m.Selected[i]
            }
        }
    }
    return m, nil
}

func (m Model) View() string {
    s := c.Yellow("What we do?") + "\n\n"
    for i, choice := range m.choices {
        cursor := " " // no cursor
        if m.cursor == i {
            cursor = c.Green(">") // cursor!
        }

        checked := " " // not Selected
        if m.Selected[i] {
            checked = c.Green("x") // Selected!
        }

        cchoise := choice
        if m.Selected[i] {
            cchoise = c.Green(cchoise)
        }

        s += fmt.Sprintf("%s " + c.Cyan("[") + "%s" + c.Cyan("]") + " %s\n", cursor, checked, cchoise)
    }

    s += "\n" + c.Yellow("Press ctrl+d to submit.") + "\n"

    return s
}