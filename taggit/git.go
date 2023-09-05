package tagit

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/rotisserie/eris"
)

type Repository struct {
	repo *git.Repository
}

func (r *Repository) CreateTag(commitID string, name string, message string) error {
	_, err := r.repo.CreateTag(name, plumbing.NewHash(commitID), &git.CreateTagOptions{
		Message: message,
	})

	return eris.Wrapf(err, "failed to create tag %s", name)
}

func (r *Repository) GetCurrentCommitID() (string, error) {
	ref, err := r.repo.Head()
	if err != nil {
		return "", eris.Wrap(err, "failed to get head ref")
	}

	return ref.Hash().String(), nil
}

func (r *Repository) GetRemoteUrl(name string) (string, error) {
	remote, err := r.repo.Remote(name)
	if err != nil {
		return "", eris.Wrap(err, "failed to get remote")
	}

	config := remote.Config()
	return config.URLs[0], nil
}

func NewRepository(path string) (*Repository, error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return nil, eris.Wrapf(err, "failed to open git repository")
	}

	return &Repository{
		repo: repo,
	}, nil
}
