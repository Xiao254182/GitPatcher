package gitlabclient

import "github.com/xanzy/go-gitlab"

func ListGroupProjects(client *gitlab.Client, groupID int) ([]*gitlab.Project, error) {
	var all []*gitlab.Project
	opt := &gitlab.ListGroupProjectsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 50,
		},
	}

	for {
		ps, resp, err := client.Groups.ListGroupProjects(groupID, opt)
		if err != nil {
			return nil, err
		}
		all = append(all, ps...)
		if resp.CurrentPage >= resp.TotalPages {
			break
		}
		opt.Page++
	}
	return all, nil
}
