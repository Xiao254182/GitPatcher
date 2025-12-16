package ui

import "fmt"

func (m *Model) View() string {
	switch m.step {

	case stepInput:
		return fmt.Sprintf(
			"GitPatcher v0.1\n\nGitLab URL:\n%s\n\nToken:\n%s\n\nGroup:\n%s\n\n[Enter] Connect\n",
			m.urlInput.View(),
			m.tokenInput.View(),
			m.groupInput.View(),
		)

	case stepProjects:
		out := "Projects:\n\n"
		for _, p := range m.state.Projects {
			out += "• " + p.PathWithNamespace + "\n"
		}
		return out
	}
	return ""
}
