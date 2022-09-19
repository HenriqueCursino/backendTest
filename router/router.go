package router

import (
	"github.com/gin-gonic/gin"
	"github.com/henriquecursino/desafioQ2/controller"
	"github.com/henriquecursino/desafioQ2/integrations"
	"github.com/henriquecursino/desafioQ2/structure"
)

func Router() {
	router := gin.Default()

	db := structure.Connect()
	repo := controller.NewRepository(db)
	inte := integrations.NewIntegration()

	controller := controller.NewController(repo, inte)

	router.POST("/createUser", controller.CreateUser)

	router.POST("/depositUser", controller.CreateAccount)
	router.POST("/transfer/:doc", controller.Transfer)
	router.PUT("/updateBalance", controller.UpdateBalance)

	router.Run(":8080")
}
