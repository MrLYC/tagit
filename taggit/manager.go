package tagit

type Tagger interface {
	CreateTag(commitID string, name string, message string) error
}

func TagProject(tagger Tagger, repo *Repository, name string, message string) (string, error) {
	commitID, err := repo.GetCurrentCommitID()
	if err != nil {
		return "", err
	}

	err = tagger.CreateTag(commitID, name, message)
	if err != nil {
		return "", err
	}

	return commitID, err
}
