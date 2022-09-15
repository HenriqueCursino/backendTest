package dto

import (
	"github.com/henriquecursino/desafioQ2/model"
)

type UserRequest struct {
	CpfCnpj    string `json:"cpf_cnpj"`
	FullName   string `json:"full_name"`
	Email      string `json:"email"`
	CategoryID int    `json:"id_category"`
	Categories model.Categories
	Password   string `json:"password"`
}
