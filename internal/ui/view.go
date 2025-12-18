package ui

import (
	"fmt"
	"strings"
)

func padRight(s string, width int) string {
	lines := strings.Split(s, "\n")
	for i, l := range lines {
		if len(l) < width {
			lines[i] = l + strings.Repeat(" ", width-len(l))
		}
	}
	return strings.Join(lines, "\n")
}

func (m Model) View() string {

	switch m.step {

	case stepLogin:
		return fmt.Sprintf(
			"GitPatcher v0.3\n\nGitLab URL:\n%s\n\nToken:\n%s\n\n[Enter] Login\n",
			m.urlInput.View(),
			m.tokenInput.View(),
		)

	case stepBrowse:
		left := m.tree.View(m.focus == FocusTree)
		mid := m.config.View(m.focus == FocusConfig)
		right := m.diff.View(m.focus == FocusDiff)

		left = padRight(left, 40)
		mid = padRight(mid, 40)

		return fmt.Sprintf(
			"GitPatcher v0.3\n\n%s │ %s │ %s\n\n← → Switch Panel | ↑↓ Enter Operate | q Quit\n",
			left,
			mid,
			right,
		)
	}

	return ""
}
