package ui

import (
	"GitPatcher/internal/state"
	"GitPatcher/internal/ui/components"

	"github.com/charmbracelet/bubbles/textinput"
)

type step int
type Focus int
type loginStep int

const (
	loginURL loginStep = iota
	loginToken
)
const (
	stepLogin step = iota
	stepBrowse
)

const (
	FocusTree Focus = iota
	FocusConfig
	FocusDiff
)

type Model struct {
	step step

	loginStep loginStep
	focus     Focus
	err       error

	urlInput   textinput.Model
	tokenInput textinput.Model

	tree   *components.Tree
	config *components.Config
	diff   *components.DiffModel

	app *state.AppState
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
		loginStep:  loginURL,
		focus:      FocusTree,
		urlInput:   url,
		tokenInput: token,
		tree:       components.NewTree(),
		config:     components.NewConfig(),
		diff:       components.NewDiff(),
		app: &state.AppState{
			SelectedProjects: make(map[int]bool),
			DryRun:           true,
		},
	}
}
