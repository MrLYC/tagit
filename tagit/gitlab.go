package tagit

import (
	"github.com/agnivade/levenshtein"
	"github.com/rotisserie/eris"
	"github.com/xanzy/go-gitlab"
)

type Gitlab struct {
	client *gitlab.Client
}

func (g *Gitlab) GetProjectByName(name string, scoringFn func(project *gitlab.Project) int) (*gitlab.Project, error) {
	projects, _, err := g.client.Projects.ListProjects(&gitlab.ListProjectsOptions{
		Search: &name,
	})

	if err != nil {
		return nil, eris.Wrapf(err, "failed to list projects")
	}

	if len(projects) == 0 {
		return nil, eris.New("project not found")
	}

	chosenProject := projects[0]
	chosenScore := scoringFn(chosenProject)

	for _, project := range projects[1:] {
		score := scoringFn(project)
		if score < chosenScore {
			chosenProject = project
			chosenScore = score
		}
	}

	return chosenProject, nil
}

func NewGitlab(token string, url string) (*Gitlab, error) {
	client, err := gitlab.NewClient(token, gitlab.WithBaseURL(url))
	if err != nil {
		return nil, eris.Wrapf(err, "failed to create gitlab client")
	}

	return &Gitlab{
		client: client,
	}, nil
}

func ProjectUrlDistance(url string) func(*gitlab.Project) int {
	return func(p *gitlab.Project) int {
		sshDistance := levenshtein.ComputeDistance(p.SSHURLToRepo, url)
		httpDistance := levenshtein.ComputeDistance(p.HTTPURLToRepo, url)

		if sshDistance < httpDistance {
			return sshDistance
		}

		return httpDistance
	}
}

func ProjectNameDistance(name string) func(*gitlab.Project) int {
	return func(p *gitlab.Project) int {
		return levenshtein.ComputeDistance(p.Name, name)
	}
}

type GitlabTagger struct {
	pid  int
	glab *Gitlab
}

func (g *GitlabTagger) CreateTag(commitID string, name string, message string) error {
	client := g.glab.client

	_, _, err := client.Tags.CreateTag(
		g.pid, &gitlab.CreateTagOptions{
			TagName: &name,
			Message: &message,
			Ref:     &commitID,
		}, nil,
	)

	return eris.Wrapf(err, "failed to create tag %s", name)
}

func NewGitlabTagger(pid int, g *Gitlab) *GitlabTagger {
	return &GitlabTagger{
		pid:  pid,
		glab: g,
	}
}
