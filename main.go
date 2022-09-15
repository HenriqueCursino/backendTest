package main

import (
	"github.com/henriquecursino/desafioQ2/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	print(err)
	router.Router()
}
