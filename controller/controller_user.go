package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/henriquecursino/desafioQ2/common"
	"github.com/henriquecursino/desafioQ2/dto"
	"github.com/henriquecursino/desafioQ2/integrations"
	"github.com/henriquecursino/desafioQ2/model"
	"github.com/henriquecursino/desafioQ2/tools"
	"gorm.io/gorm"
)

type Controller struct {
	db *gorm.DB
}

func NewController(db *gorm.DB) Controller {
	return Controller{
		db: db,
	}
}

func (ctl *Controller) CreateUser(c *gin.Context) {
	UserRequest := dto.UserRequest{}
	c.ShouldBindJSON(&UserRequest)

	documentUnmasked := tools.RemoveMask(&UserRequest.CpfCnpj)
	documentInt, _ := tools.ConvertStrToInt(documentUnmasked)

	user := model.User{
		CpfCnpj:    int64(documentInt),
		FullName:   UserRequest.FullName,
		Email:      UserRequest.Email,
		CategoryID: UserRequest.CategoryID,
		Password:   UserRequest.Password,
	}

	ctl.CreateNewUser(user)

	c.JSON(http.StatusOK, "User created sucessfully!")
}

func (ctl *Controller) CreateAccount(c *gin.Context) {

	balanceRequest := dto.AccountRequest{}
	c.ShouldBindJSON(&balanceRequest)

	documentUnmasked := tools.RemoveMask(&balanceRequest.CpfCnpj)
	documentInt, _ := tools.ConvertStrToInt(documentUnmasked)

	balance := model.Account{
		CpfCnpj: int64(documentInt),
		Balance: balanceRequest.Balance,
	}

	ctl.CreateNewAccount(balance)

	c.JSON(http.StatusOK, "deposit made successfully!")
}

func (ctl *Controller) UpdateBalance(c *gin.Context) {
	balanceRequest := dto.AccountRequest{}
	c.ShouldBindJSON(&balanceRequest)

	documentUnmasked := tools.RemoveMask(&balanceRequest.CpfCnpj)
	documentInt, _ := tools.ConvertStrToInt(documentUnmasked)

	balance := model.Account{
		CpfCnpj: int64(documentInt),
		Balance: balanceRequest.Balance,
	}

	ctl.UpdateAccontBalance(balance, documentInt)

	balanceResponse := model.Account{}
	balanceResponse.CpfCnpj = int64(documentInt)

	c.JSON(http.StatusOK, "Balance update successfully!")
}

func (ctl *Controller) Transfer(c *gin.Context) {
	documentPayerInt, _ := tools.ConvertStrToInt(c.Param("doc"))

	transferRequest := dto.TransferRequest{}
	c.ShouldBindJSON(&transferRequest)

	documentPayee := tools.RemoveMask(&transferRequest.CpfPayee)
	documentPayeeInt, _ := tools.ConvertStrToInt(documentPayee)

	accountPayer := model.Account{}
	accountReceiver := model.Account{}
	statusTransaction := model.Status{}
	userPayer := model.User{}

	ctl.db.Table("statuses").Find(&statusTransaction)
	ctl.db.Table("accounts").Where("cpf_cnpj = ?", &documentPayerInt).First(&accountPayer)
	ctl.db.Table("accounts").Where("cpf_cnpj = ?", &documentPayeeInt).First(&accountReceiver)
	ctl.db.Table("users").Where("cpf_cnpj = ?", &documentPayerInt).First(&userPayer)

	balanceError := integrations.ValidateTransfer(accountPayer.Balance, transferRequest.Value)
	if balanceError != nil {
		c.JSON(http.StatusBadRequest, balanceError.Error())
		return
	}

	sellerError := integrations.ValidateIsCommon(userPayer.CategoryID)
	if sellerError != nil {
		c.JSON(http.StatusBadRequest, sellerError.Error())
		return
	}

	transaction := model.Transactions{
		IdPayer:  accountPayer.ID,
		Account:  accountPayer,
		IdPayee:  accountReceiver.ID,
		IdStatus: common.STATUS_PENDENTE,
		Value:    transferRequest.Value,
	}

	data, errValid := integrations.ValidateTransaction()

	ctl.CreateTransaction(transaction)

	if data.Authorization {
		transaction.IdStatus = common.STATUS_CONCLUIDO
		ctl.DebitScheme(accountPayer, accountReceiver, transaction.Value)

		ctl.UpdateStatusId(transaction.ID)

		c.JSON(http.StatusOK, transaction)
	} else {
		fmt.Println("failed to authorize transaction", errValid)
	}
}

func (ctl *Controller) DebitScheme(Payer, Payee model.Account, value int) {
	newBalance := Payer.Balance - value
	ctl.RemoveMoney(int(Payer.CpfCnpj), newBalance)

	addBalance := Payee.Balance + value
	ctl.AddMoney(int(Payee.CpfCnpj), addBalance)
}
