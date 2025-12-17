package components

import (
	"fmt"

	gitlabclient "GitPatcher/internal/gitlab"
	"github.com/xanzy/go-gitlab"
)

type Tree struct {
	client *gitlab.Client

	Groups   []*gitlab.Group
	Projects map[int][]*gitlab.Project
	Expanded map[int]bool
	Cursor   int
}

func NewTree(client *gitlab.Client) *Tree {
	return &Tree{
		client:   client,
		Projects: make(map[int][]*gitlab.Project),
		Expanded: make(map[int]bool),
	}
}

func (t *Tree) LoadGroups() error {
	groups, err := gitlabclient.ListGroups(t.client)
	if err != nil {
		return err
	}
	t.Groups = groups
	return nil
}

func (t *Tree) ToggleGroup(idx int) error {
	g := t.Groups[idx]
	if t.Expanded[g.ID] {
		t.Expanded[g.ID] = false
		return nil
	}

	if _, ok := t.Projects[g.ID]; !ok {
		ps, err := gitlabclient.ListGroupProjects(t.client, g.ID)
		if err != nil {
			return err
		}
		t.Projects[g.ID] = ps
	}
	t.Expanded[g.ID] = true
	return nil
}

func (t *Tree) View() string {
	out := "Groups / Projects\n\n"
	for i, g := range t.Groups {
		prefix := "▶"
		if t.Expanded[g.ID] {
			prefix = "▼"
		}
		cursor := " "
		if i == t.Cursor {
			cursor = ">"
		}
		out += fmt.Sprintf("%s %s %s\n", cursor, prefix, g.FullPath)

		if t.Expanded[g.ID] {
			for _, p := range t.Projects[g.ID] {
				out += fmt.Sprintf("    • %s\n", p.Path)
			}
		}
	}
	return out
}
