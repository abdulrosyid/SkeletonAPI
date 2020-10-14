package model

type Customer struct {
	CustomerNumber int64  `json:"customer_number"`
	Name           string `json:"name"`
}

type CustomerAccount struct {
	AccountNumber  int64  `json:"account_number"`
	Name           string `json:"name"`
	CustomerNumber int64  `json:"customer_number"`
	Balance        int64  `json:"balance"`
}

type TransferRequest struct {
	FromAccountNumber int `json:"from_account_number"`
	ToAccountNumber   int `json:"to_account_number"`
	Amount            int `json:"amount"`
}
