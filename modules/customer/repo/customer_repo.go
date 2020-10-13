package repo

import (
	"SkeletonAPI/modules/customer/model"
	"database/sql"
	"fmt"
)

type CustomerRepo struct {
	DB *sql.DB
}

func NewCustomerRepo(db *sql.DB) *CustomerRepo {
	return &CustomerRepo{DB: db}
}

func (c *CustomerRepo) GetCustomerSaldo(accountNumber int) ResultRepository {
	output := ResultRepository{}
	customerAccount := model.CustomerAccount{}

	sq := `select a.account_number, c."name", a.balance 
			from customer c inner join account a on c.customer_number = a.customer_number where a.account_number = $1`

	stmt, err := c.DB.Prepare(sq)
	if err != nil {
		output.Error = err
		return output
	}

	err = stmt.QueryRow(accountNumber).Scan(&customerAccount.AccountNumber, &customerAccount.Name,
		&customerAccount.Balance)
	if err != nil {
		output.Error = err
		return output
	}

	output.Result = customerAccount

	return ResultRepository{Result: output.Result}
}

func (c *CustomerRepo) TransferSaldo(fromAccountNumber, toAccountNumber, balance int) ResultRepository {
	output := ResultRepository{}
	output.Result = model.CustomerAccount{}

	sq := `update account set balance = %s $1 where account_number = $2`

	tx, err := c.DB.Begin()
	if err != nil {
		output.Error = err
		return output
	}

	sqSubs := fmt.Sprintf(sq, "balance -")
	stmt, err := tx.Prepare(sqSubs)
	if err != nil {
		tx.Rollback()
		output.Error = err
		return output
	}

	_, err = stmt.Exec(balance, fromAccountNumber)
	if err != nil {
		tx.Rollback()
		output.Error = err
		return output
	}

	sqAdd := fmt.Sprintf(sq, "balance +")
	stmt, err = tx.Prepare(sqAdd)
	if err != nil {
		tx.Rollback()
		output.Error = err
		return output
	}

	_, err = stmt.Exec(balance, toAccountNumber)
	if err != nil {
		tx.Rollback()
		output.Error = err
		return output
	}

	defer stmt.Close()

	tx.Commit()

	return ResultRepository{Result: output}
}
