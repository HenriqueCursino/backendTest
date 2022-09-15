package model

type Account struct {
	ID      int   `gorm:"primaryKey; autoIncrement" json:"id"`
	CpfCnpj int64 `json:"cpf_cnpj"`
	User    User  `gorm:"foreignKey:CpfCnpj" json:"-"`
	Balance int   `json:"balance"`
}
