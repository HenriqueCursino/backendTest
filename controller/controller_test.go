package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/henriquecursino/desafioQ2/dto"
	"github.com/henriquecursino/desafioQ2/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewUser(t *testing.T) {
	mockUserBody := dto.UserRequest{
		CpfCnpj:    "630.652.468-13",
		FullName:   "Geraldo Gael Andre Galvao",
		Email:      "geraldogaelgalvao@mesquita.not.br",
		CategoryID: 1,
		Password:   "kpdXcIGIk2",
	}
	mockUserRepository := model.User{
		CpfCnpj:    63065246813,
		FullName:   "Geraldo Gael Andre Galvao",
		Email:      "geraldogaelgalvao@mesquita.not.br",
		CategoryID: 1,
		Password:   "kpdXcIGIk2",
	}

	t.Run("Success - Should return StatusCode 200 (OK)", func(t *testing.T) {
		var repository = new(TestRepositoryMock)
		// mocka o retorno da repository, com as informaçoes já formatadas
		repository.Mock.On("CreateNewUser", mockUserRepository).Return(nil)
		controller := NewController(repository)

		mockUserJSON, _ := json.Marshal(mockUserBody)
		mockUserBuffer := bytes.NewBuffer(mockUserJSON)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		req, _ := http.NewRequest(http.MethodPost, "/", mockUserBuffer)
		ctx.Request = req

		controller.CreateUser(ctx)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
	t.Run("Failed - Should return StatusCode 400 (BadRequest)", func(t *testing.T) {
		var repository = new(TestRepositoryMock)
		// mocka o retorno da repository, com as informaçoes já formatadas
		repository.Mock.On("CreateNewUser", mockUserRepository).Return(errors.New("failed to create user!"))
		controller := NewController(repository)

		mockUserJSON, _ := json.Marshal(mockUserBody)
		mockUserBuffer := bytes.NewBuffer(mockUserJSON)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		req, _ := http.NewRequest(http.MethodPost, "/", mockUserBuffer)
		ctx.Request = req

		controller.CreateUser(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})
}

func TestCreateAccount(t *testing.T) {
	mockAccountBody := dto.AccountRequest{
		CpfCnpj: "630.652.468-13",
		Balance: 200,
	}
	mockAccountRepository := model.Account{
		CpfCnpj: 63065246813,
		Balance: 200,
	}
	t.Run("Success - Should return StatusCode 200 (OK)", func(t *testing.T) {
		repository := new(TestRepositoryMock)
		repository.Mock.On("CreateNewAccount", mockAccountRepository).Return(nil)

		controller := NewController(repository)

		mockAccountJson, _ := json.Marshal(mockAccountBody)
		mockAccountBuffer := bytes.NewBuffer(mockAccountJson)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		req, _ := http.NewRequest(http.MethodPost, "/", mockAccountBuffer)
		ctx.Request = req

		controller.CreateAccount(ctx)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
	t.Run("Failed - Should return StatusCode 400 (Bad Request)", func(t *testing.T) {
		repository := new(TestRepositoryMock)
		repository.Mock.On("CreateNewAccount", mockAccountRepository).Return(errors.New("failed to create account!"))

		controller := NewController(repository)

		mockAccountJson, _ := json.Marshal(mockAccountBody)
		mockAccountBuffer := bytes.NewBuffer(mockAccountJson)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		req, _ := http.NewRequest(http.MethodPost, "/", mockAccountBuffer)
		ctx.Request = req

		controller.CreateAccount(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})
}
