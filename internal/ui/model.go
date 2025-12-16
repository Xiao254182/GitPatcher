package ui

import (
	"GitPatcher/internal/state"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type step int

const (
	stepLogin step = iota
	stepProjects
)

type Model struct {
	step    step
	err     error
	loading bool

	urlInput   textinput.Model
	tokenInput textinput.Model

	cursor int
	state  *state.AppState
}

func NewModel() Model {
	url := textinput.New()
	url.Placeholder = "http://gitlab.example.com"
	url.Focus()

	token := textinput.New()
	token.Placeholder = "glpat-xxxx"
	token.EchoMode = textinput.EchoPassword

	return Model{
		step:       stepLogin,
		urlInput:   url,
		tokenInput: token,
		state: &state.AppState{
			Selected: make(map[int]bool),
			Branch:   "dev",
			DryRun:   true,
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
