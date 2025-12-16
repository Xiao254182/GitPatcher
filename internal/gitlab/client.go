package gitlabclient

import "github.com/xanzy/go-gitlab"

func NewClient(url, token string) (*gitlab.Client, error) {
	return gitlab.NewClient(token, gitlab.WithBaseURL(url))
}
