package ui

import (
	"GitPatcher/internal/state"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type step int

const (
	stepInput step = iota
	stepProjects
)

type Model struct {
	step step
	err  error

	urlInput   textinput.Model
	tokenInput textinput.Model
	groupInput textinput.Model

	state *state.AppState
}

func NewModel() *Model {
	url := textinput.New()
	url.Placeholder = "http://gitlab.example.com"
	url.Focus()

	token := textinput.New()
	token.Placeholder = "glpat-xxxx"
	token.EchoMode = textinput.EchoPassword

	group := textinput.New()
	group.Placeholder = "group-name"

	return &Model{
		step:       stepInput,
		urlInput:   url,
		tokenInput: token,
		groupInput: group,
		state:      &state.AppState{},
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}
