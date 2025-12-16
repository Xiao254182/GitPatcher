package state

import "github.com/xanzy/go-gitlab"

type AppState struct {
	GitlabURL   string
	GitlabToken string
	Group       string

	Projects []*gitlab.Project
}
type State struct {
	Step int
}
