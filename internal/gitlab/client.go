package gitlabclient

import (
	"strings"

	"github.com/xanzy/go-gitlab"
)

func normalizeURL(u string) string {
	u = strings.TrimSpace(u)
	u = strings.TrimRight(u, "/")
	if !strings.HasSuffix(u, "/api/v4") {
		u += "/api/v4"
	}
	return u
}

func NewClient(url, token string) (*gitlab.Client, error) {
	return gitlab.NewClient(token, gitlab.WithBaseURL(normalizeURL(url)))
}
