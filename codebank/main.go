package main

import (
	"fmt"
	"database/sql"
	"log"
	_"github.com/lib/pg"
)

func main() {
	db := setupDb()
	defer db.Close()
	
	cc := domain.NewCreditCard()
	cc.Number = "1234"
	cc.Name = "Wesley"
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

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
							"db",
							"5432",
							"postgres",
							"root",
							"codebank",
						 )
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error connecting database")
	}
	return db
}func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDb(db)
	useCase := useCase.NewUseCaseTransaction(transactionRepository)
	return useCase
}