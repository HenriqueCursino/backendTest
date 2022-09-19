package controller

import (
	"github.com/henriquecursino/desafioQ2/common"
	"github.com/henriquecursino/desafioQ2/model"
	"gorm.io/gorm"
)

type Repository interface {
	CreateNewUser(user model.User) error
	CreateNewAccount(account model.Account) error
	UpdateAccountBalance(account model.Account, where int) error
	CreateTransaction(transaction model.Transactions) error

	UpdateStatusId(where int) error
	RemoveMoney(where int64, newBalance int) error
	AddMoney(where int64, newBalance int) error
	GetAccountPayer(document int) (model.Account, error)
	GetAccountReceiver(document int) (model.Account, error)
	GetUserPayer(document int) (model.User, error)
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

func (repo *repository) UpdateAccountBalance(account model.Account, where int) error {
	err := repo.db.Table("accounts").
		Where("cpf_cnpj =?", where).
		Update("balance", account.Balance).
		Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) CreateTransaction(transaction model.Transactions) error {
	err := repo.db.Table("transactions").Create(&transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) UpdateStatusId(where int) error {
	err := repo.db.Table("transactions").
		Where("id = ?", where).
		Update("id_status", common.STATUS_CONCLUIDO).
		Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) RemoveMoney(where int64, newBalance int) error {
	err := repo.db.
		Table("accounts").
		Where("cpf_cnpj = ?", where).
		Update("balance", newBalance).
		Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) AddMoney(where int64, newBalance int) error {
	err := repo.db.
		Table("accounts").
		Where("cpf_cnpj = ?", where).
		Update("balance", newBalance).
		Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *repository) GetAccountPayer(document int) (model.Account, error) {
	accountPayer := model.Account{}
	err := repo.db.Table("accounts").Where("cpf_cnpj = ?", &document).First(&accountPayer).Error
	return accountPayer, err
}

func (repo *repository) GetAccountReceiver(document int) (model.Account, error) {
	accountReceiver := model.Account{}
	err := repo.db.Table("accounts").Where("cpf_cnpj = ?", &document).First(&accountReceiver).Error
	return accountReceiver, err
}

func (repo *repository) GetUserPayer(document int) (model.User, error) {
	accountPayer := model.User{}
	err := repo.db.Table("users").Where("cpf_cnpj = ?", document).First(&accountPayer).Error
	return accountPayer, err
}
