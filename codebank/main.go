package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/srpulga/CodeBank/domain"
	"github.com/srpulga/CodeBank/infrastructure/repository"
	"github.com/srpulga/CodeBank/usecase"
	"log"
)

func main() {
	db := setupDb()
	defer db.Close()

	cc := domain.NewCreditCard()
	cc.Number = "1234567890123456"
	cc.Name = "Pulga"
	cc.ExpirationMonth = 12
	cc.ExpirationYear = 29
	cc.CVV = 123
	cc.Limit = 10000
	cc.Balance = 0

	repo := repository.NewTransactionRepositoryDB(db)
	err := repo.CreateCreditCard(*cc)
	if err != nil {
		fmt.Println(err)
	}
}

func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDB(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)
	return useCase
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
		log.Fatal("Error Connection to Database")
	}
	return db
}
