package integrations

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Spy struct {
	arrayBytes    []byte
	expectedError error
}

func (s *Spy) GetExternalApi() ([]byte, error) {
	return s.arrayBytes, s.expectedError
}

func TestValidateTransaction(t *testing.T) {
	inte := NewIntegration()
	t.Run("sucess - authorized transaction", func(t *testing.T) {

		os.Setenv("EXTERNAL_URL", "http://run.mocky.io/v3/")
		defer os.Unsetenv("EXTERNAL_URL")

		_, err := inte.ValidateTransaction()

		assert.Nil(t, err)
	})
	t.Run("failed - transaction not authorized", func(t *testing.T) {
		_, err := inte.ValidateTransaction()

		assert.Error(t, err)
	})
}

func ValidateTransaction() {
	panic("unimplemented")
}

func TestGetExternalApi(t *testing.T) {
	inte := NewIntegration()

	t.Run("sucess - get external api", func(t *testing.T) {
		defer os.Unsetenv("EXTERNAL_URL")

		os.Setenv("EXTERNAL_URL", "https://run.mocky.io/v3/")
		_, err := inte.getExternalApi()

		assert.Nil(t, err)
	})
	t.Run("failed - request return nil", func(t *testing.T) {
		defer os.Unsetenv("EXTERNAL_URL")

		os.Setenv("EXTERNAL_URL", "")
		_, err := inte.getExternalApi()

		assert.NotNil(t, err)
	})
}

func TestValidateTransfer(t *testing.T) {
	inte := NewIntegration()
	t.Run("failed - user has enough money", func(t *testing.T) {
		mockedValue := 100
		mockedBalance := 20
		expectedError := "payer doesn't have enough money"

		receivedValue := inte.ValidateTransfer(mockedBalance, mockedValue)

		assert.Equal(t, fmt.Errorf(expectedError), receivedValue)
	})
	t.Run("sucess - user have money to make the transfer", func(t *testing.T) {
		mockedValue := 20
		mockedBalance := 100

		receivedValue := inte.ValidateTransfer(mockedBalance, mockedValue)

		assert.Equal(t, nil, receivedValue)
	})
}

func TestValidateIsCommon(t *testing.T) {
	inte := NewIntegration()
	t.Run("sucess - user(seller) can't transfer", func(t *testing.T) {
		mockedIdSeller := 1
		expectedError := "shopkeeper cannot make transfers"

		receivedValue := inte.ValidateIsCommon(mockedIdSeller)

		assert.Equal(t, fmt.Errorf(expectedError), receivedValue)
	})

	t.Run("sucess - user(common) can transfer", func(t *testing.T) {
		mockedIdCommon := 2

		receivedValue := inte.ValidateIsCommon(mockedIdCommon)

		assert.Equal(t, nil, receivedValue)
	})
}
