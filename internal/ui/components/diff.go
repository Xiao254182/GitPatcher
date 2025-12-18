package components

import tea "github.com/charmbracelet/bubbletea"

type DiffModel struct{}

func NewDiff() *DiffModel {
	return &DiffModel{}
}

func (d *DiffModel) Update(msg tea.Msg) tea.Cmd {
	return nil
}

func (d *DiffModel) View(active bool) string {
	title := "Diff"
	if active {
		title += " [ACTIVE]"
	}

	return title + `
┌──── OLD ────┐
│             │
│   (empty)   │
└─────────────┘

┌──── NEW ────┐
│             │
│   (empty)   │
└─────────────┘`
}
