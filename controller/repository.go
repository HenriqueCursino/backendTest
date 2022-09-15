package controller

import (
	"fmt"

	"github.com/henriquecursino/desafioQ2/common"
	"github.com/henriquecursino/desafioQ2/model"
	"gorm.io/gorm"
)

type Repository interface {
	CreateNewUser(user model.User) error
	CreateNewAccount(account model.Account) error
	// UpdateAccontBalance(account model.Account, where int)
	// CreateTransaction(transaction model.Transactions)
	// UpdateStatusId(where int)
	// RemoveMoney(where, newBalance int)
	// AddMoney(where, newBalance int)
	// GetAccountPayer(document int) model.Account
	// GetAccountReceiver(document int) model.Account
	// GetUserPayer(document int) model.User
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (repo *repository) CreateNewUser(user model.User) error {
	if err := repo.db.Table("users").Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repository) CreateNewAccount(account model.Account) error {
	err := repo.db.Table("accounts").Create(&account).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) UpdateAccontBalance(account model.Account, where int) {
	err := repo.db.Table("accounts").
		Where("cpf_cnpj =?", where).
		Update("balance", account.Balance).
		Error
	if err != nil {
		fmt.Println("failed to update balance", err.Error())
	}
}

func (repo *repository) CreateTransaction(transaction model.Transactions) {
	err := repo.db.Table("transactions").Create(&transaction).Error
	if err != nil {
		fmt.Println("failed to create transaction!", err.Error())
	}
}

func (repo *repository) UpdateStatusId(where int) {
	err := repo.db.Table("transactions").
		Where("id = ?", where).
		Update("id_status", common.STATUS_CONCLUIDO).
		Error
	if err != nil {
		fmt.Println("failed to create transaction!")
	}
}

func (repo *repository) RemoveMoney(where, newBalance int) {
	err := repo.db.
		Table("accounts").
		Where("cpf_cnpj = ?", where).
		Update("balance", newBalance).
		Error
	if err != nil {
		fmt.Println("Failed to update balance!", err.Error())
	}
}

func (repo *repository) AddMoney(where, newBalance int) {
	err := repo.db.
		Table("accounts").
		Where("cpf_cnpj = ?", where).
		Update("balance", newBalance).
		Error
	if err != nil {
		fmt.Println("Failed to update balance!", err.Error())
	}
}

func (repo *repository) GetAccountPayer(document int) model.Account {
	accountPayer := model.Account{}
	repo.db.Table("accounts").Where("cpf_cnpj = ?", &document).First(&accountPayer)
	return accountPayer
}

func (repo *repository) GetAccountReceiver(document int) model.Account {
	accountReceiver := model.Account{}
	repo.db.Table("accounts").Where("cpf_cnpj = ?", &document).First(&accountReceiver)
	return accountReceiver
}

func (repo *repository) GetUserPayer(document int) model.User {
	accountPayer := model.User{}
	repo.db.Table("users").Where("cpf_cnpj = ?", document).First(&accountPayer)
	return accountPayer
}
