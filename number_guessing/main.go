package main

import (
	"log/slog"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	textInput textinput.Model
	err       error
}

func returnHi() tea.Msg {
	return "hi"
}

func (m model) Init() tea.Cmd {
	return returnHi
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	default:
		return m, nil
	}
	return m, nil
}

func (m model) View() string {
	return m.msg
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter a number"
	ti.Focus()
	ti.CharLimit = 20
	ti.Width = 30

	return model{
		textInput: ti,
		err:       nil,
	}

}
func main() {

	p := tea.NewProgram()
	if _, err := p.Run(); err != nil {
		slog.Error("Error running program")
	}
}
