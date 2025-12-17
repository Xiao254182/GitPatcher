package ui

import (
	"GitPatcher/internal/state"
	"GitPatcher/internal/ui/components"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/xanzy/go-gitlab"
)

type step int

const (
	stepLogin step = iota
	stepBrowse
)

type Model struct {
	step step
	err  error

	urlInput   textinput.Model
	tokenInput textinput.Model

	tree *components.Tree
	app  *state.AppState
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
		tree:       components.NewTree(),
		app: &state.AppState{
			Projects:         make(map[int][]*gitlab.Project),
			SelectedProjects: make(map[int]bool),
			ExpandedGroups:   make(map[int]bool),
		},
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
