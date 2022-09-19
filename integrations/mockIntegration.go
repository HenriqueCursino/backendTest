package integrations

import (
	"github.com/henriquecursino/desafioQ2/dto"
	"github.com/stretchr/testify/mock"
)

type TestIntegrationMock struct {
	mock.Mock
}

func (integrations *TestIntegrationMock) ValidateTransaction() (*dto.Authorization, error) {
	args := integrations.Called()
	return args.Get(0).(*dto.Authorization), args.Error(1)
}

func (integrations *TestIntegrationMock) ValidateTransfer(payerBalance, value int) error {
	args := integrations.Called(payerBalance, value)
	getted := args.Error(0)
	return getted
}

func (integrations *TestIntegrationMock) ValidateIsCommon(payerId int) error {
	args := integrations.Called(payerId)
	return args.Error(0)
}

func (integrations *TestIntegrationMock) getExternalApi() ([]byte, error) {
	args := integrations.Called()
	return args.Get(0).([]byte), args.Error(0)
}
