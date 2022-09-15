package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/henriquecursino/desafioQ2/dto"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func mockDb() *gorm.DB {
	dbMock, _, err := sqlmock.New()
	if err != nil {
		println(err.Error())
	}

	dialector := mysql.New(mysql.Config{
		DSN:        "sqlmock_db_0",
		DriverName: "mysql",
		Conn:       dbMock,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		println(err.Error())
	}

	return db
}

func mockDbUser() *gorm.DB {
	dbMock, _, err := sqlmock.New()
	if err != nil {
		println(err.Error())
	}

	dialector := mysql.New(mysql.Config{
		DSN:        "sqlmock_db_0",
		DriverName: "mysql",
		Conn:       dbMock,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		println(err.Error())
	}

	mockUser := dto.UserRequest{
		CpfCnpj:    "630.652.468-13",
		FullName:   "Geraldo Gael Andre Galvao",
		Email:      "geraldogaelgalvao@mesquita.not.br",
		CategoryID: 1,
		Password:   "kpdXcIGIk2",
	}

	db.Table("users").Create(&mockUser)
	return db
}

func mockDbAccount() *gorm.DB {
	dbMock, _, err := sqlmock.New()
	if err != nil {
		println(err.Error())
	}

	dialector := mysql.New(mysql.Config{
		DSN:        "sqlmock_db_0",
		DriverName: "mysql",
		Conn:       dbMock,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		println(err.Error())
	}

	mockUser := dto.UserRequest{
		CpfCnpj:    "630.652.468-13",
		FullName:   "Geraldo Gael Andre Galvao",
		Email:      "geraldogaelgalvao@mesquita.not.br",
		CategoryID: 1,
		Password:   "kpdXcIGIk2",
	}

	db.Table("users").Create(&mockUser)

	mockAccount := dto.AccountRequest{
		CpfCnpj: "630.652.468-13",
		Balance: 200,
	}

	db.Table("accounts").Create(&mockAccount)
	return db
}

func TestCreateUserController(t *testing.T) {
	db := mockDb()
	t.Run("Success - Should return StatusCode 200 (OK)", func(t *testing.T) {
		controller := NewController(db)

		mockUser := dto.UserRequest{
			CpfCnpj:    "630.652.468-13",
			FullName:   "Geraldo Gael Andre Galvao",
			Email:      "geraldogaelgalvao@mesquita.not.br",
			CategoryID: 1,
			Password:   "kpdXcIGIk2",
		}

		mockUserJSON, _ := json.Marshal(mockUser)
		mockUserBuffer := bytes.NewBuffer(mockUserJSON)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		req, _ := http.NewRequest(http.MethodPost, "/", mockUserBuffer)
		ctx.Request = req

		controller.CreateUser(ctx)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
}

func TestCreateAccount(t *testing.T) {
	db := mockDbUser()

	t.Run("Success - Should return StatusCode 200 (OK)", func(t *testing.T) {
		controller := NewController(db)

		mockInfo := dto.AccountRequest{
			CpfCnpj: "630.652.468-13",
			Balance: 200,
		}

		mockAccountJSON, _ := json.Marshal(mockInfo)
		mockAccountBuffer := bytes.NewBuffer(mockAccountJSON)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		req, _ := http.NewRequest(http.MethodPost, "/", mockAccountBuffer)
		ctx.Request = req

		controller.CreateAccount(ctx)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
}

func TestUpdateBalance(t *testing.T) {
	db := mockDbAccount()
	t.Run("Success - Should return StatusCode 200 (OK)", func(t *testing.T) {
		controller := NewController(db)

		mockInfo := dto.AccountRequest{
			CpfCnpj: "630.652.468-13",
			Balance: 2000,
		}

		mockAccountJSON, _ := json.Marshal(mockInfo)
		mockAccountBuffer := bytes.NewBuffer(mockAccountJSON)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		req, _ := http.NewRequest(http.MethodPut, "/", mockAccountBuffer)
		ctx.Request = req

		controller.UpdateBalance(ctx)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
}
