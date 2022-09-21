package controller

import (
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/henriquecursino/desafioQ2/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func mockDB() (*gorm.DB, sqlmock.Sqlmock) {
	dbSql, mockSql, err := sqlmock.New()
	if err != nil {
		log.Fatal("Failed to open test connection")
	}

	dialector := mysql.New(mysql.Config{
		DriverName:                "mysql",
		DSN:                       "mysql_mock_0",
		Conn:                      dbSql,
		SkipInitializeWithVersion: true,
	})

	dbMock, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open test connection")
	}

	return dbMock, mockSql
}

func TestCreateNewUserRepository(t *testing.T) {
	db, sql := mockDB()
	mockUser := model.User{
		ID:         1,
		CpfCnpj:    11111111111,
		FullName:   gofakeit.Name(),
		Email:      gofakeit.Email(),
		CategoryID: 1,
		Password:   gofakeit.Password(true, true, true, false, false, 8),
	}
	t.Run("Sucess - Should return nil", func(t *testing.T) {

		repository := NewRepository(db)
		mockQuery := regexp.QuoteMeta("INSERT INTO `users` (`cpf_cnpj`,`full_name`,`email`,`category_id`,`password`,`id`) VALUES (?,?,?,?,?,?)")

		sql.ExpectBegin()
		sql.ExpectExec(mockQuery).
			WithArgs(mockUser.CpfCnpj, mockUser.FullName, mockUser.Email, mockUser.CategoryID, mockUser.Password, mockUser.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		sql.ExpectCommit()

		err := repository.CreateNewUser(mockUser)

		assert.Nil(t, err)
	})
	t.Run("Error - Should return error", func(t *testing.T) {
		repository := NewRepository(db)
		mockQuery := regexp.QuoteMeta("INSERT INTO `users` (`cpf_cnpj`,`full_name`,`email`,`category_id`,`password`,`id`) VALUES (?,?,?,?,?,?)")

		sql.ExpectBegin()
		sql.ExpectExec(mockQuery).
			WithArgs()
		sql.ExpectCommit()

		err := repository.CreateNewUser(mockUser)

		assert.Error(t, err)
	})
}

func TestCreateNewAccountRepository(t *testing.T) {
	db, sql := mockDB()
	mockAccount := model.Account{
		CpfCnpj: 12345678910,
		User:    model.User{},
		Balance: 100,
	}
	t.Run("Sucess - Should restun nil", func(t *testing.T) {
		repository := NewRepository(db)
		mockQuery := regexp.QuoteMeta("INSERT INTO `accounts` (`cpf_cnpj`,`balance`) VALUES (?,?)")

		sql.ExpectBegin()
		sql.ExpectExec(mockQuery).
			WithArgs(mockAccount.CpfCnpj, mockAccount.Balance).
			WillReturnResult(sqlmock.NewResult(1, 1))
		sql.ExpectCommit()

		err := repository.CreateNewAccount(mockAccount)

		assert.Nil(t, err)
	})
	t.Run("Failed - Should return error", func(t *testing.T) {
		repository := NewRepository(db)
		mockQuery := regexp.QuoteMeta("INSERT INTO `accounts` (`cpf_cnpj`,`balance`) VALUES (?,?)")

		sql.ExpectBegin()
		sql.ExpectExec(mockQuery).WithArgs()
		sql.ExpectCommit()

		err := repository.CreateNewAccount(mockAccount)

		assert.Error(t, err)
	})
}

func TestUpdateAccountBalance(t *testing.T) {
	db, sql := mockDB()
	mockUpdateBalance := model.Account{
		CpfCnpj: 12345678910,
		User:    model.User{},
		Balance: 100,
	}
	t.Run("Sucess - Should return nil", func(t *testing.T) {
		repository := NewRepository(db)
		mockQuery := regexp.QuoteMeta("UPDATE `accounts` SET `balance`=? WHERE cpf_cnpj =?")

		sql.ExpectBegin()
		sql.ExpectExec(mockQuery).WithArgs(mockUpdateBalance.Balance, mockUpdateBalance.CpfCnpj).
			WillReturnResult(sqlmock.NewResult(1, 1))
		sql.ExpectCommit()

		err := repository.UpdateAccountBalance(mockUpdateBalance, int(mockUpdateBalance.CpfCnpj))
		assert.Nil(t, err)
	})
	t.Run("Failed - Should return error", func(t *testing.T) {
		repository := NewRepository(db)
		mockQuery := regexp.QuoteMeta("UPDATE `accounts` SET `balance`=? WHERE cpf_cnpj =?")

		sql.ExpectBegin()
		sql.ExpectExec(mockQuery).WithArgs()
		sql.ExpectCommit()

		err := repository.UpdateAccountBalance(mockUpdateBalance, int(mockUpdateBalance.CpfCnpj))
		assert.Error(t, err)
	})
}

func TestCreateTransaction(t *testing.T) {
	db, sql := mockDB()
	mockTransaction := model.Transactions{
		IdPayer:  1,
		Account:  model.Account{},
		IdPayee:  2,
		IdStatus: 1,
		Value:    10,
	}
	t.Run("Sucess - Should return nil", func(t *testing.T) {
		repository := NewRepository(db)
		mockQuery := regexp.QuoteMeta("INSERT INTO `transactions` (`id_payer`,`id_payee`,`id_status`,`value`) VALUES (?,?,?,?)")

		sql.ExpectBegin()
		sql.ExpectExec(mockQuery).WithArgs(mockTransaction.IdPayer, mockTransaction.IdPayee, mockTransaction.IdStatus, mockTransaction.Value).
			WillReturnResult(sqlmock.NewResult(1, 1))
		sql.ExpectCommit()

		err := repository.CreateTransaction(mockTransaction)
		assert.Nil(t, err)
	})
	t.Run("Failed - Should return error", func(t *testing.T) {
		repository := NewRepository(db)
		mockQuery := regexp.QuoteMeta("INSERT INTO `transactions` (`id_payer`,`id_payee`,`id_status`,`value`) VALUES (?,?,?,?)")

		sql.ExpectBegin()
		sql.ExpectExec(mockQuery).WithArgs()
		sql.ExpectCommit()

		err := repository.CreateTransaction(mockTransaction)
		assert.Error(t, err)
	})
}

func TestUpdateStatusId(t *testing.T) {
	db, sql := mockDB()
	mockWhereId := 1
	t.Run("Sucess - Should return nil", func(t *testing.T) {
		repository := NewRepository(db)
		mockQuery := regexp.QuoteMeta("UPDATE `transactions` SET `id_status`=? WHERE id = ?")

		sql.ExpectBegin()
		sql.ExpectExec(mockQuery).WithArgs(2, mockWhereId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		sql.ExpectCommit()

		err := repository.UpdateStatusId(mockWhereId)
		assert.Nil(t, err)
	})
	t.Run("Failed - Should return error", func(t *testing.T) {
		repository := NewRepository(db)
		mockQuery := regexp.QuoteMeta("UPDATE `transactions` SET `id_status`=? WHERE id = ?")

		sql.ExpectBegin()
		sql.ExpectExec(mockQuery).WithArgs()
		sql.ExpectCommit()

		err := repository.UpdateStatusId(mockWhereId)
		assert.Error(t, err)
	})
}

func TestRemoveMoney(t *testing.T) {
	db, sql := mockDB()
	mockCpf := 12345678910
	mockNewBalance := 10

	t.Run("Sucess - Should return nil", func(t *testing.T) {
		repository := NewRepository(db)
		mockQuery := regexp.QuoteMeta("UPDATE `accounts` SET `balance`=? WHERE cpf_cnpj = ?")

		sql.ExpectBegin()
		sql.ExpectExec(mockQuery).WithArgs(mockNewBalance, int64(mockCpf)).
			WillReturnResult(sqlmock.NewResult(1, 1))
		sql.ExpectCommit()

		err := repository.RemoveMoney(int64(mockCpf), mockNewBalance)
		assert.Nil(t, err)
	})
	t.Run("Failed - Should return error", func(t *testing.T) {
		repository := NewRepository(db)
		mockQuery := regexp.QuoteMeta("UPDATE `accounts` SET `balance`=? WHERE cpf_cnpj = ?")

		sql.ExpectBegin()
		sql.ExpectExec(mockQuery).WithArgs()
		sql.ExpectCommit()

		err := repository.RemoveMoney(int64(mockCpf), mockNewBalance)
		assert.Error(t, err)
	})
}

func TestAddMoney(t *testing.T) {
	db, sql := mockDB()
	mockCpf := 12345678910
	mockNewBalance := 10

	t.Run("Sucess - Should return nil", func(t *testing.T) {
		repository := NewRepository(db)
		mockQuery := regexp.QuoteMeta("UPDATE `accounts` SET `balance`=? WHERE cpf_cnpj = ?")

		sql.ExpectBegin()
		sql.ExpectExec(mockQuery).WithArgs(mockNewBalance, int64(mockCpf)).
			WillReturnResult(sqlmock.NewResult(1, 1))
		sql.ExpectCommit()

		err := repository.AddMoney(int64(mockCpf), mockNewBalance)
		assert.Nil(t, err)
	})
	t.Run("Failed - Should return error", func(t *testing.T) {
		repository := NewRepository(db)
		mockQuery := regexp.QuoteMeta("UPDATE `accounts` SET `balance`=? WHERE cpf_cnpj = ?")

		sql.ExpectBegin()
		sql.ExpectExec(mockQuery).WithArgs()
		sql.ExpectCommit()

		err := repository.AddMoney(int64(mockCpf), mockNewBalance)
		assert.Error(t, err)
	})
}

func TestGetAccountPayerRepository(t *testing.T) {
	db, sql := mockDB()
	mockExpectedAccount := model.Account{
		CpfCnpj: 11122233344,
		Balance: 100,
	}
	t.Run("Success - Should return account and nil", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"CpfCnpj",
			"Balance",
		}).AddRow(
			mockExpectedAccount.CpfCnpj,
			mockExpectedAccount.Balance,
		)

		mockQuery := regexp.QuoteMeta("SELECT * FROM `accounts` WHERE cpf_cnpj = ? ORDER BY `accounts`.`id` LIMIT 1")

		sql.ExpectQuery(mockQuery).WithArgs(&mockExpectedAccount.CpfCnpj).WillReturnRows(rows)

		repository := NewRepository(db)
		_, err := repository.GetAccountPayer(int(mockExpectedAccount.CpfCnpj))

		assert.Nil(t, err)
	})
}

func TestGetAccountReceiverRepository(t *testing.T) {
	db, sql := mockDB()
	mockExpectedAccount := model.Account{
		CpfCnpj: 11122233344,
		Balance: 100,
	}
	t.Run("Success - Should return nil", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"CpfCnpj",
			"Balance",
		}).AddRow(
			mockExpectedAccount.CpfCnpj,
			mockExpectedAccount.Balance,
		)

		mockQuery := regexp.QuoteMeta("SELECT * FROM `accounts` WHERE cpf_cnpj = ? ORDER BY `accounts`.`id` LIMIT 1")

		sql.ExpectQuery(mockQuery).WithArgs(&mockExpectedAccount.CpfCnpj).WillReturnRows(rows)

		repository := NewRepository(db)
		account, _ := repository.GetAccountReceiver(int(mockExpectedAccount.CpfCnpj))

		assert.Equal(t, mockExpectedAccount, account)
	})
}

func TestGetUserPayerRepository(t *testing.T) {
	db, sql := mockDB()
	mockCpf := 12345678910
	mockCategory := model.Categories{
		ID:   0,
		Name: "",
	}

	mockExpectedUser := model.User{
		ID:         1,
		CpfCnpj:    12345678910,
		FullName:   "Henrique Cursino",
		Email:      "henrique@gmail.com",
		CategoryID: 1,
		Categories: mockCategory,
		Password:   "1234",
	}

	t.Run("Success - Should return nil", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{
			"ID",
			"CpfCnpj",
			"FullName",
			"Email",
			"CategoryID",
			"Categories.ID",
			"Categories.Name",
			"Password",
		}).AddRow(
			mockExpectedUser.ID,
			mockExpectedUser.CpfCnpj,
			mockExpectedUser.FullName,
			mockExpectedUser.Email,
			mockExpectedUser.CategoryID,
			mockCategory.ID,
			mockCategory.Name,
			mockExpectedUser.Password,
		)

		mockQuery := regexp.QuoteMeta("SELECT * FROM `users` WHERE cpf_cnpj = ? ORDER BY `users`.`id` LIMIT 1")

		sql.ExpectQuery(mockQuery).WithArgs(&mockCpf).WillReturnRows(rows)

		repository := NewRepository(db)
		user, _ := repository.GetUserPayer(mockCpf)

		assert.Equal(t, mockExpectedUser, user)
	})
}
