package embedded

import "fmt"

type (
	GHUserCard struct {
		Username string
		HTML     string
	}

	GHRepoCard struct {
		Username string
		Repo     string
		HTML     string
	}
)

const (
	userCardTemplate = `
<div class="github-card" data-github="%s" data-width="400" data-height="318" data-theme="medium"></div>
<script src="//cdn.jsdelivr.net/github-cards/latest/widget.js"></script>`

	repoCardTemplate = `
<div class="github-card" data-github="%s/%s" data-width="400" data-height="150" data-theme="default"></div>
<script src="//cdn.jsdelivr.net/github-cards/latest/widget.js"></script>`
)

func NewGHUserCard(username string) GHUserCard {
	return GHUserCard{
		Username: username,
		HTML:     fmt.Sprintf(userCardTemplate, username),
	}
}

func NewGHRepoCard(username, repo string) GHRepoCard {
	return GHRepoCard{
		Username: username,
		Repo:     repo,
		HTML:     fmt.Sprintf(repoCardTemplate, username, repo),
	}
}
