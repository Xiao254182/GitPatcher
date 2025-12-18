package state

import "github.com/xanzy/go-gitlab"

type AppState struct {
	Client *gitlab.Client

	SelectedProjects map[int]bool

	Branch   string
	FilePath string
	DryRun   bool
}
