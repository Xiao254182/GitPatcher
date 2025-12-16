package state

import "github.com/xanzy/go-gitlab"

type AppState struct {
	GitlabURL   string
	GitlabToken string

	Groups   []*gitlab.Group
	Projects []*gitlab.Project

	Selected map[int]bool

	Branch   string
	FilePath string
	DryRun   bool
}
