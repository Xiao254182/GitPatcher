package ui

import (
	gitlabclient "GitPatcher/internal/gitlab"

	tea "github.com/charmbracelet/bubbletea"
)

type msgConnected struct{}
type msgGroupsLoaded struct{}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// 全局退出
	if key, ok := msg.(tea.KeyMsg); ok {
		if key.String() == "ctrl+c" || key.String() == "q" {
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
					m.app.Client = c
					return msgConnected{}
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
				if m.tree.Cursor < len(m.tree.Flat)-1 {
					m.tree.Cursor++
				}
			case "enter":
				m.tree.Toggle()
			}
		}
	}

	switch msg.(type) {

	case msgConnected:
		return m, func() tea.Msg {
			groups, err := gitlabclient.ListGroups(m.app.Client)
			if err != nil {
				return err
			}
			m.tree.Groups = groups
			for _, g := range groups {
				ps, _ := gitlabclient.ListGroupProjects(m.app.Client, g.ID)
				m.tree.Projects[g.ID] = ps
			}
			m.tree.Build()
			return msgGroupsLoaded{}
		}

	case msgGroupsLoaded:
		m.step = stepBrowse
		return m, nil
	}

	return m, nil
}
