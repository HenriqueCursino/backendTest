package model

type Transactions struct {
	ID       int     `gorm:"primaryKey; autoIncrement" json:"id"`
	IdPayer  int     `json:"id_payer"`
	Account  Account `gorm:"foreignKey:IdPayer"`
	IdPayee  int     `json:"id_payee"`
	IdStatus int     `json:"id_status"`
	Value    int     `json:"value"`
}
