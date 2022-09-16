package integrations

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/henriquecursino/desafioQ2/common"
	"github.com/henriquecursino/desafioQ2/dto"
)

type (
	Integration struct{}

	IIntegrations interface {
		ValidateTransaction() (*dto.Authorization, error)
		ValidateTransfer(payerBalance, value int) error
		ValidateIsCommon(payerId int) error
	}
)

func (i *Integration) ValidateTransaction() (*dto.Authorization, error) {
	responseData, err := i.getExternalApi()
	if err != nil {
		return nil, err
	}

	data := dto.Authorization{}

	err = json.Unmarshal(responseData, &data)
	return &data, err
}

func (i *Integration) getExternalApi() ([]byte, error) {
	url := os.Getenv("EXTERNAL_URL")
	response, err := http.Get(url + common.END_POINT)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(response.Body)
}

func (i *Integration) ValidateTransfer(payerBalance, value int) error {
	if payerBalance < value {
		return fmt.Errorf("payer doesn't have enough money")
	}
	return nil
}

func (i *Integration) ValidateIsCommon(payerId int) error {
	if payerId == common.LOJISTA {
		return fmt.Errorf("shopkeeper cannot make transfers")
	}
	return nil
}
