package local

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/rotisserie/eris"
)

type Tagger struct {
	repo *git.Repository
}

func (t *Tagger) CreateTag(commitID string, name string, message string) error {
	_, err := t.repo.CreateTag(name, plumbing.NewHash(commitID), &git.CreateTagOptions{
		Message: message,
	})

	return eris.Wrapf(err, "failed to create tag %s", name)
}

func NewTagger(path string) (*Tagger, error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to open git repo")
	}

	return &Tagger{
		repo: repo,
	}, nil
}
