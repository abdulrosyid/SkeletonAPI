package repo

import (
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	sqlMock "gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestCustomerRepo_GetCustomerSaldo(t *testing.T) {
	t.Run("Test Get Customer Saldo", func(t *testing.T) {
		db, mock, _ := sqlMock.New()

		rows := sqlMock.NewRows([]string{"a_account_number", "c_\"name\"", "a_balance"}).
			AddRow(555001, "Bob Martin", 5000)
		mock.ExpectPrepare(`select a.account_number, c."name", a.balance
			from customer c inner join account a on c.customer_number = a.customer_number where a.account_number = \$1`).ExpectQuery().
			WithArgs(1).WillReturnRows(rows)

		cr := NewCustomerRepo(db)
		result := cr.GetCustomerSaldo(1)
		assert.NoError(t, result.Error)
	})

	t.Run("Test Get Customer Saldo (error prepare)", func(t *testing.T) {
		db, mock, _ := sqlMock.New()
		defer db.Close()

		mock.ExpectPrepare(`select a.account_number, c."name", a.balance
			from customer c inner join account a on c.customer_number = a.customer_number where a.account_number = \$1`).WillReturnError(fmt.Errorf("error"))

		cr := NewCustomerRepo(db)
		result := cr.GetCustomerSaldo(1)
		assert.Error(t, result.Error)
	})

	t.Run("Test Failed Get Customer Saldo (error ErrNoRows)", func(t *testing.T) {
		db, mock, _ := sqlMock.New()
		defer db.Close()

		mock.ExpectPrepare(`select a.account_number, c."name", a.balance
			from customer c inner join account a on c.customer_number = a.customer_number where a.account_number = \$1`).ExpectQuery().
			WithArgs(1).WillReturnError(sql.ErrNoRows)

		cr := NewCustomerRepo(db)
		result := cr.GetCustomerSaldo(1)
		assert.Error(t, result.Error)
	})
}

func TestCustomerRepo_TransferSaldo(t *testing.T) {
	t.Run("Test Success Transfer Saldo", func(t *testing.T) {
		db, mock, _ := sqlMock.New()
		defer db.Close()

		mock.ExpectBegin()
		mock.ExpectPrepare(`UPDATE .+`).ExpectExec().
			WithArgs(1, 1).
			WillReturnResult(sqlMock.NewResult(1, 1))

		mock.ExpectPrepare(`UPDATE .+`).ExpectExec().
			WithArgs(2, 2).
			WillReturnResult(sqlMock.NewResult(1, 1))

		mock.ExpectCommit()

		cr := NewCustomerRepo(db)
		result := cr.TransferSaldo(1, 1, 1)
		assert.Error(t, result.Error)
	})
}
