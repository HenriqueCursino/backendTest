package controller

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/henriquecursino/desafioQ2/model"
	"github.com/stretchr/testify/assert"
)

func mockDB() (*sql.DB, sqlmock.Sqlmock, error) {
	return sqlmock.New()
}

func TestCreateNewUserRepository(t *testing.T) {
	db, mock, _ := mockDB()

	t.Run("Sucess - Should return nil", func(t *testing.T) {
		mockUser := model.User{}

		repository := NewRepository(db)

		sqlQuery := regexp.QuoteMeta(`SELECT * FROM  "users"`)

		mock.ExpectQuery(sqlQuery)

		err := repository.CreateNewUser(mockUser)

		assert.Nil(t, err)
	})
}
