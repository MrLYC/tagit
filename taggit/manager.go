package tagit

type Tagger interface {
	CreateTag(commitID string, name string, message string) error
}

func CreateTag(tagger Tagger, commitID string, name string, message string) error {
	return tagger.CreateTag(commitID, name, message)
}
