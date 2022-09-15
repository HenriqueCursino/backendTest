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
	t.Run("sucess - authorized transaction", func(t *testing.T) {
		os.Setenv("EXTERNAL_URL", "http://run.mocky.io/v3/")
		defer os.Unsetenv("EXTERNAL_URL")

		_, err := ValidateTransaction()

		assert.Nil(t, err)
	})
	t.Run("failed - transaction not authorized", func(t *testing.T) {
		_, err := ValidateTransaction()

		assert.Error(t, err)
	})
}

func TestGetExternalApi(t *testing.T) {
	t.Run("sucess - get external api", func(t *testing.T) {
		defer os.Unsetenv("EXTERNAL_URL")

		os.Setenv("EXTERNAL_URL", "https://run.mocky.io/v3/")
		_, err := getExternalApi()

		assert.Nil(t, err)
	})
	t.Run("failed - request return nil", func(t *testing.T) {
		defer os.Unsetenv("EXTERNAL_URL")

		os.Setenv("EXTERNAL_URL", "")
		_, err := getExternalApi()

		assert.NotNil(t, err)
	})
}

func TestValidateTransfer(t *testing.T) {
	t.Run("failed - user has enough money", func(t *testing.T) {
		mockedValue := 100
		mockedBalance := 20
		expectedError := "payer doesn't have enough money"

		receivedValue := ValidateTransfer(mockedBalance, mockedValue)

		assert.Equal(t, fmt.Errorf(expectedError), receivedValue)
	})
	t.Run("sucess - user have money to make the transfer", func(t *testing.T) {
		mockedValue := 20
		mockedBalance := 100

		receivedValue := ValidateTransfer(mockedBalance, mockedValue)

		assert.Equal(t, nil, receivedValue)
	})
}

func TestValidateIsCommon(t *testing.T) {

	t.Run("sucess - user(seller) can't transfer", func(t *testing.T) {
		mockedIdSeller := 1
		expectedError := "shopkeeper cannot make transfers"

		receivedValue := ValidateIsCommon(mockedIdSeller)

		assert.Equal(t, fmt.Errorf(expectedError), receivedValue)
	})

	t.Run("sucess - user(common) can transfer", func(t *testing.T) {
		mockedIdCommon := 2

		receivedValue := ValidateIsCommon(mockedIdCommon)

		assert.Equal(t, nil, receivedValue)
	})
}
