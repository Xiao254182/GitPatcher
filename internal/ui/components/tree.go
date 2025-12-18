package components

import (
	"fmt"

	"github.com/xanzy/go-gitlab"
)

type Tree struct {
	Groups   []*gitlab.Group
	Projects map[int][]*gitlab.Project

	Expanded map[int]bool
	Checked  map[int]bool

	Cursor int
	Flat   []TreeItem
}

type TreeItem struct {
	IsGroup bool
	Group   *gitlab.Group
	Project *gitlab.Project
}

func NewTree() *Tree {
	return &Tree{
		Projects: make(map[int][]*gitlab.Project),
		Expanded: make(map[int]bool),
		Checked:  make(map[int]bool),
	}
}

func (t *Tree) Build() {
	t.Flat = nil

	for _, g := range t.Groups {
		t.Flat = append(t.Flat, TreeItem{
			IsGroup: true,
			Group:   g,
		})

		if !t.Expanded[g.ID] {
			continue
		}

		for _, p := range t.Projects[g.ID] {
			t.Flat = append(t.Flat, TreeItem{
				IsGroup: false,
				Project: p,
			})
		}
	}
}

func (t *Tree) Toggle() {
	item := t.Flat[t.Cursor]

	if item.IsGroup {
		t.Expanded[item.Group.ID] = !t.Expanded[item.Group.ID]
	} else {
		id := item.Project.ID
		t.Checked[id] = !t.Checked[id]
	}
	t.Build()
}

func (t *Tree) View(active bool) string {
	out := "Projects\n"
	if active {
		out += "[ACTIVE]\n"
	}
	out += "\n"

	for i, item := range t.Flat {
		cursor := " "
		if i == t.Cursor {
			cursor = "▶"
		}

		if item.IsGroup {
			arrow := "▶"
			if t.Expanded[item.Group.ID] {
				arrow = "▼"
			}
			out += fmt.Sprintf("%s %s %s\n", cursor, arrow, item.Group.FullName)
		} else {
			box := "☐"
			if t.Checked[item.Project.ID] {
				box = "✔"
			}
			out += fmt.Sprintf("  %s %s %s\n", cursor, box, item.Project.Path)
		}
	}
	return out
}
