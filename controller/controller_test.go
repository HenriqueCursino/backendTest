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
	"github.com/henriquecursino/desafioQ2/integrations"
	"github.com/henriquecursino/desafioQ2/model"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewUser(t *testing.T) {
	integrations := new(integrations.TestIntegrationMock)

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
		controller := NewController(repository, integrations)

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
		controller := NewController(repository, integrations)

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
		integrations := new(integrations.TestIntegrationMock)

		repository.Mock.On("CreateNewAccount", mockAccountRepository).Return(nil)

		controller := NewController(repository, integrations)

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
		integrations := new(integrations.TestIntegrationMock)

		repository.Mock.On("CreateNewAccount", mockAccountRepository).Return(assert.AnError)

		controller := NewController(repository, integrations)

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

func TestUpdateBalance(t *testing.T) {
	integrations := new(integrations.TestIntegrationMock)

	mockUpdateBalanceBody := dto.AccountRequest{
		CpfCnpj: "630.652.468-13",
		Balance: 200,
	}
	mockUpdateBalanceRepository := model.Account{
		CpfCnpj: 63065246813,
		Balance: 200,
	}
	t.Run("Success - Should return StatusCode 200 (OK)", func(t *testing.T) {
		repository := new(TestRepositoryMock)

		repository.Mock.On("UpdateAccountBalance", mockUpdateBalanceRepository).Return(nil)

		controller := NewController(repository, integrations)

		mockAccountJson, _ := json.Marshal(mockUpdateBalanceBody)
		mockAccountBuffer := bytes.NewBuffer(mockAccountJson)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		req, _ := http.NewRequest(http.MethodPut, "/", mockAccountBuffer)
		ctx.Request = req

		controller.UpdateBalance(ctx)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
	t.Run("Failed - Should return StatusCode 400 (BadRequest)", func(t *testing.T) {
		repository := new(TestRepositoryMock)
		repository.Mock.On("UpdateAccountBalance", mockUpdateBalanceRepository).Return(assert.AnError)

		controller := NewController(repository, integrations)

		mockAccountJson, _ := json.Marshal(mockUpdateBalanceBody)
		mockAccountBuffer := bytes.NewBuffer(mockAccountJson)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		req, _ := http.NewRequest(http.MethodPut, "/", mockAccountBuffer)
		ctx.Request = req

		controller.UpdateBalance(ctx)

		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})
}

func TestGetAccountPayer(t *testing.T) {
	mockIdPayerParams := "12345678910"
	mockIdPayer := 12345678910
	transactionAccount := 11122233344

	mockCategoryPayer := model.Categories{
		ID:   2,
		Name: "commum",
	}
	mockCategoryReciver := model.Categories{
		ID:   1,
		Name: "lojista",
	}

	mockPayer := model.User{
		CpfCnpj:    12345678910,
		FullName:   "Henrique Cursino",
		Email:      "henrique@gmail.com",
		CategoryID: mockCategoryPayer.ID,
		Password:   "123",
	}

	mockReciver := model.User{
		CpfCnpj:    11122233344,
		FullName:   "Guilherme Sembeneli",
		Email:      "guilherme@gmail.com",
		CategoryID: mockCategoryReciver.ID,
		Password:   "1234",
	}

	mockAccountPayer := model.Account{
		CpfCnpj: mockPayer.CpfCnpj,
		Balance: 100,
	}

	mockAccountReciver := model.Account{
		CpfCnpj: mockReciver.CpfCnpj,
		Balance: 100,
	}

	mockStatusTransactionPending := model.Status{
		ID:   1,
		Name: "Pendente",
	}

	transactionBody := dto.TransferRequest{
		CpfPayee: "111.222.333-44",
		Value:    10,
	}

	transactionRepository := model.Transactions{
		IdPayer:  mockAccountPayer.ID,
		Account:  mockAccountPayer,
		IdPayee:  mockAccountReciver.ID,
		IdStatus: mockStatusTransactionPending.ID,
		Value:    10,
	}

	auth := dto.Authorization{
		Authorization: true,
	}

	mockNewBalancePayer := 90
	mockNewBalanceReceiver := 110

	t.Run("Success - Should return StatusCode 200 (OK)", func(t *testing.T) {
		repository := new(TestRepositoryMock)
		integrations := new(integrations.TestIntegrationMock)
		repository.Mock.On("GetAccountPayer", mockIdPayer).Return(mockAccountPayer, nil)
		repository.Mock.On("GetAccountReceiver", transactionAccount).Return(mockAccountReciver, nil)
		repository.Mock.On("GetUserPayer", mockIdPayer).Return(mockPayer, nil)
		integrations.Mock.On("ValidateTransfer", mockAccountPayer.Balance, transactionBody.Value).Return(nil)
		integrations.Mock.On("ValidateIsCommon", mockPayer.CategoryID).Return(nil)
		integrations.Mock.On("ValidateTransaction").Return(&auth, nil)
		repository.Mock.On("CreateTransaction", transactionRepository).Return(nil)
		repository.Mock.On("RemoveMoney", mockPayer.CpfCnpj, mockNewBalancePayer).Return(nil)
		repository.Mock.On("AddMoney", mockReciver.CpfCnpj, mockNewBalanceReceiver).Return(nil)
		repository.Mock.On("UpdateStatusId", transactionRepository.ID).Return(nil)

		controller := NewController(repository, integrations)

		mockTransferJson, _ := json.Marshal(transactionBody)
		mockTransferBuffer := bytes.NewBuffer(mockTransferJson)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		req, _ := http.NewRequest(http.MethodPost, "/", mockTransferBuffer)
		ctx.Params = append(ctx.Params, gin.Param{Key: "doc", Value: mockIdPayerParams})
		ctx.Request = req

		controller.Transfer(ctx)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
	})
}
