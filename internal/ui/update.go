package ui

import (
	gitlabclient "GitPatcher/internal/gitlab"
	"GitPatcher/internal/ui/components"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/xanzy/go-gitlab"
)

type msgConnected struct {
	client *gitlab.Client
}

type msgGroupsLoaded struct{}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// 全局退出
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	switch m.step {
	case stepLogin:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			if msg.String() == "enter" {
				if m.urlInput.Focused() {
					m.urlInput.Blur()
					m.tokenInput.Focus()
					return m, nil
				}

				return m, func() tea.Msg {
					c, err := gitlabclient.NewClient(
						m.urlInput.Value(),
						m.tokenInput.Value(),
					)
					if err != nil {
						return err
					}
					return msgConnected{client: c}
				}
			}
		}

		m.urlInput, _ = m.urlInput.Update(msg)
		m.tokenInput, _ = m.tokenInput.Update(msg)

	case stepBrowse:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "up":
				if m.tree.Cursor > 0 {
					m.tree.Cursor--
				}
			case "down":
				if m.tree.Cursor < len(m.tree.Groups)-1 {
					m.tree.Cursor++
				}
			case "enter":
				_ = m.tree.ToggleGroup(m.tree.Cursor)
			}
		}
	}

	switch msg := msg.(type) {
	case msgConnected:
		m.client = msg.client
		m.tree = components.NewTree(m.client)
		return m, func() tea.Msg {
			_ = m.tree.LoadGroups()
			return msgGroupsLoaded{}
		}

	case msgGroupsLoaded:
		m.step = stepBrowse
		return m, nil
	}

	return m, nil
}
