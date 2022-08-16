package main

import (
	"database/sql"
	"fmt"
	"log"

	domain "github.com/RafaelDalarosa/fc-bank/domain/entities"
	"github.com/RafaelDalarosa/fc-bank/domain/usecase"
	"github.com/RafaelDalarosa/fc-bank/infra/repository"
	_ "github.com/lib/pq"
)

func main() {
	db := setupDb()
	defer db.Close()

	cc := domain.NewCreditCard()
	cc.Number = "1234"
	cc.Name = "RafaelDalarosa"
	cc.ExpirationYear = 2021
	cc.ExpirationMonth = 7
	cc.CVV = 123
	cc.Limit = 1000
	cc.Balance = 0

	repo := repository.NewTransactionRepositoryDb(db)
	err := repo.CreateCreditCard(*cc)
	if err != nil {
		fmt.Println(err)
	}
}

func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)
	return useCase
}

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"10.255.254.95",
		"5433",
		"postgres",
		"adeee44c40c89bfca362ed663ab9675b83ee16bd655cfd1dd942513d53392d44",
		"bank",
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error connection to database")
	}
	return db
}
