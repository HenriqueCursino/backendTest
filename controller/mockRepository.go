package controller

import (
	"github.com/henriquecursino/desafioQ2/model"
	"github.com/stretchr/testify/mock"
)

type TestRepositoryMock struct {
	mock.Mock
}

func (repository *TestRepositoryMock) CreateNewUser(user model.User) error {
	args := repository.Called(user)
	getted := args.Error(0)
	return getted
}
