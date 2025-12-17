package ui

import (
	"GitPatcher/internal/ui/components"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/xanzy/go-gitlab"
)

type Step int

const (
	stepLogin Step = iota
	stepBrowse
)

type Model struct {
	step Step
	err  error

	urlInput   textinput.Model
	tokenInput textinput.Model

	client *gitlab.Client
	tree   *components.Tree
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
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
