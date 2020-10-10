package repo

// ResultRepository data structure
type ResultRepository struct {
	Result interface{}
	Error  error
}

type CustomerRepository interface {
	GetCustomerSaldo(accountNumber int) ResultRepository
	TransferSaldo(fromAccountNumber, toAccountNumber, balance int) ResultRepository
}
