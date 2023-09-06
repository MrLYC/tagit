package tagit

type Tagger interface {
	CreateTag(commitID string, name string, message string) error
}

type Manager struct {
	repo *Repository
}

func (m *Manager) Repository() *Repository {
	return m.repo
}

func (m *Manager) TagHead(tagger Tagger, name string, message string) (string, error) {
	commitID, err := m.repo.GetHeadCommitID()
	if err != nil {
		return "", err
	}

	err = tagger.CreateTag(commitID, name, message)
	if err != nil {
		return "", err
	}

	return commitID, err
}

func NewManager(path string) (*Manager, error) {
	repo, err := NewRepository(path)
	if err != nil {
		return nil, err
	}

	return &Manager{
		repo: repo,
	}, nil
}
