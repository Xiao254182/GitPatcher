package components

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Config struct {
	Branch   textinput.Model
	FilePath textinput.Model
	DryRun   bool
}

func NewConfig() *Config {
	b := textinput.New()
	b.Placeholder = "branch (e.g. main)"

	f := textinput.New()
	f.Placeholder = "file path (e.g. CICD/.gitlab-ci.yml)"

	return &Config{
		Branch:   b,
		FilePath: f,
		DryRun:   true,
	}
}

func (c *Config) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "tab" {
			if c.Branch.Focused() {
				c.Branch.Blur()
				c.FilePath.Focus()
			} else {
				c.FilePath.Blur()
				c.Branch.Focus()
			}
		}
		if msg.String() == " " {
			c.DryRun = !c.DryRun
		}
	}

	var cmd tea.Cmd
	c.Branch, cmd = c.Branch.Update(msg)
	c.FilePath, _ = c.FilePath.Update(msg)
	return cmd
}

func (c *Config) View(active bool) string {
	title := "Config"
	if active {
		title += " [ACTIVE]"
	}

	mode := "Apply"
	if c.DryRun {
		mode = "Dry-run"
	}

	return fmt.Sprintf(
		"%s\n\nBranch:\n%s\n\nFile Path:\n%s\n\nMode: %s\n\n[Space] Toggle Mode\n[Tab] Switch Input\n",
		title,
		c.Branch.View(),
		c.FilePath.View(),
		mode,
	)
}
