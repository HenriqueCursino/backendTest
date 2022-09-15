package controller

import (
	"fmt"

	"github.com/henriquecursino/desafioQ2/common"
	"github.com/henriquecursino/desafioQ2/model"
)

func (ctl *Controller) CreateNewUser(user model.User) {
	err := ctl.db.Table("users").Create(&user).Error
	if err != nil {
		fmt.Println("failed to create user!", err.Error())
		return
	}
}

func (ctl *Controller) CreateNewAccount(account model.Account) {
	err := ctl.db.Table("accounts").Create(&account).Error
	if err != nil {
		fmt.Println("failed to add balance!", err.Error())
		return
	}
}

func (ctl *Controller) UpdateAccontBalance(account model.Account, where int) {
	err := ctl.db.Table("accounts").
		Where("cpf_cnpj =?", where).
		Update("balance", account.Balance).
		Error
	if err != nil {
		fmt.Println("failed to update balance", err.Error())
	}
}

func (ctl *Controller) CreateTransaction(transaction model.Transactions) {
	err := ctl.db.Table("transactions").Create(&transaction).Error
	if err != nil {
		fmt.Println("failed to create transaction!", err.Error())
	}
}

func (ctl *Controller) UpdateStatusId(where int) {
	err := ctl.db.Table("transactions").
		Where("id = ?", where).
		Update("id_status", common.STATUS_CONCLUIDO).
		Error
	if err != nil {
		fmt.Println("failed to create transaction!")
	}
}

func (ctl *Controller) RemoveMoney(where, newBalance int) {
	err := ctl.db.
		Table("accounts").
		Where("cpf_cnpj = ?", where).
		Update("balance", newBalance).
		Error
	if err != nil {
		fmt.Println("Failed to update balance!", err.Error())
	}
}

func (ctl *Controller) AddMoney(where, newBalance int) {
	err := ctl.db.
		Table("accounts").
		Where("cpf_cnpj = ?", where).
		Update("balance", newBalance).
		Error
	if err != nil {
		fmt.Println("Failed to update balance!", err.Error())
	}
}
