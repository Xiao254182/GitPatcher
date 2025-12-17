package state

import "github.com/xanzy/go-gitlab"

type AppState struct {
	Client *gitlab.Client

	Groups   []*gitlab.Group
	Projects map[int][]*gitlab.Project

	SelectedProjects map[int]bool
	ExpandedGroups   map[int]bool
}
