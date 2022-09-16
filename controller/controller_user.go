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
)

type Controller interface {
	CreateUser(c *gin.Context)
	CreateAccount(c *gin.Context)
	UpdateBalance(c *gin.Context)
	Transfer(c *gin.Context)
}

type controller struct {
	repo        Repository
	integration integrations.Integration
}

// função para receber os métodos da interface
func NewController(repo Repository) Controller {
	return &controller{
		repo: repo,
	}
}

func (ctl *controller) CreateUser(c *gin.Context) {
	UserRequest := dto.UserRequest{}
	c.ShouldBindJSON(&UserRequest)

	documentInt := treatDoc(UserRequest.CpfCnpj)

	user := model.User{
		CpfCnpj:    int64(documentInt),
		FullName:   UserRequest.FullName,
		Email:      UserRequest.Email,
		CategoryID: UserRequest.CategoryID,
		Password:   UserRequest.Password,
	}

	err := ctl.repo.CreateNewUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, "failed to create user!")
	}

	c.JSON(http.StatusOK, "User created sucessfully!")
}

func (ctl *controller) CreateAccount(c *gin.Context) {

	balanceRequest := dto.AccountRequest{}
	c.ShouldBindJSON(&balanceRequest)

	documentInt := treatDoc(balanceRequest.CpfCnpj)

	balance := model.Account{
		CpfCnpj: int64(documentInt),
		Balance: balanceRequest.Balance,
	}

	err := ctl.repo.CreateNewAccount(balance)
	if err != nil {
		c.JSON(http.StatusBadRequest, "failed to create account!")
	}

	c.JSON(http.StatusOK, "account created successfully!")
}

func (ctl *controller) UpdateBalance(c *gin.Context) {
	balanceRequest := dto.AccountRequest{}
	c.ShouldBindJSON(&balanceRequest)

	documentInt := treatDoc(balanceRequest.CpfCnpj)

	balance := model.Account{
		CpfCnpj: int64(documentInt),
		Balance: balanceRequest.Balance,
	}
	err := ctl.repo.UpdateAccountBalance(balance, documentInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Failed to update balance!")
	}

	balanceResponse := model.Account{}
	balanceResponse.CpfCnpj = int64(documentInt)

	c.JSON(http.StatusOK, "Balance update successfully!")
}

func (ctl *controller) Transfer(c *gin.Context) {
	documentPayerInt, _ := tools.ConvertStrToInt(c.Param("doc"))

	transferRequest := dto.TransferRequest{}
	c.ShouldBindJSON(&transferRequest)

	documentPayee := treatDoc(transferRequest.CpfPayee)

	accountPayer, err := ctl.repo.GetAccountPayer(documentPayerInt)
	fmt.Print(accountPayer, err)

	accountReceiver, _ := ctl.repo.GetAccountReceiver(documentPayee)
	userPayer, _ := ctl.repo.GetUserPayer(documentPayerInt)

	balanceError := ctl.integration.ValidateTransfer(accountPayer.Balance, transferRequest.Value)
	if balanceError != nil {
		c.JSON(http.StatusBadRequest, balanceError.Error())
		return
	}

	sellerError := ctl.integration.ValidateIsCommon(userPayer.CategoryID)
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

	data, errValid := ctl.integration.ValidateTransaction()
	if errValid != nil {
		c.JSON(http.StatusBadRequest, sellerError.Error())
		return
	}

	err = ctl.repo.CreateTransaction(transaction)
	if err != nil {
		c.JSON(http.StatusBadRequest, sellerError.Error())
		return
	}

	if data.Authorization {
		transaction.IdStatus = common.STATUS_CONCLUIDO
		ctl.DebitScheme(accountPayer, accountReceiver, transaction.Value)
		ctl.repo.UpdateStatusId(transaction.ID)

		c.JSON(http.StatusOK, transaction)
	} else {
		fmt.Println("failed to authorize transaction", errValid)
	}
}

func treatDoc(doc string) int {
	documentUnmasked := tools.RemoveMask(&doc)
	documentPayeeInt, _ := tools.ConvertStrToInt(documentUnmasked)
	return documentPayeeInt
}

func (ctl *controller) DebitScheme(Payer, Payee model.Account, value int) {
	newBalance := Payer.Balance - value
	ctl.repo.RemoveMoney(int(Payer.CpfCnpj), newBalance)

	addBalance := Payee.Balance + value
	ctl.repo.AddMoney(int(Payee.CpfCnpj), addBalance)
}
