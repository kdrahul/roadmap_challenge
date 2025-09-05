package main

import (
	"fmt"
	"log/slog"
	"math/rand"
	"strconv"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	redColor = lipgloss.NewStyle().
		// Bold(true).
		// Background(lipgloss.Color("#00FF00")).
		Foreground(lipgloss.Color("#FF0000"))
	mainStyle = lipgloss.NewStyle().MarginLeft(9)
)

type (
	errMsg error
)

var correctNumber int = rand.Intn(100)

type model struct {
	viewname         string
	diffLevel        int64
	guessNo          int64
	textInput        textinput.Model
	err              error
	attemptsLeft     int
	guessedCorrectly bool
	firstAttempt     bool
}

func returnHi() tea.Msg {
	return "hi"
}

func (m model) Init() tea.Cmd {
	m.guessedCorrectly = false
	m.firstAttempt = true
	return textinput.Blink
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.viewname == "init" {
				value, err := strconv.ParseInt(m.textInput.Value(), 10, 64)
				if err != nil {
					slog.Error(err.Error())
				}
				m.diffLevel = value
				switch m.diffLevel {
				case 1:
					m.attemptsLeft = 10
				case 2:
					m.attemptsLeft = 5
				case 3:
					m.attemptsLeft = 3
				}
				if value < 0 || value > 3 {
					fmt.Println("Wrong input. Try again!")
					time.Sleep(1 * time.Second)
					m.viewname = "init"
					m.textInput.Reset()
				} else {
					m.viewname = "game"
					m.guessNo = 0
					m.textInput.Reset()
				}
			} else if m.viewname == "game" {
				m.textInput.Focus()
				m.firstAttempt = false
				v, err := strconv.ParseInt(m.textInput.Value(), 10, 64)
				if err != nil {
					slog.Error(err.Error())
				}
				m.textInput.Reset()
				m.guessNo = v
				if m.guessNo != int64(correctNumber) {
					m.attemptsLeft = m.attemptsLeft - 1
					if m.attemptsLeft < 1 {
						m.viewname = "end"
					}
				} else {
					m.guessedCorrectly = true
					m.viewname = "end"
				}
				m.textInput, cmd = m.textInput.Update(msg)
				return m, cmd
			} else {

				return m, tea.Quit
			}
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	var s string
	switch m.viewname {
	case "init":
		s = fmt.Sprintf("Welcome to the Number Guessing Game!\nI'm thinking of a number between 1 and 100.\nYou have 5 chances to guess the correct number.\nPlease select the difficulty level:\n1. Easy (10 chances)\n2. Medium (5 chances)\n3. Hard (3 chances)\n\n%s\n\n%s", m.textInput.View(), "(Esc to quit)") + "\n"

	case "game":
		var level string
		switch m.diffLevel {
		case 1:
			level = "Great! You have chosen Easy (10 chances)\n"
		case 2:
			level = "Great! You have chosen Medium (5 chances)\n"
		case 3:
			level = "Great! You have chosen Hard (3 chances)\n"
		}

		attempts := fmt.Sprintf("\nAttempts left: %d\n\n", m.attemptsLeft)
		startGame := "Let's start the game\n\n"
		var attemptSummary string
		if !m.firstAttempt {
			if m.guessNo > int64(correctNumber) {
				attemptSummary = fmt.Sprintf("Wrong answer. The number is lesser than %d\n\n", m.guessNo)
			} else if m.guessNo < int64(correctNumber) {
				attemptSummary = fmt.Sprintf("Wrong answer. The number is greater than %d\n\n", m.guessNo)
			}
			startGame = ""
		}
		m.textInput.Placeholder = "Guess here"
		m.textInput.CharLimit = 3
		m.textInput.Width = 10
		// s = level + redColor.Render(attempts) + attemptSummary + startGame + m.textInput.View()
		s = level + attempts + attemptSummary + startGame + m.textInput.View()

	case "end":
		var guesser string
		if m.guessedCorrectly {
			guesser = fmt.Sprintf("Yep! You guessed it right. The correct answer was: %d\n", correctNumber)
		} else {
			guesser = fmt.Sprintf("Aww, better luck next time. The correct answer was: %d\n", correctNumber)
		}
		s = guesser + "Bye!"
	}
	return mainStyle.Render("\n" + s + "\n")
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter a number"
	ti.Focus()
	ti.CharLimit = 20
	ti.Width = 30

	return model{
		viewname:     "init",
		textInput:    ti,
		err:          nil,
		firstAttempt: true,
	}

}
func main() {

	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		slog.Error("Error running program")
	}
}
