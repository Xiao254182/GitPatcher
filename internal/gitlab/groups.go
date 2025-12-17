package gitlabclient

import "github.com/xanzy/go-gitlab"

func ListGroups(client *gitlab.Client) ([]*gitlab.Group, error) {
	var all []*gitlab.Group
	opt := &gitlab.ListGroupsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 50,
		},
	}

	for {
		groups, resp, err := client.Groups.ListGroups(opt)
		if err != nil {
			return nil, err
		}
		all = append(all, groups...)
		if resp.CurrentPage >= resp.TotalPages {
			break
		}
		opt.Page++
	}
	return all, nil
}
