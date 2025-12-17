package components

import tea "github.com/charmbracelet/bubbletea"

type ConfigModel struct{}

func NewConfig() ConfigModel {
	return ConfigModel{}
}

func (m ConfigModel) Update(msg tea.Msg) (ConfigModel, tea.Cmd) {
	return m, nil
}

func (m ConfigModel) View(active bool) string {
	title := "Config Panel"
	if active {
		title += " [ACTIVE]"
	}

	return title + `
Branch: dev ▼

File Path:
[ CICD/ ]

Mode:
(●) Dry-run  ( ) Apply

[ Preview Diff ]`
}
