package ui

import "fmt"

func (m Model) View() string {
	if m.err != nil {
		return "Error: " + m.err.Error()
	}

	switch m.step {

	case stepLogin:
		return fmt.Sprintf(
			"GitPatcher v0.1\n\nGitLab URL:\n%s\n\nToken:\n%s\n\n[Enter] Continue\n",
			m.urlInput.View(),
			m.tokenInput.View(),
		)

	case stepBrowse:
		return m.tree.View()
	}

	return ""
}
