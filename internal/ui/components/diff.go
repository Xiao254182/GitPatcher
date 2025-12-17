package components

import tea "github.com/charmbracelet/bubbletea"

type DiffModel struct{}

func NewDiff() DiffModel {
	return DiffModel{}
}

func (m DiffModel) Update(msg tea.Msg) (DiffModel, tea.Cmd) {
	return m, nil
}

func (m DiffModel) View(active bool) string {
	title := "Diff Preview"
	if active {
		title += " [ACTIVE]"
	}

	return title + `
┌──── OLD ────┐   ┌──── NEW ────┐
│             │   │             │
│             │   │             │
└─────────────┘   └─────────────┘`
}
