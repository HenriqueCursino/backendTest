package controller

import (
	"github.com/henriquecursino/desafioQ2/model"
	"github.com/stretchr/testify/mock"
)

type ControllerMock struct {
	mock.Mock
}

type TestRepositoryMock struct {
	mock.Mock
}

// pegar o primeiro erro da função
func (repository *TestRepositoryMock) CreateNewUser(user model.User) error {
	args := repository.Called(user)
	getted := args.Error(0)
	return getted
}

func (repository *TestRepositoryMock) CreateNewAccount(account model.Account) error {
	args := repository.Called(account)
	getted := args.Error(0)
	return getted
}

func (repository *TestRepositoryMock) UpdateAccountBalance(account model.Account, where int) error {
	args := repository.Called(account)
	getted := args.Error(0)
	return getted
}

func (repository *TestRepositoryMock) GetAccountPayer(where int) (model.Account, error) {
	args := repository.Called(where)
	return args.Get(0).(model.Account), args.Error(1)
}

func (repository *TestRepositoryMock) UpdateStatusId(where int) error {
	args := repository.Called(where)
	getted := args.Error(0)
	return getted
}

func (repository *TestRepositoryMock) RemoveMoney(where int64, newBalance int) error {
	args := repository.Called(where, newBalance)
	getted := args.Error(0)
	return getted
}

func (repository *TestRepositoryMock) AddMoney(where int64, newBalance int) error {
	args := repository.Called(where, newBalance)
	getted := args.Error(0)
	return getted
}

func (repository *TestRepositoryMock) GetAccountReceiver(document int) (model.Account, error) {
	args := repository.Called(document)
	return args.Get(0).(model.Account), args.Error(1)
}

func (repository *TestRepositoryMock) GetUserPayer(document int) (model.User, error) {
	args := repository.Called(document)
	return args.Get(0).(model.User), args.Error(1)
}

func (repository *TestRepositoryMock) CreateTransaction(transaction model.Transactions) error {
	args := repository.Called(transaction)
	getted := args.Error(0)
	return getted
}
