package tagit

import (
	git "github.com/go-git/go-git/v5"
	"github.com/rotisserie/eris"
)

func GetCurrentCommitID(path string) (string, error) {
	repo, err := git.PlainOpen(path)
	if err != nil {
		return "", eris.Wrap(err, "failed to open git repo")
	}

	ref, err := repo.Head()
	if err != nil {
		return "", eris.Wrap(err, "failed to get head ref")
	}

	return ref.Hash().String(), nil
}
