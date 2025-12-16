package ui

import (
	gitlabclient "GitPatcher/internal/gitlab"
	"GitPatcher/internal/state"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.step {

	case stepInput:
		switch msg := msg.(type) {

		case tea.KeyMsg:
			switch msg.String() {

			case "tab", "enter":
				if m.urlInput.Focused() {
					m.urlInput.Blur()
					m.tokenInput.Focus()
					return m, nil
				}
				if m.tokenInput.Focused() {
					m.tokenInput.Blur()
					m.groupInput.Focus()
					return m, nil
				}
				if m.groupInput.Focused() {
					// 最后一个，执行连接
					m.groupInput.Blur()

					m.state.GitlabURL = m.urlInput.Value()
					m.state.GitlabToken = m.tokenInput.Value()
					m.state.Group = m.groupInput.Value()

					return m, fetchProjects(m.state)
				}

			case "shift+tab":
				if m.groupInput.Focused() {
					m.groupInput.Blur()
					m.tokenInput.Focus()
				} else if m.tokenInput.Focused() {
					m.tokenInput.Blur()
					m.urlInput.Focus()
				}
			}
		}

		m.urlInput, _ = m.urlInput.Update(msg)
		m.tokenInput, _ = m.tokenInput.Update(msg)
		m.groupInput, _ = m.groupInput.Update(msg)

	case stepProjects:
		// v0.1 仅展示
	}

	switch msg.(type) {
	case projectsLoadedMsg:
		m.step = stepProjects
	}

	return m, nil
}

type projectsLoadedMsg struct{}

func fetchProjects(s *state.AppState) tea.Cmd {
	return func() tea.Msg {
		client, err := gitlabclient.NewClient(s.GitlabURL, s.GitlabToken)
		if err != nil {
			return err
		}
		projects, err := gitlabclient.ListGroupProjects(client, s.Group)
		if err != nil {
			return err
		}
		s.Projects = projects
		return projectsLoadedMsg{}
	}
}
