package gitlabclient

import "github.com/xanzy/go-gitlab"

func New(url, token string) (*gitlab.Client, error) {
	return gitlab.NewClient(token, gitlab.WithBaseURL(url))
}
