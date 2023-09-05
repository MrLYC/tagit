package tagit

type Tagger interface {
	CreateTag(commitID string, name string, message string) error
}

func TagProject(tagger Tagger, projectPath string, name string, message string) (string, error) {
	repo, err := NewRepository(projectPath)
	if err != nil {
		return "", err
	}

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
