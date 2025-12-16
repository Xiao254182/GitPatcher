package ui

import (
	"GitPatcher/internal/gitlab"
	"GitPatcher/internal/state"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case error:
		m.loading = false
		m.err = msg
		return m, nil

	case string:
		if msg == "loaded" {
			m.loading = false
			m.step = stepProjects
			return m, nil
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

				m.state.GitlabURL = normalizeURL(m.urlInput.Value())
				m.state.GitlabToken = m.tokenInput.Value()
				m.loading = true
				m.err = nil

				return m, fetchAll(m.state)
			}
		}

		m.urlInput, _ = m.urlInput.Update(msg)
		m.tokenInput, _ = m.tokenInput.Update(msg)
	}

	return m, nil
}

func fetchAll(s *state.AppState) tea.Cmd {
	return func() tea.Msg {
		client, err := gitlabclient.New(s.GitlabURL, s.GitlabToken)
		if err != nil {
			return err
		}
		groups, err := gitlabclient.ListGroups(client)
		if err != nil {
			return err
		}
		s.Groups = groups

		for _, g := range groups {
			ps, _ := gitlabclient.ListGroupProjects(client, g.ID)
			s.Projects = append(s.Projects, ps...)
		}
		return "loaded"
	}
}

func normalizeURL(u string) string {
	u = strings.TrimSpace(u)      // 去掉首尾空格
	u = strings.TrimRight(u, "/") // 去掉尾部斜杠
	if !strings.HasSuffix(u, "/api/v4") {
		u += "/api/v4"
	}
	return u
}
