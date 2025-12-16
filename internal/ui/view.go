package ui

import (
	"fmt"
)

func (m Model) View() string {
	switch m.step {

	case stepLogin:
		out := fmt.Sprintf(
			"GitPatcher v0.1\n\nGitLab URL:\n%s\n\nToken:\n%s\n\n[Enter] Continue\n",
			m.urlInput.View(),
			m.tokenInput.View(),
		)

		if m.loading {
			out += "\n⏳ Loading projects...\n"
		}
		if m.err != nil {
			out += "\n❌ Error: " + m.err.Error() + "\n"
		}
		return out

	case stepProjects:
		out := "Projects (Space to select, q to quit)\n\n"
		for i, p := range m.state.Projects {
			cursor := " "
			if i == m.cursor {
				cursor = ">"
			}
			check := " "
			if m.state.Selected[p.ID] {
				check = "✔"
			}
			out += fmt.Sprintf("%s [%s] %s\n", cursor, check, p.PathWithNamespace)
		}
		return out
	}
	return ""
}
