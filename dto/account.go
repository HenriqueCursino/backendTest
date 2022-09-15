package dto

import (
	"github.com/henriquecursino/desafioQ2/model"
)

type AccountRequest struct {
	CpfCnpj string `json:"cpf_cnpj"`
	User    model.User
	Balance int `json:"balance"`
}
