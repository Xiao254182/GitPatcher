package ui

import "fmt"

func (m Model) View() string {

	if m.err != nil {
		return fmt.Sprintf("❌ Error: %v\n\n[q] quit", m.err)
	}

	switch m.step {

	case stepLogin:
		return fmt.Sprintf(
			"GitPatcher v0.2\n\nGitLab URL:\n%s\n\nToken:\n%s\n\n[Enter] Continue  [q] Quit\n",
			m.urlInput.View(),
			m.tokenInput.View(),
		)

	case stepBrowse:
		return fmt.Sprintf(
			"GitPatcher v0.2\n\nGroups / Projects:\n\n%s\n\n↑↓ Move  Enter Toggle  q Quit\n",
			m.tree.View(),
		)
	}

	return ""
}
