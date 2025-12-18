package ui

import (
	gitlabclient "GitPatcher/internal/gitlab"

	tea "github.com/charmbracelet/bubbletea"
)

type msgConnected struct{}
type msgLoaded struct{}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// 全局退出
	if key, ok := msg.(tea.KeyMsg); ok {
		switch key.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "left":
			if m.focus > FocusTree {
				m.focus--
			}
		case "right":
			if m.focus < FocusDiff {
				m.focus++
			}
		}
	}

	switch m.step {

	case stepLogin:
		switch msg := msg.(type) {

		case tea.KeyMsg:
			switch msg.String() {

			case "enter":
				switch m.loginStep {

				case loginURL:
					// 从 URL → Token
					m.loginStep = loginToken
					m.urlInput.Blur()
					m.tokenInput.Focus()
					return m, nil

				case loginToken:
					// Token 填完 → 连接 GitLab
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
		}

		// 输入更新
		m.urlInput, _ = m.urlInput.Update(msg)
		m.tokenInput, _ = m.tokenInput.Update(msg)

	case stepBrowse:
		switch m.focus {

		case FocusTree:
			if key, ok := msg.(tea.KeyMsg); ok {
				switch key.String() {
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

		case FocusConfig:
			return m, m.config.Update(msg)

		case FocusDiff:
			return m, m.diff.Update(msg)
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
			return msgLoaded{}
		}

	case msgLoaded:
		m.step = stepBrowse
	}

	return m, nil
}
