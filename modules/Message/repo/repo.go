package repo

// ResultRepository data structure
type ResultRepository struct {
	Result interface{}
	Error  error
}

type MessageRepository interface {
	Add(message string) ResultRepository
	Get() ResultRepository
}
