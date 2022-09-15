package controller

import (
	"github.com/henriquecursino/desafioQ2/model"
	"github.com/stretchr/testify/mock"
)

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
