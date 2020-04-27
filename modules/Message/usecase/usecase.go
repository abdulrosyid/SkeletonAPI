package usecase

// ResultUseCase data structure
type ResultUseCase struct {
	Result interface{}
	Error  error
}

type MessageUseCase interface {
	AddMessage(message string) ResultUseCase
	GetMessage() ResultUseCase
}
