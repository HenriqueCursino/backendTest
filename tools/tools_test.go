package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveMask(t *testing.T) {

	t.Run("sucess - mask has been removed", func(t *testing.T) {
		//ASSERT
		mockedValidCpf := "218.226.670-46"
		expectedValue := "21822667046"

		//ACT
		receivedValue := RemoveMask(&mockedValidCpf)

		//ASSERT
		assert.Equal(t, expectedValue, receivedValue)
	})

	t.Run("failed - document in invalid format", func(t *testing.T) {
		mockedErrorValue := "218.l226.670-46"
		expectedValue := "218l22667046"

		//ACT
		receivedValue := RemoveMask(&mockedErrorValue)

		//ASSERT
		assert.Equal(t, expectedValue, receivedValue)
	})
}

func TestConvertStrToInt(t *testing.T) {
	t.Run("sucess - string has been converted to integer", func(t *testing.T) {
		mockedValidStr := "21822667046"
		expectedValue := 21822667046

		receivedValue, _ := ConvertStrToInt(mockedValidStr)

		assert.Equal(t, expectedValue, receivedValue)
	})

	t.Run("failed - invalid string", func(t *testing.T) {
		mockedValidStr := "21822667046a"

		_, err := ConvertStrToInt(mockedValidStr)

		assert.Error(t, err)
	})
}
